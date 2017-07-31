// Copyright 2017 The goimagehash Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package etcs

// MeanOfPixels function returns a mean of pixels.
func MeanOfPixels(pixels []float64) float64 {
	m := 0.0
	lens := len(pixels)
	if lens == 0 {
		return 0
	}

	for _, p := range pixels {
		m += p
	}

	return m / float64(lens)
}

// MedianOfPixels function returns a median value of pixels.
// It uses quick selection algorithm.
func MedianOfPixels(pixels []float64) float64 {
	tmp := make([]float64, len(pixels))
	copy(tmp, pixels)
	l := len(tmp) - 1
	pos := l / 2
	v := quickSelect(tmp, 0, l, pos)
	return v
}

func quickSelect(sequence []float64, low int, hi int, k int) float64 {
	if hi-low <= 1 {
		return sequence[k]
	}
	j := low
	sequence[j], sequence[k] = sequence[k], sequence[j]
	j++
	for i := j; i < hi; i++ {
		if sequence[i] < sequence[low] {
			sequence[j], sequence[i] = sequence[i], sequence[j]
			j++
		}
	}
	j--
	sequence[j], sequence[low] = sequence[low], sequence[j]

	if k < j {
		return quickSelect(sequence, low, j, k)
	}

	if k > j {
		return quickSelect(sequence, j+1, hi, k-j)
	}

	return sequence[j]
}
