package contract

import (
	"mime/multipart"

	"github.com/RedVentures/rvbr-2020-hackathon-time-3-backend/domain/entity"
	"github.com/diegoclair/go_utils-lib/resterrors"
	"github.com/labstack/echo"
)

// PingService holds a ping service operations
type PingService interface {
}

// UserService holds a user service operations
type UserService interface {
	SignIn(user entity.User) (entity.User, resterrors.RestErr)
}

// UploadService holds a upload service operations
type UploadService interface {
	SaveFileForUser(c echo.Context, file *multipart.FileHeader, userUUID, fileType string) (text string, restErr resterrors.RestErr)
}

// CloudVisionService holds a cloud vision service operations
type CloudVisionService interface {
	SendImageToCloudVision(filePath string) (text string, err resterrors.RestErr)
}

type HealthCheckService interface {
	FindByUser(userUUID string) ([]entity.HealthCheck, resterrors.RestErr)
	FindByUUID(userUUID string, uuid string) resterrors.RestErr
}
