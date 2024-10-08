package licensing_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/aquasecurity/trivy/pkg/licensing"
)

func TestSplitLicenses(t *testing.T) {
	tests := []struct {
		name     string
		license  string
		licenses []string
	}{
		{
			"simple list comma-separated",
			"GPL-1+,GPL-2",
			[]string{
				"GPL-1+",
				"GPL-2",
			},
		},
		{
			"simple list comma-separated",
			"GPL-1+,GPL-2,GPL-3",
			[]string{
				"GPL-1+",
				"GPL-2",
				"GPL-3",
			},
		},
		{
			"3 licenses 'or'-separated",
			"GPL-1+ or Artistic or Artistic-dist",
			[]string{
				"GPL-1+",
				"Artistic",
				"Artistic-dist",
			},
		},
		{
			"two licenses _or_ separated",
			"LGPLv3+_or_GPLv2+",
			[]string{
				"LGPLv3+",
				"GPLv2+",
			},
		},
		{
			"licenses `and`-separated",
			"BSD-3-CLAUSE and GPL-2",
			[]string{
				"BSD-3-CLAUSE",
				"GPL-2",
			},
		},
		{
			"three licenses and/or separated",
			"GPL-1+ or Artistic, and BSD-4-clause-POWERDOG",
			[]string{
				"GPL-1+",
				"Artistic",
				"BSD-4-clause-POWERDOG",
			},
		},
		{
			"two licenses with version",
			"Apache License,Version 2.0, OSET Public License version 2.1",
			[]string{
				"Apache License, Version 2.0",
				"OSET Public License version 2.1",
			},
		},
		{
			"the license starts with `ver`",
			"verbatim and BSD-4-clause",
			[]string{
				"verbatim",
				"BSD-4-clause",
			},
		},
		{
			"the license with `or later`",
			"GNU Affero General Public License v3 or later (AGPLv3+)",
			[]string{
				"GNU Affero General Public License v3 or later (AGPLv3+)",
			},
		},
		{
			"Python license exceptions",
			"GNU Library or Lesser General Public License (LGPL), Common Development and Distribution License 1.0 (CDDL-1.0), Historical Permission Notice and Disclaimer (HPND)",
			[]string{
				"GNU Library or Lesser General Public License (LGPL)",
				"Common Development and Distribution License 1.0 (CDDL-1.0)",
				"Historical Permission Notice and Disclaimer (HPND)",
			},
		},
		{
			name:    "License text",
			license: "* Permission to use this software in any way is granted without",
			licenses: []string{
				"text://* Permission to use this software in any way is granted without",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := licensing.SplitLicenses(tt.license)
			assert.Equal(t, tt.licenses, res)
		})
	}
}
