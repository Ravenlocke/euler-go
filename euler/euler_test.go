package euler

import "testing"

func TestEuler01(t *testing.T) {
	result := Euler01()
	expected := 233_168

	if result != expected {
		t.Errorf("Euler01() = %d; want %d", result, expected)
	}
}

func TestEuler02(t *testing.T) {
	result := Euler02()
	expected := 4_613_732

	if result != expected {
		t.Errorf("Euler02() = %d; want %d", result, expected)
	}

}

func TestEuler03(t *testing.T) {
	result, err := Euler03()
	if err != nil {
		t.Errorf("Euler03() raised unexpected error: %v", err)
	}
	expected := 6_857

	if result != expected {
		t.Errorf("Euler03() = %d; want %d", result, expected)
	}
}
