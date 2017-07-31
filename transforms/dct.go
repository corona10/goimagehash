// Copyright 2017 The goimagehash Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package transforms

import (
	"math"
)

// DCT1D function returns result of DCT-II.
// Follows Matlab dct().
// Implementation reference:
// https://unix4lyfe.org/dct-1d/
func DCT1D(input []float64) []float64 {
	n := len(input)
	out := make([]float64, n)
	for i := 0; i < n; i++ {
		z := 0.0
		for j := 0; j < n; j++ {
			z += input[j] * math.Cos(math.Pi*(float64(j)+0.5)*float64(i)/float64(n))
		}

		if i == 0 {
			z *= math.Sqrt(0.5)
		}
		out[i] = z * math.Sqrt(2.0/float64(n))

	}
	return out
}

// DCT2D function returns a  result of DCT2D by using the seperable property.
func DCT2D(input [][]float64, w int, h int) [][]float64 {
	output := make([][]float64, h)
	for i := range output {
		output[i] = make([]float64, w)
	}

	for i := 0; i < h; i++ {
		cols := DCT1D(input[i])
		copy(output[i], cols)
	}

	for i := 0; i < w; i++ {
		in := make([]float64, h)
		for j := 0; j < h; j++ {
			in[j] = output[j][i]
		}
		rows := DCT1D(in)
		for j := 0; j < len(rows); j++ {
			output[j][i] = rows[j]
		}
	}

	return output
}
