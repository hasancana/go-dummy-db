package dummydb

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"sort"
)

// DummyDatabase ...
type DummyDatabase struct {
	queries map[int]string
	results map[int]string
}

// AddQuery into dummy database...
func (dmd *DummyDatabase) AddQuery(query string) (index string) {
	currentIndex := len(dmd.queries) + 1
	dmd.queries[currentIndex] = query
	return fmt.Sprintf("{{ setValue %d}}", currentIndex)
}

// GetQueries into dummy database...
func (dmd DummyDatabase) GetQueries() map[int]string {
	return dmd.queries
}

// GetQueries into dummy database...
func (dmd DummyDatabase) GetInsertID(index int) string {
	return dmd.results[index]
}

// SetInsertID sets the insert id for the executed query
func (dmd *DummyDatabase) SetInsertID(index int, result string) {
	dmd.results[index] = result
}

// ParseQuery parses the query and replaces placeholders with previous query result
func (dmd DummyDatabase) ParseQuery(query string) string {
	var result bytes.Buffer
	tmpl, _ := template.New("test").Funcs(
		template.FuncMap{
			"setValue": func(val int) string {
				return dmd.GetInsertID(val)
			},
		},
	).Parse(query)

	tmpl.Execute(&result, nil)

	return result.String()
}

// ExecuteQueries executes queries which have been previously stored on the list
func (dmd DummyDatabase) ExecuteQueries() {

	// Execute query for real using 'db' postgres database connection and fetch insert id
	/*
		dbTrans, _ := db.Begin()
		dbTrans.Commit()
		dbTrans.Rollback()
		...
	*/

	queries := dmd.GetQueries()

	var keys []int
	for k := range queries {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {

		// Save insert id for the current query executed
		dmd.SetInsertID(k, "insert id of the query")

		// Parse query and replace placeholders with previous query insert IDs
		query := dmd.ParseQuery(queries[k])

		fmt.Println(query)
	}
}

// DMDatabase returns a dummy database
func DMDatabase(ctx context.Context) DummyDatabase {
	var dummydatabase DummyDatabase

	dummydatabase.queries = make(map[int]string)
	dummydatabase.results = make(map[int]string)
	/*
		if ctx.Value("DMDB") != nil {
			dmdb := ctx.Value("DMDB").(string)
			var ctxQueries map[int]string
			json.Unmarshal([]byte(dmdb), &ctxQueries)
			dummydatabase.queries = ctxQueries
		}
	*/
	return dummydatabase
}
