package main

import (
	"context"
	"net/http"

	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	logger      = initLogger()
	rdbContext  = context.TODO()
	dbContext   = context.Background()
	tokenSecret string
)

func main() {

	conf := LoadConfigFromEnv()
	tokenSecret = conf.TokenSecret

	todoDB, err := pgxpool.Connect(context.Background(), conf.DBTodo.DSN)
	failOnError(err, "failed to init todo db connection")
	// FIXME for newer version of pgx lib, you can use Ping function instead (that is added to lib from https://github.com/jackc/pgx/commit/aa8604b5c22989167e7158ecb1f6e7b8ddfebf04)
	// err = todoDB.Ping(context.Background())
	var stt int
	err = todoDB.QueryRow(context.Background(), "SELECT 1").Scan(&stt)
	failOnError(err, "failed to connect to todo db")

	storeTodo := &StoreTodo{
		DB: todoDB,
	}

	contextHandler := ContextHandler{
		Config: &conf,
		Store:  storeTodo,
	}

	h := NewHandler(&contextHandler, logger)

	r := newRouter()
	h.Register(r)

	_ = http.ListenAndServe(conf.Service.Address, r)
}

func failOnError(err error, msg string) {
	if err != nil {
		logger.Fatalw(msg, "error", err)
	}
}
