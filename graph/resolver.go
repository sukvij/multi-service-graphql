package graph

import (
	"vijju/post/service"
	userService "vijju/user/service"
)

type Resolver struct {
	UserService *userService.UserService
	PostService *service.PostService
}
