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

// UpgradeGiftRequest represents TL type `upgradeGift#e0c0953e`.
type UpgradeGiftRequest struct {
	// Identifier of the user that sent the gift
	SenderUserID int64
	// Identifier of the message with the gift in the chat with the user
	MessageID int64
	// Pass true to keep the original gift text, sender and receiver in the upgraded gift
	KeepOriginalDetails bool
}

// UpgradeGiftRequestTypeID is TL type id of UpgradeGiftRequest.
const UpgradeGiftRequestTypeID = 0xe0c0953e

// Ensuring interfaces in compile-time for UpgradeGiftRequest.
var (
	_ bin.Encoder     = &UpgradeGiftRequest{}
	_ bin.Decoder     = &UpgradeGiftRequest{}
	_ bin.BareEncoder = &UpgradeGiftRequest{}
	_ bin.BareDecoder = &UpgradeGiftRequest{}
)

func (u *UpgradeGiftRequest) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.SenderUserID == 0) {
		return false
	}
	if !(u.MessageID == 0) {
		return false
	}
	if !(u.KeepOriginalDetails == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *UpgradeGiftRequest) String() string {
	if u == nil {
		return "UpgradeGiftRequest(nil)"
	}
	type Alias UpgradeGiftRequest
	return fmt.Sprintf("UpgradeGiftRequest%+v", Alias(*u))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*UpgradeGiftRequest) TypeID() uint32 {
	return UpgradeGiftRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*UpgradeGiftRequest) TypeName() string {
	return "upgradeGift"
}

// TypeInfo returns info about TL type.
func (u *UpgradeGiftRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "upgradeGift",
		ID:   UpgradeGiftRequestTypeID,
	}
	if u == nil {
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
			Name:       "KeepOriginalDetails",
			SchemaName: "keep_original_details",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (u *UpgradeGiftRequest) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode upgradeGift#e0c0953e as nil")
	}
	b.PutID(UpgradeGiftRequestTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *UpgradeGiftRequest) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode upgradeGift#e0c0953e as nil")
	}
	b.PutInt53(u.SenderUserID)
	b.PutInt53(u.MessageID)
	b.PutBool(u.KeepOriginalDetails)
	return nil
}

// Decode implements bin.Decoder.
func (u *UpgradeGiftRequest) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode upgradeGift#e0c0953e to nil")
	}
	if err := b.ConsumeID(UpgradeGiftRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode upgradeGift#e0c0953e: %w", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *UpgradeGiftRequest) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode upgradeGift#e0c0953e to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode upgradeGift#e0c0953e: field sender_user_id: %w", err)
		}
		u.SenderUserID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode upgradeGift#e0c0953e: field message_id: %w", err)
		}
		u.MessageID = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode upgradeGift#e0c0953e: field keep_original_details: %w", err)
		}
		u.KeepOriginalDetails = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (u *UpgradeGiftRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if u == nil {
		return fmt.Errorf("can't encode upgradeGift#e0c0953e as nil")
	}
	b.ObjStart()
	b.PutID("upgradeGift")
	b.Comma()
	b.FieldStart("sender_user_id")
	b.PutInt53(u.SenderUserID)
	b.Comma()
	b.FieldStart("message_id")
	b.PutInt53(u.MessageID)
	b.Comma()
	b.FieldStart("keep_original_details")
	b.PutBool(u.KeepOriginalDetails)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (u *UpgradeGiftRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if u == nil {
		return fmt.Errorf("can't decode upgradeGift#e0c0953e to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("upgradeGift"); err != nil {
				return fmt.Errorf("unable to decode upgradeGift#e0c0953e: %w", err)
			}
		case "sender_user_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode upgradeGift#e0c0953e: field sender_user_id: %w", err)
			}
			u.SenderUserID = value
		case "message_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode upgradeGift#e0c0953e: field message_id: %w", err)
			}
			u.MessageID = value
		case "keep_original_details":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode upgradeGift#e0c0953e: field keep_original_details: %w", err)
			}
			u.KeepOriginalDetails = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetSenderUserID returns value of SenderUserID field.
func (u *UpgradeGiftRequest) GetSenderUserID() (value int64) {
	if u == nil {
		return
	}
	return u.SenderUserID
}

// GetMessageID returns value of MessageID field.
func (u *UpgradeGiftRequest) GetMessageID() (value int64) {
	if u == nil {
		return
	}
	return u.MessageID
}

// GetKeepOriginalDetails returns value of KeepOriginalDetails field.
func (u *UpgradeGiftRequest) GetKeepOriginalDetails() (value bool) {
	if u == nil {
		return
	}
	return u.KeepOriginalDetails
}

// UpgradeGift invokes method upgradeGift#e0c0953e returning error if any.
func (c *Client) UpgradeGift(ctx context.Context, request *UpgradeGiftRequest) (*UpgradeGiftResult, error) {
	var result UpgradeGiftResult

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}