package demo

import (
	"context"
	"database/sql"
	"errors"
	"reflect"
	"unicode"
)

// 最简元数据
type ModelInfo struct {
	tableName string
	fileMap map[string]*FieldInfo
}


type FieldInfo struct {
	// 列名
	columnName string
	// 字段类型
	fieldTypeName string
}


// 注册中心
type registry struct {
	models map[reflect.Type]*ModelInfo
}


// 元数据解析
func (r *registry) register(val interface{}) (*ModelInfo, error) {
	typ := reflect.TypeOf(val)
	if typ.Kind() != reflect.Struct {
		return nil, errors.New("toy-orm：非法类型")
	}

	tableName := typ.Name()  // 表名
	numField := typ.NumField()  // 列的数量
	fdInfos := make(map[string]*FieldInfo, numField)

	for i := 0; i < numField; i++ {
		curField := typ.Field(i)  // 例，{ID  int  0 [0] false}
		fdInfos[curField.Name] = &FieldInfo{
			columnName: curField.Name,
			fieldTypeName: curField.Type.Name(), 
		}
	}
	
	mi := &ModelInfo{
		tableName: underscoreName(tableName),
		fileMap: fdInfos,
	}
	return mi, nil 
}


// DB 是 sql.DB 的装饰器
type DB struct {
	db *sql.DB
	r *registry
}

func NewDB(driver string, dsn string) (*DB, error) {

	// db, err := sql.Open(driver, dsn)
	// if err != nil {
	// 	return nil, err
	// }

	res := &DB{
		// db: db,
		r : &registry{},
	}

	return res, nil
}


func (d *DB) NewSelector() *Selector {
	return &Selector{
		db: d,
	}
}

func (s *Selector) Get(ctx context.Context) (interface{}, error) {
	q, err := s.Build()
	if err != nil {
		return nil, err
	}

	row := s.db.db.QueryRowContext(ctx, q.SQL, q.Arg...)
	if row.Err() != nil {
		return nil, row.Err()
	}

	// TODO
	panic("")
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