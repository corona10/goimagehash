// Copyright 2017 The goimagehash Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goimagehash

import (
	"errors"
	"image"

	"github.com/corona10/goimagehash/etcs"
	"github.com/corona10/goimagehash/transforms"
	"github.com/nfnt/resize"
)

// AverageHash fuction returns a hash computation of average hash.
// Implementation follows
// http://www.hackerfactor.com/blog/index.php?/archives/432-Looks-Like-It.html
func AverageHash(img image.Image) (*ImageHash, error) {
	if img == nil {
		return nil, errors.New("Image object can not be nil.")
	}

	// Create 64bits hash.
	ahash := NewImageHash(0, AHash)
	resized := resize.Resize(8, 8, img, resize.Bilinear)
	pixels := transforms.Rgb2Gray(resized)
	flattens := transforms.FlattenPixels(pixels, 8, 8)
	avg := etcs.MeanOfPixels(flattens)

	for idx, p := range flattens {
		if p > avg {
			ahash.Set(idx)
		}
	}

	return ahash, nil
}

// DifferenceHash function returns a hash computation of difference hash.
// Implementation follows
// http://www.hackerfactor.com/blog/?/archives/529-Kind-of-Like-That.html
func DifferenceHash(img image.Image) (*ImageHash, error) {
	if img == nil {
		return nil, errors.New("Image object can not be nil.")
	}

	dhash := NewImageHash(0, DHash)
	resized := resize.Resize(9, 8, img, resize.Bilinear)
	pixels := transforms.Rgb2Gray(resized)
	idx := 0
	for i := 0; i < len(pixels); i++ {
		for j := 0; j < len(pixels[i])-1; j++ {
			if pixels[i][j] < pixels[i][j+1] {
				dhash.Set(idx)
			}
			idx++
		}
	}

	return dhash, nil
}

// PerceptionHash function returns a hash computation of phash.
// Implementation follows
// http://www.hackerfactor.com/blog/index.php?/archives/432-Looks-Like-It.html
func PerceptionHash(img image.Image) (*ImageHash, error) {
	if img == nil {
		return nil, errors.New("Image object can not be nil")
	}

	phash := NewImageHash(0, PHash)
	resized := resize.Resize(64, 64, img, resize.Bilinear)
	pixels := transforms.Rgb2Gray(resized)
	dct := transforms.DCT2D(pixels, 64, 64)
	flattens := transforms.FlattenPixels(dct, 8, 8)
	median := etcs.MedianOfPixels(flattens)

	for idx, p := range flattens {
		if p > median {
			phash.Set(idx)
		}
	}
	return phash, nil
}

// PerceptionHashExtend function returns phash of which the size can be set larger than uint64
// Some variable name refer to https://github.com/JohannesBuchner/imagehash/blob/master/imagehash/__init__.py
// Support 64bits phash (hashSize=8) and 256bits phash (hashSize=16)
func PerceptionHashExtend(img image.Image, hashSize int) (*ExtImageHash, error) {
	if img == nil {
		return nil, errors.New("Image object can not be nil")
	}
	imgSize := hashSize * hashSize
	resized := resize.Resize(uint(imgSize), uint(imgSize), img, resize.Bilinear)
	pixels := transforms.Rgb2Gray(resized)
	dct := transforms.DCT2D(pixels, imgSize, imgSize)
	flattens := transforms.FlattenPixels(dct, hashSize, hashSize)
	median := etcs.MedianOfPixels(flattens)

	lenOfUnit := 64
	phash := make([]uint64, imgSize/lenOfUnit)
	for idx, p := range flattens {
		indexOfArray := (imgSize - 1 - idx) / lenOfUnit
		indexOfBit := idx % lenOfUnit
		if p > median {
			phash[indexOfArray] |= 1 << uint(indexOfBit)
		}
	}
	return NewExtImageHash(phash, PHash), nil
}

// AverageHashExtend function returns ahash of which the size can be set larger than uint64
// Support 64bits ahash (hashSize=8) and 256bits ahash (hashSize=16)
func AverageHashExtend(img image.Image, hashSize int) (*ExtImageHash, error) {
	if img == nil {
		return nil, errors.New("Image object can not be nil")
	}

	imgSize := hashSize * hashSize

	resized := resize.Resize(uint(hashSize), uint(hashSize), img, resize.Bilinear)
	pixels := transforms.Rgb2Gray(resized)
	flattens := transforms.FlattenPixels(pixels, hashSize, hashSize)
	avg := etcs.MeanOfPixels(flattens)

	lenOfUnit := 64
	ahash := make([]uint64, imgSize/lenOfUnit)
	for idx, p := range flattens {
		indexOfArray := (imgSize - 1 - idx) / lenOfUnit
		indexOfBit := idx % lenOfUnit
		if p > avg {
			ahash[indexOfArray] |= 1 << uint(indexOfBit)
		}
	}
	return NewExtImageHash(ahash, AHash), nil
}
