package options

// Copyright (c) 2024 Jann Fischer <jann@mistrust.net>
// SPDX-License-Identifier: 0BSD

import (
	"fmt"
	"strconv"
)

// Bool returns the bool value for an option named n from the OptionSet.
//
// If def is not nil, its value will be returned as a default value if the
// option n is not set in the OptionSet.
//
// An error will be returned if def is nil and the option n is not set in the
// OptionSet, the value of option n is nil or the value of option n is not of
// type boolean.
func (opts *Options) Bool(n string, def *bool) (value bool, err error) {
	opts.mu.RLock()
	defer opts.mu.RUnlock()
	o, ok := opts.options[n]
	if !ok {
		if def == nil {
			return false, ErrNotFound
		} else {
			return *def, nil
		}
	}
	if o.t != OptionTypeBool {
		return false, ErrInvalidType
	}
	if o.v == nil {
		return false, ErrInvalidValue
	}
	r, ok := o.v.(*bool)
	if !ok {
		return false, ErrInvalidValue
	}
	return *r, nil
}

// MustBool returns the boolean value for the option n. If no such option is
// found, or if the option is of a different type, then the default value def
// will be returned.
//
// If you want to check for potential errors, use Bool instead.
func (opts *Options) MustBool(n string, def bool) bool {
	v, _ := opts.Bool(n, &def)
	return v
}

// ParseBool will parse the string s into a new boolean option and store
// it in the OptionSet. The string s must be in format key=value, and the value
// part must be parseable by strconv.ParsePool. Whitespace will be ignored.
//
// If parsing and storing the key value pair was successful, returns nil.
// Otherwise, returns an error indicating what went wrong.
func (opts *Options) ParseBool(s string) error {
	k, v, err := splitKeyValue(s)
	if err != nil {
		return err
	}
	return opts.setBool(k, v)
}

func (opts *Options) setBool(k string, v string) error {
	bv, err := strconv.ParseBool(v)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrInvalidValue, err)
	}
	opts.setValue(k, OptionTypeBool, &bv)
	return nil
}
