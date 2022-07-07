package usecase

import (
	"net/http"
	"user-service/constants"
	"user-service/datatransfers"
	"user-service/helpers"
	"user-service/models"
	repository "user-service/repositories"
	"user-service/utils"

	"gorm.io/gorm"
)

type UserUsecase interface {
	Create(user *models.User) (err error)
	Delete(user *models.User) (err error)
	GetAll(param *datatransfers.ListQueryParams) (users []*models.User, cnt int64, err error)
	GetByID(userID int) (user *models.User, err error)
	Login(params *datatransfers.LoginRequest) (user *models.User, err error)
	Register(params *datatransfers.RegisterRequest) (err error)
	Update(user *models.User) (err error)
}

type userUsecase struct {
	db       *gorm.DB
	userRepo repository.UserRepository
}

func NewUserUsecase(db *gorm.DB) UserUsecase {
	userRepo := repository.NewUserRepository(db)
	return &userUsecase{
		userRepo: userRepo,
		db:       db,
	}
}

func (u *userUsecase) GetAll(param *datatransfers.ListQueryParams) (users []*models.User, cnt int64, err error) {
	users, cnt, err = u.userRepo.GetAll(param)
	return
}

func (u *userUsecase) GetByID(userID int) (user *models.User, err error) {
	user, err = u.userRepo.GetByID(userID)
	return
}

func (u *userUsecase) Create(user *models.User) (err error) {
	err = u.userRepo.Create(user, u.db)
	return
}

func (u *userUsecase) Update(user *models.User) (err error) {
	err = u.userRepo.Update(user, u.db)
	return
}

func (u *userUsecase) Delete(user *models.User) (err error) {
	err = u.userRepo.Delete(user, u.db)
	return
}

func (u *userUsecase) Register(params *datatransfers.RegisterRequest) (err error) {
	userByEmail, err := u.userRepo.GetByEmail(params.Email)
	if err != nil && !utils.IsErrRecordNotFound(err) {
		return
	}

	if userByEmail != nil {
		err = &datatransfers.CustomError{
			Code:    constants.RegisterEmailNotAvailableErrCode,
			Status:  http.StatusBadRequest,
			Message: "EMAIL_IS_USED",
		}
		return
	}

	userByUsername, err := u.userRepo.GetByUsername(params.Username)
	if err != nil && !utils.IsErrRecordNotFound(err) {
		return
	}

	if userByUsername != nil {
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

	err = u.userRepo.Create(&models.User{
		Username: params.Username,
		Email:    params.Email,
		Password: hashedPassword,
	}, u.db)

	// TODO: generate token
	return
}

func (u *userUsecase) Login(params *datatransfers.LoginRequest) (user *models.User, err error) {
	user, err = u.userRepo.GetByUsername(params.Username)
	if err != nil {
		return
	}

	isPasswordMatched := helpers.CheckPasswordHash(params.Password, user.Password)
	if !isPasswordMatched {
		err = &datatransfers.CustomError{
			Code:    constants.LoginInvalidPasswordErrCode,
			Status:  http.StatusBadRequest,
			Message: "INVALID_PASSWORD",
		}
		return nil, err
	}

	// TODO: generate token
	return user, nil
}
