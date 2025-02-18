package helpers

import (
	"math"
	"sort"
)

// Calculates the next y-value using linear regression
func CalculateLinearRegression(xValues, yValues []int) float64 {
	meanX := CalculateAverage(xValues)
	meanY := CalculateAverage(yValues)

	var sumXY, sumXSquared float64

	// Calculate the slope (m) and intercept (b) for linear regression
	for i := 0; i < len(xValues); i++ {
		xValue := float64(xValues[i]) - meanX
		yValue := float64(yValues[i]) - meanY

		sumXY += xValue * yValue
		sumXSquared += xValue * xValue
	}
	// Calculate slope and intercept
	slope := sumXY / sumXSquared
	intercept := meanY - (slope * meanX)

	// Predict the next y-value using the regression equation
	nextX := float64(len(xValues)) // Next x is assumed to be one step forward
	yNextValue := slope*nextX + intercept

	return yNextValue
}

// CalculatePCC calculates the Pearson Correlation Coefficient between x and y values
func CalculatePCC(xValues, yValues []int) float64 {
	if len(xValues) < 2 || len(yValues) < 2 {
		return 0 // Not enough data to calculate correlation
	}
	meanX := CalculateAverage(xValues)
	meanY := CalculateAverage(yValues)

	var sumXY, sumXSquared, sumYSquared float64

	// Calculate the necessary sums for PCC
	for i := 0; i < len(xValues); i++ {
		xValue := float64(xValues[i]) - meanX
		yValue := float64(yValues[i]) - meanY

		sumXY += xValue * yValue
		sumXSquared += xValue * xValue
		sumYSquared += yValue * yValue
	}
	if sumXSquared == 0 || sumYSquared == 0 {
		return 0 // No variation in x or y, correlation is undefined
	}
	// Return the Pearson Correlation Coefficient
	return sumXY / (math.Sqrt(sumXSquared * sumYSquared))
}

// Calculates the average of a list of integers
func CalculateAverage(numList []int) float64 {
	if len(numList) == 0 {
		return 0 // Return 0 if the list is empty
	}

	// Calculate sum and average
	sum := 0
	for _, n := range numList {
		sum += n
	}
	return float64(sum) / float64(len(numList))
}

// Calculates the median of a list of integers
func GetMedian(numList []int) int {
	tempList := append([]int{}, numList...)
	sort.Ints(tempList)
	// If the list length is even, return the average of the two middle elements
	if len(tempList)%2 == 0 {
		return int(math.Round((float64(tempList[len(tempList)/2-1]) + float64(tempList[len(tempList)/2])) / 2))
	}
	// Return the middle element for odd-length list
	return tempList[len(tempList)/2]
}

// Calculates the standard deviation of a list of integers
func GetStandardDeviation(numList []int) float64 {
	variance := GetVariance(numList) // Calculate variance
	return math.Sqrt(variance)       // Return the square root of variance
}

// Calculates the variance of a list of integers
func GetVariance(numList []int) float64 {
	average := CalculateAverage(numList)
	var varianceSum float64
	// Sum of squared differences from the mean
	for _, n := range numList {
		diff := float64(n) - average
		varianceSum += diff * diff
	}
	// Return the variance (average of squared differences)
	return varianceSum / float64(len(numList))
}
