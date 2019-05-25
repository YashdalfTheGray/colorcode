// Package colorcode provides color space definitions
// and methods to convert between them.
package colorcode

import "fmt"

// HexCode represents an HTML color string. Implements
// fmt.Stringer
type HexCode string

// NewHexCode returns a new HexCode built out of a string
func NewHexCode(code string) HexCode {
	return HexCode(code)
}

// String returns the string representation of the HexCode
func (cs HexCode) String() string {
	return string(cs)
}

// ToRGB converts a HexCode string color into its RGB equivalent
func (cs HexCode) ToRGB() (RGB, error) {
	var r, g, b uint8

	n, err := fmt.Sscanf(string(cs), "#%2x%2x%2x", &r, &g, &b)
	if err != nil || n != 3 {
		return RGB{}, err
	}

	return RGB{r, g, b}, nil
}

// ToHSL converts a HexCode string color to its HSL equivalent
func (cs HexCode) ToHSL() (HSL, error) {
	rgb, err := cs.ToRGB()
	if err != nil {
		return HSL{}, err
	}

	return rgb.ToHSL(), nil
}

// ToHSV converts a HexCode string color to its HSV equivalent
func (cs HexCode) ToHSV() (HSV, error) {
	rgb, err := cs.ToRGB()
	if err != nil {
		return HSV{}, err
	}

	return rgb.ToHSV(), nil
}
