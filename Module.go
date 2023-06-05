/*---------------------------------------------------------------------------------
  Start your module
  Version
  Tidy
  Verify
  Vendor
  Get
  Test
---------------------------------------------------------------------------------*/

package main

func main() {
	//---------------------------------------------------------------------------------
	// Start your module

	//go mod init [module-path]

	//---------------------------------------------------------------------------------
	// Version

	//go version

	//---------------------------------------------------------------------------------
	// Tidy

	//go mod tidy [-e] [-v] [-go=version] [-compat=version]

	//The -e flag (added in Go 1.16) causes go mod vendor to attempt to proceed despite errors encountered while loading packages.
	//The -v flag causes go mod vendor to print the names of vendored modules and packages to standard error.

	//---------------------------------------------------------------------------------
	// Verify

	//go mod verify

	//---------------------------------------------------------------------------------
	// Vendor

	//go mod vendor [-e] [-v] [-o]

	//---------------------------------------------------------------------------------
	// Get

	//go get [-d] [-t] [-u] [build flags] [packages]

	//---------------------------------------------------------------------------------

}
