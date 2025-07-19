// Copyright (c) 2024 Six After, Inc
//
// This source code is licensed under the Apache 2.0 License found in the
// LICENSE file in the root directory of this source tree.

package prng

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// FuzzPRNGRead fuzzes the package-level Reader using various buffer sizes.
func Fuzz_PRNG_Read(f *testing.F) {
	f.Add(0)
	f.Add(1)
	f.Add(32)
	f.Add(64)
	f.Add(256)
	f.Add(1024)
	f.Add(4096)

	f.Fuzz(func(t *testing.T, size int) {
		t.Parallel()
		is := assert.New(t)

		// Skip absurdly large or invalid sizes
		if size < 0 || size > 1<<20 {
			t.Skip()
		}

		buf := make([]byte, size)
		n, err := Reader.Read(buf)

		is.NoError(err, "Reader.Read should not return error")
		is.Equal(size, n, "expected %d bytes from Reader.Read", size)
	})
}

// FuzzNewReader checks that a new Reader can be created and used correctly.
func Fuzz_NewReader(f *testing.F) {
	f.Add(16)
	f.Add(64)
	f.Add(512)

	f.Fuzz(func(t *testing.T, size int) {
		t.Parallel()
		is := assert.New(t)

		if size < 0 || size > 65536 {
			t.Skip()
		}

		r, err := NewReader()
		is.NoError(err, "NewReader should succeed")

		buf := make([]byte, size)
		n, err := r.Read(buf)

		is.NoError(err, "Reader.Read from NewReader should not error")
		is.Equal(size, n, "expected %d bytes from NewReader.Read", size)
	})
}
