// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package govulncheck_test

import (
	"testing"

	"github.com/StevenACoffman/invuln/external/test"
)

func TestImports(t *testing.T) {
	test.VerifyImports(t,
		"github.com/StevenACoffman/invuln/external/osv", // allowed to pull in the osv json entries
	)
}