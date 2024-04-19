package options

// Copyright (c) 2024 Jann Fischer <jann@mistrust.net>
// SPDX-License-Identifier: 0BSD

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_New(t *testing.T) {
	t.Run("Clean new options", func(t *testing.T) {
		opts := New()
		assert.Len(t, opts.options, 0)
		assert.Len(t, opts.valid, 0)
	})
	t.Run("Set some valid options", func(t *testing.T) {
		opts := New().WithOption("Foo", OptionTypeBool)
		assert.Len(t, opts.options, 0)
		assert.Len(t, opts.valid, 1)
	})
}

func Test_Parse(t *testing.T) {
	t.Run("All valid options", func(t *testing.T) {
		opts := New().WithOption("Foo", OptionTypeBool).WithOption("Bar", OptionTypeString)
		inp := []string{"Foo=true", "Bar=baz"}
		err := opts.Parse(inp)
		assert.NoError(t, err)
		assert.Len(t, opts.options, 2)
	})
	t.Run("Some invalid option", func(t *testing.T) {
		opts := New().WithOption("Foo", OptionTypeBool).WithOption("Bar", OptionTypeString)
		inp := []string{"Foo=true", "Baz=bar"}
		err := opts.Parse(inp)
		assert.ErrorIs(t, err, ErrForbidden)
	})
}
