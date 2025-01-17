/*
 * This file is subject to the terms and conditions defined in
 * file 'LICENSE.md', which is part of this source code package.
 */

package main

import (
	"fmt"
	"os"

	"github.com/unidoc/unihtml"
	"github.com/unidoc/unihtml/selector"
)

// This example is based on the html file along with embedded javascript that waits for the dom loaded and
// subsequently renders the content of the graphs.
// By default unihtml rendering system, before rendering the PDF is waiting for the event dom loaded and then renders the PDF.
// In order to render that document correctly the graph.html should wait for specific selectors that are used by the Javascript.
// It is allowed by the `WaitVisible` or `WaitReady` methods of the unihtml document which takes regular html selector
// i.e. for id's it would be `#example` and classes `.example` the second parameter defines how to match given selector.
func main() {
	if len(os.Args) != 2 {
		fmt.Println("Err: provided invalid arguments. No UniHTML server path provided")
		os.Exit(1)
	}

	// Establish connection with the UniHTML Server.
	if err := unihtml.Connect(os.Args[1]); err != nil {
		fmt.Printf("Err:  Connect failed: %v\n", err)
		os.Exit(1)
	}

	// Read the content of the graph.html file and load it to the conversion.
	htmlDocument, err := unihtml.NewDocument("graph.html")
	if err != nil {
		fmt.Printf("Err: NewDocument failed: %v\n", err)
		os.Exit(1)
	}

	// We need to wait for the highcharts-root class nodes to be visible before rendering the PDF document.
	htmlDocument.WaitVisible(".highcharts-root", selector.ByQueryAll)

	// Write the result file to PDF.
	if err = htmlDocument.WriteToFile("graph.pdf"); err != nil {
		fmt.Printf("Err: %v\n", err)
		os.Exit(1)
	}
}
