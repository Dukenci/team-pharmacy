package transport

import (
	"github.com/gin-gonic/gin"
	"github.com/itsvagapov/team-pharmacy/internal/service"
)


func RegisterRoutes(router *gin.Engine, categoryService service.CategoryService, subcategoryService service.SubcategoryService, reviewService service.ReviewService, medicineService service.MedicineService, userService service.UserService, 	cartService service.CartService, orderService service.OrderService) {
	categoryHandler := NewCategoryHandler(categoryService)
	subcategoryHandler := NewSubcategoryHandler(subcategoryService)
	reviewHandler := NewReviewHandler(reviewService)
	medicineHandler := NewMedicineHandler(medicineService)
	userHandler := NewUserHandler(userService)
	cartHandler := NewCartHandler(cartService)
	orderHandler := NewOrderHandler(orderService)

	categoryHandler.RegisterRoutes(router)
	subcategoryHandler.RegisterRoutes(router)
	reviewHandler.RegisterRoutes(router)
	medicineHandler.RegisterRoutes(router)
	userHandler.RegisterRoutes(router)
	cartHandler.RegisterRoutes(router)
	orderHandler.RegisterRoutes(router)

}
