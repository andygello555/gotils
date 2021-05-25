// A convenience wrapper for file checks and operations using paths.
package files

import (
	"os"
)

// Wrapper for finding the mode of a file.
//
// Returns the mode of the file (if it exists) and whether or not the path exists.
func exists(path string) (mode os.FileMode, exists bool) {
	fi, err := os.Stat(path)
	if err != nil {
		return mode, !os.IsNotExist(err)
	}
	return fi.Mode(), !os.IsNotExist(err)
}

// Returns whether the given path exists regardless of what mode the pointed to file has.
func Exists(path string) bool {
	_, exists := exists(path)
	return exists
}

// Checks if the given path exists and whether the path points to a file.
//
// Returns true if the path exists and the path points to a file, otherwise false.
func IsFile(path string) bool {
	mode, exists := exists(path)
	return exists && !mode.IsDir()
}

// Checks if the given path exists a directory.
//
// Returns true if the path given exists and it points to a directory, false otherwise.
func IsDir(path string) bool {
	mode, exists := exists(path)
	return mode.IsDir() && exists
}
