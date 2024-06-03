package main

import (
	"calendar/internals/handler"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

var (
	PORT string
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetPrefix("calendar: ")
	log.SetOutput(os.Stderr)

	if os.Getenv("ENV") == "dev" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	if os.Getenv("ENV") == "dev" {
		log.Println("dev mode")
		PORT = ":3000"
	} else {
		log.Println("prod mode")
		PORT = fmt.Sprintf(":%s", os.Getenv("PORT"))
	}

	// add additional code here
}

func main() {

	log.Println(PORT)

	engine := html.New("./template/", ".tpl")

	engine.Reload(false) // Optional. Default: false

	engine.AddFunc("mod", func(i, j int) int { return i % j })
	engine.AddFunc("add1", func(i int) int { return i + 1 })

	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
	})

	// serve static files
	app.Static("/assets", "./assets")
	app.Static("/files", "./files")

	// routes
	app.Get("/month", handler.GetHandleMonth)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		})
	})

	app.Get("/romain", handler.GetHandleRomain)

	app.Get("ical", handler.GetHandleIcal)

	log.Panicln(app.Listen(PORT))

}
