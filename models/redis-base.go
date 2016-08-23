package models

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

var (
	// RedisClient a available Redis Client.
	RedisClient *redis.Pool
	// RedisDB The number of DB in Redis.
	RedisDB = 2
)

// initRedis Connect RedisDB and Make it available.
func initRedis(host string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     64,
		IdleTimeout: 3 * time.Second,
		MaxActive:   99999, // max number of connections
		// TestOnBorrow: func(c redis.Conn, t time.Time) error {
		// 	_, err := c.Do("PING")

		// 	return err
		// },
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", host)
			if err != nil {
				return nil, err
			}
			return c, err
		},
	}
}

// GetRedisClient return an available Redis Client.
// Usage:
//      GetRedisClient()
//      // Get a Conn from Pool
//      rc := RedisClient.Get()
//      // Return Conn into Pool When you are Done.
//      defer rc.Close()
//
func GetRedisClient() {
	RedisClient = initRedis(":6379")
}