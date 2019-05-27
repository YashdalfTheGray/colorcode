# colorcode

A library of color spaces and conversions written in Go 🎨

## Get the library

Getting the code down is as simple as running `go get yashdalfthegray/colorcode`

## What's included?

This library comes with a few different color spaces to play around with - RGB + hex code colors, HSL, and HSV. Those are the most common ones that require conversions. Feel free to open an issue if you would like to see a color space added to this library.

## Usage

Each color space, RGB, hex, HSL and HSV give you a way to construct a color of that type by passing in the channels required. The package comes with methods that start with `New` that let you create the colors. For example, to create a new RGB color, you would use `colorcode.NewRGB(127, 127, 127)`.

Furthermore, each of the color spaces extend `fmt.Stringer` if you ever wanted a string representation of them. On instances that are returned from the `New` methods, you will find `To` methods which are used to convert between color spaces. For example, `rgbColor.ToHSV()` will return an HSV color and `hexColor.toHSL()` will return an HSL color. Some color conversions throw errors because the library wants you to know that something went wrong outside of the math that we're using to convert.

An example is below.

```golang
package main

import "github.com/yashdalfthegray/colorcode"

func main() {
  hexColor := colorcode.NewHexCode("#deadaf")
  rgbColor, err := hexColor.ToRGB()
  if err != nil {
    hslColor := rgbColor.ToHSL()
    hsvColor := hslColor.ToHSV()
  }
}
```
