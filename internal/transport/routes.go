package transport

import (
	"github.com/gin-gonic/gin"
	"github.com/itsvagapov/team-pharmacy/internal/service"
)

func RegisterRoutes(router *gin.Engine,
	 categoryService service.CategoryService, 
	 subcategoryService service.SubcategoryService, 
	 userService service.UserService,
	cartService service.CartService,
	orderService service.OrderService,
	) { // userService service.User....
	categoryHandler := NewCategoryHandler(categoryService)
	subcategoryHandler := NewSubcategoryHandler(subcategoryService)
	userHandler := NewUserHandler(userService)
	cartHandler := NewCartHandler(cartService)
	orderHandler := NewOrderHandler(orderService)

	categoryHandler.RegisterRoutes(router)
	subcategoryHandler.RegisterRoutes(router)
	userHandler.RegisterRoutes(router)
	cartHandler.RegisterRoutes(router)
	orderHandler.RegisterRoutes(router)

}
