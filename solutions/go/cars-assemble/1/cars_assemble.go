package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {

	productionRateFloat := float64(productionRate)
	var carConstructPerHour float64 = productionRateFloat * ((successRate) / 100)

	return carConstructPerHour
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {

	productionRateFloat := float64(productionRate)
	var carConstructPerMinute float64 = (productionRateFloat * ((successRate) / 100)) / 60

	return int(carConstructPerMinute)

}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {

	var priceUnitary uint = 10000
	var priceBundle10 uint = 95000
	carsCountUint := uint(carsCount)

	var cost uint = priceBundle10*(carsCountUint/10) + priceUnitary*(carsCountUint%10)

	return cost

}
