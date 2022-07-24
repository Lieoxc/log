package main

import (
	"testing"

	logs "github.com/Lieoxc/log"
	"github.com/zeromicro/go-zero/core/logx"
)

func Benchmark_infoLog(b *testing.B) {
	for i := 0; i < b.N; i++ {
		logs.Info("123")
	}
}
func Benchmark_logx(b *testing.B) {

	for i := 0; i < b.N; i++ {
		logx.Info("123")
	}
}
