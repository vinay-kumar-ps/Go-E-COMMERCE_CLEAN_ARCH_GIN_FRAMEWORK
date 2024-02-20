package routes

import (
	"ecommerce/pkg/api/handler"
	"ecommerce/pkg/api/middleware"

	"github.com/gin-gonic/gin"
)

// "golang.org/x/tools/go/callgraph/rta"
// "golang.org/x/tools/go/callgraph/rta"

	
func AdminRoutes(engine *gin.RouterGroup,
	adminHandler *handler.AdminHandler,
	inventoryHandler *handler.InventoryHandler,
	userHandler *handler.UserHandler,
	categoryHandler *handler.CategoryHandler,
	orderHandler *handler.OrderHandler,
	couponHandler *handler.CouponHandler,
	offerHandler *handler.OfferHandler) {

		engine.POST("/adminlogin",adminHandler.LoginHandler)
		// api := router.Group("/admin_panel", middleware.AuthorizationMiddleware)
	// api.GET("users", adminHandler.GetUsers)
	
engine.Use(middleware.AdminAuthMiddleware)
{
	usermanagement := engine.Group("/users")
	{
		usermanagement.GET("",adminHandler.GetUsers)
		usermanagement.PUT("/block",adminHandler.BlockUser)
		usermanagement.PUT("/unblock",adminHandler.UnBlockUser)
	}
	categorymanagement :=engine.Group("/category")
}

}
	