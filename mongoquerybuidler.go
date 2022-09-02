package mongoquerybuidler

import (
	"encoding/json"
	"fmt"
	"strings"
)

type ResultQuery struct {
	Query interface{}
}

type Params []interface{}
type Query struct {
	Query  string
	Params Params
}

func paramsConverter(params []interface{}) []interface{} {
	paramlength := len(params)
	convert := make([]interface{}, paramlength)

	for i := 0; i < paramlength; i++ {
		convert[i] = fmt.Sprintf("%v", params[i])
	}
	return convert
}

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
