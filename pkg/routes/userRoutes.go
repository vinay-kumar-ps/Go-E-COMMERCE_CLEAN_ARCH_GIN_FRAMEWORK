package routes

import (
	"ecommerce/pkg/api/handler"

	"github.com/gin-gonic/gin"
)

func UserRoutes(engine *gin.RouterGroup,
userHandler *handler.UserHandler,
otpHandler *handler.OtpHandler,
inventoryHandler *handler.InventoryHandler
orderHandler *handler.OrderHandler
carthandler *handler.CartHandler
paymenthandler *handler.PaymentHandler
wishlistHandler *handler.WishlistHandler
categoryHandler *handler.CategoryHandler
couponHandler *handler.CouponHandler



)