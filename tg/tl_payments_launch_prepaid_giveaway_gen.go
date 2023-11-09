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

// PaymentsLaunchPrepaidGiveawayRequest represents TL type `payments.launchPrepaidGiveaway#5ff58f20`.
//
// See https://core.telegram.org/method/payments.launchPrepaidGiveaway for reference.
type PaymentsLaunchPrepaidGiveawayRequest struct {
	// Peer field of PaymentsLaunchPrepaidGiveawayRequest.
	Peer InputPeerClass
	// GiveawayID field of PaymentsLaunchPrepaidGiveawayRequest.
	GiveawayID int64
	// Purpose field of PaymentsLaunchPrepaidGiveawayRequest.
	Purpose InputStorePaymentPurposeClass
}

// PaymentsLaunchPrepaidGiveawayRequestTypeID is TL type id of PaymentsLaunchPrepaidGiveawayRequest.
const PaymentsLaunchPrepaidGiveawayRequestTypeID = 0x5ff58f20

// Ensuring interfaces in compile-time for PaymentsLaunchPrepaidGiveawayRequest.
var (
	_ bin.Encoder     = &PaymentsLaunchPrepaidGiveawayRequest{}
	_ bin.Decoder     = &PaymentsLaunchPrepaidGiveawayRequest{}
	_ bin.BareEncoder = &PaymentsLaunchPrepaidGiveawayRequest{}
	_ bin.BareDecoder = &PaymentsLaunchPrepaidGiveawayRequest{}
)

func (l *PaymentsLaunchPrepaidGiveawayRequest) Zero() bool {
	if l == nil {
		return true
	}
	if !(l.Peer == nil) {
		return false
	}
	if !(l.GiveawayID == 0) {
		return false
	}
	if !(l.Purpose == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (l *PaymentsLaunchPrepaidGiveawayRequest) String() string {
	if l == nil {
		return "PaymentsLaunchPrepaidGiveawayRequest(nil)"
	}
	type Alias PaymentsLaunchPrepaidGiveawayRequest
	return fmt.Sprintf("PaymentsLaunchPrepaidGiveawayRequest%+v", Alias(*l))
}

// FillFrom fills PaymentsLaunchPrepaidGiveawayRequest from given interface.
func (l *PaymentsLaunchPrepaidGiveawayRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetGiveawayID() (value int64)
	GetPurpose() (value InputStorePaymentPurposeClass)
}) {
	l.Peer = from.GetPeer()
	l.GiveawayID = from.GetGiveawayID()
	l.Purpose = from.GetPurpose()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PaymentsLaunchPrepaidGiveawayRequest) TypeID() uint32 {
	return PaymentsLaunchPrepaidGiveawayRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*PaymentsLaunchPrepaidGiveawayRequest) TypeName() string {
	return "payments.launchPrepaidGiveaway"
}

// TypeInfo returns info about TL type.
func (l *PaymentsLaunchPrepaidGiveawayRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "payments.launchPrepaidGiveaway",
		ID:   PaymentsLaunchPrepaidGiveawayRequestTypeID,
	}
	if l == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "GiveawayID",
			SchemaName: "giveaway_id",
		},
		{
			Name:       "Purpose",
			SchemaName: "purpose",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (l *PaymentsLaunchPrepaidGiveawayRequest) Encode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't encode payments.launchPrepaidGiveaway#5ff58f20 as nil")
	}
	b.PutID(PaymentsLaunchPrepaidGiveawayRequestTypeID)
	return l.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (l *PaymentsLaunchPrepaidGiveawayRequest) EncodeBare(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't encode payments.launchPrepaidGiveaway#5ff58f20 as nil")
	}
	if l.Peer == nil {
		return fmt.Errorf("unable to encode payments.launchPrepaidGiveaway#5ff58f20: field peer is nil")
	}
	if err := l.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode payments.launchPrepaidGiveaway#5ff58f20: field peer: %w", err)
	}
	b.PutLong(l.GiveawayID)
	if l.Purpose == nil {
		return fmt.Errorf("unable to encode payments.launchPrepaidGiveaway#5ff58f20: field purpose is nil")
	}
	if err := l.Purpose.Encode(b); err != nil {
		return fmt.Errorf("unable to encode payments.launchPrepaidGiveaway#5ff58f20: field purpose: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (l *PaymentsLaunchPrepaidGiveawayRequest) Decode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't decode payments.launchPrepaidGiveaway#5ff58f20 to nil")
	}
	if err := b.ConsumeID(PaymentsLaunchPrepaidGiveawayRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode payments.launchPrepaidGiveaway#5ff58f20: %w", err)
	}
	return l.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (l *PaymentsLaunchPrepaidGiveawayRequest) DecodeBare(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't decode payments.launchPrepaidGiveaway#5ff58f20 to nil")
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode payments.launchPrepaidGiveaway#5ff58f20: field peer: %w", err)
		}
		l.Peer = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode payments.launchPrepaidGiveaway#5ff58f20: field giveaway_id: %w", err)
		}
		l.GiveawayID = value
	}
	{
		value, err := DecodeInputStorePaymentPurpose(b)
		if err != nil {
			return fmt.Errorf("unable to decode payments.launchPrepaidGiveaway#5ff58f20: field purpose: %w", err)
		}
		l.Purpose = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (l *PaymentsLaunchPrepaidGiveawayRequest) GetPeer() (value InputPeerClass) {
	if l == nil {
		return
	}
	return l.Peer
}

// GetGiveawayID returns value of GiveawayID field.
func (l *PaymentsLaunchPrepaidGiveawayRequest) GetGiveawayID() (value int64) {
	if l == nil {
		return
	}
	return l.GiveawayID
}

// GetPurpose returns value of Purpose field.
func (l *PaymentsLaunchPrepaidGiveawayRequest) GetPurpose() (value InputStorePaymentPurposeClass) {
	if l == nil {
		return
	}
	return l.Purpose
}

// PaymentsLaunchPrepaidGiveaway invokes method payments.launchPrepaidGiveaway#5ff58f20 returning error if any.
//
// See https://core.telegram.org/method/payments.launchPrepaidGiveaway for reference.
func (c *Client) PaymentsLaunchPrepaidGiveaway(ctx context.Context, request *PaymentsLaunchPrepaidGiveawayRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}