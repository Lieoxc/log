package main

import (
	"log"
	"testing"

	logs "github.com/Lieoxc/log"
)

func Benchmark_infoLog(b *testing.B) {
	for i := 0; i < b.N; i++ {
		logs.Info("123")
	}
}
func Benchmark_systemLog(b *testing.B) {

	for i := 0; i < b.N; i++ {
		log.Println("123")
	}
}
