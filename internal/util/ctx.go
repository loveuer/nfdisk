package util

import (
	"context"
	"time"
)

func Timeout(seconds ...int) context.Context {
	second := 30
	if len(seconds) > 0 && seconds[0] > 0 {
		second = seconds[0]
	}

	ctx, _ := context.WithTimeout(context.Background(), time.Duration(second)*time.Second)

	return ctx
}
