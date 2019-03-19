package fastjson

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func testValue1(t *testing.T) {
	var p Parser
	fmt.Println("I'm here.")

	p.c.mapped = true

	c, err := ioutil.ReadFile("/testdata/small.json") // just pass the file name
	if err != nil {
		t.Fatal(err)
	}

	v, err := p.Parse(b2s(c))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(v.String())
	fmt.Println(p.c.v)
}
