package seven

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func exampleLines() []string {
	return []string{
		"32T3K 765",
		"T55J5 684",
		"KK677 28",
		"KTJJT 220",
		"QQQJA 483",
	}
}

func TestOne(t *testing.T) {
	got, err := One(exampleLines())
	require.NoError(t, err)
	assert.Equal(t, 6440, got)
}

func TestTwo(t *testing.T) {
	got, err := Two(exampleLines())
	require.NoError(t, err)
	assert.Equal(t, 5905, got)
}
