package repository

import (
	"net/http"
	"strings"
	"user-service/constants"
	"user-service/datatransfers"
	"user-service/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository interface {
	Create(user *models.User, db *gorm.DB) (err error)
	Delete(user *models.User, db *gorm.DB) (err error)
	GetAll(params *datatransfers.ListQueryParams) (users []*models.User, cnt int64, err error)
	GetByEmail(email string) (user *models.User, err error)
	GetByID(userID string) (user *models.User, err error)
	Update(user *models.User, db *gorm.DB) (err error)
}
type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetAll(params *datatransfers.ListQueryParams) (users []*models.User, cnt int64, err error) {
	qs := r.db
	// TODO: filter select
	err = qs.Model(&models.User{}).Count(&cnt).Error
	if err != nil {
		err = &datatransfers.CustomError{
			Code:    constants.QueryInternalServerErrCode,
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}

	if params.Limit > 0 {
		qs = qs.Limit(params.Limit)
	}

	if params.Offset > 0 {
		qs = qs.Offset(params.Offset)
	}

	err = qs.Find(&users).Error
	if err != nil {
		err = &datatransfers.CustomError{
			Code:    constants.QueryInternalServerErrCode,
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return
}

func (r *userRepository) GetByID(userID string) (user *models.User, err error) {
	qs := r.db.Where("id = ?", userID)
	err = qs.First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = &datatransfers.CustomError{
				Code:    constants.QueryNotFoundErrCode,
				Status:  http.StatusNotFound,
				Message: err.Error(),
			}
		}

		err = &datatransfers.CustomError{
			Code:    constants.QueryInternalServerErrCode,
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return
}

func (r *userRepository) GetByEmail(email string) (user *models.User, err error) {
	qs := r.db.Where("email = ?", strings.ToLower(strings.TrimSpace(email)))
	err = qs.First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = &datatransfers.CustomError{
				Code:    constants.QueryNotFoundErrCode,
				Status:  http.StatusNotFound,
				Message: err.Error(),
			}
		}
		err = &datatransfers.CustomError{
			Code:    constants.InternalServerErrCode,
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return
}

func (r *userRepository) Create(user *models.User, db *gorm.DB) (err error) {
	err = db.Omit(clause.Associations).Model(user).Create(user).Error
	if err != nil {
		err = &datatransfers.CustomError{
			Code:    constants.InternalServerErrCode,
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return
}

func (r *userRepository) Update(user *models.User, db *gorm.DB) (err error) {
	row := db.Omit(clause.Associations).Model(user).Updates(user)
	err = row.Error
	if err != nil {
		err = &datatransfers.CustomError{
			Code:    constants.InternalServerErrCode,
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}

	if row.RowsAffected == 0 {
		err = &datatransfers.CustomError{
			Code:    constants.QueryNotFoundErrCode,
			Status:  http.StatusNotFound,
			Message: err.Error(),
		}
		return
	}
	return
}

func (r *userRepository) Delete(user *models.User, db *gorm.DB) (err error) {
	row := db.Omit(clause.Associations).Model(user).Delete(user)
	err = row.Error
	if err != nil {
		err = &datatransfers.CustomError{
			Code:    constants.InternalServerErrCode,
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}

	if row.RowsAffected == 0 {
		err = &datatransfers.CustomError{
			Code:    constants.QueryNotFoundErrCode,
			Status:  http.StatusNotFound,
			Message: err.Error(),
		}
		return
	}
	return
}
