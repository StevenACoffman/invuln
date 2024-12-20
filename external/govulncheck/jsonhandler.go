// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package govulncheck

import (
	"encoding/json"

	"io"

	"github.com/StevenACoffman/invuln/external/osv"
)

type jsonHandler struct {
	enc *json.Encoder
}

// NewJSONHandler returns a handler that writes govulncheck output as json.
func NewJSONHandler(w io.Writer) Handler {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	return &jsonHandler{enc: enc}
}

// Config writes config block in JSON to the underlying writer.
func (h *jsonHandler) Config(config *Config) error {
	return h.enc.Encode(Message{Config: config})
}

// Progress writes a progress message in JSON to the underlying writer.
func (h *jsonHandler) Progress(progress *Progress) error {
	return h.enc.Encode(Message{Progress: progress})
}

// SBOM writes the SBOM block in JSON to the underlying writer.
func (h *jsonHandler) SBOM(sbom *SBOM) error {
	return h.enc.Encode(Message{SBOM: sbom})
}

// OSV writes an osv entry in JSON to the underlying writer.
func (h *jsonHandler) OSV(entry *osv.Entry) error {
	return h.enc.Encode(Message{OSV: entry})
}

// Finding writes a finding in JSON to the underlying writer.
func (h *jsonHandler) Finding(finding *Finding) error {
	return h.enc.Encode(Message{Finding: finding})
}
