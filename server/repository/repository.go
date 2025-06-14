package repository

import (
	"bookcabin/repository/mysql"
	"bookcabin/repository/repoiface"
	"fmt"
	"io"
	"time"
)

type Repository struct {
	dbReadWriter repoiface.DBReadWriter
	io.Closer
}

type DBConfig struct {
	User            string
	Password        string
	Host            string
	Port            string
	DBName          string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

type RepoConfig struct {
	DBConfig
}

func NewBookCabinRepository(cfg *RepoConfig) (*Repository, error) {
	schemaURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
	db, err := mysql.NewDBReadWriter(&mysql.DBConfig{
		SchemaURL:       schemaURL,
		MaxOpenConns:    cfg.MaxOpenConns,
		MaxIdleConns:    cfg.MaxIdleConns,
		ConnMaxLifetime: cfg.ConnMaxLifetime,
	})
	if err != nil {
		return nil, err
	}

	return &Repository{
		dbReadWriter: db,
	}, nil
}

func (r *Repository) Close() error {
	return r.dbReadWriter.Close()
}

func (r *Repository) HealthCheck() error {
	return r.dbReadWriter.Ping()
}
