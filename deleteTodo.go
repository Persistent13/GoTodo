package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func DeleteTodo(ctx echo.Context) error {
	id := ctx.Param("id")

	if _, err := db.Exec("UPDATE todos SET isDeleted = 1 WHERE id = ?", id); err != nil {
		msg := new(Error)
		msg.Message = "Failed to delete data in database"
		msg.Code = http.StatusInternalServerError
		return ctx.JSON(http.StatusInternalServerError, msg)
	}

	return ctx.NoContent(http.StatusOK)
}
