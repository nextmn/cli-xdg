// Copyright Louis Royer and the NextMN contributors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package osrelease

import (
	"strings"
	"testing"
)

func TestFromFile(t *testing.T) {
	debian := `
PRETTY_NAME="Debian GNU/Linux 12 (bookworm)"
NAME="Debian GNU/Linux"
VERSION_ID="12"
VERSION="12 (bookworm)"
VERSION_CODENAME=bookworm
ID=debian
HOME_URL="https://www.debian.org/"
SUPPORT_URL="https://www.debian.org/support"
BUG_REPORT_URL="https://bugs.debian.org/"
`
	rel, err := FromFile(strings.NewReader(debian))
	if err != nil {
		t.Fatal(err)
	}
	if rel.Id != "debian" {
		t.Errorf("osRelease.Id = %s; want debian", rel.Id)
	}
	if rel.VersionId != "12" {
		t.Errorf("osRelease.VersionId = %s; want 12", rel.VersionId)
	}
	if rel.VersionCodename != "bookworm" {
		t.Errorf("osRelease.VersionCodename = %s; want bookworm", rel.VersionCodename)
	}

	alpine := `
NAME="Alpine Linux"
ID=alpine
VERSION_ID=3.22.1
PRETTY_NAME="Alpine Linux v3.22"
HOME_URL="https://alpinelinux.org/"
BUG_REPORT_URL="https://gitlab.alpinelinux.org/alpine/aports/-/issues"
`
	rel, err = FromFile(strings.NewReader(alpine))
	if err != nil {
		t.Fatal(err)
	}
	if rel.Id != "alpine" {
		t.Errorf("osRelease.Id = %s; want alpine", rel.Id)
	}
	if rel.VersionId != "3.22.1" {
		t.Errorf("osRelease.VersionId = %s; want 3.22.1", rel.VersionId)
	}
}
