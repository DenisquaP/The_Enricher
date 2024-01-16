package postgres

type Postgres interface {
	Connection() error
	InsertUser() (*string, error)
	SelectUser(email string) error
	MigrationsUp(url ...string) error
	Close() error
}

type postgres struct {
	Config
}

func NewPostgres() postgres {
	config := NewConfig()
	return postgres{
		config,
	}
}
