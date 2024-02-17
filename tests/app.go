package test

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	handler "github.com/shivambaku/go-web-templ-htmx-tailwind-demo/handlers"
	"github.com/shivambaku/go-web-templ-htmx-tailwind-demo/internal/database"
	"github.com/testcontainers/testcontainers-go"
)

type TestApp struct {
	Handler     handler.Handler
	dbContainer testcontainers.Container
}

func NewTest() *TestApp {
	dbContainer, db := SetupTestDatabase()

	h := handler.Handler{
		DB: database.New(db),
	}

	return &TestApp{
		Handler:     h,
		dbContainer: dbContainer,
	}
}

func (t *TestApp) Close() {
	if err := t.dbContainer.Terminate(context.Background()); err != nil {
		log.Fatalf("failed to terminate container: %s", err)
	}
}

func StructToJSONBody(param interface{}) *bytes.Reader {
	body, _ := json.Marshal(param)
	return bytes.NewReader(body)
}
