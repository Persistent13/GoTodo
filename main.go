package main

import (
	"database/sql"
	"errors"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/labstack/echo/v4"
	"log"
	_ "modernc.org/sqlite"
)

type (
	Todo struct {
		ID           uint
		Content      string
		CreatedAtUtc uint
		UpdatedAtUtc uint
		Done         bool
		IsDeleted    bool
	}

	PatchTodoPogo struct {
		ID        uint    `json:"id"`
		Content   *string `json:"content"`
		Done      *bool   `json:"done"`
		IsDeleted *bool   `json:"isDeleted"`
	}

	CreateTodoPogo struct {
		Content string `json:"content"`
	}

	Error struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
	}
)

var db *sql.DB

func main() {
	var sqlErr error

	if db, sqlErr = sql.Open("sqlite", "./todos.db"); sqlErr != nil {
		log.Fatal(sqlErr)
	}

	driver, driverErr := sqlite.WithInstance(db, &sqlite.Config{})
	if driverErr != nil {
		log.Fatal(driverErr)
	}

	m, err := migrate.NewWithDatabaseInstance("file://./migrations", "sqlite", driver)
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); nil != err && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatal(err)
	}

	log.Print("Migrations successfully in place")

	e := echo.New()
	e.POST("/api/todo", CreateTodo)
	e.GET("/api/todo", ReadTodo)
	e.PATCH("/api/todo/:id", UpdateTodo)
	e.DELETE("/api/todo/:id", DeleteTodo)
	e.File("/", "static/index.html")

	if err := e.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}
