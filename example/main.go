package main

import (
	"context"

	logs "github.com/Lieoxc/log"
)

func main() {
	val := 10
	logs.Info("hello world : %v", val)

	ctx := context.Background()
	logs.WithContext(ctx).Info("hello world")
}
