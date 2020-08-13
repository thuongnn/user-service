package redis

import (
	"fmt"
	rd "github.com/go-redis/redis/v7"
	"time"
)

var client *rd.Client

type Config struct {
	Host     string
	Port     int64
	Password string
}

func Init(c *Config) error {
	//if err := utils.TestTCPConn(fmt.Sprintf("%s:%d", c.Host, c.Port), 60, 2); err != nil {
	//	return err
	//}

	//Initializing redis
	fmt.Println("Initializing redis")
	client = rd.NewClient(&rd.Options{
		Addr:     fmt.Sprintf("%s:%d", c.Host, c.Port),
		Password: c.Password,
	})

	pong, err := client.Ping().Result()
	if err != nil {
		return err
	}

	fmt.Println("Successfully connect to redis")
	fmt.Println(pong)
	return nil
}

func Set(key string, value interface{}, expiration time.Duration) *rd.StatusCmd {
	return client.Set(key, value, expiration)
}

func Get(key string) *rd.StringCmd {
	return client.Get(key)
}

func Del(key string) *rd.IntCmd {
	return client.Del(key)
}
