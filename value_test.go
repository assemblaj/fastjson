package fastjson

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestValue1(t *testing.T) {
	var p Parser

	p.c.mapped = true

	c, err := ioutil.ReadFile("./testdata/small.json") // just pass the file name
	if err != nil {
		t.Fatal(err)
	}

	v, err := p.Parse(b2s(c))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(v.String())
	fmt.Println()

	for key, vals := range p.c.v {
		fmt.Println(key)
		fmt.Println()
		for _, val := range vals {
			fmt.Println(val.k)
			fmt.Println()

			fmt.Println(val.kvs)
			fmt.Println()

			fmt.Println(val.ps)
			fmt.Println()

		}
		fmt.Println("--------------------------")

	}
}
