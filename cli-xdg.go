// Copyright Louis Royer and the NextMN contributors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT
package clixdg

import (
	"fmt"

	"github.com/adrg/xdg"
	"github.com/urfave/cli/v3"
)

type valueSource struct {
	path string
}

func (vsc valueSource) String() string {
	return fmt.Sprintf("XDG config file %[1]q", vsc.path)
}

func (vsc valueSource) GoString() string {
	return fmt.Sprintf("valueSource{path:%[1]q}", vsc.path)
}

func (vsc valueSource) Lookup() (string, bool) {
	if xdgPath, err := xdg.SearchConfigFile(vsc.path); err == nil {
		return xdgPath, true
	}
	return "", false
}

func ConfigFile(path string) cli.ValueSource {
	return valueSource{path: path}
}
