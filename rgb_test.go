package colorcode_test

import (
	"testing"

	"github.com/YashdalfTheGray/colorcode"
)

func TestNewRGB(t *testing.T) {
	testCases := []struct {
		desc    string
		r, g, b uint8
		out     colorcode.RGB
	}{
		{
			desc: "builds a color out of rgb values",
			r:    10,
			g:    45,
			b:    234,
			out:  colorcode.RGB{10, 45, 234},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if c := colorcode.NewRGB(tC.r, tC.g, tC.b); c.String() != tC.out.String() {
				t.Errorf("Expected %s but got %s", tC.out.String(), c.String())
			}
		})
	}
}

func TestRGBString(t *testing.T) {
	testCases := []struct {
		desc string
		in   colorcode.RGB
		out  string
	}{
		{
			desc: "builds a string out of an RGB color",
			in:   colorcode.RGB{200, 90, 50},
			out:  "rgb(200, 90, 50)",
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

func TestRGBToHexCode(t *testing.T) {
	testCases := []struct {
		desc string
		in   colorcode.RGB
		out  string
	}{
		{
			desc: "converts an RGB color to hex string",
			in:   colorcode.RGB{0, 0, 0},
			out:  "#000000",
		},
		{
			desc: "converts another RGB color to hex string",
			in:   colorcode.RGB{68, 138, 255},
			out:  "#448aff",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if result := tC.in.ToHexCode(); result.String() != tC.out {
				t.Errorf("Expected %s but got %s", tC.out, result)
			}
		})
	}
}

func TestRGBToHSL(t *testing.T) {
	testCases := []struct {
		desc string
		in   colorcode.RGB
		out  string
	}{
		{
			desc: "converts a non-gray RGB color to HSL",
			in:   colorcode.RGB{13, 166, 242},
			out:  "hsl(200, 90%, 50%)",
		},
		{
			desc: "converts a gray RGB color to HSL",
			in:   colorcode.RGB{128, 128, 128},
			out:  "hsl(0, 0%, 50%)",
		},
		{
			desc: "converts a dark RGB color to HSL",
			in:   colorcode.RGB{40, 34, 90},
			out:  "hsl(246, 45%, 24%)",
		},
		{
			desc: "converts a mostly green RGB color to HSL",
			in:   colorcode.RGB{40, 150, 90},
			out:  "hsl(147, 58%, 37%)",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if result := tC.in.ToHSL(); result.String() != tC.out {
				t.Errorf("Expected %s but got %s", tC.out, result.String())
			}
		})
	}
}

func TestRGBToHSV(t *testing.T) {
	testCases := []struct {
		desc string
		in   colorcode.RGB
		out  string
	}{
		{
			desc: "converts a non-gray RGB color to HSV",
			in:   colorcode.RGB{13, 166, 242},
			out:  "hsv(200, 95%, 95%)",
		},
		{
			desc: "converts a gray RGB color to HSV",
			in:   colorcode.RGB{128, 128, 128},
			out:  "hsv(0, 0%, 50%)",
		},
		{
			desc: "converts black color to HSV",
			in:   colorcode.RGB{0, 0, 0},
			out:  "hsv(0, 0%, 0%)",
		},
		{
			desc: "converts a dark RGB color to HSV",
			in:   colorcode.RGB{40, 34, 90},
			out:  "hsv(246, 62%, 35%)",
		},
		{
			desc: "converts a mostly green RGB color to HSV",
			in:   colorcode.RGB{40, 150, 90},
			out:  "hsv(147, 73%, 59%)",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if result := tC.in.ToHSV(); result.String() != tC.out {
				t.Errorf("Expected %s but got %s", tC.out, result.String())
			}
		})
	}
}

func benchmarkRGBToHSL(be *testing.B, r, g, b uint8) {
	for i := 0; i < be.N; i++ {
		colorcode.RGB{r, g, b}.ToHSL()
	}
}

func BenchmarkRGBToHSL_NonGray(b *testing.B) {
	benchmarkRGBToHSL(b, 13, 166, 242)
}

func BenchmarkRGBToHSL_Gray(b *testing.B) {
	benchmarkRGBToHSL(b, 128, 128, 128)
}

func BenchmarkRGBToHSL_Dark(b *testing.B) {
	benchmarkRGBToHSL(b, 40, 34, 90)
}

func BenchmarkRGBToHSL_MostlyGreen(b *testing.B) {
	benchmarkRGBToHSL(b, 40, 150, 90)
}

func benchmarkRGBToHSV(be *testing.B, r, g, b uint8) {
	for i := 0; i < be.N; i++ {
		colorcode.RGB{r, g, b}.ToHSV()
	}
}

func BenchmarkRGBToHSV_NonGray(b *testing.B) {
	benchmarkRGBToHSV(b, 13, 166, 242)
}

func BenchmarkRGBToHSV_Gray(b *testing.B) {
	benchmarkRGBToHSV(b, 128, 128, 128)
}

func BenchmarkRGBToHSV_Black(b *testing.B) {
	benchmarkRGBToHSV(b, 0, 0, 0)
}

func BenchmarkRGBToHSV_Dark(b *testing.B) {
	benchmarkRGBToHSV(b, 40, 34, 90)
}

func BenchmarkRGBToHSV_MostlyGreen(b *testing.B) {
	benchmarkRGBToHSV(b, 40, 150, 90)
}
