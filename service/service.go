package service

// Services all service object injected here
type Services struct {
	User        UserService
	SocialMedia SocialMediaService
	Comment     CommentService
	Photo       PhotoService
}
