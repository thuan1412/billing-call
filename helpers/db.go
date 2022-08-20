package helpers

import (
	"calling-bill/ent"
	"context"
	"database/sql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"fmt"
	"log"
	"os"
	"time"
)

var DbClient *ent.Client

// GetDb return an ent client
func GetDb() (*ent.Client, error) {
	dbURI := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PW"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"))

	fmt.Println("db url", dbURI)
	db, err := sql.Open(dialect.Postgres, dbURI)
	if err != nil {
		fmt.Println("failed to open connection to database:", err)
		return nil, err
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)

	driver := entsql.OpenDB(dialect.Postgres, db)
	if err != nil {
		return nil, err
	}
	client := ent.NewClient(ent.Driver(driver))

	// time logging for database mutations
	client.Use(func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			start := time.Now()
			defer func() {
				log.Printf("Op=%s\tType=%s\tTime=%s\tConcreteType=%T\n", m.Op(), m.Type(), time.Since(start), m)
			}()
			return next.Mutate(ctx, m)
		})
	})
	DbClient = client
	return client, nil
}

func MigrateDb(client *ent.Client) {
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalln("failed to create schema:", err)
	}
}
