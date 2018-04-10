package utils

import (
	"log"

	"github.com/go-redis/redis"
)

// NewCli new a redis client.
func NewCli() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return client
}

// Set key to hold the string value.
//     If key already holds a value, it is overwritten, regardless of its type.
//     Any previous time to live associated with the key is discarded on successful SET operation.
func Set(cli *redis.Client, key string, value string) bool {
	err := cli.Set(key, value, 0).Err()
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

// Get the value of key.
//     If the key does not exist the special value nil is returned.
//     An error is returned if the value stored at key is not a string, because GET only handles string values.
func Get(cli *redis.Client, key string) string {
	result, err := cli.Get(key).Result()
	if err != nil {
		log.Println(err)
		return ""
	}
	return result
}

// Delete Removes the specified keys. A key is ignored if it does not exist.
func Delete(cli *redis.Client, key string) bool {
	err := cli.Del(key).Err()
	if err != nil {
		return false
	}
	return true
}

// LPush insert all the specified values at the head of the list stored at key.
//       If key does not exist, it is created as empty list before performing the push operations.
//       When key holds a value that is not a list, an error is returned.
/****************demo**********************
redis>  LPUSH mylist "world"
(integer) 1
redis>  LPUSH mylist "hello"
(integer) 2
redis>  LRANGE mylist 0 -1
1) "hello"
2) "world"
********************************************/
func LPush(cli *redis.Client, key string, value string) bool {
	err := cli.LPush(key, value).Err()
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

// LRange Returns the specified elements of the list stored at key.
//        The offsets start and stop are zero-based indexes,
//        with 0 being the first element of the list (the head of the list),
//        1 being the next element and so on.
func LRange(cli *redis.Client, key string) []string {
	result, err := cli.LRange(key, 0, -1).Result()
	if err != nil {
		log.Println(err)
		return nil
	}
	return result
}

// Close the client session.
func Close(cli *redis.Client) {
	err := cli.Close()
	if err != nil {
		log.Println(err)
	}
}
