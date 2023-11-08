package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

  database "github.com/IAmRiteshKoushik/kalaakar-backend/database"
  routes "github.com/IAmRiteshKoushik/kalaakar-backend/routes"
  middleware "github.com/IAmRiteshKoushik/kalaakar-backend/middleware"
)

func main() {

  err := godotenv.Load(".env")
  if err != nil {
    log.Fatal("Environmental variables are not accessible.")
  }
  
  port := os.Getenv("PORT")
  if port == "" {
    port = "8080"
  }

  // Setting up the gin-router
  router := gin.New()
  router.Use(gin.Logger())

  // validate routes
  routes.BuyerRoutes(router)
  routes.SellerRoutes(router)

  router.User(middleware.BuyerAuth())
  router.User(middleware.SellerAuth())

  // routes
  routes.BuyerProfileRoutes(router)
  routes.SellerProfileRoutes(router)
  routes.CartRoutes(router)
  routes.MarketplaceRoutes(router)
  routes.BuyerProductRoutes(router)
  routes.SellerProductRoutes(router)

  // to work upon later
  routes.SellerAnalyticsRoutes(router)
  routes.ProductAnalyticsRoutes(router)

}
