package go_base_libs

import (
	"os"
)

var Tool = & tool{}

type tool struct {
}

func (t * tool) GetCwd() string {
	dir, _ := os.Getwd()
	return dir
}
