package service

import (
	"github.com/RedVentures/rvbr-2020-hackathon-time-3-backend/domain/contract"
	"github.com/RedVentures/rvbr-2020-hackathon-time-3-backend/domain/entity"
	"github.com/diegoclair/go_utils-lib/resterrors"
)

type healthCheckService struct {
	svc *Service
}

func newHealthCheckService(svc *Service) contract.HealthCheckService {
	return &healthCheckService{
		svc: svc,
	}
}

func (s healthCheckService) FindByUser(userUUID string) ([]entity.HealthCheck, resterrors.RestErr) {

	healths, err := s.svc.db.HealthCheck().FindAllByUserUUID(userUUID)
	if err != nil {
		return nil, err
	}

	return healths, nil

}

func (s healthCheckService) FindByUUID(userUUID string, healthCheckUUID string) resterrors.RestErr {

	_, restErr := s.svc.db.User().GetUserByUUID(userUUID)
	if restErr != nil {
		return restErr
	}

	return nil
}
