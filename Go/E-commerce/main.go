package main

import (
	"e-commerce/controllers"
	"e-commerce/database"
	"e-commerce/middleware"
	"e-commerce/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main(){
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))

	//verify gin package
	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.Get("/addcart", app.AddToCart())
	router.Get("/removeitem", app.RemoveItem())
	router.Get("/cartcheckout", app.BuyFromCart())
	router.Get("/instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))


}