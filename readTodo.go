package main

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
)

func ReadTodo(ctx echo.Context) error {
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
		return ctx.JSON(http.StatusInternalServerError, "Failed to query database")
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var item Todo
		if err := rows.Scan(&item.ID, &item.Content, &item.CreatedAtUtc, &item.UpdatedAtUtc, &item.Done, &item.IsDeleted); err != nil {
			log.Error(err)
			return ctx.JSON(http.StatusInternalServerError, "Failed to map database object")
		}
		todos = append(todos, item)
	}

	if err := rows.Err(); err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusInternalServerError, "Failed to map database object")
	}

	return ctx.JSON(http.StatusOK, todos)
}
