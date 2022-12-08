package fs

import (
	"fmt"
	"strings"
)

type Dir struct {
	Name     string
	Parent   *Dir
	Children map[string]*Dir
	Files    []File
}

func NewDir(name string) *Dir {
	return &Dir{
		Name:     name,
		Parent:   nil,
		Children: map[string]*Dir{},
		Files:    []File{},
	}
}

func (d *Dir) Size() int {
	size := 0
	for _, file := range d.Files {
		size += file.Size
	}

	for _, child := range d.Children {
		size += child.Size()
	}

	return size
}

func (d *Dir) WithParent(p *Dir) *Dir {
	d.Parent = p
	return d
}

func (d *Dir) String() string {
	s := fmt.Sprintf("- %s (dir, size=%d)\n", d.Name, d.Size())

	for _, file := range d.Files {
		s += fmt.Sprintf("  - %s (file, size=%d)\n", file.Name, file.Size)
	}

	for _, child := range d.Children {
		for _, line := range strings.Split(strings.TrimSpace(child.String()), "\n") {
			s += fmt.Sprintf("  %s\n", line)
		}
	}

	return s
}
