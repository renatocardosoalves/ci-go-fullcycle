package main

import "testing"

func TestSoma(t *testing.T)  {
	total := soma(15, 15)
	if total != 30 {
		t.Errorf("Resultado da soma é inválido: Resultado %d", total)
	}
}

func TestSubtracao(t *testing.T)  {
	total := subtracao(15, 15)
	if total != 0 {
		t.Errorf("Resultado da subtração é inválido: Resultado %d", total)
	}
}