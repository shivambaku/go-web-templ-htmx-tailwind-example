package test

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/shivambaku/go-web-templ-htmx-tailwind-demo/internal/database"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func SetupTestDatabase() (testcontainers.Container, *sql.DB) {
	containerReq := testcontainers.ContainerRequest{
		Image:        "postgres:16",
		ExposedPorts: []string{"5432/tcp"},
		WaitingFor:   wait.ForListeningPort("5432/tcp"),
		Env: map[string]string{
			"POSTGRES_DB":       "testdb",
			"POSTGRES_PASSWORD": "postgres",
			"POSTGRES_USER":     "postgres",
		},
	}

	dbContainer, err := testcontainers.GenericContainer(
		context.Background(),
		testcontainers.GenericContainerRequest{
			ContainerRequest: containerReq,
			Started:          true,
		},
	)
	if err != nil {
		log.Fatalf("Error starting container: %s", err)
	}

	host, err := dbContainer.Host(context.Background())
	if err != nil {
		log.Fatalf("Error getting container host: %s", err)
	}

	port, err := dbContainer.MappedPort(context.Background(), "5432")
	if err != nil {
		log.Fatalf("Error getting container port: %s", err)
	}

	dbURL := fmt.Sprintf("postgres://postgres:postgres@%v:%v/testdb?sslmode=disable", host, port.Port())
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error opening database: %s", err)
	}

	err = database.RunSQLFile(db, "../sql/schema.sql")
	if err != nil {
		log.Fatalf("Error running schema.sql: %s", err)
	}
	return dbContainer, db
}
