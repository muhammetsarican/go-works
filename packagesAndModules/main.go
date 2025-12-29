package main

import (
	"fmt"

	"my-module/packagesAndModules/helper"
	"my-module/packagesAndModules/helper/rest"

	"my-module/packagesAndModules/helper2"
	helper2Rest "my-module/packagesAndModules/helper2/rest"
)

func main() {
	fmt.Println("packages and modules")
	helper.Helper1()
	rest.Rest()
	rest.Rest2()

	helper2.Helper2()
	helper2Rest.Rest()
}
