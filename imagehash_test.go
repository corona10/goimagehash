// Copyright 2017 The goimagehash Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goimagehash

import (
	"errors"
	"testing"
)

func TestNewImageHash(t *testing.T) {
	for _, tt := range []struct {
		datas    [][]uint8
		hash1    Kind
		hash2    Kind
		distance int
		err      error
	}{
		{[][]uint8{{1, 0, 1, 1}, {0, 0, 0, 0}}, Unknown, Unknown, 3, nil},
		{[][]uint8{{0, 0, 0, 0}, {0, 0, 0, 0}}, Unknown, Unknown, 0, nil},
		{[][]uint8{{0, 0, 0, 0}, {0, 0, 0, 1}}, Unknown, Unknown, 1, nil},
		{[][]uint8{{0, 0, 0, 0}, {0, 0, 0, 1}}, Unknown, AHash, -1, errors.New("Image hashes's kind should be identical")},
	} {
		data1 := tt.datas[0]
		data2 := tt.datas[1]
		hash1 := NewImageHash(0, tt.hash1)
		hash2 := NewImageHash(0, tt.hash2)

		for i := 0; i < len(data1); i++ {
			if data1[i] == 1 {
				hash1.Set(i)
			}
		}

		for i := 0; i < len(data2); i++ {
			if data2[i] == 1 {
				hash2.Set(i)
			}
		}

		dis, err := hash1.Distance(hash2)
		if dis != tt.distance {
			t.Errorf("Distance between %v and %v expected as %d but got %d.", data1, data2, tt.distance, dis)
		}
		if err != nil && err.Error() != tt.err.Error() {
			t.Errorf("Expected err %s, actual %s.", tt.err, err)
		}
	}

}
