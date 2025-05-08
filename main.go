package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func convertUnit(args []string) string {
	fs := flag.NewFlagSet("unitconverter", flag.ContinueOnError)

	var buf bytes.Buffer
	fs.SetOutput(&buf)

	valuePtr := fs.String("value", "", "The value to convert")
	fromUnitPtr := fs.String("from", "", "The original unit")
	toUnitPtr := fs.String("to", "", "The target unit")

	// Parse the arguments
	err := fs.Parse(args)
	if err != nil {
		return fmt.Sprintf("error: %v\n", err)
	}

	valueStr := *valuePtr
	fromUnit := *fromUnitPtr
	toUnit := *toUnitPtr

	if valueStr == "" || fromUnit == "" || toUnit == "" {
		return "Usage: go run main.go -value <number> -from <unit> -to <unit>)\n"
	}

	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		return "invalid input provided\n"
	}

	var result float64
	var errConv error

	switch fromUnit {
	case "C", "celsius":
		switch toUnit {
		case "F", "fahrenheit":
			result = celsiusToFahrenheit(value)
		case "K", "kelvin":
			result = celsiusToKelvin(value)
		default:
			errConv = fmt.Errorf("invalid conversion from %s to %s", fromUnit, toUnit)
		}
	case "F", "fahrenheit":
		switch toUnit {
		case "C", "celsius":
			result = fahrenheitToCelsius(value)
		case "K", "kelvin":
			result = fahrenheitToKelvin(value)
		default:
			errConv = fmt.Errorf("invalid conversion from %s to %s", fromUnit, toUnit)
		}
	case "K", "kelvin":
		switch toUnit {
		case "C", "celsius":
			result = kelvinToCelsius(value)
		case "F", "fahrenheit":
			result = kelvinToFahrenheit(value)
		default:
			errConv = fmt.Errorf("invalid conversion from %s to %s", fromUnit, toUnit)
		}
	case "km", "kilometers":
		switch toUnit {
		case "mi", "miles":
			result = kilometersToMiles(value)
		default:
			errConv = fmt.Errorf("invalid conversion from %s to %s", fromUnit, toUnit)
		}
	case "mi", "miles":
		switch toUnit {
		case "km", "kilometers":
			result = milesToKilometers(value)
		default:
			errConv = fmt.Errorf("invalid conversion from %s to %s", fromUnit, toUnit)
		}

	default:
		errConv = fmt.Errorf("invalid unit: %s", fromUnit)
	}

	if errConv != nil {
		return errConv.Error()
	}
	return fmt.Sprintf("%.2f %s is equal to %.2f %s\n", value, fromUnit, result, toUnit)
}

func main() {
	output := convertUnit(os.Args[1:])
	fmt.Print(output)
}

func celsiusToKelvin(c float64) float64 {
	return c + 273.15
}

func kelvinToCelsius(k float64) float64 {
	return k - 273.15
}

func fahrenheitToKelvin(f float64) float64 {
	return (f-32)*5/9 + 273.15
}

func kelvinToFahrenheit(k float64) float64 {
	return (k-273.15)*9/5 + 32
}

func celsiusToFahrenheit(c float64) float64 {
	return kelvinToFahrenheit(celsiusToKelvin(c))
}

func fahrenheitToCelsius(f float64) float64 {
	return kelvinToCelsius(fahrenheitToKelvin(f))
}

func kilometersToMiles(km float64) float64 {
	return km * 0.621371
}

func milesToKilometers(mi float64) float64 {
	return mi / 0.621371
}
