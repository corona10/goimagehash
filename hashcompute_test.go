// Copyright 2017 The goimagehash Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goimagehash

import (
	"image"
	"image/jpeg"
	"os"
	"testing"
)

func TestHashCompute(t *testing.T) {
	for _, tt := range []struct {
		img1     string
		img2     string
		method   func(img image.Image) (*ImageHash, error)
		name     string
		distance int
	}{
		{"_examples/sample1.jpg", "_examples/sample1.jpg", AverageHash, "AverageHash", 0},
		{"_examples/sample2.jpg", "_examples/sample2.jpg", AverageHash, "AverageHash", 0},
		{"_examples/sample3.jpg", "_examples/sample3.jpg", AverageHash, "AverageHash", 0},
		{"_examples/sample4.jpg", "_examples/sample4.jpg", AverageHash, "AverageHash", 0},
		{"_examples/sample1.jpg", "_examples/sample2.jpg", AverageHash, "AverageHash", 42},
		{"_examples/sample1.jpg", "_examples/sample3.jpg", AverageHash, "AverageHash", 4},
		{"_examples/sample1.jpg", "_examples/sample4.jpg", AverageHash, "AverageHash", 38},
		{"_examples/sample2.jpg", "_examples/sample3.jpg", AverageHash, "AverageHash", 40},
		{"_examples/sample2.jpg", "_examples/sample4.jpg", AverageHash, "AverageHash", 6},
		{"_examples/sample1.jpg", "_examples/sample1.jpg", DifferenceHash, "DifferenceHash", 0},
		{"_examples/sample2.jpg", "_examples/sample2.jpg", DifferenceHash, "DifferenceHash", 0},
		{"_examples/sample3.jpg", "_examples/sample3.jpg", DifferenceHash, "DifferenceHash", 0},
		{"_examples/sample4.jpg", "_examples/sample4.jpg", DifferenceHash, "DifferenceHash", 0},
		{"_examples/sample1.jpg", "_examples/sample2.jpg", DifferenceHash, "DifferenceHash", 43},
		{"_examples/sample1.jpg", "_examples/sample3.jpg", DifferenceHash, "DifferenceHash", 0},
		{"_examples/sample1.jpg", "_examples/sample4.jpg", DifferenceHash, "DifferenceHash", 37},
		{"_examples/sample2.jpg", "_examples/sample3.jpg", DifferenceHash, "DifferenceHash", 43},
		{"_examples/sample2.jpg", "_examples/sample4.jpg", DifferenceHash, "DifferenceHash", 16},
		{"_examples/sample1.jpg", "_examples/sample1.jpg", PerceptionHash, "PerceptionHash", 0},
		{"_examples/sample2.jpg", "_examples/sample2.jpg", PerceptionHash, "PerceptionHash", 0},
		{"_examples/sample3.jpg", "_examples/sample3.jpg", PerceptionHash, "PerceptionHash", 0},
		{"_examples/sample4.jpg", "_examples/sample4.jpg", PerceptionHash, "PerceptionHash", 0},
		{"_examples/sample1.jpg", "_examples/sample2.jpg", PerceptionHash, "PerceptionHash", 32},
		{"_examples/sample1.jpg", "_examples/sample3.jpg", PerceptionHash, "PerceptionHash", 2},
		{"_examples/sample1.jpg", "_examples/sample4.jpg", PerceptionHash, "PerceptionHash", 30},
		{"_examples/sample2.jpg", "_examples/sample3.jpg", PerceptionHash, "PerceptionHash", 34},
		{"_examples/sample2.jpg", "_examples/sample4.jpg", PerceptionHash, "PerceptionHash", 20},
	} {
		file1, err := os.Open(tt.img1)
		if err != nil {

		}
		defer file1.Close()

		file2, err := os.Open(tt.img2)
		if err != nil {
			t.Errorf("%s", err)
		}
		defer file2.Close()

		img1, err := jpeg.Decode(file1)
		if err != nil {
			t.Errorf("%s", err)
		}

		img2, err := jpeg.Decode(file2)
		if err != nil {
			t.Errorf("%s", err)
		}

		hash1, err := tt.method(img1)
		if err != nil {
			t.Errorf("%s", err)
		}
		hash2, err := tt.method(img2)
		if err != nil {
			t.Errorf("%s", err)
		}

		dis1, err := hash1.Distance(hash2)
		if err != nil {
			t.Errorf("%s", err)
		}

		dis2, err := hash2.Distance(hash1)
		if err != nil {
			t.Errorf("%s", err)
		}

		if dis1 != dis2 {
			t.Errorf("Distance should be identical %v vs %v", dis1, dis2)
		}

		if dis1 != tt.distance {
			t.Errorf("%s: Distance between %v and %v is expected %v but got %v", tt.name, tt.img1, tt.img2, tt.distance, dis1)
		}
	}
}

