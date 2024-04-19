package options

// Copyright (c) 2024 Jann Fischer <jann@mistrust.net>
// SPDX-License-Identifier: 0BSD

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ParseBool(t *testing.T) {
	opts := New()
	t.Run("Parse valid bool options", func(t *testing.T) {
		err := opts.ParseBool("Foo=true")
		assert.NoError(t, err)
		err = opts.ParseBool("Bar = false")
		assert.NoError(t, err)
	})
	t.Run("Parse invalid bool options", func(t *testing.T) {
		err := opts.ParseBool("Invalid")
		assert.Error(t, err)
		err = opts.ParseBool("Invalid=invalid")
		assert.Error(t, err)
		err = opts.ParseBool("Invalid=")
		assert.Error(t, err)
		err = opts.ParseBool("Invalid==true")
		assert.Error(t, err)
	})
	t.Run("Get valid bool options", func(t *testing.T) {
		v, err := opts.Bool("Foo", nil)
		assert.NoError(t, err)
		assert.True(t, v)
		v, err = opts.Bool("Foo", Ptr(false))
		assert.NoError(t, err)
		assert.True(t, v)
		v, err = opts.Bool("Bar", nil)
		assert.NoError(t, err)
		assert.False(t, v)
		v, err = opts.Bool("Bar", Ptr(true))
		assert.NoError(t, err)
		assert.False(t, v)
	})
	t.Run("Get invalid bool options", func(t *testing.T) {
		_, err := opts.Bool("DoesNotExist", nil)
		assert.ErrorIs(t, err, ErrNotFound)
		v, err := opts.Bool("DoesNotExist", Ptr(true))
		assert.NoError(t, err)
		assert.True(t, v)
		v, err = opts.Bool("DoesNotExist", Ptr(false))
		assert.NoError(t, err)
		assert.False(t, v)
		opts.options["InvalidType"] = Option{n: "InvalidType", t: OptionTypeString}
		_, err = opts.Bool("InvalidType", nil)
		assert.ErrorIs(t, err, ErrInvalidType)
		opts.options["NilValue"] = Option{n: "NilValue", t: OptionTypeBool}
		_, err = opts.Bool("NilValue", nil)
		assert.ErrorIs(t, err, ErrInvalidValue)
		opts.options["InvalidValue"] = Option{n: "InvalidValue", t: OptionTypeBool, v: "Foo"}
		_, err = opts.Bool("InvalidValue", nil)
		assert.ErrorIs(t, err, ErrInvalidValue)
	})
}
