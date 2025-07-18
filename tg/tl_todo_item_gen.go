// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
	_ = tdjson.Encoder{}
)

// TodoItem represents TL type `todoItem#cba9a52f`.
//
// See https://core.telegram.org/constructor/todoItem for reference.
type TodoItem struct {
	// ID field of TodoItem.
	ID int
	// Title field of TodoItem.
	Title TextWithEntities
}

// TodoItemTypeID is TL type id of TodoItem.
const TodoItemTypeID = 0xcba9a52f

// Ensuring interfaces in compile-time for TodoItem.
var (
	_ bin.Encoder     = &TodoItem{}
	_ bin.Decoder     = &TodoItem{}
	_ bin.BareEncoder = &TodoItem{}
	_ bin.BareDecoder = &TodoItem{}
)

func (t *TodoItem) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.ID == 0) {
		return false
	}
	if !(t.Title.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *TodoItem) String() string {
	if t == nil {
		return "TodoItem(nil)"
	}
	type Alias TodoItem
	return fmt.Sprintf("TodoItem%+v", Alias(*t))
}

// FillFrom fills TodoItem from given interface.
func (t *TodoItem) FillFrom(from interface {
	GetID() (value int)
	GetTitle() (value TextWithEntities)
}) {
	t.ID = from.GetID()
	t.Title = from.GetTitle()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*TodoItem) TypeID() uint32 {
	return TodoItemTypeID
}

// TypeName returns name of type in TL schema.
func (*TodoItem) TypeName() string {
	return "todoItem"
}

// TypeInfo returns info about TL type.
func (t *TodoItem) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "todoItem",
		ID:   TodoItemTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "Title",
			SchemaName: "title",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *TodoItem) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode todoItem#cba9a52f as nil")
	}
	b.PutID(TodoItemTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *TodoItem) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode todoItem#cba9a52f as nil")
	}
	b.PutInt(t.ID)
	if err := t.Title.Encode(b); err != nil {
		return fmt.Errorf("unable to encode todoItem#cba9a52f: field title: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (t *TodoItem) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode todoItem#cba9a52f to nil")
	}
	if err := b.ConsumeID(TodoItemTypeID); err != nil {
		return fmt.Errorf("unable to decode todoItem#cba9a52f: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *TodoItem) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode todoItem#cba9a52f to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode todoItem#cba9a52f: field id: %w", err)
		}
		t.ID = value
	}
	{
		if err := t.Title.Decode(b); err != nil {
			return fmt.Errorf("unable to decode todoItem#cba9a52f: field title: %w", err)
		}
	}
	return nil
}

// GetID returns value of ID field.
func (t *TodoItem) GetID() (value int) {
	if t == nil {
		return
	}
	return t.ID
}

// GetTitle returns value of Title field.
func (t *TodoItem) GetTitle() (value TextWithEntities) {
	if t == nil {
		return
	}
	return t.Title
}
