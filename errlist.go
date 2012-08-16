// Copyright 2012 The Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package errlist contains a type compatible with the error interface which handles lists of
// errors. All of the methods in this package are nil-safe; that is, calling them on nil
// pointers is expected behavior.
package errlist

import (
	"errors"
)

// An error type which supports lists of errors.
type Errlist struct {
	hd  *errnode
	tl  *errnode
	num int
}

type errnode struct {
	error
	next *errnode
}

// Create a new error list starting with
// an error created from e. If an empty
// error string is provided, a nil
// pointer is returned.
func NewString(e string) *Errlist {
	if e == "" {

	}
	var erl Errlist
	erl.hd = &errnode{errors.New(e), nil}
	erl.tl = erl.hd
	erl.num = 1
	return &erl
}

// Create a new error list starting
// with e. If a nil error is provided,
// a nil pointer is returned.
func NewError(e error) *Errlist {
	if e == nil {
		return nil
	}
	var erl Errlist
	erl.hd = &errnode{e, nil}
	erl.tl = erl.hd
	erl.num = 1
	return &erl
}

// Create an error from e and append 
// it to the error list, or if the
// list is nil, create a new list
// with the error as its first element.
// In either case, return the resultant list.
// If e is an empty error string, do not
// append an error. If AddString was called
// on a nil list and with an emtpy string
// as the argument, it returns a nil pointer.
func (erl *Errlist) AddString(e string) *Errlist {
	if erl == nil {
		return NewString(e)
	}
	if e == "" {
		return erl
	}
	ern := new(errnode)
	ern.error = errors.New(e)
	erl.tl.next = ern
	erl.tl = ern
	erl.num++
	return erl

}

// Append e to the error list,
// or if the list is nil, create
// a new list with e as its first
// element. In either case, return
// the resultant list. If e is nil,
// do not append an error. If
// AddError was called on a nil
// list and with a nil argument,
// it returns a nil pointer.
func (erl *Errlist) AddError(e error) *Errlist {
	if erl == nil {
		return NewError(e)
	}
	if e == nil {
		return erl
	}
	ern := new(errnode)
	ern.error = e
	erl.tl.next = ern
	erl.tl = ern
	erl.num++
	return erl
}

// Return a string consisting of
// each error in the list printed
// and separated by newlines, or
// an empty string if called on
// a nil pointer.
func (erl *Errlist) Error() string {
	out := ""
	if erl == nil {
		return out
	}
	for n := erl.hd; n != nil; n = n.next {
		out += n.error.Error() + "\n"
	}
	return out[:len(out)-1]
}

// Return the errors as a slice.
// If called on a nil pointer,
// returns an empty slice.
func (erl *Errlist) Slice() []error {
	if erl == nil {
		return make([]error, 0)
	}
	esl := make([]error, erl.num)
	for i, n := 0, erl.hd; i < erl.num; i, n = (i + 1), n.next {
		esl[i] = n.error
	}
	return esl
}

// Return the number of errors
// in the list, or 0 if called
// on a nil pointer.
func (erl *Errlist) Num() int {
	if erl == nil {
		return 0
	}
	return erl.num
}

// Err returns an error equivalent 
// to this error list. If the list 
// is empty, Err returns nil. 
func (erl *Errlist) Err() error {
	if erl == nil {
		return nil
	}
	if erl.num == 1 {
		return erl.hd.error
	}
	return erl
}
