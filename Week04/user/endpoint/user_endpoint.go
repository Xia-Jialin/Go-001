package endpoint

import (
	"context"

<<<<<<< HEAD
	"github.com/Xia-Jialin/Go-000/Week04/user/service"
=======
	"github.com/cty898/Go-000/Week04/user/service"
>>>>>>> ee315abcb790869e30fe090de1f8ea2ef5d6413e
	"github.com/go-kit/kit/endpoint"
)

type UserEndpoints struct {
	RegisterEndpoint endpoint.Endpoint
	LoginEndpoint    endpoint.Endpoint
}

type LoginRequest struct {
	Email    string
	Password string
}

type LoginResponse struct {
	UserInfo *service.UserInfoDTO `json:"user_info"`
}

func MakeLoginEndpoint(userService service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*LoginRequest)
		userInfo, err := userService.Login(ctx, req.Email, req.Password)
		return &LoginResponse{UserInfo: userInfo}, err

	}
}

type RegisterRequest struct {
	Username string
	Email    string
	Password string
}

type RegisterResponse struct {
	UserInfo *service.UserInfoDTO `json:"user_info"`
}

func MakeRegisterEndpoint(userService service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*RegisterRequest)
		userInfo, err := userService.Register(ctx, &service.RegisterUserVO{
			Username: req.Username,
			Password: req.Password,
			Email:    req.Email,
		})
		return &RegisterResponse{UserInfo: userInfo}, err

	}
}
