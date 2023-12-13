package six

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOne(t *testing.T) {
	lines := []string{
		"Time:      7  15   30",
		"Distance:  9  40  200",
	}
	got, err := One(lines)
	require.NoError(t, err)
	assert.Equal(t, 288, got)
}

func TestWaysToWin(t *testing.T) {
	cases := []struct {
		race Race
		want int
	}{
		{Race{7, 9}, 4},
		{Race{15, 40}, 8},
		{Race{30, 200}, 9},
		{Race{71530, 940200}, 71503},
	}
	for _, c := range cases {
		got := c.race.WaysToWin()
		assert.Equal(t, c.want, got)
	}
}
