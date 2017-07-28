// Copyright 2017 The goimagehash Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goimagehash

import (
	"errors"
)

// hashKind describes the kinds of hash.
type hashKind int

type ImageHash struct {
	hash uint64
	kind hashKind
}

const (
	Unknown hashKind = iota
	AHash            // Average Hash
	PHash            // Perceptual Hash
	DHash            // Difference Hash
	WHash            // Wavlet Hash
)

// Create a new image hash.
func NewImageHash(hash uint64, kind hashKind) *ImageHash {
	return &ImageHash{hash: hash, kind: kind}
}

// Return distance between hashes.
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

// Return hash values.
func (h *ImageHash) GetHash() uint64 {
	return h.hash
}

// Get kind of a hash.
func (h *ImageHash) GetKind() hashKind {
	return h.kind
}

// Set index of bits.
func (h *ImageHash) Set(idx int) {
	h.hash |= 1 << uint(idx)
}
