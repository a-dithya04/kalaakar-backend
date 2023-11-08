package routes

import(
  controller "github.com/IAmRiteshKoushik/kalaakar-backend/controllers"

  "github.com/gin-gonic/gin"
)

func MarketplaceRoutes(incomingRoutes *gin.Engine) {
  
  incomingRoutes.GET("buyer/marketplace", controller.GetMarketplace())
  incomingRoutes.GET("buyer/marketplace/:product_id", controller.GetProductPreview())
}
