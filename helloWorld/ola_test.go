package main

import "testing"

func TestOla(t *testing.T) {
	x := "Chris"
	resultado := Ola(x)
	esperado := "Olá, Chris"

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}
