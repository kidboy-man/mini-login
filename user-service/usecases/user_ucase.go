package usecase

import (
	"net/http"
	"net/mail"
	"user-service/constants"
	"user-service/datatransfers"
	"user-service/models"
	repository "user-service/repositories"
	"user-service/utils"

	"gorm.io/gorm"
)

type UserUsecase interface {
	Create(user *models.User) (err error)
	Delete(user *models.User) (err error)
	GetAll(param *datatransfers.ListQueryParams) (users []*models.User, cnt int64, err error)
	GetByID(userID string) (user *models.User, err error)
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

func (u *userUsecase) GetByID(userID string) (user *models.User, err error) {
	user, err = u.userRepo.GetByID(userID)
	return
}

func (u *userUsecase) Create(user *models.User) (err error) {
	err = u.userRepo.Create(user, u.db)
	return
}

func (u *userUsecase) Update(user *models.User) (err error) {
	if user.Email != "" {
		// check email format
		_, err = mail.ParseAddress(user.Email)
		if err != nil {
			err = &datatransfers.CustomError{
				Code:    constants.InvalidEmailFormatErrCode,
				Status:  http.StatusBadRequest,
				Message: "INVALID_EMAIL_ADDRESS",
			}
			return
		}

		existingUser, errGetUser := u.userRepo.GetByEmail(user.Email)
		if errGetUser != nil && !utils.IsErrRecordNotFound(errGetUser) {
			return errGetUser
		}

		if existingUser != nil {
			if existingUser.ID == user.ID {
				return
			}

			err = &datatransfers.CustomError{
				Code:    constants.UpdateUserEmailExistErrCode,
				Status:  http.StatusBadRequest,
				Message: "EMAIL_IS_USED",
			}
			return
		}
	}

	err = u.userRepo.Update(user, u.db)
	return
}

func (u *userUsecase) Delete(user *models.User) (err error) {
	err = u.userRepo.Delete(user, u.db)
	return
}
