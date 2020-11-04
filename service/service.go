package service

import (
	"net"
	"net/http"

	"github.com/RedVentures/rvbr-2020-hackathon-time-3-backend/domain"
	"github.com/RedVentures/rvbr-2020-hackathon-time-3-backend/domain/contract"
)

// Service holds the domain service repositories
type Service struct {
	db         contract.RepoManager
	httpClient *http.Client
}

// New returns a new domain Service instance
func New(db contract.RepoManager) *Service {
	svc := new(Service)
	svc.db = db

	httpDialer := new(net.Dialer)
	httpDialer.Timeout = domain.DefaultConnectionTimeout

	httpTransport := new(http.Transport)
	httpTransport.TLSHandshakeTimeout = domain.DefaultConnectionTimeout
	httpTransport.Dial = httpDialer.Dial

	svc.httpClient = new(http.Client)
	svc.httpClient.Transport = httpTransport
	svc.httpClient.Timeout = domain.DefaultConnectionTimeout

	return svc
}

//Manager defines the services aggregator interface
type Manager interface {
	UserService(svc *Service) contract.UserService
	HealthCheckService(svc *Service) contract.HealthCheckService
	UploadService(svc *Service, cloudVisionService contract.CloudVisionService) contract.UploadService
	CloudVisionService(svc *Service) contract.CloudVisionService
}

type serviceManager struct {
	svc *Service
}

// NewServiceManager return a service manager instance
func NewServiceManager() Manager {
	return &serviceManager{}
}

func (s *serviceManager) UserService(svc *Service) contract.UserService {
	return newUserService(svc)
}

func (s *serviceManager) UploadService(svc *Service, cloudVisionService contract.CloudVisionService) contract.UploadService {
	return newUploadService(svc, cloudVisionService)
}

func (s *serviceManager) CloudVisionService(svc *Service) contract.CloudVisionService {
	return newCloudVisionService(svc)
}

func (s *serviceManager) HealthCheckService(svc *Service) contract.HealthCheckService {
	return newHealthCheckService(svc)
}
