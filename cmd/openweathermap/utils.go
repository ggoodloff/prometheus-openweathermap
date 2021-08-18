package main

import "math"

func CalculateWetBulbTemperature(t, rh float64) float64 {
	// Using formula from here: https://perryweather.com/2020/04/01/what-is-wbgt-and-how-do-you-calculate-it/
	return t*math.Atan(0.151977*math.Sqrt(rh+8.313659)) + math.Atan(t+rh) - math.Atan(rh-1.676331) + 0.00391838*math.Pow(rh, 1.5)*math.Atan(0.023101*rh) - 4.686035
}
