package utils

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
