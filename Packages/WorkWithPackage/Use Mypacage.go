package main

import (
	"fmt"
	"project/module/Packages/WorkWithPackage/OtherTry"
	"project/module/Packages/WorkWithPackage/mypackage"
)

func main() {
	fmt.Println(mypackage.Do()) //Print "1"
	fmt.Println(OtherTry.Do())  // Print "1 /n 1"
}
