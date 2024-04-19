package options

// Copyright (c) 2024 Jann Fischer <jann@mistrust.net>
// SPDX-License-Identifier: 0BSD

import (
	"errors"
	"fmt"
	"strings"
	"sync"
)

type OptionType int

var ErrNotFound = errors.New("option not found")
var ErrForbidden = errors.New("option not allowed")
var ErrInvalidType = errors.New("wrong option type")
var ErrInvalidValue = errors.New("invalid value in option")
var ErrParse = errors.New("could not parse option")

const (
	OptionTypeBool    OptionType = 1
	OptionTypeString  OptionType = 2
	OptionTypeNumeric OptionType = 3
)

type Option struct {
	n string     // Name of the value
	t OptionType // Type of the value
	v any        // The value itself
}

type Options struct {
	valid   map[string]Option
	options map[string]Option
	mu      sync.RWMutex
}

// New creates a new OptionSet and returns a pointer to it
func New() *Options {
	o := &Options{
		valid:   make(map[string]Option),
		options: make(map[string]Option),
	}
	return o
}

// WithOption adds a valid option named n of type t to the OptionSet
func (opts *Options) WithOption(n string, t OptionType) *Options {
	opts.mu.Lock()
	defer opts.mu.Unlock()
	opts.valid[n] = Option{n: n, t: t}
	return opts
}

// Parse parses key/value pairs in the slice options. Each element in the slice
// needs to be present in key=value format.
//
// Parse will only parse known options that have been declared using WithOption
// before. If it encounters an unknown option, it will abort parsing and return
// an error.
func (opts *Options) Parse(options []string) error {
	for _, s := range options {
		k, v, err := splitKeyValue(s)
		if err != nil {
			return err
		}
		opts.mu.RLock()
		o, ok := opts.valid[k]
		opts.mu.RUnlock()
		if !ok {
			return fmt.Errorf("option %s: %w", s, ErrForbidden)
		}
		switch o.t {
		case OptionTypeBool:
			err = opts.setBool(k, v)
		case OptionTypeString:
			err = opts.setString(k, v)
		default:
			err = fmt.Errorf("option %s: %w", s, ErrInvalidType)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func splitKeyValue(s string) (key string, value string, err error) {
	kv := strings.SplitN(s, "=", 2)
	if len(kv) != 2 {
		return "", "", fmt.Errorf("%w: not a valid key-value string: %s", ErrParse, s)
	}
	key = strings.TrimSpace(kv[0])
	value = strings.TrimSpace(kv[1])
	return key, value, nil
}

func (opts *Options) setValue(n string, t OptionType, v any) {
	opts.mu.Lock()
	defer opts.mu.Unlock()
	opts.options[n] = Option{n: n, t: t, v: v}
}

// Ptr is helper method to return a pointer to variable V of arbitrary type T
func Ptr[T any](v T) *T {
	return &v
}
