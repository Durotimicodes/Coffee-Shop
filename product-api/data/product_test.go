package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name: "Oluwadurotimi",
		Price: 35.00,
		SKU: "abe-abc-def",
	}

	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
