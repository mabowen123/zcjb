package utility

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"time"
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
		err := g.Redis(redisConn).SetEX(ctx, GetProjectName(), time.Now().Format("2006-01-02 15:04:05"), 10)
		if err != nil {
			panic(redisConn + " redis链接异常")
		}
	}
}
