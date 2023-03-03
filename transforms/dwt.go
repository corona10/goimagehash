// Copyright 2017 The goimagehash Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package transforms

const (
	coeff1 = 0.5
	coeff2 = -0.5
)

func DWT1D(data []float64) {
	temp := make([]float64, len(data))
	half := len(data) / 2
	for i := 0; i < half; i++ {
		k := i * 2
		temp[i] = coeff1*data[k] + coeff1*data[k+1]
		temp[i+half] = coeff1*data[k] + coeff2*data[k+1]
	}
	copy(data, temp)
}

func DWT2D(data [][]float64, level int) {
	dims := len(data)
	for k := 0; k < level; k++ {
		curlvl := 1 << k
		curdims := dims / curlvl
		row := make([]float64, curdims)
		for i := 0; i < curdims; i++ {
			copy(row, data[i])
			DWT1D(row)
			copy(data[i], row)
		}
		col := make([]float64, curdims)
		for j := 0; j < curdims; j++ {
			for i := 0; i < curdims; i++ {
				col[i] = data[i][j]
			}
			DWT1D(col)
			for i := 0; i < curdims; i++ {
				data[i][j] = col[i]
			}
		}
	}
}

func IDWT1D(data []float64) {
	temp := make([]float64, len(data))
	half := len(data) / 2
	for i := 0; i < half; i++ {
		k := i * 2
		temp[k] = (coeff1*data[i] + coeff1*data[i+half]) / coeff1
		temp[k+1] = (coeff1*data[i] + coeff2*data[i+half]) / coeff1
	}
	copy(data, temp)
}

func IDWT2D(data [][]float64, level int) {
	dims := len(data)
	for k := level - 1; k >= 0; k-- {
		curlvl := 1 << k
		curdims := dims / curlvl
		col := make([]float64, curdims)
		for j := 0; j < curdims; j++ {
			for i := 0; i < curdims; i++ {
				col[i] = data[i][j]
			}
			IDWT1D(col)
			for i := 0; i < curdims; i++ {
				data[i][j] = col[i]
			}
		}
		row := make([]float64, curdims)
		for i := 0; i < curdims; i++ {
			copy(row, data[i])
			IDWT1D(row)
			copy(data[i], row)
		}
	}
}

// Floorp2 computes closest to val power of 2 (less or equal)
func Floorp2(val int) uint {
	val |= val >> 1
	val |= val >> 2
	val |= val >> 4
	val |= val >> 8
	val |= val >> 16
	val |= val >> 32
	return uint(val - (val >> 1))
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
