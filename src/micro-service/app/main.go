package main

import (
	"fmt"
	"github.com/callistaenterprise/goblog/accountservice/service" // 新增代码
)

var appName = "accountservice"

func main() {
	fmt.Printf("Starting %v\n", appName)
	service.StartWebServer("6767") // 新增代码
}