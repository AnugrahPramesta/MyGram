package repository

import (
	"log"
	"project-mygram/dto"
	"project-mygram/entity"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CommentRepository interface {
	Create(input entity.Comment) (res entity.Comment, err error)
	GetAll(ctx *gin.Context, queryparam dto.ListParam) (result []entity.Comment, total uint64, err error)
	GetByID(ID uint64) (result entity.Comment, err error)
	UpdateByID(ID uint64, input dto.CommentCreateUpdateReq) (result entity.Comment, err error)
	DeleteByID(ID uint64) (err error)
}

type commentRepo struct {
	DB *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepo{
		DB: db,
	}
}

func (repo *commentRepo) Create(input entity.Comment) (res entity.Comment, err error) {
	if err = repo.DB.Create(&input).Error; err != nil {
		log.Printf("[CommentRepository-Create] error Create new data: %+v \n", err)
		return input, err
	}

	return input, err
}

func (repo *commentRepo) GetAll(ctx *gin.Context, queryparam dto.ListParam) (result []entity.Comment, total uint64, err error) {
	offset := queryparam.Limit * (queryparam.Page - 1)
	var totaldata int64 = 0

	query := repo.DB.WithContext(ctx)

	err = query.Model(&entity.Comment{}).Count(&totaldata).Error
	if err != nil {
		log.Printf("[CommentRepository-GetAll] error count total data: %+v \n", err)
		return
	}

	total = uint64(totaldata)

	err = query.Offset(int(offset)).Limit(int(queryparam.Limit)).
		Preload(clause.Associations).Find(&result).Error
	if err != nil {
		log.Printf("[CommentRepository-GetAll] error get data: %+v \n", err)
		return
	}

	return
}

func (repo *commentRepo) GetByID(ID uint64) (result entity.Comment, err error) {
	if err = repo.DB.
		Preload(clause.Associations).Where("id = ?", ID).Take(&result).Error; err != nil {
		log.Printf("[CommentRepository-GetByID][%v] error: %+v \n", ID, err)
		return result, err
	}

	return result, err
}

func (repo *commentRepo) UpdateByID(ID uint64, input dto.CommentCreateUpdateReq) (result entity.Comment, err error) {
	data := entity.Comment{Message: input.Message}
	if input.PhotoID != 0 {
		data.PhotoID = input.PhotoID
	}

	err = repo.DB.Clauses(clause.Returning{}).Model(&result).Where("id=?", ID).Updates(data).Error
	if err != nil {
		log.Printf("[CommentRepository-UpdateByID][%v] error: %+v \n", ID, err)
		return result, err
	}

	return result, err
}

func (repo *commentRepo) DeleteByID(ID uint64) (err error) {
	Comment := entity.Comment{}
	if err = repo.DB.Where("id = ?", ID).Delete(&Comment).Error; err != nil {
		log.Printf("[CommentRepository-DeleteByID][%v] error: %+v \n", ID, err)
		return
	}

	return
}
