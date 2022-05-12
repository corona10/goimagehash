// Copyright 2017 The goimagehash Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package transforms

import (
	"math"
	"sync"
)

// DCT1D function returns result of DCT-II.
// DCT type II, unscaled. Algorithm by Byeong Gi Lee, 1984.
func DCT1D(input []float64) []float64 {
	temp := make([]float64, len(input))
	forwardTransform(input, temp, len(input))
	return input
}

func forwardTransform(input, temp []float64, Len int) {
	if Len == 1 {
		return
	}

	halfLen := Len / 2
	for i := 0; i < halfLen; i++ {
		x, y := input[i], input[Len-1-i]
		temp[i] = x + y
		temp[i+halfLen] = (x - y) / (math.Cos((float64(i)+0.5)*math.Pi/float64(Len)) * 2)
	}
	forwardTransform(temp, input, halfLen)
	forwardTransform(temp[halfLen:], input, halfLen)
	for i := 0; i < halfLen-1; i++ {
		input[i*2+0] = temp[i]
		input[i*2+1] = temp[i+halfLen] + temp[i+halfLen+1]
	}
	input[Len-2], input[Len-1] = temp[halfLen-1], temp[Len-1]
}

// DCT2D function returns a  result of DCT2D by using the seperable property.
func DCT2D(input [][]float64, w int, h int) [][]float64 {
	output := make([][]float64, h)
	for i := range output {
		output[i] = make([]float64, w)
	}

	wg := new(sync.WaitGroup)
	for i := 0; i < h; i++ {
		wg.Add(1)
		go func(i int) {
			cols := DCT1D(input[i])
			output[i] = cols
			wg.Done()
		}(i)
	}

	wg.Wait()
	for i := 0; i < w; i++ {
		wg.Add(1)
		in := make([]float64, h)
		go func(i int) {
			for j := 0; j < h; j++ {
				in[j] = output[j][i]
			}
			rows := DCT1D(in)
			for j := 0; j < len(rows); j++ {
				output[j][i] = rows[j]
			}
			wg.Done()
		}(i)
	}

	wg.Wait()
	return output
}

// DCT2DFast64 function returns a result of DCT2D by using the seperable property.
// Fast uses static DCT tables for improved performance.
func DCT2DFast64(input *[]float64) {
	if len(*input) != 4096 {
		panic("incorrect input size")
	}

	for i := 0; i < 64; i++ { // height
		DCT1DFast64((*input)[i*64 : (i*64)+64])
		//forwardTransformFast((*input)[i*64:(i*64)+64], temp[:], 64)
	}

	for i := 0; i < 64; i++ { // width
		row := [64]float64{}
		for j := 0; j < 64; j++ {
			row[j] = (*input)[i+((j)*64)]
		}
		DCT1DFast64(row[:])
		for j := 0; j < len(row); j++ {
			(*input)[i+(j*64)] = row[j]
		}
	}
}

func forwardTransformFast(input, temp []float64, Len int) {
	if Len == 1 {
		return
	}

	halfLen := Len / 2
	t := dctTables[halfLen>>1]
	for i := 0; i < halfLen; i++ {
		x, y := input[i], input[Len-1-i]
		temp[i] = x + y
		temp[i+halfLen] = (x - y) / t[i]
	}
	forwardTransformFast(temp, input, halfLen)
	forwardTransformFast(temp[halfLen:], input, halfLen)
	for i := 0; i < halfLen-1; i++ {
		input[i*2+0] = temp[i]
		input[i*2+1] = temp[i+halfLen] + temp[i+halfLen+1]
	}

	input[Len-2], input[Len-1] = temp[halfLen-1], temp[Len-1]
}
