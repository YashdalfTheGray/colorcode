package colorcode

import (
	"errors"
	"fmt"
	"math"
)

// HSV represents a color in the HSV color space. This also
// implements the fmt.Stringer interface
type HSV struct {
	H, S, V float64
}

// NewHSV takes values for H, S and V and returns an HSV type
func NewHSV(h, s, v float64) (HSV, error) {
	if h < 0 || h > 360 || s < 0 || s > 100 || v < 0 || v > 100 {
		return HSV{}, errors.New("invalid argument error")
	}
	return HSV{h, s, v}, nil
}

// String returns a string representation of the HSV object
func (c HSV) String() string {
	return fmt.Sprintf("hsv(%.0f, %.0f%%, %.0f%%)", c.H, c.S, c.V)
}

// ToRGB converts the HSV color to the RGB representation
func (c HSV) ToRGB() RGB {
	h := c.H / 60.0
	s := c.S / 100.0
	v := c.V / 100.0
	hi := int(math.Floor(h)) % 6

	f := h - math.Floor(h)
	p := 255 * v * (1 - s)
	q := 255 * v * (1 - (s * f))
	t := 255 * v * (1 - (s * (1 - f)))

	v *= 255

	switch hi {
	case 0:
		return RGB{uint8(v), uint8(t), uint8(p)}
	case 1:
		return RGB{uint8(q), uint8(v), uint8(p)}
	case 2:
		return RGB{uint8(p), uint8(v), uint8(t)}
	case 3:
		return RGB{uint8(p), uint8(q), uint8(v)}
	case 4:
		return RGB{uint8(t), uint8(p), uint8(v)}
	case 5:
		return RGB{uint8(v), uint8(p), uint8(q)}
	}
	panic(errors.New("a number modulo 6 should never be outside of range [0, 5]"))
}

// ToHexCode converts the HSV color into the HexCode representation
func (c HSV) ToHexCode() HexCode {
	return c.ToRGB().ToHexCode()
}

// ToHSL converts the HSV color into the HSL representation
func (c HSV) ToHSL() HSL {
	h := c.H
	s := c.S / 100.0
	v := c.V / 100.0
	vmin := math.Max(v, 0.01)
	l := (2 - s) * v
	lmin := (2 - s) * vmin
	sl := s * vmin

	if lmin <= l {
		sl /= lmin
	} else {
		sl /= 2 - lmin
	}

	l /= 2

	return HSL{h, sl * 100.0, l * 100.0}
}
