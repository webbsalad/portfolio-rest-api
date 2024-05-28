package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/webbsalad/portfolio-rest-api/db"
	"github.com/webbsalad/portfolio-rest-api/db/operations"
)

func GetAllItemsRouter(dbConn *db.DBConnection) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tableName := c.Params("table_name")

		filters := make(map[string]string)
		c.QueryParser(&filters)

		sortBy := c.Query("sortBy", "")

		itemsJSON, err := operations.FetchDataAsJSON(dbConn, tableName, filters, sortBy)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		if itemsJSON == "[]" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "No items found"})
		}

		return c.Status(fiber.StatusOK).SendString(itemsJSON[1 : len(itemsJSON)-1])
	}
}
