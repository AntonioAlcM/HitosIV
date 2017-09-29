package HitosIV

import (
	"testing"
	"reflect"
)

func TestHitos (t *testing.T){
	t.Log("Test Id");
	if CuantosHitos() <= 0 {
		t.Error("No milestones")
	}
}

func TestTodosHitos (t *testing.T){
	t.Log("Test Todos");
	these_milestones := Hitos()
	if reflect.TypeOf(these_milestones).String() == "Data" {
		t.Error("No milestones here")
	}
}

/*
*	Comprueba el tipo de cada uno de los hitos.
*/
func TestUno(t *testing.T){
	t.Log("Test Uno");
	for i < CuantosHitos(){
		var aux = Uno(i);
		if t.Log(reflect.TypeOf(prueba1)) != "HitosIV.Hito"{
			t.Error("No milestones here")
		}
	}

	t.Log(reflect.TypeOf(prueba1));

	if reflect.TypeOf(prueba1).String() == "Data"{
		t.Log("Ha salido por aquí");
	}else{
		t.Log("Ha salido por allá");
	}
}
