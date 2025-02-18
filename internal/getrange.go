package internal

import (
	"bufio"
	"fmt"
	"guessit/internal/helpers"
	"math"
	"os"
	"strconv"
)

// Calculates the lower and upper range based on the input data
func GetRange() {
	var yValue int
	var xValues, yValues []int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		yValue, _ = strconv.Atoi(line)

		xValues = append(xValues, len(xValues))
		yValues = append(yValues, yValue)

		// Check if the current yValue is an extreme outlier
		if helpers.IsExtreme(yValue, yValues) {

			xValues, yValues = helpers.RemoveOutlier(xValues, yValues, yValue)
		}

		pcc := helpers.CalculatePCC(xValues, yValues)

		median := helpers.GetMedian(yValues)
		standardDeviation := helpers.GetStandardDeviation(yValues)

		// Predict the next y-value using linear regression
		yNextValue := helpers.CalculateLinearRegression(xValues, yValues)
		margin := 18 // Define a margin around the predicted value

		// Return calculated range based on the data
		if len(yValues) == 0 {
			fmt.Printf("0 0\n")
		}
		if len(yValues) == 1 {
			fmt.Printf("%d %d\n", yValues[0], yValues[0])
			continue
		}
		// If the PCC is 0 or NaN, return the range based on median and standard deviation
		if pcc == 0 || math.IsNaN(pcc) {
			fmt.Printf("%d %d\n", median-int(standardDeviation), median+int(standardDeviation))
		}

		// Return the range based on linear regression and margin

		fmt.Printf("%d %d\n", int(yNextValue)-int(margin), int(yNextValue)+int(margin))
	}
}
