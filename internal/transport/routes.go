package transport

import (
	"github.com/gin-gonic/gin"
	"github.com/itsvagapov/team-pharmacy/internal/service"
)

func RegisterRoutes(router *gin.Engine, categoryService service.CategoryService, subcategoryService service.SubcategoryService) { // userService service.User....
	categoryHandler := NewCategoryHandler(categoryService)
	subcategoryHandler := NewSubcategoryHandler(subcategoryService)
	// userHandler

	categoryHandler.RegisterRoutes(router)
	subcategoryHandler.RegisterRoutes(router)
	// user.regist
}
