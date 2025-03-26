package main

import (
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fibre.App) {
	app.Get("/:url", routes.ResolveURL)
	app.Post("/api/v1", routes.ShortenURL)
}

func main() {
	
	app := fibre.New()
}