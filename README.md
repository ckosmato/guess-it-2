
# Data Analysis with Linear Regression and Outlier Detection

This Go program analyzes a series of data points, detecting extreme outliers, performing linear regression, and calculating ranges for predicted values. The program takes input data via standard input, processes it, and outputs calculated ranges (lower and upper bounds) based on the provided data.

## Features

- **Outlier Detection:** Detects extreme outliers based on mean and standard deviation. Outliers are automatically removed from the dataset.
- **Linear Regression:** Predicts the next value in the sequence based on linear regression.
- **Range Calculation:** Calculates the predicted range around the next value using a margin, and prints the lower and upper bounds.
- **Pearson Correlation Coefficient:** Computes the correlation between the x and y values to assess the relationship between them.

## Requirements

- Go 1.18 or later

## Usage

1. Run the program:
   ```bash
   go run main.go
   ```

2. The program expects input through standard input. For each line, enter an integer value representing the y-value, and the program will automatically assign x-values based on the order.

3. The program will output the calculated range in the form:
   ```
   <lower_bound> <upper_bound>
   ```

4. Input can be provided via terminal, and the program will continue to process and output the results for each new input line.

### Example Input:
```
5
8
12
3
15
```

### Example Output:
```
5 10
6 14
7 16
6 14
8 15
```

## Functions

### `GetRange(xValues, yValues []int) (int, int)`
Calculates the lower and upper range based on the provided `xValues` and `yValues`.

### `CalculateLinearRegression(xValues, yValues []int) float64`
Predicts the next `y` value using linear regression based on the `xValues` and `yValues`.

### `CalculatePCC(xValues, yValues []int) float64`
Calculates the Pearson Correlation Coefficient (PCC) to measure the relationship between `xValues` and `yValues`.

### `CalculateAverage(numList []int) float64`
Calculates the average of a list of integers.

### `GetMedian(numList []int) int`
Calculates the median of a list of integers.

### `IsExtreme(num int, numList []int) bool`
Checks if a given number is an extreme outlier based on the mean and standard deviation.

### `RemoveOutlier(xValues, yValues []int, outlier int) ([]int, []int)`
Removes an outlier from both `xValues` and `yValues` lists.

### `GetStandardDeviation(numList []int) float64`
Calculates the standard deviation of a list of integers.

### `GetVariance(numList []int) float64`
Calculates the variance of a list of integers.

## Author

Christos Kosmatos
