// Vektor project main.go
package main

type Vektor3d struct {
	x, y, z float64
}

func main() {}

func add(a Vektor3d, b Vektor3d) Vektor3d {
	s := Vektor3d{a.x + b.x, a.y + b.y, a.z + b.z}
	return s
}

func sub(a Vektor3d, b Vektor3d) Vektor3d {
	s := Vektor3d{a.x - b.x, a.y - b.y, a.z - b.z}
	return s
}

func mult(a Vektor3d, b float64) Vektor3d {
	s := Vektor3d{a.x * b, a.y * b, a.z * b}
	return s
}

func skalar(a Vektor3d, b Vektor3d) float64 {
	s := a.x*b.x + a.y*b.y + a.z*b.z
	return s
}
