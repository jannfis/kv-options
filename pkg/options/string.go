package options

// Copyright (c) 2024 Jann Fischer <jann@mistrust.net>
// SPDX-License-Identifier: 0BSD

// String returns the string value for an option named n from the
// OptionSet. The returned string may be empty.
//
// If def is not nil, its value will be returned as a default value if the
// option n is not set in the OptionSet.
//
// An error will be returned if def is nil and the option n is not set in the
// OptionSet, the value of option n is nil or the value of option n is not of
// type string.
func (opts *Options) String(name string, def *string) (value string, err error) {
	opts.mu.RLock()
	defer opts.mu.RUnlock()
	o, ok := opts.options[name]
	if !ok {
		if def == nil {
			return "", ErrNotFound
		} else {
			return *def, nil
		}
	}
	if o.t != OptionTypeString {
		return "", ErrInvalidType
	}
	if o.v == nil {
		return "", ErrInvalidValue
	}
	r, ok := o.v.(*string)
	if !ok {
		return "", ErrInvalidValue
	}
	return *r, nil
}

func (opts *Options) ParseString(s string) error {
	k, v, err := splitKeyValue(s)
	if err != nil {
		return err
	}
	return opts.setString(k, v)
}

func (opts *Options) setString(k, v string) error {
	opts.setValue(k, OptionTypeString, &v)
	return nil
}
