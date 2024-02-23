package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	redis "github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

func InitRedis() *redis.Client {

	redis_host := os.Getenv("REDIS_HOST")
	redis_pass := os.Getenv("REDIS_PASS")

	client := redis.NewClient(&redis.Options{
		Addr:     redis_host,
		Password: redis_pass,
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	return client
}

func Query(rdb *redis.Client, query *gorm.DB, key string) interface{} {

	data := GetData(rdb, key)
	if data != nil {
		return data
	}

	query_data, _ := MarshalBinary(query)

	if SetData(rdb, key, query_data) {
		return query_data
	}

	return nil
}

func GetData(rdb *redis.Client, key string) interface{} {

	op2 := rdb.Get(key)

	if err := op2.Err(); err != nil {
		fmt.Printf("unable to GET data. error: %v", err)
		return nil
	}

	res, err := op2.Result()

	if err != nil {
		fmt.Printf("unable to GET data. error: %v", err)
		return nil
	}

	log.Println("get operation success. result:", res)

	return res
}

func SetData(rdb *redis.Client, key string, data interface{}) bool {

	err := rdb.Set(key, data, 1*time.Hour).Err()
	// if there has been an error setting the value
	// handle the error
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

func MarshalBinary(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}
