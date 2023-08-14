package main

import "testing"

func TestSoma(t *testing.T) {

	total := soma(15, 15)
	if total != 30 {
		t.Error("Resultado é errado, valor esperado era ")
	}
}

func TestSub(t *testing.T) {

	total := sub(3, 3)
	if total != 0 {
		t.Error("Resultado é errado, valor esperado era ")
	}
}

func TestPrint(t *testing.T) {

	total := main()
	if total != "" {
		t.Error("Resultado é errado, valor esperado era ")
	}
}
