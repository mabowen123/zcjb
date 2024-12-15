package ierr

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

const COMMON = 100001

func Common(msg string) error {
	return gerror.NewCode(gcode.New(COMMON, msg, nil))
}
