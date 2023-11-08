package routes

import(
  controller "github.com/IAmRiteshKoushik/kalaakar-backend/controllers"

  "github.com/gin-gonic/gin"
)

func CartRoutes(incomingRoutes *gin.Engine) {
  incomingRoutes.GET("buyer/cart", controller.GetCart())
  incomingRoutes.POST("buyer/cart/", controller.AddToCart())
  incomingRoutes.POST("buyer/cart/checkout", controller.Checkout())
  incomingRoutes.DELETE("buyer/cart", controller.RemoveItemFromCart())
  incomingRoutes.DELETE("buyer/cart", controller.ClearCart())

}
