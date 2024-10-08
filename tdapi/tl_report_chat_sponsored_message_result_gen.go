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

// ReportChatSponsoredMessageResultOk represents TL type `reportChatSponsoredMessageResultOk#754f721f`.
type ReportChatSponsoredMessageResultOk struct {
}

// ReportChatSponsoredMessageResultOkTypeID is TL type id of ReportChatSponsoredMessageResultOk.
const ReportChatSponsoredMessageResultOkTypeID = 0x754f721f

// construct implements constructor of ReportChatSponsoredMessageResultClass.
func (r ReportChatSponsoredMessageResultOk) construct() ReportChatSponsoredMessageResultClass {
	return &r
}

// Ensuring interfaces in compile-time for ReportChatSponsoredMessageResultOk.
var (
	_ bin.Encoder     = &ReportChatSponsoredMessageResultOk{}
	_ bin.Decoder     = &ReportChatSponsoredMessageResultOk{}
	_ bin.BareEncoder = &ReportChatSponsoredMessageResultOk{}
	_ bin.BareDecoder = &ReportChatSponsoredMessageResultOk{}

	_ ReportChatSponsoredMessageResultClass = &ReportChatSponsoredMessageResultOk{}
)

func (r *ReportChatSponsoredMessageResultOk) Zero() bool {
	if r == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (r *ReportChatSponsoredMessageResultOk) String() string {
	if r == nil {
		return "ReportChatSponsoredMessageResultOk(nil)"
	}
	type Alias ReportChatSponsoredMessageResultOk
	return fmt.Sprintf("ReportChatSponsoredMessageResultOk%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ReportChatSponsoredMessageResultOk) TypeID() uint32 {
	return ReportChatSponsoredMessageResultOkTypeID
}

// TypeName returns name of type in TL schema.
func (*ReportChatSponsoredMessageResultOk) TypeName() string {
	return "reportChatSponsoredMessageResultOk"
}

// TypeInfo returns info about TL type.
func (r *ReportChatSponsoredMessageResultOk) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "reportChatSponsoredMessageResultOk",
		ID:   ReportChatSponsoredMessageResultOkTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (r *ReportChatSponsoredMessageResultOk) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode reportChatSponsoredMessageResultOk#754f721f as nil")
	}
	b.PutID(ReportChatSponsoredMessageResultOkTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *ReportChatSponsoredMessageResultOk) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode reportChatSponsoredMessageResultOk#754f721f as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *ReportChatSponsoredMessageResultOk) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode reportChatSponsoredMessageResultOk#754f721f to nil")
	}
	if err := b.ConsumeID(ReportChatSponsoredMessageResultOkTypeID); err != nil {
		return fmt.Errorf("unable to decode reportChatSponsoredMessageResultOk#754f721f: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *ReportChatSponsoredMessageResultOk) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode reportChatSponsoredMessageResultOk#754f721f to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (r *ReportChatSponsoredMessageResultOk) EncodeTDLibJSON(b tdjson.Encoder) error {
	if r == nil {
		return fmt.Errorf("can't encode reportChatSponsoredMessageResultOk#754f721f as nil")
	}
	b.ObjStart()
	b.PutID("reportChatSponsoredMessageResultOk")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (r *ReportChatSponsoredMessageResultOk) DecodeTDLibJSON(b tdjson.Decoder) error {
	if r == nil {
		return fmt.Errorf("can't decode reportChatSponsoredMessageResultOk#754f721f to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("reportChatSponsoredMessageResultOk"); err != nil {
				return fmt.Errorf("unable to decode reportChatSponsoredMessageResultOk#754f721f: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// ReportChatSponsoredMessageResultFailed represents TL type `reportChatSponsoredMessageResultFailed#7f1f9bc6`.
type ReportChatSponsoredMessageResultFailed struct {
}

// ReportChatSponsoredMessageResultFailedTypeID is TL type id of ReportChatSponsoredMessageResultFailed.
const ReportChatSponsoredMessageResultFailedTypeID = 0x7f1f9bc6

// construct implements constructor of ReportChatSponsoredMessageResultClass.
func (r ReportChatSponsoredMessageResultFailed) construct() ReportChatSponsoredMessageResultClass {
	return &r
}

// Ensuring interfaces in compile-time for ReportChatSponsoredMessageResultFailed.
var (
	_ bin.Encoder     = &ReportChatSponsoredMessageResultFailed{}
	_ bin.Decoder     = &ReportChatSponsoredMessageResultFailed{}
	_ bin.BareEncoder = &ReportChatSponsoredMessageResultFailed{}
	_ bin.BareDecoder = &ReportChatSponsoredMessageResultFailed{}

	_ ReportChatSponsoredMessageResultClass = &ReportChatSponsoredMessageResultFailed{}
)

func (r *ReportChatSponsoredMessageResultFailed) Zero() bool {
	if r == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (r *ReportChatSponsoredMessageResultFailed) String() string {
	if r == nil {
		return "ReportChatSponsoredMessageResultFailed(nil)"
	}
	type Alias ReportChatSponsoredMessageResultFailed
	return fmt.Sprintf("ReportChatSponsoredMessageResultFailed%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ReportChatSponsoredMessageResultFailed) TypeID() uint32 {
	return ReportChatSponsoredMessageResultFailedTypeID
}

// TypeName returns name of type in TL schema.
func (*ReportChatSponsoredMessageResultFailed) TypeName() string {
	return "reportChatSponsoredMessageResultFailed"
}

// TypeInfo returns info about TL type.
func (r *ReportChatSponsoredMessageResultFailed) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "reportChatSponsoredMessageResultFailed",
		ID:   ReportChatSponsoredMessageResultFailedTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (r *ReportChatSponsoredMessageResultFailed) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode reportChatSponsoredMessageResultFailed#7f1f9bc6 as nil")
	}
	b.PutID(ReportChatSponsoredMessageResultFailedTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *ReportChatSponsoredMessageResultFailed) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode reportChatSponsoredMessageResultFailed#7f1f9bc6 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *ReportChatSponsoredMessageResultFailed) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode reportChatSponsoredMessageResultFailed#7f1f9bc6 to nil")
	}
	if err := b.ConsumeID(ReportChatSponsoredMessageResultFailedTypeID); err != nil {
		return fmt.Errorf("unable to decode reportChatSponsoredMessageResultFailed#7f1f9bc6: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *ReportChatSponsoredMessageResultFailed) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode reportChatSponsoredMessageResultFailed#7f1f9bc6 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (r *ReportChatSponsoredMessageResultFailed) EncodeTDLibJSON(b tdjson.Encoder) error {
	if r == nil {
		return fmt.Errorf("can't encode reportChatSponsoredMessageResultFailed#7f1f9bc6 as nil")
	}
	b.ObjStart()
	b.PutID("reportChatSponsoredMessageResultFailed")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (r *ReportChatSponsoredMessageResultFailed) DecodeTDLibJSON(b tdjson.Decoder) error {
	if r == nil {
		return fmt.Errorf("can't decode reportChatSponsoredMessageResultFailed#7f1f9bc6 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("reportChatSponsoredMessageResultFailed"); err != nil {
				return fmt.Errorf("unable to decode reportChatSponsoredMessageResultFailed#7f1f9bc6: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// ReportChatSponsoredMessageResultOptionRequired represents TL type `reportChatSponsoredMessageResultOptionRequired#34fc24b2`.
type ReportChatSponsoredMessageResultOptionRequired struct {
	// Title for the option choice
	Title string
	// List of available options
	Options []ReportOption
}

// ReportChatSponsoredMessageResultOptionRequiredTypeID is TL type id of ReportChatSponsoredMessageResultOptionRequired.
const ReportChatSponsoredMessageResultOptionRequiredTypeID = 0x34fc24b2

// construct implements constructor of ReportChatSponsoredMessageResultClass.
func (r ReportChatSponsoredMessageResultOptionRequired) construct() ReportChatSponsoredMessageResultClass {
	return &r
}

// Ensuring interfaces in compile-time for ReportChatSponsoredMessageResultOptionRequired.
var (
	_ bin.Encoder     = &ReportChatSponsoredMessageResultOptionRequired{}
	_ bin.Decoder     = &ReportChatSponsoredMessageResultOptionRequired{}
	_ bin.BareEncoder = &ReportChatSponsoredMessageResultOptionRequired{}
	_ bin.BareDecoder = &ReportChatSponsoredMessageResultOptionRequired{}

	_ ReportChatSponsoredMessageResultClass = &ReportChatSponsoredMessageResultOptionRequired{}
)

func (r *ReportChatSponsoredMessageResultOptionRequired) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Title == "") {
		return false
	}
	if !(r.Options == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *ReportChatSponsoredMessageResultOptionRequired) String() string {
	if r == nil {
		return "ReportChatSponsoredMessageResultOptionRequired(nil)"
	}
	type Alias ReportChatSponsoredMessageResultOptionRequired
	return fmt.Sprintf("ReportChatSponsoredMessageResultOptionRequired%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ReportChatSponsoredMessageResultOptionRequired) TypeID() uint32 {
	return ReportChatSponsoredMessageResultOptionRequiredTypeID
}

// TypeName returns name of type in TL schema.
func (*ReportChatSponsoredMessageResultOptionRequired) TypeName() string {
	return "reportChatSponsoredMessageResultOptionRequired"
}

// TypeInfo returns info about TL type.
func (r *ReportChatSponsoredMessageResultOptionRequired) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "reportChatSponsoredMessageResultOptionRequired",
		ID:   ReportChatSponsoredMessageResultOptionRequiredTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Title",
			SchemaName: "title",
		},
		{
			Name:       "Options",
			SchemaName: "options",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *ReportChatSponsoredMessageResultOptionRequired) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode reportChatSponsoredMessageResultOptionRequired#34fc24b2 as nil")
	}
	b.PutID(ReportChatSponsoredMessageResultOptionRequiredTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *ReportChatSponsoredMessageResultOptionRequired) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode reportChatSponsoredMessageResultOptionRequired#34fc24b2 as nil")
	}
	b.PutString(r.Title)
	b.PutInt(len(r.Options))
	for idx, v := range r.Options {
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare reportChatSponsoredMessageResultOptionRequired#34fc24b2: field options element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *ReportChatSponsoredMessageResultOptionRequired) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode reportChatSponsoredMessageResultOptionRequired#34fc24b2 to nil")
	}
	if err := b.ConsumeID(ReportChatSponsoredMessageResultOptionRequiredTypeID); err != nil {
		return fmt.Errorf("unable to decode reportChatSponsoredMessageResultOptionRequired#34fc24b2: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *ReportChatSponsoredMessageResultOptionRequired) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode reportChatSponsoredMessageResultOptionRequired#34fc24b2 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode reportChatSponsoredMessageResultOptionRequired#34fc24b2: field title: %w", err)
		}
		r.Title = value
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode reportChatSponsoredMessageResultOptionRequired#34fc24b2: field options: %w", err)
		}

		if headerLen > 0 {
			r.Options = make([]ReportOption, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value ReportOption
			if err := value.DecodeBare(b); err != nil {
				return fmt.Errorf("unable to decode bare reportChatSponsoredMessageResultOptionRequired#34fc24b2: field options: %w", err)
			}
			r.Options = append(r.Options, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (r *ReportChatSponsoredMessageResultOptionRequired) EncodeTDLibJSON(b tdjson.Encoder) error {
	if r == nil {
		return fmt.Errorf("can't encode reportChatSponsoredMessageResultOptionRequired#34fc24b2 as nil")
	}
	b.ObjStart()
	b.PutID("reportChatSponsoredMessageResultOptionRequired")
	b.Comma()
	b.FieldStart("title")
	b.PutString(r.Title)
	b.Comma()
	b.FieldStart("options")
	b.ArrStart()
	for idx, v := range r.Options {
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode reportChatSponsoredMessageResultOptionRequired#34fc24b2: field options element with index %d: %w", idx, err)
		}
		b.Comma()
	}
	b.StripComma()
	b.ArrEnd()
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (r *ReportChatSponsoredMessageResultOptionRequired) DecodeTDLibJSON(b tdjson.Decoder) error {
	if r == nil {
		return fmt.Errorf("can't decode reportChatSponsoredMessageResultOptionRequired#34fc24b2 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("reportChatSponsoredMessageResultOptionRequired"); err != nil {
				return fmt.Errorf("unable to decode reportChatSponsoredMessageResultOptionRequired#34fc24b2: %w", err)
			}
		case "title":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode reportChatSponsoredMessageResultOptionRequired#34fc24b2: field title: %w", err)
			}
			r.Title = value
		case "options":
			if err := b.Arr(func(b tdjson.Decoder) error {
				var value ReportOption
				if err := value.DecodeTDLibJSON(b); err != nil {
					return fmt.Errorf("unable to decode reportChatSponsoredMessageResultOptionRequired#34fc24b2: field options: %w", err)
				}
				r.Options = append(r.Options, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode reportChatSponsoredMessageResultOptionRequired#34fc24b2: field options: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetTitle returns value of Title field.
func (r *ReportChatSponsoredMessageResultOptionRequired) GetTitle() (value string) {
	if r == nil {
		return
	}
	return r.Title
}

// GetOptions returns value of Options field.
func (r *ReportChatSponsoredMessageResultOptionRequired) GetOptions() (value []ReportOption) {
	if r == nil {
		return
	}
	return r.Options
}

// ReportChatSponsoredMessageResultAdsHidden represents TL type `reportChatSponsoredMessageResultAdsHidden#e8eade1e`.
type ReportChatSponsoredMessageResultAdsHidden struct {
}

// ReportChatSponsoredMessageResultAdsHiddenTypeID is TL type id of ReportChatSponsoredMessageResultAdsHidden.
const ReportChatSponsoredMessageResultAdsHiddenTypeID = 0xe8eade1e

// construct implements constructor of ReportChatSponsoredMessageResultClass.
func (r ReportChatSponsoredMessageResultAdsHidden) construct() ReportChatSponsoredMessageResultClass {
	return &r
}

// Ensuring interfaces in compile-time for ReportChatSponsoredMessageResultAdsHidden.
var (
	_ bin.Encoder     = &ReportChatSponsoredMessageResultAdsHidden{}
	_ bin.Decoder     = &ReportChatSponsoredMessageResultAdsHidden{}
	_ bin.BareEncoder = &ReportChatSponsoredMessageResultAdsHidden{}
	_ bin.BareDecoder = &ReportChatSponsoredMessageResultAdsHidden{}

	_ ReportChatSponsoredMessageResultClass = &ReportChatSponsoredMessageResultAdsHidden{}
)

func (r *ReportChatSponsoredMessageResultAdsHidden) Zero() bool {
	if r == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (r *ReportChatSponsoredMessageResultAdsHidden) String() string {
	if r == nil {
		return "ReportChatSponsoredMessageResultAdsHidden(nil)"
	}
	type Alias ReportChatSponsoredMessageResultAdsHidden
	return fmt.Sprintf("ReportChatSponsoredMessageResultAdsHidden%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ReportChatSponsoredMessageResultAdsHidden) TypeID() uint32 {
	return ReportChatSponsoredMessageResultAdsHiddenTypeID
}

// TypeName returns name of type in TL schema.
func (*ReportChatSponsoredMessageResultAdsHidden) TypeName() string {
	return "reportChatSponsoredMessageResultAdsHidden"
}

// TypeInfo returns info about TL type.
func (r *ReportChatSponsoredMessageResultAdsHidden) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "reportChatSponsoredMessageResultAdsHidden",
		ID:   ReportChatSponsoredMessageResultAdsHiddenTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (r *ReportChatSponsoredMessageResultAdsHidden) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode reportChatSponsoredMessageResultAdsHidden#e8eade1e as nil")
	}
	b.PutID(ReportChatSponsoredMessageResultAdsHiddenTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *ReportChatSponsoredMessageResultAdsHidden) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode reportChatSponsoredMessageResultAdsHidden#e8eade1e as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *ReportChatSponsoredMessageResultAdsHidden) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode reportChatSponsoredMessageResultAdsHidden#e8eade1e to nil")
	}
	if err := b.ConsumeID(ReportChatSponsoredMessageResultAdsHiddenTypeID); err != nil {
		return fmt.Errorf("unable to decode reportChatSponsoredMessageResultAdsHidden#e8eade1e: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *ReportChatSponsoredMessageResultAdsHidden) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode reportChatSponsoredMessageResultAdsHidden#e8eade1e to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (r *ReportChatSponsoredMessageResultAdsHidden) EncodeTDLibJSON(b tdjson.Encoder) error {
	if r == nil {
		return fmt.Errorf("can't encode reportChatSponsoredMessageResultAdsHidden#e8eade1e as nil")
	}
	b.ObjStart()
	b.PutID("reportChatSponsoredMessageResultAdsHidden")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (r *ReportChatSponsoredMessageResultAdsHidden) DecodeTDLibJSON(b tdjson.Decoder) error {
	if r == nil {
		return fmt.Errorf("can't decode reportChatSponsoredMessageResultAdsHidden#e8eade1e to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("reportChatSponsoredMessageResultAdsHidden"); err != nil {
				return fmt.Errorf("unable to decode reportChatSponsoredMessageResultAdsHidden#e8eade1e: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// ReportChatSponsoredMessageResultPremiumRequired represents TL type `reportChatSponsoredMessageResultPremiumRequired#770c2ed0`.
type ReportChatSponsoredMessageResultPremiumRequired struct {
}

// ReportChatSponsoredMessageResultPremiumRequiredTypeID is TL type id of ReportChatSponsoredMessageResultPremiumRequired.
const ReportChatSponsoredMessageResultPremiumRequiredTypeID = 0x770c2ed0

// construct implements constructor of ReportChatSponsoredMessageResultClass.
func (r ReportChatSponsoredMessageResultPremiumRequired) construct() ReportChatSponsoredMessageResultClass {
	return &r
}

// Ensuring interfaces in compile-time for ReportChatSponsoredMessageResultPremiumRequired.
var (
	_ bin.Encoder     = &ReportChatSponsoredMessageResultPremiumRequired{}
	_ bin.Decoder     = &ReportChatSponsoredMessageResultPremiumRequired{}
	_ bin.BareEncoder = &ReportChatSponsoredMessageResultPremiumRequired{}
	_ bin.BareDecoder = &ReportChatSponsoredMessageResultPremiumRequired{}

	_ ReportChatSponsoredMessageResultClass = &ReportChatSponsoredMessageResultPremiumRequired{}
)

func (r *ReportChatSponsoredMessageResultPremiumRequired) Zero() bool {
	if r == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (r *ReportChatSponsoredMessageResultPremiumRequired) String() string {
	if r == nil {
		return "ReportChatSponsoredMessageResultPremiumRequired(nil)"
	}
	type Alias ReportChatSponsoredMessageResultPremiumRequired
	return fmt.Sprintf("ReportChatSponsoredMessageResultPremiumRequired%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ReportChatSponsoredMessageResultPremiumRequired) TypeID() uint32 {
	return ReportChatSponsoredMessageResultPremiumRequiredTypeID
}

// TypeName returns name of type in TL schema.
func (*ReportChatSponsoredMessageResultPremiumRequired) TypeName() string {
	return "reportChatSponsoredMessageResultPremiumRequired"
}

// TypeInfo returns info about TL type.
func (r *ReportChatSponsoredMessageResultPremiumRequired) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "reportChatSponsoredMessageResultPremiumRequired",
		ID:   ReportChatSponsoredMessageResultPremiumRequiredTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (r *ReportChatSponsoredMessageResultPremiumRequired) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode reportChatSponsoredMessageResultPremiumRequired#770c2ed0 as nil")
	}
	b.PutID(ReportChatSponsoredMessageResultPremiumRequiredTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *ReportChatSponsoredMessageResultPremiumRequired) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode reportChatSponsoredMessageResultPremiumRequired#770c2ed0 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *ReportChatSponsoredMessageResultPremiumRequired) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode reportChatSponsoredMessageResultPremiumRequired#770c2ed0 to nil")
	}
	if err := b.ConsumeID(ReportChatSponsoredMessageResultPremiumRequiredTypeID); err != nil {
		return fmt.Errorf("unable to decode reportChatSponsoredMessageResultPremiumRequired#770c2ed0: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *ReportChatSponsoredMessageResultPremiumRequired) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode reportChatSponsoredMessageResultPremiumRequired#770c2ed0 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (r *ReportChatSponsoredMessageResultPremiumRequired) EncodeTDLibJSON(b tdjson.Encoder) error {
	if r == nil {
		return fmt.Errorf("can't encode reportChatSponsoredMessageResultPremiumRequired#770c2ed0 as nil")
	}
	b.ObjStart()
	b.PutID("reportChatSponsoredMessageResultPremiumRequired")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (r *ReportChatSponsoredMessageResultPremiumRequired) DecodeTDLibJSON(b tdjson.Decoder) error {
	if r == nil {
		return fmt.Errorf("can't decode reportChatSponsoredMessageResultPremiumRequired#770c2ed0 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("reportChatSponsoredMessageResultPremiumRequired"); err != nil {
				return fmt.Errorf("unable to decode reportChatSponsoredMessageResultPremiumRequired#770c2ed0: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// ReportChatSponsoredMessageResultClassName is schema name of ReportChatSponsoredMessageResultClass.
const ReportChatSponsoredMessageResultClassName = "ReportChatSponsoredMessageResult"

// ReportChatSponsoredMessageResultClass represents ReportChatSponsoredMessageResult generic type.
//
// Example:
//
//	g, err := tdapi.DecodeReportChatSponsoredMessageResult(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tdapi.ReportChatSponsoredMessageResultOk: // reportChatSponsoredMessageResultOk#754f721f
//	case *tdapi.ReportChatSponsoredMessageResultFailed: // reportChatSponsoredMessageResultFailed#7f1f9bc6
//	case *tdapi.ReportChatSponsoredMessageResultOptionRequired: // reportChatSponsoredMessageResultOptionRequired#34fc24b2
//	case *tdapi.ReportChatSponsoredMessageResultAdsHidden: // reportChatSponsoredMessageResultAdsHidden#e8eade1e
//	case *tdapi.ReportChatSponsoredMessageResultPremiumRequired: // reportChatSponsoredMessageResultPremiumRequired#770c2ed0
//	default: panic(v)
//	}
type ReportChatSponsoredMessageResultClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() ReportChatSponsoredMessageResultClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool

	EncodeTDLibJSON(b tdjson.Encoder) error
	DecodeTDLibJSON(b tdjson.Decoder) error
}

// DecodeReportChatSponsoredMessageResult implements binary de-serialization for ReportChatSponsoredMessageResultClass.
func DecodeReportChatSponsoredMessageResult(buf *bin.Buffer) (ReportChatSponsoredMessageResultClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case ReportChatSponsoredMessageResultOkTypeID:
		// Decoding reportChatSponsoredMessageResultOk#754f721f.
		v := ReportChatSponsoredMessageResultOk{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ReportChatSponsoredMessageResultClass: %w", err)
		}
		return &v, nil
	case ReportChatSponsoredMessageResultFailedTypeID:
		// Decoding reportChatSponsoredMessageResultFailed#7f1f9bc6.
		v := ReportChatSponsoredMessageResultFailed{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ReportChatSponsoredMessageResultClass: %w", err)
		}
		return &v, nil
	case ReportChatSponsoredMessageResultOptionRequiredTypeID:
		// Decoding reportChatSponsoredMessageResultOptionRequired#34fc24b2.
		v := ReportChatSponsoredMessageResultOptionRequired{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ReportChatSponsoredMessageResultClass: %w", err)
		}
		return &v, nil
	case ReportChatSponsoredMessageResultAdsHiddenTypeID:
		// Decoding reportChatSponsoredMessageResultAdsHidden#e8eade1e.
		v := ReportChatSponsoredMessageResultAdsHidden{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ReportChatSponsoredMessageResultClass: %w", err)
		}
		return &v, nil
	case ReportChatSponsoredMessageResultPremiumRequiredTypeID:
		// Decoding reportChatSponsoredMessageResultPremiumRequired#770c2ed0.
		v := ReportChatSponsoredMessageResultPremiumRequired{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ReportChatSponsoredMessageResultClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode ReportChatSponsoredMessageResultClass: %w", bin.NewUnexpectedID(id))
	}
}

// DecodeTDLibJSONReportChatSponsoredMessageResult implements binary de-serialization for ReportChatSponsoredMessageResultClass.
func DecodeTDLibJSONReportChatSponsoredMessageResult(buf tdjson.Decoder) (ReportChatSponsoredMessageResultClass, error) {
	id, err := buf.FindTypeID()
	if err != nil {
		return nil, err
	}
	switch id {
	case "reportChatSponsoredMessageResultOk":
		// Decoding reportChatSponsoredMessageResultOk#754f721f.
		v := ReportChatSponsoredMessageResultOk{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ReportChatSponsoredMessageResultClass: %w", err)
		}
		return &v, nil
	case "reportChatSponsoredMessageResultFailed":
		// Decoding reportChatSponsoredMessageResultFailed#7f1f9bc6.
		v := ReportChatSponsoredMessageResultFailed{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ReportChatSponsoredMessageResultClass: %w", err)
		}
		return &v, nil
	case "reportChatSponsoredMessageResultOptionRequired":
		// Decoding reportChatSponsoredMessageResultOptionRequired#34fc24b2.
		v := ReportChatSponsoredMessageResultOptionRequired{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ReportChatSponsoredMessageResultClass: %w", err)
		}
		return &v, nil
	case "reportChatSponsoredMessageResultAdsHidden":
		// Decoding reportChatSponsoredMessageResultAdsHidden#e8eade1e.
		v := ReportChatSponsoredMessageResultAdsHidden{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ReportChatSponsoredMessageResultClass: %w", err)
		}
		return &v, nil
	case "reportChatSponsoredMessageResultPremiumRequired":
		// Decoding reportChatSponsoredMessageResultPremiumRequired#770c2ed0.
		v := ReportChatSponsoredMessageResultPremiumRequired{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ReportChatSponsoredMessageResultClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode ReportChatSponsoredMessageResultClass: %w", tdjson.NewUnexpectedID(id))
	}
}

// ReportChatSponsoredMessageResult boxes the ReportChatSponsoredMessageResultClass providing a helper.
type ReportChatSponsoredMessageResultBox struct {
	ReportChatSponsoredMessageResult ReportChatSponsoredMessageResultClass
}

// Decode implements bin.Decoder for ReportChatSponsoredMessageResultBox.
func (b *ReportChatSponsoredMessageResultBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode ReportChatSponsoredMessageResultBox to nil")
	}
	v, err := DecodeReportChatSponsoredMessageResult(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.ReportChatSponsoredMessageResult = v
	return nil
}

// Encode implements bin.Encode for ReportChatSponsoredMessageResultBox.
func (b *ReportChatSponsoredMessageResultBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.ReportChatSponsoredMessageResult == nil {
		return fmt.Errorf("unable to encode ReportChatSponsoredMessageResultClass as nil")
	}
	return b.ReportChatSponsoredMessageResult.Encode(buf)
}

// DecodeTDLibJSON implements bin.Decoder for ReportChatSponsoredMessageResultBox.
func (b *ReportChatSponsoredMessageResultBox) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("unable to decode ReportChatSponsoredMessageResultBox to nil")
	}
	v, err := DecodeTDLibJSONReportChatSponsoredMessageResult(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.ReportChatSponsoredMessageResult = v
	return nil
}

// EncodeTDLibJSON implements bin.Encode for ReportChatSponsoredMessageResultBox.
func (b *ReportChatSponsoredMessageResultBox) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil || b.ReportChatSponsoredMessageResult == nil {
		return fmt.Errorf("unable to encode ReportChatSponsoredMessageResultClass as nil")
	}
	return b.ReportChatSponsoredMessageResult.EncodeTDLibJSON(buf)
}
