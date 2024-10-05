package make

import (
	"log"
	"os"
	routes "github.com/ante-neh/url_shortener/internal/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)


func main(){
	err := godotenv.Load() 
	if err != nil{
		log.Fatal("Unable to get environmental variable")
	}
	port := os.Getenv("PORT")
	app := fiber.New()
	app.Use(logger.New())
	err = app.Listen(port) 

	if err != nil{
		log.Fatal("Server down")
	}
}


func setupRoutes(app *fiber.App){
	app.Get("/:url", routes.ResolveUrl)
	app.Post("/api/v1", routes.ShortenUrl)
}