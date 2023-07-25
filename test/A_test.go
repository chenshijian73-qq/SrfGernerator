package test

import (
	"flag"
	"fmt"
	"testing"
)

var (
	input  string
	output string
)

func init() {
	flag.StringVar(&input, "in", "", " (optional; default is stdin)")
	flag.StringVar(&output, "out", "", "(optional; default is stdout)")
}

func Test_A(t *testing.T) {
	A()
}
func A() {
	fmt.Println("test")
}
