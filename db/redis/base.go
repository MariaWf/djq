package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"mimi/djq/config"
	"sync"
)

var client *redis.Client
var once sync.Once

func Get() *redis.Client {
	once.Do(func() {
		client = redis.NewClient(&redis.Options{
			Addr:     config.Get("redis_address"),
			Password: config.Get("redis_password"), // no password set
			DB:       0,                            // use default DB
		})
	})

	//client := redis.NewClient(&redis.Options{
	//	Addr:     "localhost:6379",
	//	Password: "", // no password set
	//	DB:       0, // use default DB
	//	PoolSize:5,
	//})

	return client
}

func Close(client *redis.Client) error {
	return nil
}

func getClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return client
}

func ExampleNewClient() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
}

func ExampleClient() {
	client := getClient()
	err := client.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := client.Get("name").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exists")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("name", val2)
	}
	// Output: key value
	// key2 does not exists
}
