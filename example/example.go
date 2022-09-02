package example

import (
	"fmt"

	buidler "github.com/cluster05/mongoquerybuidler"
)

func main() {

	query := buidler.Query{
		// pass ~number
		Query: `{
				"boolean":~1,
				"int":~2,
				"float":~3,
				"array":~4,
				"object":~5,
				"string":~6
			}`,
		// respective value for ~number
		Params: buidler.Params{
			false, // boolean
			10,    // int
			12.34, // flaot
			`[{ "foo": "bar" }, { "foo": 1 }, { "foo": false }]`, // array  put `[]` >> JSON string array
			`{ "foo": { "nested": true } }`,                      // object put `{}` >> JSON string object
			`"string"`,                                           // string put string `"string"` format
		},
	}

	resultQuery, err := query.QueryBuilder()
	if err != nil {
		fmt.Println(err)
		return
	}

	resultQuery.Print() //	print the created query in string
	// resultQuery.Query										//	query to pass

}
