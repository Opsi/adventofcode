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

func TestTwo(t *testing.T) {
	lines := []string{
		"LR",
		"11A = (11B, XXX)",
		"11B = (XXX, 11Z)",
		"11Z = (11B, XXX)",
		"22A = (22B, XXX)",
		"22B = (22C, 22C)",
		"22C = (22Z, 22Z)",
		"22Z = (22B, 22B)",
		"XXX = (XXX, XXX)",
	}
	got, err := Two(lines)
	require.NoError(t, err)
	assert.Equal(t, 6, got)
}
