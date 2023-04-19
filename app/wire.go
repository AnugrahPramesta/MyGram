package app

import (
	"project-mygram/repository"
	"project-mygram/service"

	"gorm.io/gorm"
)

func WiringRepository(db *gorm.DB) *repository.Repositories {
	return &repository.Repositories{
		User:        repository.NewUserRepository(db),
		SocialMedia: repository.NewSocialMediaRepository(db),
		Comment:     repository.NewCommentRepository(db),
		Photo:       repository.NewPhotoRepository(db),
	}
}

func WiringService(repo *repository.Repositories) *service.Services {
	return &service.Services{
		User:        service.NewUserService(repo.User),
		SocialMedia: service.NewSocialMediaService(repo.SocialMedia),
		Comment:     service.NewCommentService(repo.Comment),
		Photo:       service.NewPhotoService(repo.Photo),
	}
}
