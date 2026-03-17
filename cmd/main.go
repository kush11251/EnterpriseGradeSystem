package main

import (
	"fmt"
	"github.com/EnterpriseGradeSystem/pkg/controllers"
)

func main() {
	controllers.Init()
	fmt.Println("Enterprise Grade System started")
}