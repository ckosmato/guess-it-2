package helpers

// Checks if a number is an extreme outlier based on mean and standard deviation
func IsExtreme(num int, numList []int) bool {
	if len(numList) < 2 {
		return false // Not enough data to detect outliers
	}

	mean := CalculateAverage(numList)
	stdDev := GetStandardDeviation(numList)

	// Calculate lower and upper thresholds based on 3 standard deviations from the mean
	lowerThreshold := mean - float64(3)*stdDev
	upperThreshold := mean + float64(3)*stdDev

	// Check if the number falls outside the thresholds
	return float64(num) < lowerThreshold || float64(num) > upperThreshold
}

// Removes a specific outlier from the x and y value lists
func RemoveOutlier(xValues, yValues []int, outlier int) ([]int, []int) {
	for i, n := range yValues {
		if n == outlier {
			// Remove the outlier by excluding it from both lists
			return append(xValues[:i], xValues[i+1:]...), append(yValues[:i], yValues[i+1:]...)
		}
	}
	return xValues, yValues
}
