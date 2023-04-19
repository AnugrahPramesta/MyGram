package repository

import (
	"log"
	"project-mygram/dto"
	"project-mygram/entity"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PhotoRepository interface {
	Create(input entity.Photo) (res entity.Photo, err error)
	GetAll(ctx *gin.Context, queryparam dto.ListParam) (result []entity.Photo, total uint64, err error)
	GetByID(ID uint64) (result entity.Photo, err error)
	UpdateByID(ID uint64, input dto.PhotoCreateUpdateReq) (result entity.Photo, err error)
	DeleteByID(ID uint64) (err error)
}

type photoRepo struct {
	DB *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) PhotoRepository {
	return &photoRepo{
		DB: db,
	}
}

func (repo *photoRepo) Create(input entity.Photo) (res entity.Photo, err error) {
	if err = repo.DB.Create(&input).Error; err != nil {
		log.Printf("[PhotoRepository-Create] error Create new data: %+v \n", err)
		return input, err
	}

	return input, err
}

func (repo *photoRepo) GetAll(ctx *gin.Context, queryparam dto.ListParam) (result []entity.Photo, total uint64, err error) {
	offset := queryparam.Limit * (queryparam.Page - 1)
	var totaldata int64 = 0

	query := repo.DB.WithContext(ctx)

	err = query.Model(&entity.Photo{}).Count(&totaldata).Error
	if err != nil {
		log.Printf("[PhotoRepository-GetAll] error count total data: %+v \n", err)
		return
	}

	total = uint64(totaldata)

	err = query.Offset(int(offset)).Limit(int(queryparam.Limit)).
		Preload(clause.Associations, func(db *gorm.DB) *gorm.DB {
			return db.Select("Users.id", "Users.username", "Users.email", "Users.age",
				"Users.created_at", "Users.updated_at")
		}).
		Find(&result).Error
	if err != nil {
		log.Printf("[PhotoRepository-GetAll] error get data: %+v \n", err)
		return
	}

	return
}

func (repo *photoRepo) GetByID(ID uint64) (result entity.Photo, err error) {
	if err = repo.DB.
		Preload(clause.Associations, func(db *gorm.DB) *gorm.DB {
			return db.Select("Users.id", "Users.username", "Users.email", "Users.age",
				"Users.created_at", "Users.updated_at")
		}).Where("id = ?", ID).Take(&result).Error; err != nil {
		log.Printf("[PhotoRepository-GetByID][%v] error: %+v \n", ID, err)
		return result, err
	}
	result.User.Password = ""
	return result, err
}

func (repo *photoRepo) UpdateByID(ID uint64, input dto.PhotoCreateUpdateReq) (result entity.Photo, err error) {
	data := entity.Photo{Title: input.Title, PhotoUrl: input.PhotoUrl, Caption: input.Caption}
	err = repo.DB.Clauses(clause.Returning{}).Model(&result).Where("id=?", ID).Updates(data).Error
	if err != nil {
		log.Printf("[PhotoRepository-UpdateByID][%v] error: %+v \n", ID, err)
		return result, err
	}

	return result, err
}

func (repo *photoRepo) DeleteByID(ID uint64) (err error) {
	Photo := entity.Photo{}
	if err = repo.DB.Where("id = ?", ID).Delete(&Photo).Error; err != nil {
		log.Printf("[PhotoRepository-DeleteByID][%v] error: %+v \n", ID, err)
		return
	}

	return
}
