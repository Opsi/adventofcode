package three

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/exp/maps"
)

func TestParseField(t *testing.T) {
	lines := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}
	field, err := parseField(lines)
	require.NoError(t, err)
	assert.Zero(t, field.Cells[0][3])
	assert.Zero(t, field.Cells[9][9])
	assert.Negative(t, field.Cells[1][3])
	assert.Negative(t, field.Cells[4][3])
	assert.Negative(t, field.Cells[8][5])
	assert.Len(t, field.IDValues, 10)
	values := maps.Values(field.IDValues)
	assert.Contains(t, values, 467)
	assert.Contains(t, values, 114)
	assert.Contains(t, values, 35)
	assert.Contains(t, values, 633)
	assert.Contains(t, values, 617)
	assert.Contains(t, values, 58)
	assert.Contains(t, values, 592)
	assert.Contains(t, values, 755)
	assert.Contains(t, values, 664)
	assert.Contains(t, values, 598)

}

func TestOne(t *testing.T) {
	lines := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}
	got, err := One(lines)
	require.NoError(t, err)
	assert.Equal(t, 4361, got)
}
