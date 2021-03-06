[![Build Status](https://travis-ci.com/YashdalfTheGray/colorcode.svg?branch=master)](https://travis-ci.com/YashdalfTheGray/colorcode)
[![tests](https://github.com/YashdalfTheGray/colorcode/actions/workflows/tests.yml/badge.svg?branch=master)](https://github.com/YashdalfTheGray/colorcode/actions/workflows/tests.yml)
[![Coverage Status](https://coveralls.io/repos/github/YashdalfTheGray/colorcode/badge.svg?branch=master)](https://coveralls.io/github/YashdalfTheGray/colorcode?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/yashdalfthegray/colorcode)](https://goreportcard.com/report/github.com/yashdalfthegray/colorcode)
[![Documentation](https://godoc.org/github.com/yashdalfthegray/colorcode?status.svg)](http://godoc.org/github.com/yashdalfthegray/colorcode)

# colorcode

A library of color spaces and conversions written in Go 🎨

## Get the library

Getting the code down is as simple as running `go get github.com/yashdalfthegray/colorcode` and then importing into your code with `import "github.com/yashdalfthegray/colorcode"`.

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

## Contributing

Feel free to use Github to open issues and create PRs against this repository. There are templates that you can use to describe issues and PRs as needed. Please follow the guidelines in the [contribution guide](.github/CONTRIBUTING.md) for development. This project uses Go 1.12 with standard tooling and Go modules.
