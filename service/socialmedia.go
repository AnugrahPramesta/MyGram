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

type SocialMediaService interface {
	Create(input entity.SocialMedia) (res dto.SocialMediaCreateUpdateResponse, statusCode int, err error)
	GetAll(ctx *gin.Context, queryparam dto.ListParam) (result dto.SocialMediaListRes, statusCode int, err error)
	GetByID(ID uint64) (result entity.SocialMedia, statusCode int, err error)
	UpdateByID(ID uint64, input dto.SocialMediaUpdateReq) (result dto.SocialMediaCreateUpdateResponse, statusCode int, err error)
	DeleteByID(ID uint64) (statusCode int, err error)
}

type socialMediaSrv struct {
	SocialMediaRepository repository.SocialMediaRepository
}

func NewSocialMediaService(SocialMediaRepo repository.SocialMediaRepository) SocialMediaService {
	return &socialMediaSrv{
		SocialMediaRepository: SocialMediaRepo,
	}
}

func (srv *socialMediaSrv) Create(input entity.SocialMedia) (res dto.SocialMediaCreateUpdateResponse, statusCode int, err error) {
	temp, err := srv.SocialMediaRepository.Create(input)
	if err != nil {
		log.Printf("[SocialMediaService-Create] error create data: %+v \n", err)
		return res, http.StatusInternalServerError, err
	}

	res = dto.SocialMediaCreateUpdateResponse{
		ID:             temp.ID,
		Name:           temp.Name,
		SocialMediaUrl: temp.SocialMediaUrl,
		UserID:         temp.UserID,
		CreatedAt:      temp.CreatedAt,
		UpdatedAt:      temp.UpdatedAt,
	}

	return res, http.StatusCreated, nil
}

func (srv *socialMediaSrv) GetAll(ctx *gin.Context, queryparam dto.ListParam) (result dto.SocialMediaListRes, statusCode int, err error) {
	result.Limit = queryparam.Limit
	result.Page = queryparam.Page

	result.Data, result.Total, err = srv.SocialMediaRepository.GetAll(ctx, queryparam)
	if err != nil {
		log.Printf("[SocialMediaService-GetAll] error get data repo: %+v \n", err)
		return result, http.StatusInternalServerError, err
	}

	result.LastPage = uint64(math.Ceil(float64(result.Total) / float64(queryparam.Limit)))

	return result, http.StatusOK, nil
}

func (srv *socialMediaSrv) GetByID(ID uint64) (result entity.SocialMedia, statusCode int, err error) {
	result, err = srv.SocialMediaRepository.GetByID(ID)
	if errors.Is(err, gorm.ErrRecordNotFound) || result.ID == 0 {
		return result, http.StatusNotFound, err
	}

	if err != nil {
		log.Printf("[SocialMediaService-GetByID] error get data repo: %+v \n", err)
		return result, http.StatusInternalServerError, err
	}

	return result, http.StatusOK, nil
}

func (srv *socialMediaSrv) UpdateByID(ID uint64, input dto.SocialMediaUpdateReq) (result dto.SocialMediaCreateUpdateResponse, statusCode int, err error) {
	sm, err := srv.SocialMediaRepository.GetByID(ID)
	if errors.Is(err, gorm.ErrRecordNotFound) || sm.ID == 0 {
		return result, http.StatusNotFound, err
	}

	if err != nil {
		log.Printf("[SocialMediaService-UpdateByID] error get data repo: %+v \n", err)
		return result, http.StatusInternalServerError, err
	}

	temp, err := srv.SocialMediaRepository.UpdateByID(ID, input)
	if err != nil {
		log.Printf("[SocialMediaService-UpdateByID] error update data repo: %+v \n", err)
		return result, http.StatusInternalServerError, err
	}

	result = dto.SocialMediaCreateUpdateResponse{
		ID:             temp.ID,
		Name:           temp.Name,
		SocialMediaUrl: temp.SocialMediaUrl,
		UserID:         temp.UserID,
		CreatedAt:      temp.CreatedAt,
		UpdatedAt:      temp.UpdatedAt,
	}

	return result, http.StatusOK, nil
}

func (srv *socialMediaSrv) DeleteByID(ID uint64) (statusCode int, err error) {
	sm, err := srv.SocialMediaRepository.GetByID(ID)
	if errors.Is(err, gorm.ErrRecordNotFound) || sm.ID == 0 {
		return http.StatusNotFound, err
	}

	if err != nil {
		log.Printf("[SocialMediaService-DeleteByID] error get data repo: %+v \n", err)
		return http.StatusInternalServerError, err
	}

	err = srv.SocialMediaRepository.DeleteByID(ID)
	if err != nil {
		log.Printf("[SocialMediaService-DeleteByID] error delete data repo: %+v \n", err)
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
