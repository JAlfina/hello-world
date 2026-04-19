package commission

func CheckAmount(amount int) bool {
	return amount >= 500 && amount <= 15000000
}

func CalculateCommission(amount int, isAlif bool) int {
	if isAlif {
		return 0
	}
	return amount * 29 / 10000 // 0.29%
}

func CalculateTotal(amount int, commission int) int {
	return amount + commission
}
