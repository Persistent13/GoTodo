package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func UpdateTodo(ctx echo.Context) error {
	idStr := ctx.Param("id")
	if id, err := strconv.ParseUint(idStr, 10, 64); err != nil {
		return ctx.JSON(http.StatusInternalServerError, "id must be a positive integer")
	} else {
		var todo PatchTodoPogo
		if err := ctx.Bind(&todo); err != nil {
			return ctx.JSON(http.StatusInternalServerError, "Invalid JSON")
		}
		todo.ID = uint(id)

		queryPrefix := "UPDATE todos SET "
		query := queryPrefix
		var params []interface{}

		if todo.Done != nil {
			query += "done = ?, "
			params = append(params, todo.Done)
		}
		if todo.Content != nil {
			query += "content = ?, "
			params = append(params, todo.Content)
		}
		if todo.IsDeleted != nil {
			query += "isDeleted = ?, "
			params = append(params, todo.IsDeleted)
		}
		if query == queryPrefix {
			return ctx.NoContent(http.StatusUnprocessableEntity)
		}
		query = query[:len(query)-2] // Remove the trailing comma and space
		query += " WHERE id = ?"
		params = append(params, id)

		if _, err := db.Exec(query, params...); err != nil {
			return ctx.JSON(http.StatusInternalServerError, "Failed to update data")
		}

		return ctx.NoContent(http.StatusOK)
	}
}
