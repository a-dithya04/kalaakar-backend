package routes

import(
  controller "github.com/IAmRiteshKoushik/kalaakar-backend/controllers"

  "github.com/gin-gonic/gin"
)

func BuyerProductRoutes(incomingRoutes *gin.Engine) {
  incomingRoutes.GET("", controller.GetProduct())
}

func SellerProductRoutes(incomingRoutes *gin.Engine) {
   
}
