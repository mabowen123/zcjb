package utility

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

func InitDB() {
	for dbConn, _ := range GetDatabaseConn() {
		err := g.DB(dbConn).PingMaster()
		if err != nil {
			panic(dbConn + " 数据库链接异常")
		}
	}
}

func InitRedis() {
	ctx := context.Background()
	for redisConn, _ := range GetRedisConn() {
		_, err := g.Redis(redisConn).Conn(ctx)
		if err != nil {
			panic(redisConn + " redis链接异常")
		}
	}
}
