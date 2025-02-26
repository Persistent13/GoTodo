package main

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"time"
)

func CreateTodo(ctx echo.Context) error {
	var result sql.Result
	var row int64
	var createErr error
	todo := new(CreateTodoPogo)

	if err := ctx.Bind(todo); err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Invalid JSON")
	}

	todoTime := time.Now().Unix()
	if result, createErr = db.Exec("INSERT INTO todos (content, createdAtUtc, updatedAtUtc, done, isDeleted) VALUES ($1, $2, $3, $4, $5)", todo.Content, todoTime, todoTime, false, false); createErr != nil {
		log.Error(createErr)
		return ctx.JSON(http.StatusInternalServerError, "Failed to insert data")
	}
	if row, createErr = result.LastInsertId(); createErr != nil {
		log.Error(createErr)
		return ctx.JSON(http.StatusInternalServerError, "Failed to insert data")
	}

	return ctx.JSON(http.StatusOK, row)
}
