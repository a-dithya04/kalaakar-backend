package routes

import(
  controller "github.com/IAmRiteshKoushik/kalaakar-backend/controllers"

  "github.com/gin-gonic/gin"
)

func BuyerProfileRoutes(incomingRoutes *gin.Engine) {
  
  incomingRoutes.GET("/buyer/profile/:buyer_id", controller.GetBuyerProfile())
  incomingRoutes.POST("buyer/profile/dp", controller.UpdateBuyerProfileImage())
  incomingRoutes.POST("buyer/profile", controller.UpdateBuyerDetails())
  incomingRoutes.POST("buyer/profile/status", controller.DisableBuyerAccount())
  incomingRoutes.DELETE("buyer/profile", controller.DeleteBuyerProfileImage())
  incomingRoutes.DELETE("buyer/profile", controller.DeleteBuyerAccount())
}

