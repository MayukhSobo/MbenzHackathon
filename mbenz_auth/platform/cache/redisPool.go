package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	//"github.com/gomodule/redigo/redis"
	//"log"
	//"os"
)

var pool *redis.Client

func init() {
	pool = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "mbenzREDIS@123",
		PoolSize: 12000,
		MinIdleConns: 80,
	})
}

func Ping() error {
	return pool.Ping(context.Background()).Err()
}


func Set(key string, val string) error {
	// get conn and put back when exit from method
	return  pool.Set(context.Background(), key, val, 0).Err()
}

func Remove(key string) error {
	return pool.Del(context.Background(), key).Err()
}
//
//func Get(key string) (string, error) {
//	// get conn and put back when exit from method
//	conn := pool.Get()
//	defer func(conn redis.Conn) {
//		err := conn.Close()
//		if err != nil {
//			panic(err)
//		}
//	}(conn)
//
//	s, err := redis.String(conn.Do("GET", key))
//	if err != nil {
//		log.Printf("ERROR: fail get key %s, error %s", key, err.Error())
//		return "", err
//	}
//
//	return s, nil
//}
//
//func SAdd(key string, val string) error {
//	// get conn and put back when exit from method
//	conn := pool.Get()
//	defer func(conn redis.Conn) {
//		err := conn.Close()
//		if err != nil {
//			panic(err)
//		}
//	}(conn)
//
//	_, err := conn.Do("SADD", key, val)
//	if err != nil {
//		log.Printf("ERROR: fail add val %s to set %s, error %s", val, key, err.Error())
//		return err
//	}
//
//	return nil
//}
//
//func SMembers(key string) ([]string, error) {
//	// get conn and put back when exit from method
//	conn := pool.Get()
//	defer func(conn redis.Conn) {
//		err := conn.Close()
//		if err != nil {
//			panic(err)
//		}
//	}(conn)
//
//	s, err := redis.Strings(conn.Do("SMEMBERS", key))
//	if err != nil {
//		log.Printf("ERROR: fail get set %s , error %s", key, err.Error())
//		return nil, err
//	}
//
//	return s, nil
//}
//
