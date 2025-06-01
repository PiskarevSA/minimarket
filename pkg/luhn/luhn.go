package luhn

// isValid checks if an integer order number is valid using the Luhn algorithm
func IsValid(orderNumber int) bool {
	if orderNumber < 10 {
		return false // Needs at least one digit + a check digit
	}

	checkDigit := orderNumber % 10
	number := orderNumber / 10

	// Extract digits in reverse (from right to left, excluding the check digit)
	var digits []int
	for number > 0 {
		digits = append(digits, number%10)
		number /= 10
	}

	sum := 0
	for i := range digits {
		d := digits[i]
		if i%2 == 0 {
			d *= 2
			if d > 9 {
				d -= 9
			}
		}
		sum += d
	}

	calculatedCheck := (10 - (sum % 10)) % 10
	return checkDigit == calculatedCheck
}
