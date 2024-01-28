// main_test.go

package main

import "testing"

// TestOlaMundo é um teste unitário para a função OlaMundo.
func TestOlaMundo(t *testing.T) {
	resultado := OlaMundo()
	esperado := "Olá, mundo!"

	if resultado != esperado {
		t.Fatalf("Falha no teste. Resultado: '%s', Esperado: '%s'", resultado, esperado)
	}
}
