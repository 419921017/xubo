package pkg1

import (
	"fmt"
	"xubo/chapter08/pkginit/pkg2"
)

func ExecPkg1() {
	fmt.Println("Exec Pkg1")
	pkg2.ExecPkg2()
}

func init() {
	fmt.Println("pkg1 init")
}
