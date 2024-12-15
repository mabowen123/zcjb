package redis

import (
	"fmt"
	"server/utility"
)

func GetKey(str string, a ...any) string {
	if len(a) > 0 {
		str = fmt.Sprintf(str, a...)
	}
	prefix := utility.GetProjectName()
	return prefix + ":" + str
}
