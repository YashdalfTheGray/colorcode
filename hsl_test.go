package colorcode_test

import (
	"testing"

	"github.com/YashdalfTheGray/colorcode"
)

func TestNewHSL(t *testing.T) {
	testCases := []struct {
		desc    string
		h, s, l float64
		out     string
		err     bool
	}{
		{
			desc: "builds a new HSL out of given numbers",
			h:    250.4,
			s:    65,
			l:    35,
			out:  colorcode.HSL{250.4, 65, 35}.String(),
			err:  false,
		},
		{
			desc: "errors when numbers are out of bounds",
			h:    378,
			s:    65,
			l:    35,
			out:  colorcode.HSL{0, 0, 0}.String(),
			err:  true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			color, err := colorcode.NewHSL(tC.h, tC.s, tC.l)
			if !tC.err && err != nil {
				t.Errorf("Expected to not have error but got error %s", err.Error())
			} else if tC.err && err == nil {
				t.Errorf("Expected error but got nil")
			} else if color.String() != tC.out {
				t.Errorf("Expected %s but got %s", tC.out, color.String())
			}
		})
	}
}

func TestHSLString(t *testing.T) {
	testCases := []struct {
		desc string
		in   colorcode.HSL
		out  string
	}{
		{
			desc: "builds a string out of an HSL color",
			in:   colorcode.HSL{200, 90, 50},
			out:  "hsl(200, 90%, 50%)",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if tC.in.String() != tC.out {
				t.Errorf("Expected %s but got %s", tC.out, tC.in.String())
			}
		})
	}
}

func TestHSLToRGB(t *testing.T) {
	testCases := []struct {
		desc string
		in   colorcode.HSL
		out  string
	}{
		{
			desc: "converts a non-gray HSL color to RGB",
			in:   colorcode.HSL{200, 90, 50},
			out:  "rgb(12, 165, 242)",
		},
		{
			desc: "converts a gray HSL color to RGB",
			in:   colorcode.HSL{0, 0, 50},
			out:  "rgb(127, 127, 127)",
		},
		{
			desc: "converts a low luminosity HSL color to RGB",
			in:   colorcode.HSL{256, 100, 15},
			out:  "rgb(20, 0, 76)",
		},
		{
			desc: "converts a green HSL color to RGB",
			in:   colorcode.HSL{90, 100, 15},
			out:  "rgb(38, 76, 0)",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := tC.in.ToRGB()
			if result.String() != tC.out {
				t.Errorf("Expected %s but got %s", tC.out, result.String())
			}
		})
	}
}

func TestHSLToHexCode(t *testing.T) {
	testCases := []struct {
		desc string
		in   colorcode.HSL
		out  string
	}{
		{
			desc: "converts a non-gray HSL color to RGB",
			in:   colorcode.HSL{200, 90, 50},
			out:  "#0ca5f2",
		},
		{
			desc: "converts a gray HSL color to RGB",
			in:   colorcode.HSL{0, 0, 50},
			out:  "#7f7f7f",
		},
		{
			desc: "converts a low luminosity HSL color to RGB",
			in:   colorcode.HSL{256, 100, 15},
			out:  "#14004c",
		},
		{
			desc: "converts a green HSL color to RGB",
			in:   colorcode.HSL{90, 100, 15},
			out:  "#264c00",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := tC.in.ToHexCode()
			if result.String() != tC.out {
				t.Errorf("Expected %s but got %s", tC.out, result.String())
			}
		})
	}
}

func TestHSLToHSV(t *testing.T) {
	testCases := []struct {
		desc string
		in   colorcode.HSL
		out  string
	}{
		{
			desc: "converts an HSL color into HSV",
			in:   colorcode.HSL{197, 84, 63},
			out:  "hsv(197, 66%, 94%)",
		},
		{
			desc: "converts a dull HSL color into HSV",
			in:   colorcode.HSL{197, 84, 30},
			out:  "hsv(197, 91%, 55%)",
		},
		{
			desc: "converts a dull HSL color into HSV",
			in:   colorcode.HSL{197, 84, 0},
			out:  "hsv(197, 91%, 0%)",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := tC.in.ToHSV()
			if result.String() != tC.out {
				t.Errorf("Expected %s but got %s", tC.out, result.String())
			}
		})
	}
}

func benchmarkHSLToRGB(b *testing.B, h, s, l float64) {
	for i := 0; i < b.N; i++ {
		colorcode.HSL{h, s, l}.ToRGB()
	}
}

func BenchmarkHSLToRGB_NonGray(b *testing.B) {
	benchmarkHSLToRGB(b, 200, 90, 50)
}

func BenchmarkHSLToRGB_Gray(b *testing.B) {
	benchmarkHSLToRGB(b, 0, 0, 50)
}

func BenchmarkHSLToRGB_LowLuminosity(b *testing.B) {
	benchmarkHSLToRGB(b, 256, 100, 15)
}

func BenchmarkHSLToRGB_Green(b *testing.B) {
	benchmarkHSLToRGB(b, 90, 100, 15)
}

func benchmarkHSLToHSV(b *testing.B, h, s, l float64) {
	for i := 0; i < b.N; i++ {
		colorcode.HSL{h, s, l}.ToRGB()
	}
}

func BenchmarkHSLToHSV_AColor(b *testing.B) {
	benchmarkHSLToHSV(b, 197, 84, 63)
}

func BenchmarkHSLToHSV_ADullColor(b *testing.B) {
	benchmarkHSLToHSV(b, 197, 84, 30)
}

func BenchmarkHSLToHSV_AnotherDullColor(b *testing.B) {
	benchmarkHSLToHSV(b, 197, 84, 0)
}
