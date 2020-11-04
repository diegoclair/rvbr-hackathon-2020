package healthcheckroute

import (
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/IQ-tech/go-mapper"
	"github.com/RedVentures/rvbr-2020-hackathon-time-3-backend/domain/contract"
	"github.com/RedVentures/rvbr-2020-hackathon-time-3-backend/server/viewmodel"
	"github.com/diegoclair/go_utils-lib/resterrors"
	"github.com/labstack/echo"
)

var (
	instance *Controller
	once     sync.Once
)

//Controller is a interface to interact with services
type Controller struct {
	healthCheckService contract.HealthCheckService
	mapper             mapper.Mapper
}

//NewController to handle requests
func NewController(healthCheckService contract.HealthCheckService, mapper mapper.Mapper) *Controller {
	once.Do(func() {
		instance = &Controller{
			healthCheckService: healthCheckService,
			mapper:             mapper,
		}
	})
	return instance
}

func (c Controller) HandleGetAllHealthChecks(context echo.Context) error {

	userUUID := context.Param("user-uuid")

	healthChecks, err := c.healthCheckService.FindByUser(userUUID)
	if err != nil {
		return context.JSON(err.StatusCode(), err)
	}

	response := viewmodel.ParseHealthChecksResponse(healthChecks)

	return context.JSON(http.StatusOK, response)
}

func (c Controller) HandleGetHealthCheckByUUID(context echo.Context) error {

	userUUID := context.Param("user-uuid")

	healthCheckUUID := context.Param("uuid")

	err := c.healthCheckService.FindByUUID(userUUID, healthCheckUUID)
	if err != nil {
		return context.JSON(err.StatusCode(), err)
	}

	docFileBytes, errRead := ioutil.ReadFile("./download/5ad2bf31-157b-11eb-952c-0242ac120002.pdf")
	if errRead != nil {
		return context.JSON(404, resterrors.NewNotFoundError("Not Found"))
	}

	return context.JSON(http.StatusOK, docFileBytes)
}
