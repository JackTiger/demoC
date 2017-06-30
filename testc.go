package testc

import (
    "github.com/pelletier/go-toml"
     "fmt"
)

func MethodC() {
    strTest := "test"
	toml.Load(strTest)
    fmt.Println("MethodC")
}