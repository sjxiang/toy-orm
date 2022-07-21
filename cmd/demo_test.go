package demo

import (
	"fmt"
	// "reflect"
	"testing"

)


func TestSelector_Build(t *testing.T) {
	
	type TestModel struct {
		ID int
		Name string
	}
	
	db, _ := NewDB("mysql", "admin:admin123@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	output, _ := db.r.register(TestModel{})
	fmt.Printf("%+v", output)
}

// &{tableName:TestModel fileMap:map[ID:0xc00005e420 Name:0xc00005e440]}
// 元数据