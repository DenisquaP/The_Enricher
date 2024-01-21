package postgres

import (
	"context"
	"enricher/database/models"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5"
)

type Row struct {
	UserID int
	models.User
}

var rowSlice []Row

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

func (p PostgresDB) DeleteUser(user_id int) error {
	query := fmt.Sprintf("DELETE FROM users WHERE user_id = %v", user_id)

	_, err := p.client.Exec(p.ctx, query)
	if err != nil {
		return fmt.Errorf("unable to delete an entry: %v", err)
	}

	return nil
}

func (p PostgresDB) GetUsers() ([]Row, error) {
	query := "SELECT * FROM users"

	res, err := p.client.Query(p.ctx, query)
	if err != nil {
		return nil, err
	}

	for res.Next() {
		var r Row
		err := res.Scan(&r.UserID, &r.Name, &r.Surname, &r.Patronymic, &r.Age, &r.Gender, &r.Nationality)
		if err != nil {
			log.Fatal(err)
		}
		rowSlice = append(rowSlice, r)
	}

	return rowSlice, nil
}

func (p PostgresDB) GetUsersByFilter(filterTag, filter string) ([]Row, error) {
	query := fmt.Sprintf("SELECT * FROM users WHERE %v = %v", filterTag, filter)

	res, err := p.client.Query(p.ctx, query)
	if err != nil {
		return nil, err
	}

	for res.Next() {
		var r Row
		err := res.Scan(&r.UserID, &r.Name, &r.Surname, &r.Patronymic, &r.Age, &r.Gender, &r.Nationality)
		if err != nil {
			log.Fatal(err)
		}
		rowSlice = append(rowSlice, r)
	}

	return rowSlice, nil
}

func (p PostgresDB) Close() error {
	err := p.client.Close(p.ctx)
	if err != nil {
		return err
	}

	return nil
}

func (p *PostgresDB) MigrationsUp(url ...string) error {
	var sourceURL string
	if url == nil {
		sourceURL = "file://database/migrations"
		fmt.Println(os.ReadDir("database/migrations"))
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
