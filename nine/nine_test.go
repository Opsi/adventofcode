package nine

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOne(t *testing.T) {
	lines := []string{
		"0 3 6 9 12 15",
		"1 3 6 10 15 21",
		"10 13 16 21 30 45",
	}
	cv, err := One(lines)
	require.NoError(t, err)
	assert.Equal(t, 114, cv)
}

func TestTwo(t *testing.T) {
	lines := []string{
		"0 3 6 9 12 15",
		"1 3 6 10 15 21",
		"10 13 16 21 30 45",
	}
	cv, err := Two(lines)
	require.NoError(t, err)
	assert.Equal(t, 2, cv)
}

func TestLineValueOne(t *testing.T) {
	cases := []struct {
		line string
		want int
	}{
		{"0 3 6 9 12 15", 18},
		{"1 3 6 10 15 21", 28},
		{"10 13 16 21 30 45", 68},
	}
	for _, c := range cases {
		got, err := lineValueOne(c.line)
		require.NoError(t, err)
		assert.Equal(t, c.want, got)
	}
}

func TestNextSequence(t *testing.T) {
	cases := []struct {
		sequence string
		want     string
	}{
		{"0 3 6 9 12 15", "3 3 3 3 3"},
		{"3 3 3 3 3", "0 0 0 0"},
		{"1 3 6 10 15 21", "2 3 4 5 6"},
		{"2 3 4 5 6", "1 1 1 1"},
		{"1 1 1 1", "0 0 0"},
		{"10 13 16 21 30 45 68", "3 3 5 9 15 23"},
		{"3 3 5 9 15 23", "0 2 4 6 8"},
		{"0 2 4 6 8", "2 2 2 2"},
		{"2 2 2 2", "0 0 0"},
	}
	for _, c := range cases {
		asInts, err := toInts(c.sequence)
		require.NoError(t, err)
		got, err := nextSequence(asInts)
		require.NoError(t, err)
		gotStr := asString(got)
		assert.Equal(t, c.want, gotStr)
	}
}
