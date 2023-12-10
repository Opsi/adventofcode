package one

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFirstDigitOne(t *testing.T) {
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
		got, err := firstDigitOne(c.line)
		require.NoError(t, err)
		assert.Equal(t, c.want, got)
	}
}

func TestLastDigitOne(t *testing.T) {
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
		got, err := lastDigitOne(c.line)
		require.NoError(t, err)
		assert.Equal(t, c.want, got)
	}
}

func TestLineValueOne(t *testing.T) {
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
		got, err := lineValueOne(c.line)
		require.NoError(t, err)
		assert.Equal(t, c.want, got)
	}
}

func TestOneExample(t *testing.T) {
	document := `
	1abc2
	pqr3stu8vwx
	a1b2c3d4e5f
	treb7uchet
	`
	cv, err := One(document)
	require.NoError(t, err)
	assert.Equal(t, 142, cv)
}

func TestFirstDigitTwo(t *testing.T) {
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
		{"two1nine", 2},
		{"eightwothree", 8},
		{"abcone2threexyz", 1},
		{"xtwone3four", 2},
		{"4nineeightseven2", 4},
		{"zoneight234", 1},
		{"7pqrstsixteen", 7},
		{"eightwo", 8},
	}
	for _, c := range cases {
		got, err := firstDigitTwo(c.line)
		require.NoError(t, err)
		assert.Equal(t, c.want, got)
	}
}

func TestLastDigitTwo(t *testing.T) {
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
		{"two1nine", 9},
		{"eightwothree", 3},
		{"abcone2threexyz", 3},
		{"xtwone3four", 4},
		{"4nineeightseven2", 2},
		{"zoneight234", 4},
		{"7pqrstsixteen", 6},
		{"eightwo", 2},
	}
	for _, c := range cases {
		got, err := lastDigitTwo(c.line)
		require.NoError(t, err)
		assert.Equal(t, c.want, got)
	}
}

func TestLineValueTwo(t *testing.T) {
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
		{"two1nine", 29},
		{"eightwothree", 83},
		{"abcone2threexyz", 13},
		{"xtwone3four", 24},
		{"4nineeightseven2", 42},
		{"zoneight234", 14},
		{"7pqrstsixteen", 76},
		{"eightwo", 82},
	}
	for _, c := range cases {
		got, err := lineValueTwo(c.line)
		require.NoError(t, err)
		assert.Equal(t, c.want, got)
	}
}

func TestTwoExample(t *testing.T) {
	document := `
	two1nine
	eightwothree
	abcone2threexyz
	xtwone3four
	4nineeightseven2
	zoneight234
	7pqrstsixteen
	`
	cv, err := Two(document)
	require.NoError(t, err)
	assert.Equal(t, 281, cv)
}
