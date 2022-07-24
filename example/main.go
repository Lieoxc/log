package main

import (
	"context"

	logs "github.com/Lieoxc/log"
)

func main() {
	logs.Info("hello world")

	ctx := context.Background()
	logs.WithContext(ctx).Info("hello world")
}
