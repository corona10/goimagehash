// Copyright 2017 The goimagehash Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goimagehash

import (
	"errors"
	"fmt"
)

// Kind describes the kinds of hash.
type Kind int

// ImageHash is a struct of hash computation.
type ImageHash struct {
	hash uint64
	kind Kind
}

const (
	// Unknown is a enum value of the unknown hash.
	Unknown Kind = iota
	// AHash is a enum value of the average hash.
	AHash
	//PHash is a enum value of the perceptual hash.
	PHash
	// DHash is a enum value of the difference hash.
	DHash
	// WHash is a enum value of the wavelet hash.
	WHash
)

// NewImageHash function creates a new image hash.
func NewImageHash(hash uint64, kind Kind) *ImageHash {
	return &ImageHash{hash: hash, kind: kind}
}

// Distance method returns a distance between two hashes.
func (h *ImageHash) Distance(other *ImageHash) (int, error) {
	if h.GetKind() != other.GetKind() {
		return -1, errors.New("Image hashes's kind should be identical")
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
func (h *ImageHash) GetKind() Kind {
	return h.kind
}

// Set method sets a bit of index.
func (h *ImageHash) Set(idx int) {
	h.hash |= 1 << uint(idx)
}

const strFmt = "%1s:%016x"

// ImageHashFromString returns an image hash from a hex representation
func ImageHashFromString(s string) (*ImageHash, error) {
	var kindStr string
	var hash uint64
	_, err := fmt.Sscanf(s, strFmt, &kindStr, &hash)
	if err != nil {
		return nil, errors.New("Couldn't parse string " + s)
	}

	kind := Unknown
	switch kindStr {
	case "a":
		kind = AHash
	case "p":
		kind = PHash
	case "d":
		kind = DHash
	case "w":
		kind = WHash
	}
	return NewImageHash(hash, kind), nil
}

// ToString returns a hex representation of the hash
func (h *ImageHash) ToString() string {
	kindStr := ""
	switch h.kind {
	case AHash:
		kindStr = "a"
	case PHash:
		kindStr = "p"
	case DHash:
		kindStr = "d"
	case WHash:
		kindStr = "w"
	}
	return fmt.Sprintf(strFmt, kindStr, h.hash)
}
