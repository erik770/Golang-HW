package uniq

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateFlagsOK(t *testing.T) {
	res := ValidateFlags(Flags{
		count:      false,
		duplicate:  true,
		unique:     false,
		ignoreReg:  true,
		fieldsSkip: 2,
		skipChars:  43,
	})

	assert.Equal(t, nil, res)
}

func TestValidateFlagsERR(t *testing.T) {
	res := ValidateFlags(Flags{
		count:      true,
		duplicate:  true,
		unique:     false,
		ignoreReg:  true,
		fieldsSkip: 2,
		skipChars:  5,
	})

	assert.Equal(t, errors.New("NO VALIDE FLAGS"), res)
}
func TestIgnoringOptionsReg(t *testing.T) {
	res := IgnoringOption([]string{"qwerty", "qWeRtY", "qweRty"}, true, -1, -1)
	exp := []string{"qwerty", "qwerty", "qwerty"}

	assert.Equal(t, exp, res)
}

func TestIgnoringOptionsSkip(t *testing.T) {
	res := IgnoringOption([]string{"qwerty", "hello", "world"}, false, 1, 2)
	exp := []string{"llo", "world"}

	assert.Equal(t, exp, res)
}

func TestCreateCounterMap(t *testing.T) {
	res := createCounterMap([]string{"qwerty", "qwerty", "qwerty", "qwerty", "hi", "hello", "hello", "world"})
	exp := map[string]int{"qwerty": 3, "hi": 0, "hello": 1, "world": 0}

	assert.Equal(t, exp, res)
}

func TestUniqCount(t *testing.T) {
	res := Uniq([]string{"qwerty", "qwerty", "qwerty", "qwerty", "hi", "hello", "hello", "world"},
		Flags{
			count:      true,
			duplicate:  false,
			unique:     false,
			ignoreReg:  false,
			fieldsSkip: 0,
			skipChars:  0,
		})
	exp := []string{"4 qwerty", "1 hi", "2 hello", "1 world"}

	assert.Equal(t, exp, res)
}

func TestUniqDuplicate(t *testing.T) {
	res := Uniq([]string{"aabb", "aabb", "bbaa", "ccss", "ccss", "ddaa"},
		Flags{
			count:      false,
			duplicate:  true,
			unique:     false,
			ignoreReg:  false,
			fieldsSkip: 0,
			skipChars:  0,
		})
	exp := []string{"aabb", "ccss"}

	assert.Equal(t, exp, res)
}

func TestUniqUniq(t *testing.T) {
	res := Uniq([]string{"aabb", "aabb", "bbaa", "ccss", "ccss", "ddaa"},
		Flags{
			count:      false,
			duplicate:  false,
			unique:     true,
			ignoreReg:  false,
			fieldsSkip: 0,
			skipChars:  0,
		})
	exp := []string{"bbaa", "ddaa"}

	assert.Equal(t, exp, res)
}

func TestUniqNoflag(t *testing.T) {
	res := Uniq([]string{"aabb", "aabb", "bbaa", "ccss", "ccss", "aabb", "aabb", "ddaa"},
		Flags{
			count:      false,
			duplicate:  false,
			unique:     false,
			ignoreReg:  false,
			fieldsSkip: 0,
			skipChars:  0,
		})
	exp := []string{"aabb", "bbaa", "ccss", "aabb", "ddaa"}

	assert.Equal(t, exp, res)
}

func TestUniqMix(t *testing.T) {
	res := Uniq([]string{"aaBb", "Aabb", "bbaa", "ccss", "cCSs", "aabb", "AaBb", "ddaa"},
		Flags{
			count:      true,
			duplicate:  false,
			unique:     false,
			ignoreReg:  true,
			fieldsSkip: 1,
			skipChars:  3,
		})
	exp := []string{"1 b", "1 bbaa", "2 ccss", "2 aabb", "1 ddaa"}

	assert.Equal(t, exp, res)
}
