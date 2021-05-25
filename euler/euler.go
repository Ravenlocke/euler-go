package euler

import (
	"errors"
	"fmt"
)

func is_palindrome(str string) bool {
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		if str[i] != str[j] {
			return false
		}
	}
	return true
}

func MaxI(a, b int) int {
	if a < b {
		return b
	}
	return a
}

type PrimeFactors struct {
	Factors map[int]int
}

func (p PrimeFactors) LowestCommonMultiple(other PrimeFactors) PrimeFactors {
	factors := make(map[int]int)

	for k, v := range p.Factors {
		factors[k] = v
	}
	for k, v := range other.Factors {
		factors[k] = MaxI(v, factors[k])
	}

	return PrimeFactors{factors}
}

func NewPrimeFactors(n int) PrimeFactors {
	factors := make(map[int]int)

	for n%2 == 0 {
		factors[2] += 1
		n /= 2
	}

	for counter := 3; true; counter += 2 {
		for n%counter == 0 {
			factors[counter] += 1
			n /= counter
		}
		if n == 1 {
			break
		}
	}

	return PrimeFactors{factors}

}

func Euler01() int {
	total := 0
	for i := 1; i < 1_000; i++ {
		if i%3 == 0 || i%5 == 0 {
			total += i
		}
	}

	return total
}

func Euler02() int {
	a, b := 1, 2
	total := 0

	for a < 4_000_000 {
		if a%2 == 0 {
			total += a
		}
		a, b = b, a+b
	}
	return total

}

func Euler03() (int, error) {
	n := 600_851_475_143

	// Special case for two.
	for n%2 == 0 {
		n /= 2
	}
	if n == 1 {
		return 2, nil
	}

	// Start at three, increment by two.
	for counter := 3; true; counter += 2 {
		for n%counter == 0 {
			n /= counter
		}
		if n == 1 {
			return counter, nil
		}
	}

	return -1, errors.New("this should never be reached")
}

func Euler04() int {
	highest_palindromic_product := 0

	for i := 100; i < 1_000; i++ {
		for j := i; j < 1_000; j++ {
			product := i * j
			if product < highest_palindromic_product {
				continue
			}
			as_string := fmt.Sprint(product)
			if is_palindrome(as_string) {
				highest_palindromic_product = product
			}
		}
	}
	return highest_palindromic_product
}

func Euler05() int {
	current := NewPrimeFactors(1)

	for i := 2; i < 21; i++ {
		next := NewPrimeFactors(i)
		current = current.LowestCommonMultiple(next)
	}

	result := 1
	for k, v := range current.Factors {
		for i := 0; i < v; i++ {
			result *= k
		}
	}
	return result
}
