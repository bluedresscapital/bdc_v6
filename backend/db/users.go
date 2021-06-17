package db

import (
	"context"
	"fmt"
	"github.com/bluedresscapital/bdc_v6/backend/graph/model"
	"log"

	//"github.com/cockroachdb/cockroach-go/v2/crdb/crdbpgx"
	"github.com/jackc/pgx/v4"
)

func FetchUserByUsername(username string) (*model.User, error) {
	return nil, nil
}

func SeedUsers() {
	// its postgres://user:pwd@hostname:port/db?params
	config, err := pgx.ParseConfig("postgres://bdc:bdc@127.0.0.1:26257/bdc?sslmode=require")
	if err != nil {
		log.Fatal("error configuring the database: ", err)
	}

	// Connect to the "bank" database.
	conn, err := pgx.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}
	defer conn.Close(context.Background())

	// Create the "accounts" table.
	if _, err := conn.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			username VARCHAR(50) UNIQUE NOT NULL,
			password VARCHAR(50) NOT NULL
		)
	`); err != nil {
		log.Fatal(err)
	}

	if _, err := conn.Exec(context.Background(),
		"DELETE FROM users"); err != nil {
		log.Fatal(err)
	}

	// Insert two rows into the "accounts" table.
	if _, err := conn.Exec(context.Background(),
		"INSERT INTO users (username, password) VALUES ('jason', '123'), ('ralles', '234')"); err != nil {
		log.Fatal(err)
	}

	//// Print out the balances.
	//rows, err := conn.Query(context.Background(), "SELECT id, balance FROM accounts")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer rows.Close()
	//fmt.Println("Initial balances:")
	//for rows.Next() {
	//	var id, balance int
	//	if err := rows.Scan(&id, &balance); err != nil {
	//		log.Fatal(err)
	//	}
	//	fmt.Printf("%d %d\n", id, balance)
	//}

	fmt.Println("hello")
}
