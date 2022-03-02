package uniq

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestValidateFlags(t *testing.T) {
	cases := map[string]struct {
		in       Flags
		expected error
	}{
		"OK flags": {
			in: Flags{
				count:      false,
				duplicate:  true,
				unique:     false,
				ignoreReg:  true,
				fieldsSkip: 2,
				skipChars:  43,
			},
			expected: nil,
		},
		"ERR flags": {
			in: Flags{
				count:      true,
				duplicate:  true,
				unique:     false,
				ignoreReg:  true,
				fieldsSkip: 2,
				skipChars:  5,
			},
			expected: errors.New("NO VALIDE FLAGS"),
		},
		"ERR flags neg numbers": {
			in: Flags{
				count:      true,
				duplicate:  true,
				unique:     false,
				ignoreReg:  true,
				fieldsSkip: -2,
				skipChars:  5,
			},
			expected: errors.New("NO VALIDE FLAGS"),
		},
	}

	for name, tc := range cases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			actual := ValidateFlags(tc.in)

			assert.Equal(t, tc.expected, actual)
		})
		log.Println("SUCCESS")
	}
}

func TestIgnoringOptionsReg(t *testing.T) {
	res := registerIgnoreOption([]string{"qwerty", "qWeRtY", "qweRty"})
	exp := []string{"qwerty", "qwerty", "qwerty"}

	assert.Equal(t, exp, res)
}

func TestIgnoringOptionsSkipFields(t *testing.T) {
	res := stringsIgnoreOption([]string{"qwerty", "hello", "world"}, 1)
	exp := []string{"hello", "world"}

	assert.Equal(t, exp, res)
}

func TestIgnoringOptionsSkipChars(t *testing.T) {
	res := charsIgnoreOption([]string{"qwerty", "hello", "world"}, 8)
	exp := []string{"llo", "world"}

	assert.Equal(t, exp, res)
}

func TestCreateCounterMap(t *testing.T) {
	res := createCounterMap([]string{"qwerty", "qwerty", "qwerty", "qwerty", "hi", "hello", "hello", "world"})
	exp := map[string]int{"qwerty": 3, "hi": 0, "hello": 1, "world": 0}

	assert.Equal(t, exp, res)
}

func TestUniq(t *testing.T) {
	type input struct {
		inputStrings []string
		flags        Flags
	}
	cases := map[string]struct {
		in       input
		expected []string
	}{
		"Uniq count option": {
			in: input{
				[]string{"qwerty", "qwerty", "qwerty", "qwerty", "hi", "hello", "hello", "world"},
				Flags{
					count:      true,
					duplicate:  false,
					unique:     false,
					ignoreReg:  false,
					fieldsSkip: 0,
					skipChars:  0,
				},
			},
			expected: []string{"4 qwerty", "1 hi", "2 hello", "1 world"},
		},
		"Uniq only duplicate option": {
			in: input{
				[]string{"aabb", "aabb", "bbaa", "ccss", "ccss", "ddaa"},
				Flags{
					count:      false,
					duplicate:  true,
					unique:     false,
					ignoreReg:  false,
					fieldsSkip: 0,
					skipChars:  0,
				},
			},
			expected: []string{"aabb", "ccss"},
		},
		"Uniq uniq option": {
			in: input{
				[]string{"aabb", "aabb", "bbaa", "ccss", "ccss", "ddaa"},
				Flags{
					count:      false,
					duplicate:  false,
					unique:     true,
					ignoreReg:  false,
					fieldsSkip: 0,
					skipChars:  0,
				},
			},
			expected: []string{"bbaa", "ddaa"},
		},
		"Uniq no flag option": {
			in: input{
				[]string{"aabb", "aabb", "bbaa", "ccss", "ccss", "aabb", "aabb", "ddaa"},
				Flags{
					count:      false,
					duplicate:  false,
					unique:     false,
					ignoreReg:  false,
					fieldsSkip: 0,
					skipChars:  0,
				},
			},
			expected: []string{"aabb", "bbaa", "ccss", "aabb", "ddaa"},
		},
		"Uniq count mix options": {
			in: input{
				[]string{"aaBb", "Aabb", "bbaa", "ccss", "cCSs", "aabb", "AaBb", "ddaa"},
				Flags{
					count:      true,
					duplicate:  false,
					unique:     false,
					ignoreReg:  true,
					fieldsSkip: 1,
					skipChars:  3,
				},
			},
			expected: []string{"1 b", "1 bbaa", "2 ccss", "2 aabb", "1 ddaa"},
		},
	}

	for name, tc := range cases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			actual := Uniq(tc.in.inputStrings, tc.in.flags)

			assert.Equal(t, tc.expected, actual)
		})
		log.Println("SUCCESS")
	}
}
