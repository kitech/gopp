package cgopp

import "testing"

func TestFC1(t *testing.T) {
	TestLitffi2callz()
	TestLitffi3callz()
}

func TestFCBM2(t *testing.T) {
	BMLitffi2callz()
}
func TestFCBM3(t *testing.T) {
	BMLitffi3callz()
}
