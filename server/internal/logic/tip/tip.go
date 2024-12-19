package tip

import (
	"context"
	"server/internal/param/tip"
	"server/internal/service"
)

type sTip struct {
}

func NewTip() *sTip {
	return &sTip{}
}
func init() {
	service.RegisterTip(NewTip())
}
func (*sTip) List(ctx context.Context, page, size int) (list []*tip.List, err error) {
	return nil, err
}
