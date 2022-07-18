package routes
//if the file's name is 'route', the package name is 'route' too

import(
	"e-commerce/controllers"
	"github.com/gin-gonic/gin"

)

func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.POST("/users/signup", controllers.SignUp())
	incomingRoutes.POST("/users/login", controllers.Login())
	incomingRoutes.POST("/admin/addproduct", controllers.ProductViewrAdmin())
	incomingRoutes.GET("/users/productview", controllers.SearchProduct())
	incomingRoutes.GET("/users/search", controllers.SearchProductByQuery())


}