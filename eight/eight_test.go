package eight

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func exampleLines() []string {
	return []string{
		"RL",
		"AAA = (BBB, CCC)",
		"BBB = (DDD, EEE)",
		"CCC = (ZZZ, GGG)",
		"DDD = (DDD, DDD)",
		"EEE = (EEE, EEE)",
		"GGG = (GGG, GGG)",
		"ZZZ = (ZZZ, ZZZ)",
	}
}

func TestOne(t *testing.T) {
	got, err := One(exampleLines())
	require.NoError(t, err)
	assert.Equal(t, 2, got)
}

func TestOneExtra(t *testing.T) {
	lines := []string{
		"LLR",
		"AAA = (BBB, BBB)",
		"BBB = (AAA, ZZZ)",
		"ZZZ = (ZZZ, ZZZ)",
	}
	got, err := One(lines)
	require.NoError(t, err)
	assert.Equal(t, 6, got)
}
