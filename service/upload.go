package service

import (
	"io"
	"mime/multipart"
	"os"

	"github.com/RedVentures/rvbr-2020-hackathon-time-3-backend/domain"
	"github.com/RedVentures/rvbr-2020-hackathon-time-3-backend/domain/contract"
	"github.com/diegoclair/go_utils-lib/logger"
	"github.com/diegoclair/go_utils-lib/resterrors"
	"github.com/labstack/echo"
)

type uploadService struct {
	svc                *Service
	cloudVisionService contract.CloudVisionService
}

//newUploadService return a new instance of the service
func newUploadService(svc *Service, cloudVisionService contract.CloudVisionService) contract.UploadService {
	return &uploadService{
		svc:                svc,
		cloudVisionService: cloudVisionService,
	}
}

func (s *uploadService) SaveFileForUser(c echo.Context, file *multipart.FileHeader, userUUID, fileType string) (text string, restErr resterrors.RestErr) {

	_, restErr = s.svc.db.User().GetUserByUUID(userUUID)
	if restErr != nil {
		return text, restErr
	}

	var path string

	// Upload the file to specific destination.
	if fileType == domain.TypePrescriptionFile {
		path = "upload/" + userUUID + "/prescriptions"
	}

	//create folder
	_, err := os.Stat("/" + path)
	if os.IsNotExist(err) {

		errDir := os.MkdirAll(path, os.FileMode(0770))
		if errDir != nil {
			logger.Error("SaveFileForUser: Error to create folder", errDir)
			return text, resterrors.NewInternalServerError("Error to save file")
		}
	}

	src, err := file.Open()
	if err != nil {
		logger.Error("SaveFileForUser: Error to open the file", err)
		return text, resterrors.NewInternalServerError("Error to save file")
	}
	defer src.Close()

	// Destination
	pathWithName := path + "/" + file.Filename

	dst, err := os.Create(pathWithName)
	if err != nil {
		logger.Error("SaveFileForUser: Error to create folder", err)
		return text, resterrors.NewInternalServerError("Error to save file")
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		logger.Error("SaveFileForUser: Error to copy folder", err)
		return text, resterrors.NewInternalServerError("Error to save file")
	}

	// pathWithName := path + "/" + file.Filename

	// saveErr := c.SaveUploadedFile(file, pathWithName)
	// if saveErr != nil {
	// 	logger.Error("SaveFileForUser: Error to save file", saveErr)
	// 	err := resterrors.NewInternalServerError("Error to save file")
	// 	return text, err
	// }

	text, err = s.cloudVisionService.SendImageToCloudVision(path)
	if err != nil {
		logger.Error("SaveFileForUser: Error to read file", err)
		return text, resterrors.NewInternalServerError("Error to read file")
	}

	return text, nil
}
