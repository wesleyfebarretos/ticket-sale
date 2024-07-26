package test_init

import (
	"log"
	"os"
	"path"
	"runtime"
)

/*
   By default, running go test executes a package’s tests with the current working
   directory set to that package’s path.

   If you execute an `os.Getwd()` from within a test function, you'll see that
   your working directory is the package directory. The same directory is set when
   starting the configuration (regardless of the template).

   There are two ways here:

   1. Use relative paths to testdata in tests (not quite the correct and convenient way, but it works).
   2. Use a special trick with Go's runtime package.

   The trick is actually very simple and is to get the current executing and add `..` to the project root.

   After that, just import the package into any of the test files:

   package main_test

   import (
   _ "project/testing_init"
   )
*/

func init() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "../../../..")
	os.Chdir(dir)
	log.Printf("Setting test root folder to: %v\n", dir)
}
