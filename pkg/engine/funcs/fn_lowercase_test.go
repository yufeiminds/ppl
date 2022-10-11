// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

package funcs

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLowercase(t *testing.T) {
	cases := []struct {
		name, pl, in string
		expected     interface{}
		fail         bool
		outkey       string
	}{
		{
			name: "normal",
			pl: `json(_, a.third)
lowercase(a.third)
`,
			in:       `{"a":{"first":2.3,"second":2,"third":"aBC","forth":true},"age":47}`,
			expected: "abc",
			outkey:   "a.third",
		},
		{
			name: "normal",
			pl: `json(_, a.third)
lowercase(a.third)
`,
			in:       `{"a":{"first":2.3,"second":2,"third":"aBC","forth":true,"age":"WWW"},"age":"wWW"}`,
			expected: nil,
			outkey:   "a.age",
		},
		{
			name: "normal",
			pl: `json(_, a.third)
lowercase(a.third)
`,
			in:       `{"a":{"first":2.3,"second":2,"third":"aBC","forth":true,"age":"WWW"},"age":"wWW"}`,
			expected: nil,
			outkey:   "a.forth",
		},
		{
			name: "normal",
			pl: `json(_, a.third)
lowercase(a.third)
`,
			in:       `{"a":{"first":"222SSd","second":2,"third":"aBC","forth":true,"age":"WWW"},"age":"wWW"}`,
			expected: nil,
			outkey:   "a.first",
		},
		{
			name: "normal",
			pl: `json(_, a.first)
lowercase(a.first)
`,
			in:       `{"a":{"first":"SSd","second":2,"third":"aBC","forth":true,"age":"WWW"},"age":"wWW"}`,
			expected: "ssd",
			outkey:   "a.first",
		},
		{
			name: "normal",
			pl: `json(_, age)
lowercase(age)
`,
			in:       `{"a":{"first":"SSd","second":2,"third":"aBC","forth":true,"age":"WWW"},"age":"wWW"}`,
			expected: "www",
			outkey:   "age",
		},
	}
	for idx, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			runner, err := NewTestingRunner(tc.pl)
			if err != nil {
				if tc.fail {
					t.Logf("[%d]expect error: %s", idx, err)
				} else {
					t.Errorf("[%d] failed: %s", idx, err)
				}
				return
			}

			_, _, f, _, _, err := runScript(runner, "test", nil, map[string]interface{}{
				"message": tc.in,
			}, time.Now())

			assert.Equal(t, nil, err)

			t.Log(f)
			v := f[tc.outkey]
			// assert.Equal(t, nil, err)
			assert.Equal(t, tc.expected, v)

			t.Logf("[%d] PASS", idx)
		})
	}
}