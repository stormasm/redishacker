package redisc

import (
	"fmt"
)

func Delkey(key string) error {
	c := getRedisConn()
	defer c.Close()

	_, err := c.Do("DEL", key)
	fmt.Println("Deleted key", key)
	return err
}
