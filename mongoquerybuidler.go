package mongoquerybuidler

import (
	"encoding/json"
	"fmt"
	"strings"
)

// type ResultQuery is contain Query of type interface{}
// if you pass the object({}) to it, it will return interface{}
// if you pass the array([]) to it, it will return interface{} of type []interface{}
type ResultQuery struct {
	Query interface{}
}

// type Params is type []interface{} is used to pass the query params
type Params []interface{}

// type Query us contain Query of type string and Params of type Params
// Query of string
// 	Query string pass the {} or [] in string format with ~number
//  Params is []interface{}
// 	Usees :
//	int 10
//  float 12.34
//  boolean false,true
//  string use `"string"` in it
//  object pass json stringified json
//  array pass json stringigied array
type Query struct {
	Query  string
	Params Params
}

// param converter is used to take []interface value from Param in Query and convert it to string for processing
func paramsConverter(params []interface{}) []interface{} {
	paramlength := len(params)
	convert := make([]interface{}, paramlength)

	for i := 0; i < paramlength; i++ {
		convert[i] = fmt.Sprintf("%v", params[i])
	}
	return convert
}

// bindParamsWithQuery takes query and params and bind ~number with query
func bindParamsWithQuery(query string, params []interface{}) (string, error) {
	length := len(params)
	for i := 0; i < length; i++ {
		replace := fmt.Sprintf("~%v", i+1)
		query = strings.Replace(query, replace, params[i].(string), -1)
	}
	return query, nil
}

func (q *Query) QueryBuilder() (ResultQuery, error) {
	var mapQuery ResultQuery

	validParams := paramsConverter(q.Params)
	bindedQuery, err := bindParamsWithQuery(q.Query, validParams)

	if err != nil {
		return mapQuery, err
	}

	if err := json.Unmarshal([]byte(bindedQuery), &mapQuery.Query); err != nil {
		return mapQuery, fmt.Errorf("invalid json format : %v", err.Error())
	}

	return mapQuery, nil
}

func (rq *ResultQuery) Print() {
	result, _ := json.MarshalIndent(rq.Query, "", " ")
	fmt.Printf("%s", result)
}
