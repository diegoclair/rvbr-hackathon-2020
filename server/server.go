package server

import (
	"github.com/IQ-tech/go-mapper"
	"github.com/RedVentures/rvbr-2020-hackathon-time-3-backend/server/routes/healthcheckroute"
	"github.com/RedVentures/rvbr-2020-hackathon-time-3-backend/server/routes/pingroute"
	"github.com/RedVentures/rvbr-2020-hackathon-time-3-backend/server/routes/uploadroute"
	"github.com/RedVentures/rvbr-2020-hackathon-time-3-backend/server/routes/userroute"
	"github.com/RedVentures/rvbr-2020-hackathon-time-3-backend/service"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type controller struct {
	pingController        *pingroute.Controller
	userController        *userroute.Controller
	uploadController      *uploadroute.Controller
	healthcheckController *healthcheckroute.Controller
}

//InitServer to initialize the server
func InitServer(svc *service.Service) *echo.Echo {
	mapper := mapper.New()
	svm := service.NewServiceManager()
	srv := echo.New()

	userService := svm.UserService(svc)
	healthCheckService := svm.HealthCheckService(svc)
	cloudVisionService := svm.CloudVisionService(svc)
	uploadService := svm.UploadService(svc, cloudVisionService)

	//CORS
	srv.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	return setupRoutes(srv, &controller{
		pingController:        pingroute.NewController(),
		userController:        userroute.NewController(userService, mapper),
		uploadController:      uploadroute.NewController(uploadService),
		healthcheckController: healthcheckroute.NewController(healthCheckService, mapper),
	})
}

//setupRoutes - Register and instantiate the routes
func setupRoutes(srv *echo.Echo, s *controller) *echo.Echo {

	pingroute.NewRouter(s.pingController, srv).RegisterRoutes()
	userroute.NewRouter(s.userController, srv).RegisterRoutes()
	uploadroute.NewRouter(s.uploadController, srv).RegisterRoutes()
	healthcheckroute.NewRouter(s.healthcheckController, srv).RegisterRoutes()

	return srv
}
