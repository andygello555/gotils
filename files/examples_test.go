package files

import (
	"fmt"
	"io/ioutil"
	"os"
)

// Checking if temporary files and directories created are said to exist by the Exists function.
func ExampleExists() {
	// We create a temporary file and a temporary directory to test the function.
	tmpFile, _ := ioutil.TempFile("", "tmpFile")
	tmpDir, _ := ioutil.TempDir("", "tmpDir")

	// Checking tmpFile
	fmt.Println("Checking file (before removal):", Exists(tmpFile.Name()))
	// Removing the file should result in "false" from Exists
	_ = os.Remove(tmpFile.Name())
	fmt.Println("Checking file (after removal):", Exists(tmpFile.Name()))

	// Checking tmpDir
	fmt.Println("Checking directory (before removal):", Exists(tmpDir))
	// Removing the file should result in "false" from Exists
	_ = os.RemoveAll(tmpDir)
	fmt.Println("Checking directory (after removal):", Exists(tmpDir))

	// Output:
	// Checking file (before removal): true
	// Checking file (after removal): false
	// Checking directory (before removal): true
	// Checking directory (after removal): false
}

// Checking if a temporary file made exists and comparing outputs to a temporary directory.
func ExampleIsFile() {
	// We create a temporary file and a temporary directory to test the function.
	tmpFile, _ := ioutil.TempFile("", "tmpFile")
	tmpDir, _ := ioutil.TempDir("", "tmpDir")

	// Checking tmpFile
	fmt.Println("Checking file (before removal):", IsFile(tmpFile.Name()))
	// Removing the file should result in "false" from IsFile
	_ = os.Remove(tmpFile.Name())
	fmt.Println("Checking file (after removal):", IsFile(tmpFile.Name()))

	// Checking tmpDir (both should be false)
	fmt.Println("Checking directory (before removal):", IsFile(tmpDir))
	_ = os.RemoveAll(tmpDir)
	fmt.Println("Checking directory (after removal):", IsFile(tmpDir))

	// Output:
	// Checking file (before removal): true
	// Checking file (after removal): false
	// Checking directory (before removal): false
	// Checking directory (after removal): false
}

// Checking if a temporary directory made exists and comparing outputs to a temporary file.
func ExampleIsDir() {
	// We create a temporary file and a temporary directory to test the function.
	tmpFile, _ := ioutil.TempFile("", "tmpFile")
	tmpDir, _ := ioutil.TempDir("", "tmpDir")

	// Checking tmpDir
	fmt.Println("Checking directory (before removal):", IsDir(tmpDir))
	// Removing the directory should result in "false" from IsDir
	_ = os.RemoveAll(tmpDir)
	fmt.Println("Checking directory (after removal):", IsDir(tmpDir))

	// Checking tmpFile (both should be false)
	fmt.Println("Checking file (before removal):", IsDir(tmpFile.Name()))
	_ = os.Remove(tmpFile.Name())
	fmt.Println("Checking file (after removal):", IsDir(tmpFile.Name()))

	// Output:
	// Checking directory (before removal): true
	// Checking directory (after removal): false
	// Checking file (before removal): false
	// Checking file (after removal): false
}
