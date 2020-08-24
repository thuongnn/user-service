package main

import (
	"fmt"
	"github.com/astaxie/beego"
	consulapi "github.com/hashicorp/consul/api"
	"log"
	"os"
	"user-service/config"
	"user-service/src/api"
	"user-service/src/common/redis"
	"user-service/src/dao"
)

func main() {
	config.Init()

	// init redis connect
	if err := redis.Init(config.RedisConfig()); err != nil {
		log.Fatalf("failed to initialize redis: %v", err)
	}

	// init postgres database connect
	if err := dao.InitDatabase(config.DatabaseConfig()); err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}

	registerRoutes()
	registerServiceWithConsul()
	beego.Run(config.AppPort())
}

// RegisterRoutes for User Service
func registerRoutes() {
	beego.Router("/healthcheck", &api.HomeAPI{}, "get:Get")
	beego.Router("/token/refresh", &api.AuthAPI{}, "post:RefreshToken")
	beego.Router("/sign-up", &api.AuthAPI{}, "post:Register")
	beego.Router("/sign-in", &api.AuthAPI{}, "post:Login")
	beego.Router("/sign-out", &api.AuthAPI{}, "post:Logout")
	beego.Router("/current", &api.UserAPI{}, "get:Get")
}

func registerServiceWithConsul() {
	consulConfig := consulapi.DefaultConfig()
	consul, err := consulapi.NewClient(consulConfig)
	if err != nil {
		log.Fatalln(err)
	}

	registration := new(consulapi.AgentServiceRegistration)

	registration.ID = "user-service"
	registration.Name = "user-service"
	address := hostname()
	registration.Address = address

	p := config.GetUserServicePort()
	registration.Port = p
	registration.Check = new(consulapi.AgentServiceCheck)
	registration.Check.HTTP = fmt.Sprintf("http://%s:%v/healthcheck", address, p)
	registration.Check.Interval = "5s"
	registration.Check.Timeout = "3s"
	err = consul.Agent().ServiceRegister(registration)
	if err != nil {
		log.Fatal(err)
	}
}

func hostname() string {
	hn, err := os.Hostname()
	if err != nil {
		log.Fatalln(err)
	}
	return hn
}
