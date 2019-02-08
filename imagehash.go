// Copyright 2017 The goimagehash Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goimagehash

import (
	"encoding/binary"
	"encoding/hex"
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

// ExtImageHash is a struct of big hash computation.
type ExtImageHash struct {
	hash []uint64
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

	lhash := h.GetHash()
	rhash := other.GetHash()

	hamming := lhash ^ rhash
	return popcnt(hamming), nil
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

// NewExtImageHash function creates a new big hash
func NewExtImageHash(hash []uint64, kind Kind) *ExtImageHash {
	return &ExtImageHash{hash: hash, kind: kind}
}

// Distance method returns a distance between two big hashes
func (h *ExtImageHash) Distance(other *ExtImageHash) (int, error) {
	if h.GetKind() != other.GetKind() {
		return -1, errors.New("Extended Image hashes's kind should be identical")
	}

	lHash := h.GetHash()
	rHash := other.GetHash()
	if len(lHash) != len(rHash) {
		return -1, errors.New("Extended Image hashes's size should be identical")
	}

	var distance int
	for idx, lh := range lHash {
		rh := rHash[idx]
		hamming := lh ^ rh
		distance += popcnt(hamming)
	}
	return distance, nil
}

// GetHash method returns a big hash value
func (h *ExtImageHash) GetHash() []uint64 {
	return h.hash
}

// GetKind method returns a kind of big hash
func (h *ExtImageHash) GetKind() Kind {
	return h.kind
}

const extStrFmt = "%1s:%s"

// ExtImageHashFromString returns a big hash from a hex representation
func ExtImageHashFromString(s string) (*ExtImageHash, error) {
	var kindStr string
	var hashStr string
	_, err := fmt.Sscanf(s, extStrFmt, &kindStr, &hashStr)
	if err != nil {
		return nil, errors.New("Couldn't parse string " + s)
	}

	hexBytes, err := hex.DecodeString(hashStr)
	if err != nil {
		return nil, err
	}

	var hash []uint64
	lenOfByte := 8
	for i := 0; i < len(hexBytes)/lenOfByte; i++ {
		startIndex := i * lenOfByte
		endIndex := startIndex + lenOfByte
		hashUint64 := binary.BigEndian.Uint64(hexBytes[startIndex:endIndex])
		hash = append(hash, hashUint64)
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
	return NewExtImageHash(hash, kind), nil
}

// ToString returns a hex representation of big hash
func (h *ExtImageHash) ToString() string {
	var hexBytes []byte
	for _, hash := range h.hash {
		hashBytes := make([]byte, 8)
		binary.BigEndian.PutUint64(hashBytes, hash)
		hexBytes = append(hexBytes, hashBytes...)
	}
	hexStr := hex.EncodeToString(hexBytes)

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
	return fmt.Sprintf(extStrFmt, kindStr, hexStr)
}
