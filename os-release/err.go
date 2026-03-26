// Copyright Louis Royer and the NextMN contributors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package osrelease

import (
	"errors"
)

var (
	ErrNoOsReleaseFile = errors.New("cannot open os-release file")
	ErrCannotParseFile = errors.New("cannot parse os-release file")
)
