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

// BotsReorderPreviewMediasRequest represents TL type `bots.reorderPreviewMedias#b627f3aa`.
//
// See https://core.telegram.org/method/bots.reorderPreviewMedias for reference.
type BotsReorderPreviewMediasRequest struct {
	// Bot field of BotsReorderPreviewMediasRequest.
	Bot InputUserClass
	// LangCode field of BotsReorderPreviewMediasRequest.
	LangCode string
	// Order field of BotsReorderPreviewMediasRequest.
	Order []InputMediaClass
}

// BotsReorderPreviewMediasRequestTypeID is TL type id of BotsReorderPreviewMediasRequest.
const BotsReorderPreviewMediasRequestTypeID = 0xb627f3aa

// Ensuring interfaces in compile-time for BotsReorderPreviewMediasRequest.
var (
	_ bin.Encoder     = &BotsReorderPreviewMediasRequest{}
	_ bin.Decoder     = &BotsReorderPreviewMediasRequest{}
	_ bin.BareEncoder = &BotsReorderPreviewMediasRequest{}
	_ bin.BareDecoder = &BotsReorderPreviewMediasRequest{}
)

func (r *BotsReorderPreviewMediasRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Bot == nil) {
		return false
	}
	if !(r.LangCode == "") {
		return false
	}
	if !(r.Order == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *BotsReorderPreviewMediasRequest) String() string {
	if r == nil {
		return "BotsReorderPreviewMediasRequest(nil)"
	}
	type Alias BotsReorderPreviewMediasRequest
	return fmt.Sprintf("BotsReorderPreviewMediasRequest%+v", Alias(*r))
}

// FillFrom fills BotsReorderPreviewMediasRequest from given interface.
func (r *BotsReorderPreviewMediasRequest) FillFrom(from interface {
	GetBot() (value InputUserClass)
	GetLangCode() (value string)
	GetOrder() (value []InputMediaClass)
}) {
	r.Bot = from.GetBot()
	r.LangCode = from.GetLangCode()
	r.Order = from.GetOrder()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BotsReorderPreviewMediasRequest) TypeID() uint32 {
	return BotsReorderPreviewMediasRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*BotsReorderPreviewMediasRequest) TypeName() string {
	return "bots.reorderPreviewMedias"
}

// TypeInfo returns info about TL type.
func (r *BotsReorderPreviewMediasRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "bots.reorderPreviewMedias",
		ID:   BotsReorderPreviewMediasRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Bot",
			SchemaName: "bot",
		},
		{
			Name:       "LangCode",
			SchemaName: "lang_code",
		},
		{
			Name:       "Order",
			SchemaName: "order",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *BotsReorderPreviewMediasRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode bots.reorderPreviewMedias#b627f3aa as nil")
	}
	b.PutID(BotsReorderPreviewMediasRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *BotsReorderPreviewMediasRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode bots.reorderPreviewMedias#b627f3aa as nil")
	}
	if r.Bot == nil {
		return fmt.Errorf("unable to encode bots.reorderPreviewMedias#b627f3aa: field bot is nil")
	}
	if err := r.Bot.Encode(b); err != nil {
		return fmt.Errorf("unable to encode bots.reorderPreviewMedias#b627f3aa: field bot: %w", err)
	}
	b.PutString(r.LangCode)
	b.PutVectorHeader(len(r.Order))
	for idx, v := range r.Order {
		if v == nil {
			return fmt.Errorf("unable to encode bots.reorderPreviewMedias#b627f3aa: field order element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode bots.reorderPreviewMedias#b627f3aa: field order element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *BotsReorderPreviewMediasRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode bots.reorderPreviewMedias#b627f3aa to nil")
	}
	if err := b.ConsumeID(BotsReorderPreviewMediasRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode bots.reorderPreviewMedias#b627f3aa: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *BotsReorderPreviewMediasRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode bots.reorderPreviewMedias#b627f3aa to nil")
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode bots.reorderPreviewMedias#b627f3aa: field bot: %w", err)
		}
		r.Bot = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode bots.reorderPreviewMedias#b627f3aa: field lang_code: %w", err)
		}
		r.LangCode = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode bots.reorderPreviewMedias#b627f3aa: field order: %w", err)
		}

		if headerLen > 0 {
			r.Order = make([]InputMediaClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeInputMedia(b)
			if err != nil {
				return fmt.Errorf("unable to decode bots.reorderPreviewMedias#b627f3aa: field order: %w", err)
			}
			r.Order = append(r.Order, value)
		}
	}
	return nil
}

// GetBot returns value of Bot field.
func (r *BotsReorderPreviewMediasRequest) GetBot() (value InputUserClass) {
	if r == nil {
		return
	}
	return r.Bot
}

// GetLangCode returns value of LangCode field.
func (r *BotsReorderPreviewMediasRequest) GetLangCode() (value string) {
	if r == nil {
		return
	}
	return r.LangCode
}

// GetOrder returns value of Order field.
func (r *BotsReorderPreviewMediasRequest) GetOrder() (value []InputMediaClass) {
	if r == nil {
		return
	}
	return r.Order
}

// MapOrder returns field Order wrapped in InputMediaClassArray helper.
func (r *BotsReorderPreviewMediasRequest) MapOrder() (value InputMediaClassArray) {
	return InputMediaClassArray(r.Order)
}

// BotsReorderPreviewMedias invokes method bots.reorderPreviewMedias#b627f3aa returning error if any.
//
// See https://core.telegram.org/method/bots.reorderPreviewMedias for reference.
func (c *Client) BotsReorderPreviewMedias(ctx context.Context, request *BotsReorderPreviewMediasRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}