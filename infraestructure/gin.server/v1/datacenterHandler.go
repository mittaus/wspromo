package v1

import (
	"log"
	"net/http"

	application "dev.azure.com/spsa/wspromo/application"
	"dev.azure.com/spsa/wspromo/domain"
	"dev.azure.com/spsa/wspromo/shared"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type DatacenterHandler struct {
	app application.Datacenter
}

func NewDatacenterHandler(myapp application.Datacenter) *DatacenterHandler {
	return &DatacenterHandler{
		app: myapp,
	}
}

func (s *DatacenterHandler) GetAll(ctx *gin.Context) {
	var p shared.Pagination
	p.Limit = ctx.DefaultQuery("limit", "-1")
	p.OffSet = ctx.DefaultQuery("offset", "-1")
	datacenters, err := s.app.GetAll(&p)
	if len(datacenters) == 0 || err != nil {
		ctx.Status(http.StatusNoContent)
		return
	}
	stringResponse := shared.ResponseService("success", "", "List Of Datacenters", datacenters)
	ctx.JSON(http.StatusOK, shared.StringResponseToResponseObj(stringResponse))
	return
}

func (s *DatacenterHandler) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	if _, err := uuid.FromString(id); err != nil {
		var auxInterface interface{}
		stringResponse := shared.ResponseService("Failure", shared.FAIL_CONVERTED_UUID_TO_STRING, "Error trying convert uuid to string", auxInterface)
		ctx.JSON(http.StatusBadRequest, shared.StringResponseToResponseObj(stringResponse))
		return
	}

	datacenter, err := s.app.GetByID(id)
	if datacenter == nil || err != nil {
		var auxInterface interface{}
		stringResponse := shared.ResponseService("Failure", shared.ID_NOT_FOUND, "Id not found", auxInterface)
		ctx.JSON(http.StatusNotFound, shared.StringResponseToResponseObj(stringResponse))
		return
	}

	stringResponse := shared.ResponseService("success", "", "Detail Of Datacenter", datacenter)
	ctx.JSON(http.StatusOK, shared.StringResponseToResponseObj(stringResponse))
	return
}

func (s *DatacenterHandler) Store(ctx *gin.Context) {
	var datacenter domain.Datacenter
	if err := ctx.ShouldBindJSON(&datacenter); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println("Parametros ingresados json: ", datacenter)
	response := s.app.Store(&datacenter)
	log.Println(response)
	ctx.String(http.StatusOK, "Mensaje ok")
}

func (s *DatacenterHandler) DeleteByID(ctx *gin.Context) {
	id := ctx.Param("id")
	if _, err := uuid.FromString(id); err != nil {
		var auxInterface interface{}
		stringResponse := shared.ResponseService("Failure", shared.FAIL_CONVERTED_UUID_TO_STRING, "Error trying convert uuid to string", auxInterface)
		ctx.JSON(http.StatusBadRequest, shared.StringResponseToResponseObj(stringResponse))
		return
	}

	response := s.app.DeleteByID(id)

	if response != nil {
		ctx.JSON(http.StatusBadRequest, shared.StringResponseToResponseObj(response.Error()))
		return
	} else {
		var auxInterface interface{}
		stringResponse := shared.ResponseService("success", "", "Datacenter deleted", auxInterface)
		ctx.JSON(http.StatusOK, shared.StringResponseToResponseObj(stringResponse))
		return
	}
}

func (s *DatacenterHandler) UpdateByID(ctx *gin.Context) {
	id := ctx.Param("id")
	if _, err := uuid.FromString(id); err != nil {
		var auxInterface interface{}
		stringResponse := shared.ResponseService("Failure", shared.FAIL_CONVERTED_UUID_TO_STRING, "Error trying convert uuid to string", auxInterface)
		ctx.JSON(http.StatusBadRequest, shared.StringResponseToResponseObj(stringResponse))
		return
	}

	var datacenter domain.Datacenter
	response := s.app.UpdateByID(id, &datacenter)

	if response != nil {
		ctx.JSON(http.StatusBadRequest, shared.StringResponseToResponseObj(response.Error()))
		return
	} else {
		var auxInterface interface{}
		stringResponse := shared.ResponseService("success", "", "Datacenter Updated", auxInterface)
		ctx.JSON(http.StatusOK, shared.StringResponseToResponseObj(stringResponse))
		return
	}
}
