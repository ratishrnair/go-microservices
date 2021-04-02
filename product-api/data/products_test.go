package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:        "Latte",
		Description: "Frothy Milky Coffee",
		Price:       2,
		SKU:         "abc-qwer-wde",
	}
	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
