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
  router.BuyerRoutes(router)
  router.SellerRoutes(router)

  router.User(middleware.BuyerAuth())
  router.User(middleware.SellerAuth())

  // routes
  router.BuyerProfileRoutes(router)
  router.SellerProfileRoutes(router)
  router.CartRoutes(router)
  router.MarketplaceRoutes(router)
  router.BuyerProductRoutes(router)
  router.SellerProductRoutes(router)

  // to work upon later
  router.SellerAnalyticsRouter(router)
  router.ProductAnalyticsRouter(router)

}
