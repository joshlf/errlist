// Copyright 2012 The Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/joshlf13/errlist"
	"os"
	"strconv"
)

func main() {
	var e *errlist.Errlist

	// Create every other file so half
	// will exist and half will not
	for i := 0; i < 10; i += 2 {
		_, err := os.Stat(strconv.Itoa(i))
		if os.IsNotExist(err) {
			_, err := os.Create(strconv.Itoa(i))

			// nil errors are ignored;
			// adding them does nothing
			e = e.AddError(err)
		}
	}

	for i := 0; i < 5; i++ {
		_, err := os.Open(strconv.Itoa(i))

		// nil errors are ignored;
		// adding them does nothing
		e = e.AddError(err)
	}

	for i := 5; i < 10; i++ {
		_, err := os.Open(strconv.Itoa(i))
		if err != nil {
			// You can add by string as well
			e = e.AddString(err.Error())
		} else {
			// empty error strings are ignored; 
			// adding them does nothing
			e = e.AddString("")
		}
	}

	fmt.Printf("%v errors\n\n", e.Num())
	fmt.Println("Here they are printed directly:")
	fmt.Println(e)
	fmt.Println()
	fmt.Println("Here they are printed from a slice:")
	sl := e.Slice()
	for _, e := range sl {
		fmt.Println(e)
	}
}
