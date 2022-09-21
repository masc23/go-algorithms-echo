package mathematics

import (
	"fmt"
	"math"
)

func evaluateEquation(input string) string {
	return input
}

func runRegulaFalsi(input string) string {
	lowerBound := float64(0)
	upperBound := float64(1)
	epsilon := 5e-15
	iterations := 100

	result := regulaFalsi(sampleEquation, lowerBound, upperBound, epsilon, iterations)

	return fmt.Sprintf("%.15f", result)
}

func regulaFalsi(f func(float64) float64, lowerBoundX float64, upperBoundX float64, epsilon float64, numberOfIterations int) float64 {

	lowerBoundY := f(lowerBoundX)
	upperBoundY := f(upperBoundX)

	side := 0

	currentX := float64(0)

	for i := 0; i < numberOfIterations; i++ {
		currentX = (lowerBoundY*upperBoundX - upperBoundY*lowerBoundX) / (lowerBoundY - upperBoundY)

		if math.Abs(upperBoundX-lowerBoundX) < epsilon*math.Abs(upperBoundX+lowerBoundX) {
			break
		}

		currentY := f(currentX)

		if currentY*upperBoundY > 0 {
			upperBoundX = currentX
			upperBoundY = currentY

			if side == -1 {
				lowerBoundY /= 2
			}

			side = -1

		} else if currentY*lowerBoundY > 0 {
			lowerBoundX = currentX
			lowerBoundY = currentY

			if side == 1 {
				upperBoundY /= 2
			}

			side = 1

		} else {
			break
		}
	}

	return currentX
}

func sampleEquation(x float64) float64 {
	return math.Cos(x) - x*x*x
}
