// Package thousandsep is a simple package to format float64 and int64 number into strings with thoudsand separated (currency) format
package thousandsep

import (
	"math"
	"strconv"
)

// Config is needed for the separator and decimal separator settings
type Config struct {
	Separator        string // to be used as thousand separator
	DecimalSeparator string // to be used as decimal separator
}

// ThousandSeparatorFloat will return the float number with thousand separated (currency) format using specified config
func ThousandSeparatorFloat(num float64, decimalCount int64, config Config) string {
	sign := ""
	if num < 0 {
		sign = "-"
		num = math.Abs(num)
	}

	numParts := []string{}
	var numInInteger int64

	if decimalCount != 0 {
		numInInteger = int64(num * math.Pow(10, float64(decimalCount)))
		numParts = append(numParts, strconv.FormatInt(numInInteger%int64(math.Pow(10, float64(decimalCount))), 10))
		numInInteger = numInInteger / int64(math.Pow(10, float64(decimalCount)))
	} else {
		numInInteger = int64(num)
		numParts = append(numParts, "0")
	}

	for numInInteger >= 1000 {
		onePart := ""
		mod := numInInteger % 1000
		if mod < 10 {
			onePart = "00" + strconv.FormatInt(mod, 10)
		} else if mod < 100 {
			onePart = "0" + strconv.FormatInt(mod, 10)
		} else {
			onePart = strconv.FormatInt(mod, 10)
		}
		numParts = append(numParts, onePart)
		numInInteger = numInInteger / 1000
	}

	numParts = append(numParts, strconv.FormatInt(numInInteger%1000, 10))

	// handle decimals
	decimals := ""
	if decimalCount > 0 {
		decimals = config.DecimalSeparator + decimals
		for i := int64(len([]rune(numParts[0]))); i < decimalCount; i++ {
			decimals = decimals + "0"
		}
		decimals = decimals + numParts[0]
	}

	numberString := ""
	for i := 1; i < len(numParts); i++ {
		numberString = numParts[i] + numberString
		if i < len(numParts)-1 {
			numberString = config.Separator + numberString
		}
	}

	return sign + numberString + decimals
}

// ThousandSeparatorInt will return the int number with thousand separated (currency) format using specified config
func ThousandSeparatorInt(num int64, config Config) string {
	return ThousandSeparatorFloat(float64(num), 0, config)
}
