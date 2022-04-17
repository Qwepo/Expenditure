package db

import (
	config "app/internal/configs"
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

type DB interface {
	dbPayment
}

type Clietn interface {
	DB
	Close() error
}
type clietn struct {
	*pgx.Conn
}

func Open(conf *config.Config) (Clietn, error) {
	dbconf, err := pgx.ParseConfig(fmt.Sprintf("user=%s password=qwerty host=localhost port=%s dbname=%s sslmode=disable", conf.Db.Username, conf.Db.Port, conf.Db.DBname))
	if err != nil {
		return nil, err
	}
	conn, err := pgx.ConnectConfig(context.TODO(), dbconf)
	if err != nil {
		return nil, err
	}
	return &clietn{conn}, nil
}

func (db *clietn) Close() error {
	return db.Conn.Close(context.TODO())
}
