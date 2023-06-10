package main


func TestSoma(t *testing.T){

	total := Soma(15,15)

	if total !=30{
		t.Errorf("resultado da soma e invalido: Resultado %d. Esperado: %d", total, 30)
	}
}