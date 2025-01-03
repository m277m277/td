// Code generated by gotdgen, DO NOT EDIT.

package tdapi

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

// TransferGiftRequest represents TL type `transferGift#7f379a4e`.
type TransferGiftRequest struct {
	// Identifier of the user that sent the gift
	SenderUserID int64
	// Identifier of the message with the upgraded gift in the chat with the user
	MessageID int64
	// Identifier of the user that will receive the gift
	ReceiverUserID int64
	// The amount of Telegram Stars required for the transfer
	StarCount int64
}

// TransferGiftRequestTypeID is TL type id of TransferGiftRequest.
const TransferGiftRequestTypeID = 0x7f379a4e

// Ensuring interfaces in compile-time for TransferGiftRequest.
var (
	_ bin.Encoder     = &TransferGiftRequest{}
	_ bin.Decoder     = &TransferGiftRequest{}
	_ bin.BareEncoder = &TransferGiftRequest{}
	_ bin.BareDecoder = &TransferGiftRequest{}
)

func (t *TransferGiftRequest) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.SenderUserID == 0) {
		return false
	}
	if !(t.MessageID == 0) {
		return false
	}
	if !(t.ReceiverUserID == 0) {
		return false
	}
	if !(t.StarCount == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *TransferGiftRequest) String() string {
	if t == nil {
		return "TransferGiftRequest(nil)"
	}
	type Alias TransferGiftRequest
	return fmt.Sprintf("TransferGiftRequest%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*TransferGiftRequest) TypeID() uint32 {
	return TransferGiftRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*TransferGiftRequest) TypeName() string {
	return "transferGift"
}

// TypeInfo returns info about TL type.
func (t *TransferGiftRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "transferGift",
		ID:   TransferGiftRequestTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "SenderUserID",
			SchemaName: "sender_user_id",
		},
		{
			Name:       "MessageID",
			SchemaName: "message_id",
		},
		{
			Name:       "ReceiverUserID",
			SchemaName: "receiver_user_id",
		},
		{
			Name:       "StarCount",
			SchemaName: "star_count",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *TransferGiftRequest) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode transferGift#7f379a4e as nil")
	}
	b.PutID(TransferGiftRequestTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *TransferGiftRequest) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode transferGift#7f379a4e as nil")
	}
	b.PutInt53(t.SenderUserID)
	b.PutInt53(t.MessageID)
	b.PutInt53(t.ReceiverUserID)
	b.PutInt53(t.StarCount)
	return nil
}

// Decode implements bin.Decoder.
func (t *TransferGiftRequest) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode transferGift#7f379a4e to nil")
	}
	if err := b.ConsumeID(TransferGiftRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode transferGift#7f379a4e: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *TransferGiftRequest) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode transferGift#7f379a4e to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode transferGift#7f379a4e: field sender_user_id: %w", err)
		}
		t.SenderUserID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode transferGift#7f379a4e: field message_id: %w", err)
		}
		t.MessageID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode transferGift#7f379a4e: field receiver_user_id: %w", err)
		}
		t.ReceiverUserID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode transferGift#7f379a4e: field star_count: %w", err)
		}
		t.StarCount = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (t *TransferGiftRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if t == nil {
		return fmt.Errorf("can't encode transferGift#7f379a4e as nil")
	}
	b.ObjStart()
	b.PutID("transferGift")
	b.Comma()
	b.FieldStart("sender_user_id")
	b.PutInt53(t.SenderUserID)
	b.Comma()
	b.FieldStart("message_id")
	b.PutInt53(t.MessageID)
	b.Comma()
	b.FieldStart("receiver_user_id")
	b.PutInt53(t.ReceiverUserID)
	b.Comma()
	b.FieldStart("star_count")
	b.PutInt53(t.StarCount)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (t *TransferGiftRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if t == nil {
		return fmt.Errorf("can't decode transferGift#7f379a4e to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("transferGift"); err != nil {
				return fmt.Errorf("unable to decode transferGift#7f379a4e: %w", err)
			}
		case "sender_user_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode transferGift#7f379a4e: field sender_user_id: %w", err)
			}
			t.SenderUserID = value
		case "message_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode transferGift#7f379a4e: field message_id: %w", err)
			}
			t.MessageID = value
		case "receiver_user_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode transferGift#7f379a4e: field receiver_user_id: %w", err)
			}
			t.ReceiverUserID = value
		case "star_count":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode transferGift#7f379a4e: field star_count: %w", err)
			}
			t.StarCount = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetSenderUserID returns value of SenderUserID field.
func (t *TransferGiftRequest) GetSenderUserID() (value int64) {
	if t == nil {
		return
	}
	return t.SenderUserID
}

// GetMessageID returns value of MessageID field.
func (t *TransferGiftRequest) GetMessageID() (value int64) {
	if t == nil {
		return
	}
	return t.MessageID
}

// GetReceiverUserID returns value of ReceiverUserID field.
func (t *TransferGiftRequest) GetReceiverUserID() (value int64) {
	if t == nil {
		return
	}
	return t.ReceiverUserID
}

// GetStarCount returns value of StarCount field.
func (t *TransferGiftRequest) GetStarCount() (value int64) {
	if t == nil {
		return
	}
	return t.StarCount
}

// TransferGift invokes method transferGift#7f379a4e returning error if any.
func (c *Client) TransferGift(ctx context.Context, request *TransferGiftRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}