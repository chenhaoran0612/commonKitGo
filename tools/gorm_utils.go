package tools

import (
	"reflect"
	"strings"
)

//GetColumnNameMap 获取数据库 PO Struct 的字段名和库里面的列名对应 map
//Example:
//	po := po.TAccount{}
//	columnMap := tools.GetColumnNameMap(po)
//
//	for k, v := range columnMap {
//		fmt.Println("field name:", k)
//		fmt.Println("field column value:", v)
//	}
func GetColumnNameMap(structs interface{}) map[string]string  {

	typeOfAccount := reflect.TypeOf(structs)
	valueOfAccount := reflect.ValueOf(structs)

	fieldCount := valueOfAccount.NumField()

	columnNameMap := make(map[string]string)
	for i:=0; i < fieldCount; i++{
		fieldType := typeOfAccount.Field(i) // field type

		gormValue := fieldType.Tag.Get("gorm")

		column := ""
		for _, s := range strings.Split(gormValue, ";") {
			if strings.Contains(s, "column") {
				column = strings.Split(s, ":")[1]
				break
			}
		}

		if column != "" {
			columnNameMap[fieldType.Name] = column
		}
	}

	return columnNameMap

}