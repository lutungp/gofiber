package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/olekukonko/tablewriter"
)

func printtable() {
	data := [][]string{
		[]string{"A", "The Good", "500"},
		[]string{"B", "The Very very Bad Man", "288"},
		[]string{"C", "The Ugly", "120"},
		[]string{"D", "The Gopher", "800"},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Sign", "Rating"})

	for _, v := range data {
		table.Append(v)
	}

	table.Render()

	return c.SendString("print table")
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/printtable", func(c *fiber.Ctx) error {
		return printtable()
	})

	app.Listen(":3000")
}
