package main

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"time"
)

type (
	Todo struct {
		ID           uint
		Content      string
		CreatedAtUtc uint
		UpdatedAtUtc uint
		Done         bool
		IsDeleted    bool
	}

	PatchTodoPogo struct {
		ID        *uint   `json:"id"`
		Content   *string `json:"content"`
		Done      *bool   `json:"done"`
		IsDeleted *bool   `json:"isDeleted"`
	}

	CreateTodoPogo struct {
		Content *string `json:"content"`
	}
)

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("sqlite3", "./todos.db")
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.POST("/api/todo", CreateTodo)
	e.GET("/api/todo", ReadTodo)
	e.PATCH("/api/todo/:id", UpdateTodo)
	e.DELETE("/api/todo/:id", DeleteTodo)
	e.GET("/", PrintHelp)

	if err = e.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}

func PrintHelp(ctx echo.Context) error {
	html := `
<!DOCTYPE html>
<html lang="en">

<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>Todos API Demo</title>
  <style>
    body{max-width:650px;margin:40px auto;padding:0 10px;font:18px/1.5 -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, "Noto Sans", sans-serif, "Apple Color Emoji", "Segoe UI Emoji", "Segoe UI Symbol", "Noto Color Emoji";color:#444}h1,h2,h3{line-height:1.2}@media (prefers-color-scheme: dark){body{color:#c9d1d9;background:#0d1117}a:link{color:#58a6ff}a:visited{color:#8e96f0}}
  </style>
</head>

<body>
  <h1>How to use this</h1>
  <p>
  Paths include:
  <ol>
    <li>POST http://localhost:8080/api/todo</li>
    <li>GET http://localhost:8080/api/todo</li>
    <li>PATCH http://localhost:8080/api/todo/:id</li>
    <li>DELETE http://localhost:8080/api/todo/:id</li>
  </ol>
  </p>
  <p>
    The JSON structure for POST http://localhost:8080/api/todo is: 
  </p>
  <p>
    {
      "content": "Hello World!"
    }
  </p>
  <p>
    The JSON structure for PATCH http://localhost:8080/api/todo/:id is: 
  </p>
  <p>
    {
      "id", "123"
      "content": "Hello (New) World!",
      "Done": "false",
      "IsDeleted": "true"
    }
  </p>
</body>

</html>
`
	return ctx.HTML(http.StatusOK, html)
}

func CreateTodo(ctx echo.Context) error {
	var result sql.Result
	var row int64
	todo := new(CreateTodoPogo)

	if err = ctx.Bind(todo); err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Invalid JSON")
	}

	todoTime := time.Now().Unix()
	if result, err = db.Exec("INSERT INTO todos VALUES ($1, $2, $3, $4, $5)", todo.Content, todoTime, todoTime, false, false); err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to insert data")
	}
	if row, err = result.LastInsertId(); err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to insert data")
	}

	return ctx.JSON(http.StatusOK, row)
}

func ReadTodo(ctx echo.Context) error {
	count := ctx.QueryParams().Get("count")

	var query string
	var rows *sql.Rows
	if count != "" {
		query = "SELECT * FROM todos WHERE IsDeleted = 0 LIMIT ?"
	} else {
		query = "SELECT * FROM todos WHERE IsDeleted = 0 LIMIT 5"
	}

	if rows, err = db.Query(query, count); err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to query database")
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var item Todo
		if err = rows.Scan(&item.ID, &item.Content); err != nil {
			return ctx.JSON(http.StatusInternalServerError, "Failed to map database object")
		}
	}

	err = rows.Err()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to map database object")
	}

	return ctx.JSON(http.StatusOK, todos)
}

func UpdateTodo(ctx echo.Context) error {
	id := ctx.Param("id")
	var todo PatchTodoPogo
	if err := ctx.Bind(&todo); err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Invalid JSON")
	}
	if todo.ID == nil {
		return ctx.JSON(http.StatusInternalServerError, "Invalid todo ID")
	}

	query := "UPDATE todo SET"
	var params []interface{}

	if todo.Done != nil {
		query += "Done = ?, "
	}
	if todo.Content != nil {
		query += "Content = ?"
	}
	if todo.IsDeleted != nil {
		query += "IsDeleted = ?"
	}
	query += query[:len(query)-2] // Remove the trailing comma and space
	query += " WHERE id = ?"
	params = append(params, id)

	if _, err = db.Exec(query, params...); err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to update data")
	}

	return ctx.NoContent(http.StatusOK)
}

func DeleteTodo(ctx echo.Context) error {
	id := ctx.Param("id")

	if _, err = db.Exec("DELETE FROM todos WHERE id = ?", id); err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to delete data")
	}

	return ctx.NoContent(http.StatusOK)
}
