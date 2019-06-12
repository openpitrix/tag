package mysql

import "github.com/jinzhu/gorm"

func Connect(url string, logOn bool) *gorm.DB {
	conn, err := gorm.Open("mysql", url)
	if err == nil {
		conn.LogMode(logOn)
		conn.SingularTable(true)
	} else {
		panic(err)
	}
	return conn
}

type QueryOption struct {
	Conditions map[string]interface{}
	Columns    string
	Preload    []string
	GroupBy    string
	Having     map[string]interface{}
	OrderBy    string
	Offset     int
	Limit      int
}

func (option *QueryOption) GetWhere() (query string, args []interface{}) {
	glue := ""
	for key, value := range option.Conditions {
		query += glue + key
		if values, assert := value.([]interface{}); assert {
			args = append(args, values...)
		} else {
			args = append(args, value)
		}
		glue = " AND "
	}
	return
}

func (option *QueryOption) GetHaving() (query string, arg interface{}) {
	for key, value := range option.Having {
		query = key
		arg = value
		break
	}
	return
}

type Query struct {
	Fn func(*QueryOption)
}

func Columns(cols string) Query {
	return Query{func(ops *QueryOption) {
		ops.Columns = cols
	}}
}

func Conditions(c map[string]interface{}) Query {
	return Query{func(ops *QueryOption) {
		ops.Conditions = c
	}}
}

func OrderBy(o string) Query {
	return Query{func(ops *QueryOption) {
		ops.OrderBy = o
	}}
}

func GroupBy(g string) Query {
	return Query{func(ops *QueryOption) {
		ops.GroupBy = g
	}}
}

func Having(h map[string]interface{}) Query {
	return Query{func(ops *QueryOption) {
		ops.Having = h
	}}
}

func Preload(pload []string) Query {
	return Query{func(ops *QueryOption) {
		ops.Preload = pload
	}}
}

func Paginator(page, size int) Query {
	return Query{func(ops *QueryOption) {
		ops.Offset = (page - 1) * size
		if ops.Offset < 0 {
			ops.Offset = 0
		}
		ops.Limit = size
	}}
}
