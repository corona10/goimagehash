// Copyright 2017 The goimagehash Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goimagehash

import (
	"errors"
)

// hashKind describes the kinds of hash.
type hashKind int

// ImageHash is a struct of hash computation.
type ImageHash struct {
	hash uint64
	kind hashKind
}

const (
	Unknown hashKind = iota // Unknown Hash
	AHash                   // Average Hash
	PHash                   // Perceptual Hash
	DHash                   // Difference Hash
	WHash                   // Wavelet Hash
)

// NewImageHash function creates a new image hash.
func NewImageHash(hash uint64, kind hashKind) *ImageHash {
	return &ImageHash{hash: hash, kind: kind}
}

// Distance method returns a distance between two hashes.
func (h *ImageHash) Distance(other *ImageHash) (int, error) {
	if h.GetKind() != other.GetKind() {
		return -1, errors.New("Image hashes's kind should be identical.")
	}

	diff := 0
	lhash := h.GetHash()
	rhash := other.GetHash()

	hamming := lhash ^ rhash
	for hamming != 0 {
		diff += int(hamming & 1)
		hamming >>= 1
	}

	return diff, nil
}

// GetHash method returns a 64bits hash value.
func (h *ImageHash) GetHash() uint64 {
	return h.hash
}

// GetKind method returns a kind of image hash.
func (h *ImageHash) GetKind() hashKind {
	return h.kind
}

// Set method sets a bit of index.
func (h *ImageHash) Set(idx int) {
	h.hash |= 1 << uint(idx)
}
