// Copyright © 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause
package crypt

import (
	"bytes"
	"testing"
)

func TestSecretSharer(t *testing.T) {
	message := "this is some test message to be broken and reconstructed"
	secret := []byte(message)

	n := 10
	k := 3

	ss := NewSecretSharerRandField(1024, n, k)

	shares := ss.BreakSecret(secret)

	shares_slice := shares[:k]
	tmp := shares_slice[0]
	shares_slice[0] = shares_slice[1]
	shares_slice[1] = tmp

	data, err := ss.ReconstructSecret(shares_slice)

	if err != nil {
		t.Fatalf("Failed to reconstruct secret: %s", err.Error())
	}

	if data == nil {
		t.Fatal("Secret sharer test failed - returned nil")
	}

	if !bytes.Equal(secret, data) {
		t.Fatal("Reconstructed data differs from secret")
	}
}
