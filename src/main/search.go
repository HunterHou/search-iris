package main

import (
	"github.com/lxn/walk"
	_ "github.com/lxn/walk/declarative"
)

func main() {
	wwm, _ :=walk.NewMainWindow()
	wwm.Run()
}
