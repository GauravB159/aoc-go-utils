package aocutils

func IsIntPrime(number int) bool {
	for i := 2; i*i <= number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}
