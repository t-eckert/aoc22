package parser_test

import (
	"dayseven/fs"
	"dayseven/parser"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	cases := []struct {
		input    string
		expected fs.Dir
	}{
		{
			input:    `$ cd /`,
			expected: *fs.NewDir("/"),
		},
		{
			input: `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d`,
			expected: func() fs.Dir {
				d := *fs.NewDir("/")
				d.Children = map[string]*fs.Dir{
					"a": fs.NewDir("a"),
					"d": fs.NewDir("d"),
				}
				d.Files = []fs.File{
					fs.NewFile("14848514 b.txt"),
					fs.NewFile("8504156 c.dat"),
				}
				return d
			}(),
		},
	}

	for _, tc := range cases {
		t.Run(tc.input, func(t *testing.T) {
			actual := parser.Parse(strings.Split(tc.input, "\n"))

			if tc.expected.Name != actual.Name {
				t.Fail()
			}
		})
	}

}
