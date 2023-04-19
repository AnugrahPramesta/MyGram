package service

import (
	"errors"
	"log"
	"math"
	"net/http"
	"project-mygram/dto"
	"project-mygram/entity"
	"project-mygram/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PhotoService interface {
	Create(input entity.Photo) (res dto.PhotoCreateUpdateResponse, statusCode int, err error)
	GetAll(ctx *gin.Context, queryparam dto.ListParam) (result dto.PhotoListRes, statusCode int, err error)
	GetByID(ID uint64) (result entity.Photo, statusCode int, err error)
	UpdateByID(ID uint64, input dto.PhotoCreateUpdateReq) (result dto.PhotoCreateUpdateResponse, statusCode int, err error)
	DeleteByID(ID uint64) (statusCode int, err error)
}

type photoSrv struct {
	PhotoRepository repository.PhotoRepository
}

func NewPhotoService(PhotoRepo repository.PhotoRepository) PhotoService {
	return &photoSrv{
		PhotoRepository: PhotoRepo,
	}
}

func (srv *photoSrv) Create(input entity.Photo) (res dto.PhotoCreateUpdateResponse, statusCode int, err error) {
	temp, err := srv.PhotoRepository.Create(input)
	if err != nil {
		log.Printf("[PhotoService-Create] error create data: %+v \n", err)
		return res, http.StatusInternalServerError, err
	}

	res = dto.PhotoCreateUpdateResponse{
		ID:        temp.ID,
		UserID:    temp.UserID,
		Title:     temp.Title,
		PhotoUrl:  temp.PhotoUrl,
		Caption:   temp.Caption,
		CreatedAt: temp.CreatedAt,
		UpdatedAt: temp.UpdatedAt,
	}

	return res, http.StatusCreated, nil
}

func (srv *photoSrv) GetAll(ctx *gin.Context, queryparam dto.ListParam) (result dto.PhotoListRes, statusCode int, err error) {
	result.Limit = queryparam.Limit
	result.Page = queryparam.Page

	result.Data, result.Total, err = srv.PhotoRepository.GetAll(ctx, queryparam)
	if err != nil {
		log.Printf("[PhotoService-GetAll] error get data repo: %+v \n", err)
		return result, http.StatusInternalServerError, err
	}

	result.LastPage = uint64(math.Ceil(float64(result.Total) / float64(queryparam.Limit)))

	return result, http.StatusOK, nil
}

func (srv *photoSrv) GetByID(ID uint64) (result entity.Photo, statusCode int, err error) {
	result, err = srv.PhotoRepository.GetByID(ID)
	if errors.Is(err, gorm.ErrRecordNotFound) || result.ID == 0 {
		return result, http.StatusNotFound, err
	}

	if err != nil {
		log.Printf("[PhotoService-GetByID] error get data repo: %+v \n", err)
		return result, http.StatusInternalServerError, err
	}

	return result, http.StatusOK, nil
}

func (srv *photoSrv) UpdateByID(ID uint64, input dto.PhotoCreateUpdateReq) (result dto.PhotoCreateUpdateResponse, statusCode int, err error) {
	sm, err := srv.PhotoRepository.GetByID(ID)
	if errors.Is(err, gorm.ErrRecordNotFound) || sm.ID == 0 {
		return result, http.StatusNotFound, err
	}

	if err != nil {
		log.Printf("[PhotoService-UpdateByID] error get data repo: %+v \n", err)
		return result, http.StatusInternalServerError, err
	}

	temp, err := srv.PhotoRepository.UpdateByID(ID, input)
	if err != nil {
		log.Printf("[PhotoService-UpdateByID] error update data repo: %+v \n", err)
		return result, http.StatusInternalServerError, err
	}

	result = dto.PhotoCreateUpdateResponse{
		ID:        temp.ID,
		UserID:    temp.UserID,
		Title:     temp.Title,
		PhotoUrl:  temp.PhotoUrl,
		Caption:   temp.Caption,
		CreatedAt: temp.CreatedAt,
		UpdatedAt: temp.UpdatedAt,
	}

	return result, http.StatusOK, nil
}

func (srv *photoSrv) DeleteByID(ID uint64) (statusCode int, err error) {
	sm, err := srv.PhotoRepository.GetByID(ID)
	if errors.Is(err, gorm.ErrRecordNotFound) || sm.ID == 0 {
		return http.StatusNotFound, err
	}

	if err != nil {
		log.Printf("[PhotoService-DeleteByID] error get data repo: %+v \n", err)
		return http.StatusInternalServerError, err
	}

	err = srv.PhotoRepository.DeleteByID(ID)
	if err != nil {
		log.Printf("[PhotoService-DeleteByID] error delete data repo: %+v \n", err)
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
