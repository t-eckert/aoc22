package parser

import (
	"dayseven/fs"
	"strings"
)

func Parse(readout []string) *fs.Dir {
	// Every instruction set will being with a command.
	dir, _ := parseCommand(nil, readout)

	for dir.Parent != nil {
		dir = dir.Parent
	}

	return dir
}

func parse(dir *fs.Dir, readout []string) (*fs.Dir, []string) {
	if len(readout) == 0 {
		return dir, readout
	}

	switch readout[0][0] {
	case '$':
		return parseCommand(dir, readout)
	case 'd':
		return parseDir(dir, readout)
	default:
		return parseFile(dir, readout)
	}
}

func parseCommand(dir *fs.Dir, readout []string) (*fs.Dir, []string) {
	switch readout[0][:4] {
	case "$ cd":
		split := strings.Split(readout[0], " ")
		name := split[len(split)-1]
		if name == "/" {
			// This is the beginning of the readout.
			return parse(fs.NewDir("/"), readout[1:])
		} else if name == ".." {
			return parse(dir.Parent, readout[1:])
		} else {
			return parse(dir.Children[name], readout[1:])
		}
	case "$ ls":
		// No-op
		return parse(dir, readout[1:])
	}

	return parse(dir, readout[1:])
}

func parseDir(dir *fs.Dir, readout []string) (*fs.Dir, []string) {
	split := strings.Split(readout[0], " ")
	name := split[1]
	dir.Children[name] = fs.NewDir(name).WithParent(dir)
	return parse(dir, readout[1:])
}

func parseFile(dir *fs.Dir, readout []string) (*fs.Dir, []string) {
	dir.Files = append(dir.Files, fs.NewFile(readout[0]))
	return parse(dir, readout[1:])
}
