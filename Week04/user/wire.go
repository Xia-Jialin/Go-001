//+build wireinject

package main

import (
	"context"
	"net/http"

<<<<<<< HEAD
	"github.com/Xia-Jialin/Go-000/Week04/user/dao"
	"github.com/Xia-Jialin/Go-000/Week04/user/service"
	"github.com/Xia-Jialin/Go-000/Week04/user/transport"
=======
	"github.com/cty898/Go-000/Week04/user/dao"
	"github.com/cty898/Go-000/Week04/user/service"
	"github.com/cty898/Go-000/Week04/user/transport"
>>>>>>> ee315abcb790869e30fe090de1f8ea2ef5d6413e
	"github.com/google/wire"
	"github.com/gorilla/mux"
)

func InitHttpHandler(userDAO dao.UserDAO, ctx context.Context) http.Handler {
	wire.Build(service.MakeUserServiceImpl, NewUserEndpoints, transport.MakeHttpHandler)
	return &mux.Router{}
}
