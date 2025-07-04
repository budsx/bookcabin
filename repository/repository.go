package repository

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/budsx/bookcabin/repository/mysql"
	"github.com/budsx/bookcabin/repository/repoiface"
)

type Repository struct {
	DBReadWriter repoiface.DBReadWriter
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
		DBReadWriter: db,
	}, nil
}

func (r *Repository) Close() error {
	return r.DBReadWriter.Close()
}

func (r *Repository) HealthCheck() error {
	return r.DBReadWriter.Ping(context.Background())
}
