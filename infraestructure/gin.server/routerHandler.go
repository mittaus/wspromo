package server

import (
	v1 "dev.azure.com/spsa/wspromo/infraestructure/gin.server/v1"
	"github.com/gin-gonic/gin"
)

type RouterHandler struct {
	datacenterHandler v1.DatacenterHandler
	// articleHandler articles.IArticleManager
}

func NewRouterHandler(datacenterHandler v1.DatacenterHandler) RouterHandler {
	return RouterHandler{
		datacenterHandler: datacenterHandler,
	}
}

func (rH RouterHandler) SetRoutes(r *gin.Engine) {

	// Group : v1
	apiV1 := r.Group("ws-rest-cem-dp-createpromovar/api/v1")

	datacenterRoutes := apiV1.Group("/datacenters")
	{

		datacenterRoutes.GET("/", rH.datacenterHandler.GetAll)
		datacenterRoutes.GET("/:id", rH.datacenterHandler.GetByID)
		datacenterRoutes.POST("/", rH.datacenterHandler.Store)
		datacenterRoutes.DELETE("/:id", rH.datacenterHandler.DeleteByID)
		datacenterRoutes.PATCH("/:id", rH.datacenterHandler.UpdateByID)

	}

}
