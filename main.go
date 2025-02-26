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
)

var db *sql.DB

func main() {
	var sqlErr error
	if db, sqlErr = sql.Open("sqlite", "./todos.db"); sqlErr != nil {
		log.Fatal(sqlErr)
	}

	if driver, driverErr := sqlite.WithInstance(db, &sqlite.Config{}); driverErr != nil {
		log.Fatal(driverErr)
	} else {
		if m, err := migrate.NewWithDatabaseInstance("file://./migrations", "sqlite", driver); err != nil {
			log.Fatal(err)
		} else {
			if err := m.Up(); nil != err && !errors.Is(err, migrate.ErrNoChange) {
				log.Fatal(err)
			} else {
				log.Print("Migrations successfully in place")
			}
		}
	}

	e := echo.New()
	e.POST("/api/todo", CreateTodo)
	e.GET("/api/todo", ReadTodo)
	e.PATCH("/api/todo/:id", UpdateTodo)
	e.DELETE("/api/todo/:id", DeleteTodo)
	e.GET("/", PrintHelp)

	if err := e.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}