func NilHashComputeTest(t *testing.T) {
	hash, err := AverageHash(nil)
	if err == nil {
		t.Errorf("Error should be got.")
	}
	if hash != nil {
		t.Errorf("Nil hash should be got. but got %v", hash)
	}

	hash, err = DifferenceHash(nil)
	if err == nil {
		t.Errorf("Error should be got.")
	}
	if hash != nil {
		t.Errorf("Nil hash should be got. but got %v", hash)
	}

	hash, err = PerceptionHash(nil)
	if err == nil {
		t.Errorf("Error should be got.")
	}
	if hash != nil {
		t.Errorf("Nil hash should be got. but got %v", hash)
	}
}

func BenchmarkDistanceIdentical(b *testing.B) {
	h1 := &ImageHash{hash: 0xe48ae53c05e502f7}
	h2 := &ImageHash{hash: 0xe48ae53c05e502f7}

	for i := 0; i < b.N; i++ {
		h1.Distance(h2)
	}
}

func BenchmarkDistanceDifferent(b *testing.B) {
	h1 := &ImageHash{hash: 0xe48ae53c05e502f7}
	h2 := &ImageHash{hash: 0x678be53815e510f7} // 8 bits flipped

	for i := 0; i < b.N; i++ {
		h1.Distance(h2)
	}
}

func TestExtImageHashCompute(t *testing.T) {
	for _, tt := range []struct {
		img1     string
		img2     string
		hashSize int
		name     string
		distance int
	}{
		{"_examples/sample1.jpg", "_examples/sample1.jpg", 8, "PerceptionHashExtend", 0},
		{"_examples/sample2.jpg", "_examples/sample2.jpg", 8, "PerceptionHashExtend", 0},
		{"_examples/sample3.jpg", "_examples/sample3.jpg", 8, "PerceptionHashExtend", 0},
		{"_examples/sample4.jpg", "_examples/sample4.jpg", 8, "PerceptionHashExtend", 0},
		{"_examples/sample1.jpg", "_examples/sample2.jpg", 8, "PerceptionHashExtend", 32},
		{"_examples/sample1.jpg", "_examples/sample3.jpg", 8, "PerceptionHashExtend", 2},
		{"_examples/sample1.jpg", "_examples/sample4.jpg", 8, "PerceptionHashExtend", 30},
		{"_examples/sample2.jpg", "_examples/sample3.jpg", 8, "PerceptionHashExtend", 34},
		{"_examples/sample2.jpg", "_examples/sample4.jpg", 8, "PerceptionHashExtend", 20},
		{"_examples/sample1.jpg", "_examples/sample1.jpg", 16, "PerceptionHashExtend", 0},
		{"_examples/sample2.jpg", "_examples/sample2.jpg", 16, "PerceptionHashExtend", 0},
		{"_examples/sample3.jpg", "_examples/sample3.jpg", 16, "PerceptionHashExtend", 0},
		{"_examples/sample4.jpg", "_examples/sample4.jpg", 16, "PerceptionHashExtend", 0},
	} {
		file1, err := os.Open(tt.img1)
		if err != nil {
			t.Errorf("%s", err)
		}
		defer file1.Close()

		file2, err := os.Open(tt.img2)
		if err != nil {
			t.Errorf("%s", err)
		}
		defer file2.Close()

		img1, err := jpeg.Decode(file1)
		if err != nil {
			t.Errorf("%s", err)
		}

		img2, err := jpeg.Decode(file2)
		if err != nil {
			t.Errorf("%s", err)
		}

		hash1, err := PerceptionHashExtend(img1, tt.hashSize)
		if err != nil {
			t.Errorf("%s", err)
		}
		hash2, err := PerceptionHashExtend(img2, tt.hashSize)
		if err != nil {
			t.Errorf("%s", err)
		}

		dis1, err := hash1.Distance(hash2)
		if err != nil {
			t.Errorf("%s", err)
		}

		dis2, err := hash2.Distance(hash1)
		if err != nil {
			t.Errorf("%s", err)
		}

		if dis1 != dis2 {
			t.Errorf("Distance should be identical %v vs %v", dis1, dis2)
		}

		if dis1 != tt.distance {
			t.Errorf("%s: Distance between %v and %v is expected %v but got %v", tt.name, tt.img1, tt.img2, tt.distance, dis1)
		}

		if tt.hashSize == 8 {
			hash0, err := PerceptionHash(img1)
			if err != nil {
				t.Errorf("%s", err)
			}
			hex0 := hash0.ToString()
			hex1 := hash1.ToString()
			if hex0 != hex1 {
				t.Errorf("Hex is expected %v but got %v", hex0, hex1)
			}
		}
	}
}

func BenchmarkExtImageHashDistanceDifferent(b *testing.B) {
	h1 := &ExtImageHash{hash: []uint64{0xe48ae53c05e502f7}}
	h2 := &ExtImageHash{hash: []uint64{0x678be53815e510f7}} // 8 bits flipped

	for i := 0; i < b.N; i++ {
		h1.Distance(h2)
	}
}
