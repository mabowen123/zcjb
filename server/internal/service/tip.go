// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"server/internal/param/tip"
)

type (
	ITip interface {
		List(ctx context.Context, page int, size int) (list []*tip.List, err error)
	}
)

var (
	localTip ITip
)

func Tip() ITip {
	if localTip == nil {
		panic("implement not found for interface ITip, forgot register?")
	}
	return localTip
}

func RegisterTip(i ITip) {
	localTip = i
}
