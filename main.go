package main

import (
	"fmt"
	"unicode"
	"toy-orm/cmd"
)

func main() {

	// 用户操作字段名 FirstName，但 Mysql 中使用的是列名 first_name
	expr := demo.C("FirstName").EQ(1)
	output, _ := (&demo.Selector{}).From("Test_Table").Where(expr).Build()
	
	fmt.Println(output)
	fmt.Println(underscoreName("FindByIndex"))
}


// underscoreName 驼峰转下划线，命名 
// 'FindByIndex' => 'find_by_index'
func underscoreName(tableName string) string {
	var buf []byte
	for i, v := range tableName {
		if unicode.IsUpper(v) {
			if i != 0 {
				buf = append(buf, '_')
			}
			buf = append(buf, byte(unicode.ToLower(v)))
		} else {
			buf = append(buf, byte(v))
		}

	}
	return string(buf)
}