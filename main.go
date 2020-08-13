package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"user-service/config"
	"user-service/src/api"
	"user-service/src/common/redis"
	"user-service/src/dao"
)

func main() {
	config.Init()

	// init redis connect
	if err := redis.Init(config.RedisConfig()); err != nil {
		fmt.Printf("failed to initialize redis: %v", err)
	}

	// init postgres database connect
	if err := dao.InitDatabase(config.DatabaseConfig()); err != nil {
		fmt.Printf("failed to initialize database: %v", err)
	}

	registerRoutes()
	beego.Run(config.AppPort())
}

// RegisterRoutes for User Service
func registerRoutes() {
	beego.Router("/", &api.HomeAPI{}, "get:Get")
	beego.Router("/token/refresh", &api.AuthAPI{}, "post:RefreshToken")
	beego.Router("/sign-up", &api.AuthAPI{}, "post:Register")
	beego.Router("/sign-in", &api.AuthAPI{}, "post:Login")
	beego.Router("/sign-out", &api.AuthAPI{}, "post:Logout")
	beego.Router("/users/current", &api.UserAPI{}, "get:Get")
}
