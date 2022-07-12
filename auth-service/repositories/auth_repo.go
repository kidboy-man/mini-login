package repository

import (
	"auth-service/constants"
	"auth-service/datatransfers"
	"auth-service/models"
	"net/http"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AuthRepository interface {
	Create(auth *models.Auth, db *gorm.DB) (err error)
	Delete(auth *models.Auth, db *gorm.DB) (err error)
	GetAll(params *datatransfers.ListQueryParams) (auths []*models.Auth, cnt int64, err error)
	GetByID(authID int) (auth *models.Auth, err error)
	GetByUsername(username string) (auth *models.Auth, err error)
	Update(auth *models.Auth, db *gorm.DB) (err error)
}
type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) GetAll(params *datatransfers.ListQueryParams) (auths []*models.Auth, cnt int64, err error) {
	qs := r.db
	if params.IsPublic {
		qs = qs.Where("is_active = ?", true)
	}

	err = qs.Model(&models.Auth{}).Count(&cnt).Error
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

	err = qs.Find(&auths).Error
	if err != nil {
		err = &datatransfers.CustomError{
			Code:    constants.QueryInternalServerErrCode,
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return
}

func (r *authRepository) GetByID(authID int) (auth *models.Auth, err error) {
	qs := r.db.Where("id = ?", authID)
	err = qs.First(&auth).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = &datatransfers.CustomError{
				Code:    constants.QueryNotFoundErrCode,
				Status:  http.StatusNotFound,
				Message: err.Error(),
			}
			return nil, err
		}

		err = &datatransfers.CustomError{
			Code:    constants.QueryInternalServerErrCode,
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		return nil, err
	}
	return
}

func (r *authRepository) GetByUsername(username string) (auth *models.Auth, err error) {
	qs := r.db.Where("username = ?", username)
	err = qs.First(&auth).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = &datatransfers.CustomError{
				Code:    constants.QueryNotFoundErrCode,
				Status:  http.StatusNotFound,
				Message: err.Error(),
			}
			return nil, err
		}
		err = &datatransfers.CustomError{
			Code:    constants.InternalServerErrCode,
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		return nil, err
	}
	return
}

func (r *authRepository) Create(auth *models.Auth, db *gorm.DB) (err error) {
	err = db.Omit(clause.Associations).Model(auth).Create(auth).Error
	if err != nil {
		err = &datatransfers.CustomError{
			Code:    constants.InternalServerErrCode,
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}

	}
	return
}

func (r *authRepository) Update(auth *models.Auth, db *gorm.DB) (err error) {
	row := db.Omit(clause.Associations).Model(auth).Updates(auth)
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

func (r *authRepository) Delete(auth *models.Auth, db *gorm.DB) (err error) {
	row := db.Omit(clause.Associations).Model(auth).Delete(auth)
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
