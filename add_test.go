package main 

import "testing"

func TestAdd(t *testing.T) {
	a := Vektor3d{0, 2, 3}
	b := Vektor3d{3, 5, 4}
	var res Vektor3d = add(a, b)
	if !(res.x == 3 && res.y == 7 && res.z == 7) {
		t.Error("Fehler gefunden")
	}
}

func TestSub(t *testing.T) {
	a := Vektor3d{0, 2, 3}
	b := Vektor3d{3, 5, 4}
	var res Vektor3d = sub(a, b)
	if !(res.x == -3 && res.y == -3 && res.z == -1) {
		t.Error("Fehler gefunden")
	}
}

func TestMult(t *testing.T) {
	a := Vektor3d{0, 2, 3}
	b := 4.0
	var res Vektor3d = mult(a, b)
	if !(res.x == 0 && res.y == 8 && res.z == 12) {
		t.Error("Fehler gefunden")
	}
}

func TestSkalar(t *testing.T) {
	a := Vektor3d{0, 2, 3}
	b := Vektor3d{3, 5, 4}
	var res float64 = skalar(a, b)
	if !(res == 22.0) {
		t.Error("Fehler gefunden")
	}
}
