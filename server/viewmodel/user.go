package viewmodel

import (
	"github.com/RedVentures/rvbr-2020-hackathon-time-3-backend/domain/entity"
	"github.com/diegoclair/go_utils-lib/resterrors"
)

type SignInUserRequest struct {
	Email    string `json:"login"`
	Password string `json:"password"`
}

func (vm SignInUserRequest) Validate() resterrors.RestErr {
	var err resterrors.RestErr
	if emptyString(vm.Email) {
		err = resterrors.NewBadRequestError("Error empty field email in User Login Request")
		return err
	}

	if emptyString(vm.Password) {
		err = resterrors.NewBadRequestError("Error empty field password in User Login Request")
		return err
	}

	return nil
}

func (vm SignInUserRequest) Parse() entity.User {
	return entity.User{
		Email:    vm.Email,
		Password: vm.Password,
	}
}

type SignInUserResponse struct {
	UUID string `json:"uuid"`
}

func ParseSignInUserResponse(user entity.User) SignInUserResponse {
	return SignInUserResponse{
		UUID: user.UUID,
	}
}
