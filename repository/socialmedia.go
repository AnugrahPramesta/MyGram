package repository

import (
	"log"
	"project-mygram/dto"
	"project-mygram/entity"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SocialMediaRepository interface {
	Create(input entity.SocialMedia) (res entity.SocialMedia, err error)
	GetAll(ctx *gin.Context, queryparam dto.ListParam) (result []entity.SocialMedia, total uint64, err error)
	GetByID(ID uint64) (result entity.SocialMedia, err error)
	UpdateByID(ID uint64, input dto.SocialMediaUpdateReq) (result entity.SocialMedia, err error)
	DeleteByID(ID uint64) (err error)
}

type socialMediaRepo struct {
	DB *gorm.DB
}

func NewSocialMediaRepository(db *gorm.DB) SocialMediaRepository {
	return &socialMediaRepo{
		DB: db,
	}
}

func (repo *socialMediaRepo) Create(input entity.SocialMedia) (res entity.SocialMedia, err error) {
	if err = repo.DB.Create(&input).Error; err != nil {
		log.Printf("[SocialMediaRepository-Create] error Create new data: %+v \n", err)
		return input, err
	}

	return input, err
}

func (repo *socialMediaRepo) GetAll(ctx *gin.Context, queryparam dto.ListParam) (result []entity.SocialMedia, total uint64, err error) {
	offset := queryparam.Limit * (queryparam.Page - 1)
	var totaldata int64 = 0

	query := repo.DB.WithContext(ctx)

	err = query.Model(&entity.SocialMedia{}).Count(&totaldata).Error
	if err != nil {
		log.Printf("[SocialMediaRepository-GetAll] error count total data: %+v \n", err)
		return
	}

	total = uint64(totaldata)

	err = query.Offset(int(offset)).Limit(int(queryparam.Limit)).
		Preload(clause.Associations, func(db *gorm.DB) *gorm.DB {
			return db.Select("Users.id", "Users.username", "Users.email", "Users.age",
				"Users.created_at", "Users.updated_at")
		}).Find(&result).Error
	if err != nil {
		log.Printf("[SocialMediaRepository-GetAll] error get data: %+v \n", err)
		return
	}

	return
}

func (repo *socialMediaRepo) GetByID(ID uint64) (result entity.SocialMedia, err error) {
	if err = repo.DB.
		Preload(clause.Associations, func(db *gorm.DB) *gorm.DB {
			return db.Select("Users.id", "Users.username", "Users.email", "Users.age",
				"Users.created_at", "Users.updated_at")
		}).
		Where("id = ?", ID).Take(&result).Error; err != nil {
		log.Printf("[SocialMediaRepository-GetByID][%v] error: %+v \n", ID, err)
		return result, err
	}

	return result, err
}

func (repo *socialMediaRepo) UpdateByID(ID uint64, input dto.SocialMediaUpdateReq) (result entity.SocialMedia, err error) {
	err = repo.DB.Clauses(clause.Returning{}).Model(&result).Where("id=?", ID).Updates(entity.SocialMedia{Name: input.Name, SocialMediaUrl: input.SocialMediaUrl}).Error
	if err != nil {
		log.Printf("[SocialMediaRepository-UpdateByID][%v] error: %+v \n", ID, err)
		return result, err
	}

	return result, err
}

func (repo *socialMediaRepo) DeleteByID(ID uint64) (err error) {
	socialMedia := entity.SocialMedia{}
	if err = repo.DB.Where("id = ?", ID).Delete(&socialMedia).Error; err != nil {
		log.Printf("[SocialMediaRepository-DeleteByID][%v] error: %+v \n", ID, err)
		return
	}

	return
}
