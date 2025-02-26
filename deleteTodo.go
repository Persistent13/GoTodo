package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func DeleteTodo(ctx echo.Context) error {
	id := ctx.Param("id")

	if _, err := db.Exec("DELETE FROM todos WHERE id = ?", id); err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to delete data")
	}

	return ctx.NoContent(http.StatusOK)
}
