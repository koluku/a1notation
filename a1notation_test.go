package a1notation

import (
	"fmt"
	"testing"
)

func TestFromIndex(t *testing.T) {
	cases := []struct {
		row  int
		col  int
		want string
	}{
		{
			row:  0,
			col:  0,
			want: "A1",
		},
		{
			row:  2,
			col:  99,
			want: "CV3",
		},
	}
	for _, tt := range cases {
		tt := tt
		t.Run(fmt.Sprintf("row: %d, col: %d: %s", tt.row, tt.col, tt.want), func(t *testing.T) {
			t.Parallel()
			got := FromIndex(tt.row, tt.col)
			if got != tt.want {
				t.Errorf("want = %s, but got = %s", tt.want, got)
			}
		})
	}

}

func TestColumnLetterFrom(t *testing.T) {
	cases := []struct {
		col  int
		want string
	}{
		{
			col:  0,
			want: "A",
		},
		{
			col:  25,
			want: "Z",
		},
		{
			col:  26,
			want: "AA",
		},
		{
			col:  27,
			want: "AB",
		},
		{
			col:  99,
			want: "CV",
		},
		{
			col:  701,
			want: "ZZ",
		},
		{
			col:  702,
			want: "AAA",
		},
		{
			col:  18277,
			want: "ZZZ",
		},
		{
			col:  18278,
			want: "AAAA",
		},
	}
	for _, tt := range cases {
		tt := tt
		t.Run(fmt.Sprintf("col: %d: %s", tt.col, tt.want), func(t *testing.T) {
			t.Parallel()
			got := ColumnLetterFrom(tt.col)
			if got != tt.want {
				t.Errorf("want = %s, but got = %s", tt.want, got)
			}
		})
	}

}

func TestToIndex(t *testing.T) {
	cases := []struct {
		alpha   string
		wantRow int
		wantCol int
	}{
		{
			alpha:   "A1",
			wantRow: 0,
			wantCol: 0,
		},
		{
			alpha:   "Z2",
			wantRow: 1,
			wantCol: 25,
		},
		{
			alpha:   "AA12",
			wantRow: 11,
			wantCol: 26,
		},
		{
			alpha:   "AB78",
			wantRow: 77,
			wantCol: 27,
		},
		{
			alpha:   "CV100",
			wantRow: 99,
			wantCol: 99,
		},
		{
			alpha:   "ZZ1001",
			wantRow: 1000,
			wantCol: 701,
		},
		{
			alpha:   "AAA23",
			wantRow: 22,
			wantCol: 702,
		},
		{
			alpha:   "ZZZ655",
			wantRow: 654,
			wantCol: 18277,
		},
		{
			alpha:   "AAAA64",
			wantRow: 63,
			wantCol: 18278,
		},
	}
	for _, tt := range cases {
		tt := tt
		t.Run(fmt.Sprintf("alpha: %s: %d, %d", tt.alpha, tt.wantRow, tt.wantCol), func(t *testing.T) {
			t.Parallel()
			gotRow, gotCol := ToIndex(tt.alpha)
			if gotRow != tt.wantRow || gotCol != tt.wantCol {
				t.Errorf("want = %d, %d, but got = %d, %d", tt.wantRow, tt.wantCol, gotRow, gotCol)
			}
		})
	}

}
