package main

import (
	"math"
	"testing"
)

const float64EqualityThreshold = 1e-3 // Suitably small tolerance

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) <= float64EqualityThreshold
}

func TestCelsiusToFahrenheit(t *testing.T) {
	tests := []struct {
		input    float64
		expected float64
	}{
		{0, 32},
		{100, 212},
		{-40, -40},
		{25, 77},
	}

	for _, test := range tests {
		actual := celsiusToFahrenheit(test.input)
		if !almostEqual(actual, test.expected) {
			t.Errorf("celsiusToFahrenheit(%f) = %f, expected %f", test.input, actual, test.expected)
		}
	}
}

func TestFahrenheitToCelsius(t *testing.T) {
	tests := []struct {
		input    float64
		expected float64
	}{
		{32, 0},
		{212, 100},
		{-40, -40},
		{77, 25},
	}

	for _, test := range tests {
		actual := fahrenheitToCelsius(test.input)
		if !almostEqual(actual, test.expected) {
			t.Errorf("fahrenheitToCelsius(%f) = %f, expected %f", test.input, actual, test.expected)
		}
	}
}

func TestCelsiusToKelvin(t *testing.T) {
	tests := []struct {
		input    float64
		expected float64
	}{
		{0, 273.15},
		{100, 373.15},
		{-40, 233.15},
		{25, 298.15},
	}

	for _, test := range tests {
		actual := celsiusToKelvin(test.input)
		if !almostEqual(actual, test.expected) {
			t.Errorf("celsiusToKelvin(%f) = %f, expected %f", test.input, actual, test.expected)
		}
	}
}

func TestKelvinToCelsius(t *testing.T) {
	tests := []struct {
		input    float64
		expected float64
	}{
		{273.15, 0},
		{373.15, 100},
		{233.15, -40},
		{298.15, 25},
	}

	for _, test := range tests {
		actual := kelvinToCelsius(test.input)
		if !almostEqual(actual, test.expected) {
			t.Errorf("kelvinToCelsius(%f) = %f, expected %f", test.input, actual, test.expected)
		}
	}
}

func TestFahrenheitToKelvin(t *testing.T) {
	tests := []struct {
		input    float64
		expected float64
	}{
		{32, 273.15},
		{212, 373.15},
		{-40, 233.15},
		{77, 298.15},
	}

	for _, test := range tests {
		actual := fahrenheitToKelvin(test.input)
		if !almostEqual(actual, test.expected) {
			t.Errorf("fahrenheitToKelvin(%f) = %f, expected %f", test.input, actual, test.expected)
		}
	}
}

func TestKelvinToFahrenheit(t *testing.T) {
	tests := []struct {
		input    float64
		expected float64
	}{
		{273.15, 32},
		{373.15, 212},
		{233.15, -40},
		{298.15, 77},
	}

	for _, test := range tests {
		actual := kelvinToFahrenheit(test.input)
		if !almostEqual(actual, test.expected) {
			t.Errorf("kelvinToFahrenheit(%f) = %f, expected %f", test.input, actual, test.expected)
		}
	}
}

func TestKilometersToMiles(t *testing.T) {
	tests := []struct {
		input    float64
		expected float64
	}{
		{0, 0},
		{1, 0.621371},
		{10, 6.21371},
	}

	for _, test := range tests {
		actual := kilometersToMiles(test.input)
		if !almostEqual(actual, test.expected) {
			t.Errorf("kilometersToMiles(%f) = %f, expected %f", test.input, actual, test.expected)
		}
	}
}

func TestMilesToKilometers(t *testing.T) {
	tests := []struct {
		input    float64
		expected float64
	}{
		{0, 0},
		{1, 1.609344},
		{10, 16.09344},
	}

	for _, test := range tests {
		actual := milesToKilometers(test.input)
		if !almostEqual(actual, test.expected) {
			t.Errorf("milesToKilometers(%f) = %f, expected %f", test.input, actual, test.expected)
		}
	}
}
func TestConvertUnit(t *testing.T) {
	tests := []struct {
		args     []string
		expected string
	}{
		{[]string{"-value", "25", "-from", "celsius", "-to", "fahrenheit"}, "25.00 celsius is equal to 77.00 fahrenheit\n"},
		{[]string{"-value", "10", "-from", "km", "-to", "miles"}, "10.00 km is equal to 6.21 miles\n"},
		{[]string{"-value", "invalid", "-from", "celsius", "-to", "fahrenheit"}, "invalid input provided\n"},
		{[]string{"-value", "25", "-from", "celsius"}, "Usage: go run main.go -value <number> -from <unit> -to <unit>)\n"},
		{[]string{"-value", "25", "-from", "celsius", "-to", "kelvin"}, "25.00 celsius is equal to 298.15 kelvin\n"},
		{[]string{"-value", "25", "-from", "invalid", "-to", "fahrenheit"}, "invalid unit: invalid"},
	}

	for _, test := range tests {
		output := convertUnit(test.args)
		if output != test.expected {
			t.Errorf("For args %v, expected output:\n%q\nGot:\n%q", test.args, test.expected, output)
		}
	}
}
