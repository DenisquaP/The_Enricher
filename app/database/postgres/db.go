package postgres

import (
	"context"
	"enricher/database/models"
	"fmt"

	"github.com/golang-migrate/migrate"
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
	query := `INSERT INTO users (name, surname, patronymic, age, gender, nationality) VALUES (@name, @surname, @patronymic, @age, @gender, @nationality)`
	args := pgx.NamedArgs{
		"name":        user.Name,
		"surname":     user.Surname,
		"patronymic":  user.Patronymic,
		"age":         user.Age,
		"gender":      user.Gender,
		"nationality": user.Nationality,
	}

	_, err := p.client.Exec(p.ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to insert row: %v", err)
	}

	return nil
}

func (p PostgresDB) UpdateUser(row, newValue string, user_id int) error {
	query := fmt.Sprintf("UPDATE users SET %v = '%v' WHERE user_id = %v", row, newValue, user_id)

	_, err := p.client.Exec(p.ctx, query)
	if err != nil {
		return fmt.Errorf("unable to change value: %v", err)
	}

	return nil
}

func (p PostgresDB) Close() error {
	err := p.client.Close(p.ctx)
	if err != nil {
		return err
	}

	return nil
}

func (p PostgresDB) migrationsUp(url ...string) error {
	var sourceURL string
	if url == nil {
		sourceURL = "file://database/migrations/up"
	} else {
		sourceURL = url[0]
	}
	m, err := migrate.New(sourceURL, p.url)
	if err != nil {
		return err
	}
	if err = m.Up(); err != nil {
		return err
	}

	return nil
}
