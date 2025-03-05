package todo

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"gotodo/constants"
	"net/http"
	"time"
)

func CreateTodo(ctx echo.Context) error {
	var row int64
	db := ctx.Get(constants.DbContextKey).(*sql.DB)
	todo := new(CreateTodoPogo)

	if err := ctx.Bind(todo); err != nil {
		msg := new(Error)
		msg.Code = http.StatusBadRequest
		msg.Message = "Invalid JSON body"
		return ctx.JSON(http.StatusBadRequest, msg)
	}

	todoTime := time.Now().Unix()
	result, createErr := db.Exec("INSERT INTO todos (content, createdAtUtc, updatedAtUtc, done, isDeleted) VALUES ($1, $2, $3, $4, $5)", todo.Content, todoTime, todoTime, false, false)
	if createErr != nil {
		log.Error(createErr)
		msg := new(Error)
		msg.Code = http.StatusInternalServerError
		msg.Message = "Failed to create data in database"
		return ctx.JSON(http.StatusInternalServerError, msg)
	}
	if row, createErr = result.LastInsertId(); createErr != nil {
		log.Error(createErr)
		msg := new(Error)
		msg.Code = http.StatusInternalServerError
		msg.Message = "Failed to create data in database"
		return ctx.JSON(http.StatusInternalServerError, msg)
	}

	todoInfo := struct {
		ID int64 `json:"id"`
	}{
		ID: row,
	}
	return ctx.JSON(http.StatusOK, todoInfo)
}
