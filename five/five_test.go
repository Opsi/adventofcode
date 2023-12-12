package five

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOne(t *testing.T) {
	lines := []string{
		"seeds: 79 14 55 13",
		"seed-to-soil map:",
		"50 98 2",
		"52 50 48",
		"soil-to-fertilizer map:",
		"0 15 37",
		"37 52 2",
		"39 0 15",
		"fertilizer-to-water map:",
		"49 53 8",
		"0 11 42",
		"42 0 7",
		"57 7 4",
		"water-to-light map:",
		"88 18 7",
		"18 25 70",
		"light-to-temperature map:",
		"45 77 23",
		"81 45 19",
		"68 64 13",
		"temperature-to-humidity map:",
		"0 69 1",
		"1 0 69",
		"humidity-to-location map:",
		"60 56 37",
		"56 93 4",
	}
	got, err := One(lines)
	require.NoError(t, err)
	assert.Equal(t, 35, got)
}

func TestMapRangeApplyRange(t *testing.T) {
	cases := []struct {
		name          string
		got           Range
		wantApplied   []Range
		wantUnapplied []Range
	}{
		{
			name:          "complete overlap",
			got:           Range{0, 10},
			wantApplied:   []Range{{50, 10}},
			wantUnapplied: nil,
		}, {
			name:          "perfect left overlap",
			got:           Range{0, 4},
			wantApplied:   []Range{{50, 4}},
			wantUnapplied: nil,
		}, {
			name:          "perfect right overlap",
			got:           Range{2, 8},
			wantApplied:   []Range{{52, 8}},
			wantUnapplied: nil,
		}, {
			name:          "complete within",
			got:           Range{3, 3},
			wantApplied:   []Range{{53, 3}},
			wantUnapplied: nil,
		}, {
			name:          "complete outside",
			got:           Range{-5, 17},
			wantApplied:   []Range{{50, 10}},
			wantUnapplied: []Range{{-5, 5}, {10, 2}},
		}, {
			name:          "no overlap",
			got:           Range{20, 5},
			wantApplied:   nil,
			wantUnapplied: []Range{{20, 5}},
		}, {
			name:          "one overlap left",
			got:           Range{-1, 2},
			wantApplied:   []Range{{50, 1}},
			wantUnapplied: []Range{{-1, 1}},
		}, {
			name:          "one overlap right",
			got:           Range{9, 3},
			wantApplied:   []Range{{59, 1}},
			wantUnapplied: []Range{{10, 2}},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			mapRange := MapRange{
				From:  Range{0, 10},
				Delta: 50,
			}
			gotApplied, gotUnapplied := mapRange.ApplyRange(c.got)
			assert.ElementsMatch(t, c.wantApplied, gotApplied)
			assert.ElementsMatch(t, c.wantUnapplied, gotUnapplied)
		})
	}
}

func TestTwo(t *testing.T) {
	lines := []string{
		"seeds: 79 14 55 13",
		"seed-to-soil map:",
		"50 98 2",
		"52 50 48",
		"soil-to-fertilizer map:",
		"0 15 37",
		"37 52 2",
		"39 0 15",
		"fertilizer-to-water map:",
		"49 53 8",
		"0 11 42",
		"42 0 7",
		"57 7 4",
		"water-to-light map:",
		"88 18 7",
		"18 25 70",
		"light-to-temperature map:",
		"45 77 23",
		"81 45 19",
		"68 64 13",
		"temperature-to-humidity map:",
		"0 69 1",
		"1 0 69",
		"humidity-to-location map:",
		"60 56 37",
		"56 93 4",
	}
	got, err := Two(lines)
	require.NoError(t, err)
	assert.Equal(t, 46, got)
}
