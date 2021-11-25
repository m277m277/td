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

// SearchStickerSetsRequest represents TL type `searchStickerSets#bf7d307b`.
type SearchStickerSetsRequest struct {
	// Query to search for
	Query string
}

// SearchStickerSetsRequestTypeID is TL type id of SearchStickerSetsRequest.
const SearchStickerSetsRequestTypeID = 0xbf7d307b

// Ensuring interfaces in compile-time for SearchStickerSetsRequest.
var (
	_ bin.Encoder     = &SearchStickerSetsRequest{}
	_ bin.Decoder     = &SearchStickerSetsRequest{}
	_ bin.BareEncoder = &SearchStickerSetsRequest{}
	_ bin.BareDecoder = &SearchStickerSetsRequest{}
)

func (s *SearchStickerSetsRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Query == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SearchStickerSetsRequest) String() string {
	if s == nil {
		return "SearchStickerSetsRequest(nil)"
	}
	type Alias SearchStickerSetsRequest
	return fmt.Sprintf("SearchStickerSetsRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SearchStickerSetsRequest) TypeID() uint32 {
	return SearchStickerSetsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SearchStickerSetsRequest) TypeName() string {
	return "searchStickerSets"
}

// TypeInfo returns info about TL type.
func (s *SearchStickerSetsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "searchStickerSets",
		ID:   SearchStickerSetsRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Query",
			SchemaName: "query",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SearchStickerSetsRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode searchStickerSets#bf7d307b as nil")
	}
	b.PutID(SearchStickerSetsRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SearchStickerSetsRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode searchStickerSets#bf7d307b as nil")
	}
	b.PutString(s.Query)
	return nil
}

// Decode implements bin.Decoder.
func (s *SearchStickerSetsRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode searchStickerSets#bf7d307b to nil")
	}
	if err := b.ConsumeID(SearchStickerSetsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode searchStickerSets#bf7d307b: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SearchStickerSetsRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode searchStickerSets#bf7d307b to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode searchStickerSets#bf7d307b: field query: %w", err)
		}
		s.Query = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SearchStickerSetsRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode searchStickerSets#bf7d307b as nil")
	}
	b.ObjStart()
	b.PutID("searchStickerSets")
	b.FieldStart("query")
	b.PutString(s.Query)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SearchStickerSetsRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode searchStickerSets#bf7d307b to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("searchStickerSets"); err != nil {
				return fmt.Errorf("unable to decode searchStickerSets#bf7d307b: %w", err)
			}
		case "query":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode searchStickerSets#bf7d307b: field query: %w", err)
			}
			s.Query = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetQuery returns value of Query field.
func (s *SearchStickerSetsRequest) GetQuery() (value string) {
	return s.Query
}

// SearchStickerSets invokes method searchStickerSets#bf7d307b returning error if any.
func (c *Client) SearchStickerSets(ctx context.Context, query string) (*StickerSets, error) {
	var result StickerSets

	request := &SearchStickerSetsRequest{
		Query: query,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}