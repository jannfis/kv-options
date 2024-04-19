package options

// Copyright (c) 2024 Jann Fischer <jann@mistrust.net>
// SPDX-License-Identifier: 0BSD

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ParseString(t *testing.T) {
	opts := New()
	err := opts.ParseString("Foo=true")
	assert.NoError(t, err)
	err = opts.ParseString("Bar = false")
	assert.NoError(t, err)
	err = opts.ParseString("Invalid=")
	assert.NoError(t, err)
	err = opts.ParseString("Invalid==true")
	assert.NoError(t, err)

	t.Run("Parse invalid string options", func(t *testing.T) {
		opts := New()
		err := opts.ParseString("Invalid")
		assert.Error(t, err)
		err = opts.ParseString("")
		assert.Error(t, err)
	})
	t.Run("Get valid string options", func(t *testing.T) {
		v, err := opts.String("Foo", nil)
		assert.NoError(t, err)
		assert.Equal(t, "true", v)
		v, err = opts.String("Foo", Ptr("bar"))
		assert.NoError(t, err)
		assert.Equal(t, "true", v)
		v, err = opts.String("Bar", nil)
		assert.NoError(t, err)
		assert.Equal(t, "false", v)
		v, err = opts.String("Bar", Ptr("true"))
		assert.NoError(t, err)
		assert.Equal(t, "false", v)
	})
	t.Run("Get invalid string options", func(t *testing.T) {
		_, err := opts.String("DoesNotExist", nil)
		assert.ErrorIs(t, err, ErrNotFound)
		v, err := opts.String("DoesNotExist", Ptr("true"))
		assert.NoError(t, err)
		assert.Equal(t, "true", v)
		v, err = opts.String("DoesNotExist", Ptr("false"))
		assert.NoError(t, err)
		assert.Equal(t, "false", v)
		opts.options["InvalidType"] = Option{n: "InvalidType", t: OptionTypeBool}
		_, err = opts.String("InvalidType", nil)
		assert.ErrorIs(t, err, ErrInvalidType)
		opts.options["NilValue"] = Option{n: "NilValue", t: OptionTypeString}
		_, err = opts.String("NilValue", nil)
		assert.ErrorIs(t, err, ErrInvalidValue)
		opts.options["InvalidValue"] = Option{n: "InvalidValue", t: OptionTypeString, v: "Foo"}
		_, err = opts.String("InvalidValue", nil)
		assert.ErrorIs(t, err, ErrInvalidValue)
	})
}
