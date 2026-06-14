package transport

import (
	"github.com/gin-gonic/gin"
	"github.com/itsvagapov/team-pharmacy/internal/service"
)

func RegisterRoutes(router *gin.Engine, categoryService service.CategoryService, subcategoryService service.SubcategoryService, reviewService service.ReviewService) {
	categoryHandler := NewCategoryHandler(categoryService)
	subcategoryHandler := NewSubcategoryHandler(subcategoryService)
	reviewHandler := NewReviewHandler(reviewService)

	categoryHandler.RegisterRoutes(router)
	subcategoryHandler.RegisterRoutes(router)
	reviewHandler.RegisterRoutes(router)
}
