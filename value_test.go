package fastjson

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestValue1(t *testing.T) {
	p := MappedParser()

	c, err := ioutil.ReadFile("./testdata/citm_catalog.json") // just pass the file name
	if err != nil {
		t.Fatal(err)
	}

	v, err := p.Parse(b2s(c))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(v.String())
	fmt.Println()

	for value, contexts := range p.ValueMap() {
		fmt.Println("Value: ")
		fmt.Println(value)
		fmt.Println()
		fmt.Println("Contexts: ")
		for _, context := range contexts {
			fmt.Println("Context ID: ")
			fmt.Println(context.k)
			fmt.Println()

			fmt.Println("MetaData")
			fmt.Println(context.kvs)
			fmt.Println()

			fmt.Println("Outer Contexts")
			fmt.Println("parent: ")
			fmt.Println(context.p)
			fmt.Println()
			fmt.Println("Ancestors: ")
			fmt.Println(context.ans)

			fmt.Println("=====================")
			fmt.Println("=====================")

			fmt.Println("=====================")

		}
		fmt.Println("--------------------------")
	}
}
