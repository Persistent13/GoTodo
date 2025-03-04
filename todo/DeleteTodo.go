package todo

import (
	"awesomeProject/constants"
	"database/sql"
	"github.com/labstack/echo/v4"
	"net/http"
)

func DeleteTodo(ctx echo.Context) error {
	db := ctx.Get(constants.DbContextKey).(*sql.DB)
	id := ctx.Param("id")

	if _, err := db.Exec("UPDATE todos SET isDeleted = 1 WHERE id = ?", id); err != nil {
		msg := new(Error)
		msg.Message = "Failed to delete data in database"
		msg.Code = http.StatusInternalServerError
		return ctx.JSON(http.StatusInternalServerError, msg)
	}

	return ctx.NoContent(http.StatusOK)
}
