package routes

import(
  controller "github.com/IAmRiteshKoushik/kalaakar-backend/controllers"

  "github.com/gin-gonic/gin"
)

func BuyerRoutes(incomingRoutes *gin.Engine) {
  incomingRoutes.POST("/login/buyer", controller.BuyerLogin())
  incomingRoutes.POST("/signup/buyer", controller.BuyerSignUp())
}

func SellerRoutes(incomingRoutes *gin.Engine) {
  incomingRoutes.POST("/login/seller", controller.SellerLogin())
  incomingRoutes.POST("signup/seller", controller.SellerSignUp())

}
