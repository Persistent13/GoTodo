package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

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
