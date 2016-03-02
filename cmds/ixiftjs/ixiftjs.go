//
// ixiftjs.go - A JavaScript wrapper around the ixift package. It is designed to be helpful in developing
// and exploring IxIF related services.
//
// @author R. S. Doiel, <rsdoiel@caltech.edu>
//
//
// Copyright (c) 2016, Caltech
// All rights not granted herein are expressly reserved by Caltech.
//
// Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.
//
// 3. Neither the name of the copyright holder nor the names of its contributors may be used to endorse or promote products derived from this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
package main

import (
	"flag"
	"fmt"
	"os"

	// Caltech Library Packages
	"github.com/caltechlibrary/ixift"
)

var (
	showHelp    bool
	showVersion bool
)

func init() {
	flag.BoolVar(&showHelp, "h", showHelp, "display help message")
	flag.BoolVar(&showVersion, "v", showVersion, "display version information")
}

func main() {
	flag.Parse()

	if showHelp == true {
		fmt.Println(`
 USAGE: ixiftjs [OPTIONS] [JAVASCRIPT_FILES]
		`)
		flag.PrintDefaults()
		fmt.Printf("\nVersion %s\n", ixift.Version)
		os.Exit(0)
	}

	if showVersion == true {
		fmt.Printf("Version %s\n", ixift.Version)
		os.Exit(0)
	}

	jsArgs := flag.Args()
	os.Exit(ixift.Repl(jsArgs))
}
