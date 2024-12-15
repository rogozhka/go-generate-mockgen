package example

import (
	"os"
)

// It seems there is a bug in the mockgen -mock_names option,
// which is ignored if there are multiple directives/interfaces in the file.
//
// Please, do place each dependency in a separate file
// if you want to get MockCamelCaseName instead of MockcamelCaseName.

// fileOperations describes imaginary fs wrapper.
// this case explains how to use prebuild docker image.
//
//go:generate docker run --rm -v ${PWD}:/w rogozhka/go-generate-mockgen -line=$GOLINE -source=$GOFILE -package=$GOPACKAGE
type fileOperations interface {
	Stat(path string) (os.FileInfo, error)
	Remove(path string) error
	RemoveAll(path string) error
	MkdirAll(path string, perm os.FileMode) error
}
