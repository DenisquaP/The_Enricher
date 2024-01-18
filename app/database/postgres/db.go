package postgres

import (
	"context"
	"enricher/database/models"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Postgres interface {
	Connection()
	InsertUser() error
	SelectUser(email string) error
	MigrationsUp(url ...string) error
	Close() error
}

type PostgresDB struct {
	Config
	url    string
	ctx    context.Context
	client *pgx.Conn
}

func NewPostgres() (PostgresDB, error) {
	config := NewConfig()
	url := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", config.User, config.Password, config.Host, config.Port, config.DbName)

	return PostgresDB{
		Config: config,
		url:    url,
		ctx:    context.Background(),
	}, nil
}

func (p *PostgresDB) Connection() error {
	db, err := pgx.Connect(p.ctx, p.url)
	if err != nil {
		return err
	}

	p.client = db
	return nil
}

func (p PostgresDB) InsertUser(user models.User) error {
	// line := fmt.Sprintf("INSERT INTO users (name, surname, patronymic, age, gender) values (%s, %s, %s, %d, %s)", user.Name, user.Surname, user.Patronymic, user.Age, user.Gender)
	query := `INSERT INTO users (name, surname, patronymic, age, gender) VALUES (@name, @surname, @patronymic, @age, @gender)`
	args := pgx.NamedArgs{
		"name":       user.Name,
		"surname":    user.Surname,
		"patronymic": user.Patronymic,
		"age":        user.Age,
		"gender":     user.Gender,
	}

	_, err := p.client.Exec(p.ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to insert row: %v", err)
	}
	// tr, err := p.client.Begin(p.ctx)
	// if err != nil {
	// 	return err
	// }
	// defer tr.Rollback(p.ctx)

	// if err := tr.QueryRow(p.ctx, line); err != nil {
	// 	return errors.New(fmt.Sprint(err))
	// }

	// err = tr.Commit(p.ctx)
	// if err != nil {
	// 	return err
	// }

	return nil

}

func (p PostgresDB) Close() error {
	err := p.client.Close(p.ctx)
	if err != nil {
		return err
	}

	return nil
}
