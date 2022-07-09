package usecase

import (
	"auth-service/constants"
	"auth-service/datatransfers"
	"auth-service/helpers"
	"auth-service/models"
	repository "auth-service/repositories"
	"auth-service/utils"
	"net/http"

	"gorm.io/gorm"
)

type AuthUsecase interface {
	Create(auth *models.Auth) (err error)
	Delete(auth *models.Auth) (err error)
	GetAll(param *datatransfers.ListQueryParams) (auths []*models.Auth, cnt int64, err error)
	GetByID(authID int) (auth *models.Auth, err error)
	Login(params *datatransfers.AuthRequest) (auth *models.Auth, err error)
	Register(params *datatransfers.AuthRequest) (err error)
	Update(auth *models.Auth) (err error)
}

type authUsecase struct {
	db       *gorm.DB
	authRepo repository.AuthRepository
}

func NewAuthUsecase(db *gorm.DB) AuthUsecase {
	authRepo := repository.NewAuthRepository(db)
	return &authUsecase{
		authRepo: authRepo,
		db:       db,
	}
}

func (u *authUsecase) GetAll(param *datatransfers.ListQueryParams) (auths []*models.Auth, cnt int64, err error) {
	auths, cnt, err = u.authRepo.GetAll(param)
	return
}

func (u *authUsecase) GetByID(authID int) (auth *models.Auth, err error) {
	auth, err = u.authRepo.GetByID(authID)
	return
}

func (u *authUsecase) Create(auth *models.Auth) (err error) {
	err = u.authRepo.Create(auth, u.db)
	return
}

func (u *authUsecase) Update(auth *models.Auth) (err error) {
	err = u.authRepo.Update(auth, u.db)
	return
}

func (u *authUsecase) Delete(auth *models.Auth) (err error) {
	err = u.authRepo.Delete(auth, u.db)
	return
}

func (u *authUsecase) Register(params *datatransfers.AuthRequest) (err error) {
	existingAuth, err := u.authRepo.GetByUsername(params.Username)
	if err != nil && !utils.IsErrRecordNotFound(err) {
		return
	}

	if existingAuth != nil {
		err = &datatransfers.CustomError{
			Code:    constants.RegisterUsernameNotAvailableErrCode,
			Status:  http.StatusBadRequest,
			Message: "USERNAME_IS_TAKEN",
		}
		return
	}

	hashedPassword, err := helpers.HashPassword(params.Password)
	if err != nil {
		return
	}

	err = u.authRepo.Create(&models.Auth{
		Username: params.Username,
		Password: hashedPassword,
	}, u.db)

	// TODO: generate token
	return
}

func (u *authUsecase) Login(params *datatransfers.AuthRequest) (auth *models.Auth, err error) {
	auth, err = u.authRepo.GetByUsername(params.Username)
	if err != nil {
		return
	}

	isPasswordMatched := helpers.CheckPasswordHash(params.Password, auth.Password)
	if !isPasswordMatched {
		err = &datatransfers.CustomError{
			Code:    constants.LoginInvalidPasswordErrCode,
			Status:  http.StatusBadRequest,
			Message: "INVALID_PASSWORD",
		}
		return nil, err
	}

	// TODO: generate token
	return auth, nil
}
