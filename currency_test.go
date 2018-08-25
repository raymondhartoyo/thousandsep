package thousandsep

import (
	"testing"
)

func TestThousandSeparatorFloat(t *testing.T) {
	config := Config{
		Separator:        ".",
		DecimalSeparator: ",",
	}
	// should produce correct string representation if number consist of zeroes
	output := ThousandSeparatorFloat(1000000, 2, config)
	if output != "1.000.000,00" {
		t.Error("should produce correct string representation if number consist of zeroes")
	}
	output = ThousandSeparatorFloat(1020030, 2, config)
	if output != "1.020.030,00" {
		t.Error("should produce correct string representation if number consist of zeroes")
	}

	// should produce correct string representation
	output = ThousandSeparatorFloat(1234567, 2, config)
	if output != "1.234.567,00" {
		t.Error("should produce correct string representation")
	}
	// should produce correct string representation without decimal
	output = ThousandSeparatorFloat(1234567.89, 0, config)
	if output != "1.234.567" {
		t.Error("should produce correct string representation without decimal")
	}
	// should produce correct string representation with decimals
	output = ThousandSeparatorFloat(1234567.89, 1, config)
	if output != "1.234.567,8" {
		t.Error("should produce correct string representation with decimals")
	}

	// should produce correct string representation (negative number)
	output = ThousandSeparatorFloat(-1234567, 2, config)
	if output != "-1.234.567,00" {
		t.Error("should produce correct string representation (negative number)")
	}
	// should produce correct string representation without decimal (negative number)
	output = ThousandSeparatorFloat(-1234567.89, 0, config)
	if output != "-1.234.567" {
		t.Error("should produce correct string representation without decimal (negative number)")
	}
	// should produce correct string representation with decimals (negative number)
	output = ThousandSeparatorFloat(-1234567.89, 1, config)
	if output != "-1.234.567,8" {
		t.Error("should produce correct string representation with decimals (negative number)")
	}
}

func TestThousandSeparatorInt(t *testing.T) {
	config := Config{
		Separator: ".",
	}
	// should produce correct string representation if number consist of zeroes
	output := ThousandSeparatorInt(1000000, config)
	if output != "1.000.000" {
		t.Error("should produce correct string representation if number consist of zeroes")
	}
	// should produce correct string representation
	output = ThousandSeparatorInt(1234567, config)
	if output != "1.234.567" {
		t.Error("should produce correct string representation")
	}
	// should produce correct string representation (negative number)
	output = ThousandSeparatorInt(-1234567, config)
	if output != "-1.234.567" {
		t.Error("should produce correct string representation (negative number)")
	}
}
