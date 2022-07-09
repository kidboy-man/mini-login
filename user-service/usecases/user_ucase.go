package usecase

import (
	"user-service/datatransfers"
	"user-service/models"
	repository "user-service/repositories"

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
	err = u.userRepo.Update(user, u.db)
	return
}

func (u *userUsecase) Delete(user *models.User) (err error) {
	err = u.userRepo.Delete(user, u.db)
	return
}
