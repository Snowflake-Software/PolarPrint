package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/snowflake-software/polarprint/pkg/handlers"
)

func MetaRouter(app fiber.Router) {
	app.Get("/", handlers.RenderUI())
}
