package main

import (
	"log"

	"github.com/ghaidafasya24/bp-tubes/config"

	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2/middleware/cors"


	"github.com/ghaidafasya24/bp-tubes/url"


	"github.com/gofiber/fiber/v2"
	
)

func main() {
	site := fiber.New(config.Iteung)
	site.Use(cors.New(config.Cors))
	url.Web(site)
	log.Fatal(site.Listen(musik.Dangdut()))
}
