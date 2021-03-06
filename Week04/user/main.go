package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

<<<<<<< HEAD
	"github.com/Xia-Jialin/Go-000/Week04/user/dao"
	"github.com/Xia-Jialin/Go-000/Week04/user/endpoint"
	"github.com/Xia-Jialin/Go-000/Week04/user/redis"
	"github.com/Xia-Jialin/Go-000/Week04/user/service"
)

//NewUserEndpoints New User Endpoints
=======
	"github.com/cty898/Go-000/Week04/user/dao"
	"github.com/cty898/Go-000/Week04/user/endpoint"
	"github.com/cty898/Go-000/Week04/user/redis"
	"github.com/cty898/Go-000/Week04/user/service"
)

>>>>>>> ee315abcb790869e30fe090de1f8ea2ef5d6413e
func NewUserEndpoints(userService service.UserService) *endpoint.UserEndpoints {
	userEndpoints := &endpoint.UserEndpoints{
		RegisterEndpoint: endpoint.MakeRegisterEndpoint(userService),
		LoginEndpoint:    endpoint.MakeLoginEndpoint(userService),
	}
	return userEndpoints
}

func main() {

	var (
		// 服务地址和服务名
		servicePort = flag.Int("service.port", 10086, "service port")

		//waitTime = flag.Int("wait.time", 10, "wait time")

		mysqlAddr = flag.String("mysql.addr", "127.0.0.1", "mysql addr")

		mysqlPort = flag.String("mysql.port", "3306", "mysql port")

		redisAddr = flag.String("redis.addr", "127.0.0.1", "redis addr")

		redisPort = flag.String("redis.port", "6379", "redis port")
	)

	flag.Parse()

	ctx := context.Background()
	errChan := make(chan error)

	err := dao.InitMysql(*mysqlAddr, *mysqlPort, "root", "123456", "user")
	if err != nil {
		log.Fatal(err)
	}

	err = redis.InitRedis(*redisAddr, *redisPort, "")
	if err != nil {
		log.Fatal(err)
	}

	//使用 Wire 构建依赖
	r := InitHttpHandler(&dao.UserDAOImpl{}, ctx)

	go func() {
		errChan <- http.ListenAndServe(":"+strconv.Itoa(*servicePort), r)
	}()

	go func() {
		// 监控系统信号，等待 ctrl + c 系统信号通知服务关闭
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	error := <-errChan
	log.Println(error)

}
