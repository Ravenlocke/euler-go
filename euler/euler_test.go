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

func TestEuler04(t *testing.T) {
	result := Euler04()
	expected := 906_609

	if result != expected {
		t.Errorf("Euler04() = %d; want %d", result, expected)
	}
}

func TestNewPrimeFactors(t *testing.T) {
	expected := map[int]int{2: 2}
	result := NewPrimeFactors(4)

	for k, v := range result.Factors {
		expected_v := expected[k]
		if v != expected_v {
			t.Errorf("NewPrimeFactors(4).Factors = %v; want %v", result.Factors, expected)
		}

	}
}

func TestEuler05(t *testing.T) {
	expected := 232_792_560
	result := Euler05()

	if result != expected {
		t.Errorf("Euler05() = %d; want %d", result, expected)
	}
}
