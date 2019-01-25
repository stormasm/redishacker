package redisc

import (
	"fmt"
	"strconv"
	"github.com/garyburd/redigo/redis"
)

func Sscan(key string, newData chan<- float64) error {

	var (
		myid       string
		cursor     int64
		items      []string
		total      int
		count      int
	)

	c := getRedisConn()
	defer c.Close()

	fmt.Println("Processing Redis Set Key", key)

	for {
		values, err := redis.Values(c.Do("SSCAN", key, cursor))

		if err != nil {
			fmt.Println("sscan error on redis.Values")
		}

		values, err = redis.Scan(values, &cursor, &items)
		if err != nil {
			fmt.Println("sscan error on redis.Scan")
		}

		for num, item := range items {
			// Grab the ID
			fmt.Println(num, item);
			myid = item
			f, err := strconv.ParseFloat(myid, 64)
			if err != nil {
				fmt.Println("strconv.ParseFloat error in redis.Scan")
			}
			newData <- f
		}
		total = total + len(items)
		count = count + 1
		if cursor == 0 {
			break
		}
	}
	fmt.Println("Total = ", total)
	return nil
}
