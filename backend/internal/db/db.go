package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

// connection pgx
// insert user, get all, get UserById/Username

// may not be same as es module's user struct
// might have extra fields here (the non searchable ones)
type User struct {
	Username string  `json:"username"`
	Age      int     `json:"age"`
	Gender   string  `json:"gender"`
	Lat      float64 `json:"location_lat"`
	Lon      float64 `json:"location_lon"`
}

func ConnectPG(ctx context.Context) *pgxpool.Pool {
	pool, err := pgxpool.New(ctx, "postgres://user:passw@localhost:5432/dbname")
	if err != nil {
		log.Fatal("couldn't connect to postgres", err)
	}

	if err = pool.Ping(ctx); err != nil {
		// shall we try having exponential backoff pinging ?
		log.Fatal("copuldnt connect to DB temporarily", err)
	}

	log.Print("connected to postgres successfully")
	return pool

}

func InsertUser(user User, ctx context.Context, pool *pgxpool.Pool) (err error) {

	conn, err := pool.Acquire(ctx)
	if err != nil {
		log.Print("error acquiring conn", err)
		return err
	}
	defer conn.Release()

	stmt_name := "insert_user"
	_, err = conn.Conn().Prepare(ctx, stmt_name, "INSERT INTO users (username, age, gender, location_lat, location_lon) VALUES ($1, $2, $3, $4, $5) RETURNING id")

	if err != nil {
		return fmt.Errorf("error preparing statement: %w", err)
	}

	_, err = conn.Exec(ctx, stmt_name, user.Username, user.Age, user.Gender, user.Lat, user.Lon)
	if err != nil {
		return err
	}
	return nil

}
