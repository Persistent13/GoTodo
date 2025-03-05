package todo

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"gotodo/constants"
	"net/http"
)

func ReadTodo(ctx echo.Context) error {
	db := ctx.Get(constants.DbContextKey).(*sql.DB)
	count := ctx.QueryParams().Get("count")

	var query string
	var rows *sql.Rows
	var queryErr error
	if count != "" {
		query = "SELECT * FROM todos WHERE isDeleted = 0 LIMIT ?"
	} else {
		query = "SELECT * FROM todos WHERE isDeleted = 0 LIMIT 5"
	}

	if rows, queryErr = db.Query(query, count); queryErr != nil {
		log.Error(queryErr)
		msg := new(Error)
		msg.Message = "Failed to query database"
		msg.Code = http.StatusInternalServerError
		return ctx.JSON(http.StatusInternalServerError, msg)
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var item Todo
		if err := rows.Scan(&item.ID, &item.Content, &item.CreatedAtUtc, &item.UpdatedAtUtc, &item.Done, &item.IsDeleted); err != nil {
			log.Error(err)
			msg := new(Error)
			msg.Message = "Failed to map database object"
			msg.Code = http.StatusInternalServerError
			return ctx.JSON(http.StatusInternalServerError, msg)
		}
		todos = append(todos, item)
	}

	if err := rows.Err(); err != nil {
		log.Error(err)
		msg := new(Error)
		msg.Message = "Failed to map database object"
		msg.Code = http.StatusInternalServerError
		return ctx.JSON(http.StatusInternalServerError, msg)
	}

	return ctx.JSON(http.StatusOK, todos)
}
