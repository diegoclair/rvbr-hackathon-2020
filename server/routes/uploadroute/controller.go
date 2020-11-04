package uploadroute

import (
	"log"
	"net/http"
	"sync"

	"github.com/RedVentures/rvbr-2020-hackathon-time-3-backend/domain"
	"github.com/RedVentures/rvbr-2020-hackathon-time-3-backend/domain/contract"
	"github.com/diegoclair/go_utils-lib/logger"
	"github.com/diegoclair/go_utils-lib/resterrors"
	"github.com/gin-gonic/gin"
	"github.com/labstack/echo"
)

var (
	instance *Controller
	once     sync.Once
)

//Controller holds user handlers
type Controller struct {
	uploadService contract.UploadService
}

//NewController to handle requests
func NewController(uploadService contract.UploadService) *Controller {
	once.Do(func() {
		instance = &Controller{
			uploadService: uploadService,
		}
	})
	return instance
}

// handlePrescriptionUpload - handle a Prescription Upload request
func (s *Controller) handlePrescriptionUpload(c echo.Context) error {

	// single file
	file, fileErr := c.FormFile("file")
	if fileErr != nil {
		logger.Error("handlePrescriptionUpload: Error to get file from request", fileErr)
		err := resterrors.NewBadRequestError(fileErr.Error())
		return c.JSON(err.StatusCode(), err)
	}
	log.Println("Filename: ", file.Filename)

	userUUID := c.Param("user_uuid")

	text, saveErr := s.uploadService.SaveFileForUser(c, file, userUUID, domain.TypePrescriptionFile)
	if saveErr != nil {
		return c.JSON(saveErr.StatusCode(), saveErr)
	}

	return c.JSON(http.StatusOK, gin.H{"fileText": text})
}
