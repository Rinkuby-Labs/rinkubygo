// Copyright (C) 2019-2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package customrawdb

import (
	"testing"

	"github.com/ava-labs/libevm/core/rawdb"
	"github.com/stretchr/testify/require"
)

func TestParseStateScheme(t *testing.T) {
	db := rawdb.NewMemoryDatabase()

	// Pass-through to rawdb for hash scheme on a fresh empty DB.
	scheme, err := ParseStateScheme(rawdb.HashScheme, db)
	require.NoError(t, err)
	require.Equal(t, rawdb.HashScheme, scheme)
}
