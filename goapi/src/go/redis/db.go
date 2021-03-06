package redis

import (
	"github.com/gomodule/redigo/redis"
	"log"
	"os"
)



var conn redis.Conn


func newPool() *redis.Pool {
	return &redis.Pool{
		// Maximum number of idle connections in the pool.
		MaxIdle: 80,
		// max number of connections
		MaxActive: 12000,
		// Dial is an application supplied function for creating and
		// configuring a connection.
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", os.Getenv("REDIS_ADDR"))
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},

	}
}

func getConnection() redis.Conn {
	if conn == nil {
		conn = newPool().Get()
	}
	return conn
}

func Set(key string, value string) error {
	c := getConnection()
	log.Printf("SET %s %s", key, value)
	_, err := c.Do("SET", key, value)
	if err != nil {
		return err
	}
	return nil
}
func MustSet(key string, value string) {
	err := Set(key, value)
	if err != nil {
		panic(err)
	}
}

func Setex(key string, expire int64, value string) error {
	c := getConnection()
	log.Printf("SETEX %s %d %s", key, expire, value)
	_, err := c.Do("SETEX", key, expire, value)
	if err != nil {
		return err
	}
	return nil
}

func MustSetex(key string, expire int64, value string) {
	err := Setex(key, expire, value)
	if err != nil {
		panic(err)
	}
}

func Get(key string) (string, error) {
	c := getConnection()
	log.Printf("GET %s\n", key)
	r, err := c.Do("GET", key)
	if err != nil {
		return "", err
	}
	if r == nil {
		return "", nil
	}
	return redis.String(r, err)
}

func MustGet(key string) string {
	s, err := Get(key)
	if err != nil {
		panic(err)
	}
	return s
}

func Del(key string) (int64, error) {
	c := getConnection()
	log.Printf("DEL %s", key)
	return redis.Int64(c.Do("DEL", key))
}

func MustDel(key string) int64 {
	keys, err := Del(key)
	if err != nil {
		panic(err)
	}
	return keys
}

