package main

import (
	"fmt"

	//"github.com/GoesToEleven/GolangTraining/02_package/icomefromalaska"
	someAlias "github.com/GoesToEleven/GolangTraining/02_package/icomefromalaska"
	"github.com/GoesToEleven/GolangTraining/02_package/stringutil"
)

func main() {
	fmt.Println(stringutil.Reverse("!oG ,olleH"))
	fmt.Println(stringutil.MyName)
	fmt.Println(someAlias.BearName)
}
