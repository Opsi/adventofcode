package adventofcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFirstDigit(t *testing.T) {
	cases := []struct {
		line string
		want int
	}{
		{"a3", 3},
		{"4a", 4},
		{"a5a", 5},
		{"67", 6},
		{"9", 9},
		{"1abc2", 1},
		{"pqr3stu8vwx", 3},
		{"a1b2c3d4e5f", 1},
		{"treb7uchet", 7},
	}
	for _, c := range cases {
		got, err := firstDigit(c.line)
		require.NoError(t, err)
		assert.Equal(t, c.want, got)
	}
}

func TestLastDigit(t *testing.T) {
	cases := []struct {
		line string
		want int
	}{
		{"a3", 3},
		{"4a", 4},
		{"a5a", 5},
		{"67", 7},
		{"9", 9},
		{"1abc2", 2},
		{"pqr3stu8vwx", 8},
		{"a1b2c3d4e5f", 5},
		{"treb7uchet", 7},
	}
	for _, c := range cases {
		got, err := lastDigit(c.line)
		require.NoError(t, err)
		assert.Equal(t, c.want, got)
	}
}

func TestLineValue(t *testing.T) {
	cases := []struct {
		line string
		want int
	}{
		{"1", 11},
		{"a3", 33},
		{"4a", 44},
		{"a5a", 55},
		{"67", 67},
		{"9", 99},
		{"1abc2", 12},
		{"pqr3stu8vwx", 38},
		{"a1b2c3d4e5f", 15},
		{"treb7uchet", 77},
	}
	for _, c := range cases {
		got, err := lineValue(c.line)
		require.NoError(t, err)
		assert.Equal(t, c.want, got)
	}
}

func TestOneOneExample(t *testing.T) {
	document := `
	1abc2
	pqr3stu8vwx
	a1b2c3d4e5f
	treb7uchet
	`
	cv, err := OneOne(document)
	require.NoError(t, err)
	assert.Equal(t, 142, cv)
}
