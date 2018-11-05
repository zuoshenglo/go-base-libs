package go_base_libs

import (
	"github.com/garyburd/redigo/redis"
)

var workRedis Redis

type Redis struct {
	Protocol  string
	Address   string
	Db        int64
	Password  string
	User      string
	RedisConn redis.Conn
}

// 初始化redis的链接
func NewRedisConn(address string, db int, password string) (Redis, error) {
	c, err := redis.Dial("tcp", address, redis.DialDatabase(db), redis.DialPassword(password))
	if err != nil {
		return workRedis, err
	}
	workRedis.RedisConn = c
	return workRedis, nil
}

// 设置redis的地址
func (r *Redis) SetAddress(address string) *Redis {
	r.Address = address
	return r
}

// 设置redis的用户
func (r *Redis) SetUser(user string) *Redis {
	r.User = user
	return r
}

// 设置redis的密码
func (r *Redis) SetPassword(password string) *Redis {
	r.Password = password
	return r
}

// 设置使用的redis的Db
func (r *Redis) SetDb(db int64) *Redis {
	r.Db = db
	return r
}

//断开链接
func (r *Redis) DisConn() error {
	if err := r.RedisConn.Close(); err != nil {
		return err
	}
	return nil
}
