package print

import (
	"bytes"
	"fmt"
	"html/template"
	"math/rand"
)

/*
	How to use:

	print.Printf(
		`String is: {{.StringData}}, Int is: {{.IntData}}, Bool is: {{.BoolData}}`,
		print.Add("StringData", "stringData", "IntData", 9999, "BoolData", true)
	)

*/

const (
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

// randomName
func randomName() string {
	b := make([]byte, 10)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

// toString
func toString(p interface{}) string {
	switch p.(type) {
	case string:
		return p.(string)
	}

	return fmt.Sprintf("%s", p)
}

// Add converts a list of interfaces to map[string]interface{}
func Add(p ...interface{}) map[string]interface{} {
	out := map[string]interface{}{}
	if len(p)%2 != 0 {
		fmt.Println("print.Add: Wrong count parameters. It should be len(...) %% 2 == 0")
		return out
	}

	lastkey := ""
	for i, v := range p {
		v := v
		if i%(2) == 0 {
			lastkey = toString(v)
		} else {
			out[lastkey] = v
		}
	}
	return out
}

// Printf prints to stdout with template.
func Printf(tmpl string, data map[string]interface{}) {

	t, err := template.New(randomName()).Parse(tmpl)
	if err != nil {
		fmt.Printf("print.Printf: %+v\n", err)
		return
	}

	buf := &bytes.Buffer{}
	if err = t.Execute(buf, data); err != nil {
		fmt.Printf("print.Printf: %+v\n", err)
		return
	}

	fmt.Printf(buf.String())
}
