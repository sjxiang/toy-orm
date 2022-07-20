package demo

import (
	"reflect"
)

// 最简元数据
type ModelInfo struct {
	tableName string
	fileMap map[string]*FieldInfo
}


type FieldInfo struct {
	// 列
	columnName string
}


// 注册中心
type registry struct {
	models map[reflect.Type]*ModelInfo
}

func (r *registry) register(val interface{}) error {
	mi := &ModelInfo{}
	r.models[reflect.TypeOf(val)] = mi
	return nil 


}


type DB struct {
	r *registry
}

func NewDB() (*DB, error) {
	res := &DB{
		r : &registry{},
	}

	return res, nil
}









func (d *DB) NewSelector() *Selector {
	return &Selector{
		db: d,
	}
}


// 元数据解析



// 驼峰转字符串命名
