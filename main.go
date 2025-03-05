package main

import (
	"database/sql"
	"errors"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/labstack/echo/v4"
	_ "github.com/labstack/echo/v4/middleware"
	"gotodo/constants"
	"gotodo/todo"
	"log"
	_ "modernc.org/sqlite"
)

func dbMiddleware(db *sql.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			ctx.Set(constants.DbContextKey, db)
			return next(ctx)
		}
	}
}

func main() {
	db, err := sql.Open("sqlite", "./todos.db")
	if err != nil {
		log.Fatal(err)
	}

	driver, err := sqlite.WithInstance(db, &sqlite.Config{})
	if err != nil {
		log.Fatal(err)
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
	e.Use(dbMiddleware(db))
	e.POST("/api/todo", todo.CreateTodo)
	e.GET("/api/todo", todo.ReadTodo)
	e.PATCH("/api/todo/:id", todo.UpdateTodo)
	e.DELETE("/api/todo/:id", todo.DeleteTodo)
	e.File("/", "static/index.html")

	if err := e.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}
