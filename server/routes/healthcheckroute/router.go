package healthcheckroute

import "github.com/labstack/echo"

// HealthCheckRouter holds the user handlers
type HealthCheckRouter struct {
	ctrl   *Controller
	router *echo.Echo
}

// NewRouter returns a new HealthCheckRouter instance
func NewRouter(ctrl *Controller, router *echo.Echo) *HealthCheckRouter {
	return &HealthCheckRouter{
		ctrl:   ctrl,
		router: router,
	}
}

//RegisterRoutes is a routers map of health check requests
func (r *HealthCheckRouter) RegisterRoutes() {
	r.router.GET("/user/:user-uuid/health-checks", r.ctrl.HandleGetAllHealthChecks)
	r.router.GET("/user/:user-uuid/health-checks/:uuid", r.ctrl.HandleGetHealthCheckByUUID)
}
