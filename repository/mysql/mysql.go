package mysql

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

type dbReadWriter struct {
	db *sql.DB
}

type DBConfig struct {
	SchemaURL       string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

func NewDBReadWriter(dbConfig *DBConfig) (*dbReadWriter, error) {
	db, err := sql.Open("mysql", dbConfig.SchemaURL)
	if err != nil {
		logrus.WithError(err).Error("Failed to connect to MySQL")
		return nil, err
	}

	db.SetMaxOpenConns(dbConfig.MaxOpenConns)
	db.SetMaxIdleConns(dbConfig.MaxIdleConns)
	db.SetConnMaxLifetime(dbConfig.ConnMaxLifetime)

	err = db.Ping()
	if err != nil {
		logrus.WithError(err).Error("Failed to ping MySQL database")
		db.Close()
		return nil, err
	}

	return &dbReadWriter{db: db}, nil
}

func (d *dbReadWriter) Close() error {
	if d.db != nil {
		if err := d.db.Close(); err != nil {
			return err
		}
	}
	return nil
}

func (d *dbReadWriter) Ping(ctx context.Context) error {
	if d.db == nil {
		return sql.ErrConnDone
	}
	return d.db.PingContext(ctx)
}

func (d *dbReadWriter) rollbackTx(tx *sql.Tx) error {
	if err := tx.Rollback(); err != nil && err != sql.ErrTxDone {
		return err
	}
	return nil
}

func (d *dbReadWriter) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	return d.db.BeginTx(ctx, opts)
}
