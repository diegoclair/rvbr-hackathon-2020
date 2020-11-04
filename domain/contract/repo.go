package contract

import (
	"github.com/RedVentures/rvbr-2020-hackathon-time-3-backend/domain/entity"
	"github.com/diegoclair/go_utils-lib/resterrors"
)

//RepoManager defines the repository aggregator interface
type RepoManager interface {
	Ping() PingRepo
	User() UserRepo
	HealthCheck() HealthCheckRepo
}

// PingRepo defines the data set for ping
type PingRepo interface{}

// UserRepo defines the data set for user
type UserRepo interface {
	FindByEmailAndPassword(email string, password string) (entity.User, resterrors.RestErr)
	GetUserByUUID(userUUID string) (user entity.User, err resterrors.RestErr)
}

type HealthCheckRepo interface {
	FindAllByUserUUID(userUUID string) ([]entity.HealthCheck, resterrors.RestErr)
}
