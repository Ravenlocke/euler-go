package euler

import (
	"errors"
	"fmt"
	"strings"
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

func AbsI(a int) int {
	if a < 1 {
		return -a
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

func Euler06() int {
	sum_square := 0
	square_sum := 0

	for i := 1; i <= 100; i++ {
		sum_square += i * i
		square_sum += i
	}

	square_sum *= square_sum

	result := AbsI(square_sum - sum_square)
	return result
}

func Euler07() int {
	primes := []int{2}

outer:
	for counter := 3; true; counter++ {
		for _, prime := range primes {
			if counter%prime == 0 {
				continue outer
			}

		}
		primes = append(primes, counter)
		if len(primes) == 10_001 {
			break outer
		}
	}

	return primes[10_000]

}

func Euler08() int {
	number := `73167176531330624919225119674426574742355349194934
	96983520312774506326239578318016984801869478851843
	85861560789112949495459501737958331952853208805511
	12540698747158523863050715693290963295227443043557
	66896648950445244523161731856403098711121722383113
	62229893423380308135336276614282806444486645238749
	30358907296290491560440772390713810515859307960866
	70172427121883998797908792274921901699720888093776
	65727333001053367881220235421809751254540594752243
	52584907711670556013604839586446706324415722155397
	53697817977846174064955149290862569321978468622482
	83972241375657056057490261407972968652414535100474
	82166370484403199890008895243450658541227588666881
	16427171479924442928230863465674813919123162824586
	17866458359124566529476545682848912883142607690042
	24219022671055626321111109370544217506941658960408
	07198403850962455444362981230987879927244284909188
	84580156166097919133875499200524063689912560717606
	05886116467109405077541002256983155200055935729725
	71636269561882670428252483600823257530420752963450
	`

	// Extract characters from the `number` string that are numbers and convert
	// to int.
	digits := `1234567890`
	number_list := []int{}

	for _, char := range number {
		if strings.ContainsRune(digits, char) {
			number_list = append(number_list, int(char)-48)
		}
	}

	window_size := 13
	n_windows := len(number_list) - window_size + 1
	highest_product := 0

	for i := 0; i < n_windows; i++ {
		slice := number_list[i : i+window_size]
		product := 1
		for _, i := range slice {
			product *= i
		}
		if product > highest_product {
			highest_product = product
		}
	}

	return highest_product
}

func Euler09() (int, error) {
	for a := 1; a < 1_001; a++ {
		for b := a + 1; b < (1_001 - a); b++ {
			c := 1000 - (a + b)
			if a*a+b*b == c*c {
				return a * b * c, nil
			}
		}
	}

	return -1, errors.New("this should never be reached")
}
