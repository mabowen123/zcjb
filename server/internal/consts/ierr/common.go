package ierr

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

var (
	AUTH_ERR_CODE = gcode.New(401, "登录失效", nil)
	AUTH_ERR      = gerror.NewCode(AUTH_ERR_CODE)
)
