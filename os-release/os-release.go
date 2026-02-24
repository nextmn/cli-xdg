// Copyright Louis Royer and the NextMN contributors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package osrelease

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strings"
)

// OsRelease contains operating system identification data,
// programmatic fields only (see `OS-RELEASE(5)`)
type OsRelease struct {
	// Operating system identifier, excluding any version information.
	// If not set, a default `ID=linux` may be used
	Id string
	// List of operating system identifiers that are closely related
	// to the local operating system in regards
	// to packaging and programming interface
	IdLike []string

	// Specific variant or edition of the operating system
	VariantId string

	// Operating system version,
	// excluding any OS name information or release code name
	VersionId string
	// Operating system release code name,
	// excluding any OS name information or release version
	VersionCodename string

	// System image originally used as the installation base
	// Optional field
	BuildId string

	// Specific image of the operating system
	// Optional field
	ImageId string
	// OS image version
	ImageVersion string
}

// Escape field value
func escape(field string) (string, error) {
	if len(field) == 0 {
		return "", ErrCannotParseFile
	}
	switch field[0] {
	case '\'':
		field = strings.Trim(field, "'")
	case '"':
		field = strings.Trim(field, "\"")
	}

	builder := strings.Builder{}
	esc := false
	for _, b := range field {
		if b == '\\' && !esc {
			esc = true
		} else if strings.ContainsRune("$'\"\\`", b) == esc {
			// Shell special characters ("$", quotes, backslash, backtick) must be escaped
			builder.WriteRune(b)
			esc = false
		} else {
			return "", ErrCannotParseFile
		}
	}
	return builder.String(), nil
}

// Create OsRelease from a file
func FromFile(file io.Reader) (*OsRelease, error) {
	osRelease := OsRelease{}
	scanner := bufio.NewScanner(file)
	var err error
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") {
			continue
		}
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		key, value, found := strings.Cut(line, "=")
		if !found {
			return nil, ErrCannotParseFile
		}
		switch key {
		case "ID":
			osRelease.Id, err = escape(value)
		case "ID_LIKE":
			if tmp, err2 := escape(value); err2 == nil {
				osRelease.IdLike = strings.Fields(tmp)
			} else {
				err = err2
			}
		case "VARIANT_ID":
			osRelease.VariantId, err = escape(value)
		case "VERSION_ID":
			osRelease.VersionId, err = escape(value)
		case "VERSION_CODENAME":
			osRelease.VersionCodename, err = escape(value)
		case "BUILD_ID":
			osRelease.BuildId, err = escape(value)
		case "IMAGE_ID":
			osRelease.ImageId, err = escape(value)
		case "IMAGE_VERSION":
			osRelease.ImageVersion, err = escape(value)
		default:
			err = nil
		}
		if err != nil {
			return &OsRelease{}, err
		}
	}
	if osRelease.Id == "" {
		osRelease.Id = "linux"
	}
	return &osRelease, nil
}

// Create OsRelease for the current operating system
func New() (*OsRelease, error) {
	// main file is /etc/os-release (symlinks followed)
	f, err := os.Open("/etc/os-release")
	if err != nil {
		// fallback on /usr/lib/os-release
		f, err = os.Open("/usr/lib/os-release")
		if err != nil {
			return nil, errors.Join(ErrNoOsReleaseFile, err)
		}
	}
	defer f.Close()
	return FromFile(f)
}
