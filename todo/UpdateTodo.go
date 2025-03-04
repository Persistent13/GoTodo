package todo

import (
	"awesomeProject/constants"
	"database/sql"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func UpdateTodo(ctx echo.Context) error {
	db := ctx.Get(constants.DbContextKey).(*sql.DB)
	idStr := ctx.Param("id")
	if id, err := strconv.ParseUint(idStr, 10, 64); err != nil {
		msg := new(Error)
		msg.Code = http.StatusBadRequest
		msg.Message = "id must be a positive integer"
		return ctx.JSON(http.StatusBadRequest, msg)
	} else {
		var todo PatchTodoPogo
		if err := ctx.Bind(&todo); err != nil {
			msg := new(Error)
			msg.Code = http.StatusBadRequest
			msg.Message = "Invalid JSON body"
			return ctx.JSON(http.StatusBadRequest, msg)
		}
		todo.ID = uint(id)

		// This breakup of the SQL code is done to support patching.
		// By checking which values were set, we can build the query to update only what was sent.
		// At the end, we check to see if the query was updated, if not, that means nothing to patch and we
		// then inform the consumer that this was a bad request.
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
			msg := new(Error)
			msg.Code = http.StatusInternalServerError
			msg.Message = "Failed to update data in database"
			return ctx.JSON(http.StatusInternalServerError, msg)
		}

		return ctx.NoContent(http.StatusOK)
	}
}
