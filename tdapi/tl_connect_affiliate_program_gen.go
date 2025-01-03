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

// ConnectAffiliateProgramRequest represents TL type `connectAffiliateProgram#6306d72c`.
type ConnectAffiliateProgramRequest struct {
	// The affiliate to which the affiliate program will be connected
	Affiliate AffiliateTypeClass
	// Identifier of the bot, which affiliate program is connected
	BotUserID int64
}

// ConnectAffiliateProgramRequestTypeID is TL type id of ConnectAffiliateProgramRequest.
const ConnectAffiliateProgramRequestTypeID = 0x6306d72c

// Ensuring interfaces in compile-time for ConnectAffiliateProgramRequest.
var (
	_ bin.Encoder     = &ConnectAffiliateProgramRequest{}
	_ bin.Decoder     = &ConnectAffiliateProgramRequest{}
	_ bin.BareEncoder = &ConnectAffiliateProgramRequest{}
	_ bin.BareDecoder = &ConnectAffiliateProgramRequest{}
)

func (c *ConnectAffiliateProgramRequest) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Affiliate == nil) {
		return false
	}
	if !(c.BotUserID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ConnectAffiliateProgramRequest) String() string {
	if c == nil {
		return "ConnectAffiliateProgramRequest(nil)"
	}
	type Alias ConnectAffiliateProgramRequest
	return fmt.Sprintf("ConnectAffiliateProgramRequest%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ConnectAffiliateProgramRequest) TypeID() uint32 {
	return ConnectAffiliateProgramRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ConnectAffiliateProgramRequest) TypeName() string {
	return "connectAffiliateProgram"
}

// TypeInfo returns info about TL type.
func (c *ConnectAffiliateProgramRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "connectAffiliateProgram",
		ID:   ConnectAffiliateProgramRequestTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Affiliate",
			SchemaName: "affiliate",
		},
		{
			Name:       "BotUserID",
			SchemaName: "bot_user_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ConnectAffiliateProgramRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode connectAffiliateProgram#6306d72c as nil")
	}
	b.PutID(ConnectAffiliateProgramRequestTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ConnectAffiliateProgramRequest) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode connectAffiliateProgram#6306d72c as nil")
	}
	if c.Affiliate == nil {
		return fmt.Errorf("unable to encode connectAffiliateProgram#6306d72c: field affiliate is nil")
	}
	if err := c.Affiliate.Encode(b); err != nil {
		return fmt.Errorf("unable to encode connectAffiliateProgram#6306d72c: field affiliate: %w", err)
	}
	b.PutInt53(c.BotUserID)
	return nil
}

// Decode implements bin.Decoder.
func (c *ConnectAffiliateProgramRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode connectAffiliateProgram#6306d72c to nil")
	}
	if err := b.ConsumeID(ConnectAffiliateProgramRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode connectAffiliateProgram#6306d72c: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ConnectAffiliateProgramRequest) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode connectAffiliateProgram#6306d72c to nil")
	}
	{
		value, err := DecodeAffiliateType(b)
		if err != nil {
			return fmt.Errorf("unable to decode connectAffiliateProgram#6306d72c: field affiliate: %w", err)
		}
		c.Affiliate = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode connectAffiliateProgram#6306d72c: field bot_user_id: %w", err)
		}
		c.BotUserID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *ConnectAffiliateProgramRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode connectAffiliateProgram#6306d72c as nil")
	}
	b.ObjStart()
	b.PutID("connectAffiliateProgram")
	b.Comma()
	b.FieldStart("affiliate")
	if c.Affiliate == nil {
		return fmt.Errorf("unable to encode connectAffiliateProgram#6306d72c: field affiliate is nil")
	}
	if err := c.Affiliate.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode connectAffiliateProgram#6306d72c: field affiliate: %w", err)
	}
	b.Comma()
	b.FieldStart("bot_user_id")
	b.PutInt53(c.BotUserID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *ConnectAffiliateProgramRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode connectAffiliateProgram#6306d72c to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("connectAffiliateProgram"); err != nil {
				return fmt.Errorf("unable to decode connectAffiliateProgram#6306d72c: %w", err)
			}
		case "affiliate":
			value, err := DecodeTDLibJSONAffiliateType(b)
			if err != nil {
				return fmt.Errorf("unable to decode connectAffiliateProgram#6306d72c: field affiliate: %w", err)
			}
			c.Affiliate = value
		case "bot_user_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode connectAffiliateProgram#6306d72c: field bot_user_id: %w", err)
			}
			c.BotUserID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetAffiliate returns value of Affiliate field.
func (c *ConnectAffiliateProgramRequest) GetAffiliate() (value AffiliateTypeClass) {
	if c == nil {
		return
	}
	return c.Affiliate
}

// GetBotUserID returns value of BotUserID field.
func (c *ConnectAffiliateProgramRequest) GetBotUserID() (value int64) {
	if c == nil {
		return
	}
	return c.BotUserID
}

// ConnectAffiliateProgram invokes method connectAffiliateProgram#6306d72c returning error if any.
func (c *Client) ConnectAffiliateProgram(ctx context.Context, request *ConnectAffiliateProgramRequest) (*ConnectedAffiliateProgram, error) {
	var result ConnectedAffiliateProgram

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}