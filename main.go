package main

import (
	"fmt"

	"toy-orm/cmd"
)

func main() {

	// 用户操作字段名 ID，但 Mysql 中使用的是列名 id
	expr := demo.C("ID").EQ(1)
	output, _ := (&demo.Selector{}).From("Test_Table").Where(expr).Build()
	
	fmt.Println(output)  // &{SELECT * FROM `Test_Table` WHERE ID = ? ; []}
}

