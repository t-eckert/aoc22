package fs

import (
	"strconv"
	"strings"
)

type File struct {
	Name string
	Size int
}

func NewFile(readout string) File {
	split := strings.Split(readout, " ")
	name := split[1]
	size, err := strconv.Atoi(split[0])
	if err != nil {
		panic(err.Error())
	}

	return File{
		Name: name,
		Size: size,
	}
}
