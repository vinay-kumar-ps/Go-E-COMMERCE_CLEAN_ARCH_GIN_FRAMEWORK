package api

import (
	_ "ecommerce/cmd/api/docs"
	handler "ecommerce/pkg/api/handler"

	"ecommerce/pkg/routes"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// http server for the web application
type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHttp(categoryHandler *handler.CategoryHandler,
	inventoryHandler *handler.InventoryHandler,
	userHandler *handler.UserHandler,
	otpHandler *handler.OtpHandler,
	adminHandler *handler.AdminHandler,
	cartHandler *handler.CartHandler,
	orderHandler *handler.OrderHandler,
	paymentHandler *handler.PaymentHandler,
	wishlistHandler *handler.WishlistHandler,
	offerHandler *handler.OfferHandler,
	couponHandler *handler.CouponHandler) *ServerHTTP {
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.LoadHTMLFiles("pkg/templates/*.html")
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	routes.UserRoutes(engine.Group("/users"), userHandler, otpHandler, inventoryHandler, cartHandler, orderHandler, couponHandler, paymentHandler, wishlistHandler)
	routes.AdminRoutes(engine.Group("/admin"), adminHandler, categoryHandler, inventoryHandler, orderHandler, paymentHandler, offerHandler, couponHandler)
	routes.InventoryRoutes(engine.Group("/products"), inventoryHandler)

	return &ServerHTTP{
		engine: engine,
	}
}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":3000")
}
