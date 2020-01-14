package redisdb

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var pool redis.Pool
func init()  {      //init 用于初始化一些参数，先于main执行
    pool = redis.Pool{
        MaxIdle:     16,
        MaxActive:   32,
        IdleTimeout: 120,
        Dial: func() (redis.Conn, error) {
            return redis.Dial("tcp", "193.112.207.30:6379",redis.DialPassword("liu123456"),redis.DialDatabase(0))
        },
    }
}
//GetConnect 获取一个与redis的连接
func GetConnect()(redis.Conn){
	return pool.Get()
}

// SetDo 设置30 分钟过期的key value
func SetDo(key,value string){
    c := GetConnect()
    defer c.Close()
    c.Do("SET",key,value)
    c.Do("EXPIRE", key, 30*60) //设置30分钟过期  
}


// GetDo 有key获取value
func GetDo(key string)(value string,err error){
    c := GetConnect()
    defer c.Close()
    value,err = redis.String(c.Do("GET", key))
    return
}


// DelDo 删除
func DelDo(key string){
    c := GetConnect()
    defer c.Close()
    c.Do("DEL",key)
}

// UpDo 修改
func UpDo(key,value string){
    c := GetConnect()
    defer c.Close()
    time,err := c.Do("TTL",key)
    if err!=nil{
        fmt.Println("a")
        return
    }
    c.Do("SETEX",key,time,value)
}

