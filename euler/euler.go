package euler

import "errors"

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

	primes := make(map[int]struct{})

	// Special case for two.
	for n%2 == 0 {
		if _, ok := primes[n]; !ok {
			primes[n] = struct{}{}
		}
		n /= 2
	}
	if n == 1 {
		return 2, nil
	}

	// Start at three, increment by two.
	for counter := 3; true; counter++ {
		for n%counter == 0 {
			if _, ok := primes[n]; !ok {
				primes[n] = struct{}{}
			}
			n /= counter
		}

		if n == 1 {
			return counter, nil
		}
	}

	return -1, errors.New("this should never be reached")
}
