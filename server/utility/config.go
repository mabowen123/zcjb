package utility

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

var ctx = gctx.New()

func GetProjectName() string {
	return getConfig("config", "server.projectName").String()
}

func getConfig(fileName string, node string) *g.Var {
	config, _ := g.Cfg(fileName).Get(ctx, node)
	return config
}

func GetJwtKey() []byte {
	key := getConfig("config", "server.jwtKey").String()

	if key == "" {
		key = GetProjectName()
	}

	return []byte(key)
}

func GetDatabaseConn() map[string]interface{} {
	data := getConfig("config", "database").Map()
	delete(data, "logger")
	return data
}

func GetRedisConn() map[string]interface{} {
	return getConfig("config", "redis").Map()
}
