// Copyright 2017 The goimagehash Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package transforms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const EPS = 1e-5

func TestDWT1D(t *testing.T) {
	tests := []struct {
		name string
		data []float64
		want []float64
	}{
		{name: "vector",
			data: []float64{3.4, 8.4, 6.5, 9.9, 7.6, 9.3, 6.6, 5.2},
			want: []float64{5.9, 8.2, 8.45, 5.9, -2.5, -1.7, -0.85, 0.7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DWT1D(tt.data)
			assert.InDeltaSlice(t, tt.want, tt.data, EPS)
		})
	}
}

func Test_IDWT1D(t *testing.T) {
	tests := []struct {
		name string
		data []float64
		want []float64
	}{
		{name: "vector inverse",
			data: []float64{5.9, 8.2, 8.45, 5.9, -2.5, -1.7, -0.85, 0.7},
			want: []float64{3.4, 8.4, 6.5, 9.9, 7.6, 9.3, 6.6, 5.2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			IDWT1D(tt.data)
			assert.InDeltaSlice(t, tt.want, tt.data, EPS)
		})
	}
}

func TestDWT2D(t *testing.T) {
	tests := []struct {
		name  string
		data  [][]float64
		want  [][]float64
		level int
	}{
		{name: "matrix level1",
			data: [][]float64{
				{3.4, 8.4, 6.5, 9.9},
				{7.6, 9.3, 6.6, 5.2},
				{5.1, 5.9, 2.3, 3.7},
				{8.3, 0.9, 3.6, 0.2}},
			want: [][]float64{
				{7.175, 7.05, -1.675, -0.5},
				{5.05, 2.45, 1.65, 0.5},
				{-1.275, 1.15, -0.825, -1.2},
				{0.45, 0.55, -2.05, -1.2},
			},
			level: 1,
		},
		{name: "matrix level2",
			data: [][]float64{
				{3.4, 8.4, 6.5, 9.9},
				{7.6, 9.3, 6.6, 5.2},
				{5.1, 5.9, 2.3, 3.7},
				{8.3, 0.9, 3.6, 0.2}},
			want: [][]float64{
				{5.43125, 0.68125, -1.675, -0.5},
				{1.68125, -0.61875, 1.65, 0.5},
				{-1.275, 1.15, -0.825, -1.2},
				{0.45, 0.55, -2.05, -1.2},
			},
			level: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DWT2D(tt.data, tt.level)
			for i := range tt.data {
				assert.InDeltaSlice(t, tt.want[i], tt.data[i], EPS)
			}
		})
	}
}

func Test_IDWT2D(t *testing.T) {
	tests := []struct {
		name  string
		data  [][]float64
		want  [][]float64
		level int
	}{
		{name: "matrix inverse level1",
			want: [][]float64{
				{3.4, 8.4, 6.5, 9.9},
				{7.6, 9.3, 6.6, 5.2},
				{5.1, 5.9, 2.3, 3.7},
				{8.3, 0.9, 3.6, 0.2}},
			data: [][]float64{
				{7.175, 7.05, -1.675, -0.5},
				{5.05, 2.45, 1.65, 0.5},
				{-1.275, 1.15, -0.825, -1.2},
				{0.45, 0.55, -2.05, -1.2},
			},
			level: 1,
		},
		{name: "matrix inverse level2",
			want: [][]float64{
				{3.4, 8.4, 6.5, 9.9},
				{7.6, 9.3, 6.6, 5.2},
				{5.1, 5.9, 2.3, 3.7},
				{8.3, 0.9, 3.6, 0.2}},
			data: [][]float64{
				{5.43125, 0.68125, -1.675, -0.5},
				{1.68125, -0.61875, 1.65, 0.5},
				{-1.275, 1.15, -0.825, -1.2},
				{0.45, 0.55, -2.05, -1.2},
			},
			level: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			IDWT2D(tt.data, tt.level)
			for i := range tt.data {
				assert.InDeltaSlice(t, tt.want[i], tt.data[i], EPS)
			}
		})
	}
}

func Test_Floorp2(t *testing.T) {
	tests := []struct {
		name string
		val  int
		want uint
	}{
		{
			name: "1023",
			val:  1023,
			want: 512,
		},
		{
			name: "1024",
			val:  1024,
			want: 1024,
		},
		{
			name: "1025",
			val:  1025,
			want: 1024,
		},
		{
			name: "32",
			val:  32,
			want: 32,
		},
		{
			name: "9999",
			val:  9999,
			want: 8192,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Floorp2(tt.val))
		})
	}
}

func Test_min(t *testing.T) {
	tests := []struct {
		name string
		x    int
		y    int
		want int
	}{
		{
			name: "0 vs 0",
			x:    0,
			y:    0,
			want: 0,
		},
		{
			name: "1 vs 0",
			x:    1,
			y:    0,
			want: 0,
		},
		{
			name: "0 vs 1",
			x:    0,
			y:    1,
			want: 0,
		},
		{
			name: "-1 vs 1",
			x:    -1,
			y:    1,
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Min(tt.x, tt.y), "Min(%v, %v)", tt.x, tt.y)
		})
	}
}
