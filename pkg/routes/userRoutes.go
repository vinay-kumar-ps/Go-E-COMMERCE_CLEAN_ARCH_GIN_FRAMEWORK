package routes

import (
	"ecommerce/pkg/api/handlers"
	"ecommerce/pkg/api/middleware"

	"github.com/gin-gonic/gin"
	// "github.com/unidoc/unipdf/v3/core/security"
)

func UserRoutes(engine *gin.RouterGroup,
userHandler *handler.UserHandler,
otpHandler *handler.OtpHandler,
inventoryHandler *handler.InventoryHandler,
orderHandler *handler.OrderHandler,
carthandler *handler.CartHandler,
paymenthandler *handler.PaymentHandler,
wishlistHandler *handler.WishlistHandler,
categoryHandler *handler.CategoryHandler,
couponHandler *handler.CouponHandler) {
	
	engine.POST("/signup",userHandler.UserSignUp)
	engine.POST("/login",userHandler.LoginHandler)
	engine.GET("/forgot_password",userHandler.ForgotPasswordSend)
	engine.POST("/forgot-password",userHandler.ForgotPasswordVerifyAndChange)

	engine.POST("/otplogin",otpHandler.SendOTP)
	engine.POST("/verifyotp",otpHandler.VerifyOTP)

	payment :=engine.Group("/payment")
	{
		payment.GET("/razorpay",paymenthandler.MakePaymentRazorPay)
		payment.GET("/update_status",paymenthandler.VerifyPayment)

	}
	engine.Use(middleware.UserAuthMiddleware)
	{
		engine.GET("/banners",categoryHandler.GetBannersForUsers)

		search :=engine.Group("/search")
		{
			search.GET("/",inventoryHandler.SearchProducts)
		}
		home :=engine.Group("/home")
		{
			home.GET("/products",inventoryHandler.ListProductsForUser)
			home.GET("/products/details",inventoryHandler.ShowIndividualProducts)
			home.POST("/add-to-cart",carthandler.AddToCart)
			home.POST("/wishlist/add",wishlistHandler.AddToWishlist)
		}
		categorymanagement :=engine.Group("/category")
		{
			categorymanagement.GET("",categoryHandler.GetCategory)
			categorymanagement.GET("/c/products",categoryHandler.GetProductDetailsInACategory)
		}
		profile :=engine.Group("/profile")
		{
			profile.GET("/details",userHandler.GetUserDetails)
			profile.POST("/address",userHandler.AddAddress)
			profile.GET("/address",userHandler.GetAddresses)
			profile.GET("/reference-link",userHandler.GetMyReferenceLink)

			orders:= profile.Group("/orders")
			{
				orders.GET("",orderHandler.GetOrders)
				orders.GET("/:id",orderHandler.GetIndividualOrderDetails)
				orders.DELETE("",orderHandler.CancelOrder)
				orders.PUT("/return",orderHandler.ReturnOrder)
			}
			edit :=profile.Group("/edit")
			{
				edit.PUT("/name",userHandler.EditName)
				edit.PUT("/email",userHandler.EditEmail)
				edit.PUT("/phone",userHandler.EditPhone)
			}
			security :=profile.Group("/security")
			{
				security.PUT("/change-password",userHandler.ChangePassword)
			}
		}
		cart :=engine.Group("/cart")
		{
			cart.GET("/",userHandler.GetCart)
			cart.DELETE("/remove",userHandler.RemoveFromCart)
			cart.PUT("/ypdateQuantitiy/plus",userHandler.UpdateQuantityAdd)
			cart.PUT("/updateQuantiy/minus",userHandler.UpdateQuantityLess)
			//hellloo
		}
		wishlist := engine.Group("/wishlist")
		{
			wishlist.GET("/",wishlistHandler.GetWishList)
			wishlist.DELETE("/remove",wishlistHandler.RemoveFromWishlist)

		}
		checkout := engine.Group("/check-out")
		{
            checkout.GET("",carthandler.CheckOut)
			checkout.POST("order",orderHandler.OrderItemsFromCart)
		}
		engine.GET("/coupon",couponHandler.GetAllCoupons)
	

	}

}

