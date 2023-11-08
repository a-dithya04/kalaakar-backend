package routes

import(
  controller "github.com/IAmRiteshKoushik/kalaakar-backend/controllers"

  "github.com/gin-gonic/gin"
)

func ProductAnalyticsRoutes(incomingRoutes *gin.Engine) {

}

func SellerAnalyticsRoutes(incomingRoutes *gin.Engine) {
  
  incomingRoutes.GET("seller/sales", controller.GetWeeklySale())
  incomingRoutes.GET("seller/sales", controller.GetMonthlySale())
  incomingRoutes.GET("seller/sales", controller.GetMostSoldProducts())
}
