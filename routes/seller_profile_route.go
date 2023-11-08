package routes

import(
  controller "github.com/IAmRiteshKoushik/kalaakar-backend/controllers"

  "github.com/gin-gonic/gin"
)

func SellerProfileRoutes(incomingRoutes *gin.Engine) {
  
  incomingRoutes.GET("", controller.GetSellerProfile())
  incomingRoutes.POST("", controller.UpdateSellerProfileImage())
  incomingRoutes.POST("", controller.UpdateSellerBannerImage())
  incomingRoutes.POST("", controller.DisableSellerAccount())
  incomingRoutes.DELETE("", controller.DeleteSellerProfileImage())
  incomingRoutes.DELETE("", controller.DeleteSellerBannerImage())
  incomingRoutes.DELETE("", controller.DeleteSellerAccount())
}
