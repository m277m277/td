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

// PaymentsStarsRevenueStats represents TL type `payments.starsRevenueStats#c92bb73b`.
//
// See https://core.telegram.org/constructor/payments.starsRevenueStats for reference.
type PaymentsStarsRevenueStats struct {
	// RevenueGraph field of PaymentsStarsRevenueStats.
	RevenueGraph StatsGraphClass
	// Status field of PaymentsStarsRevenueStats.
	Status StarsRevenueStatus
	// UsdRate field of PaymentsStarsRevenueStats.
	UsdRate float64
}

// PaymentsStarsRevenueStatsTypeID is TL type id of PaymentsStarsRevenueStats.
const PaymentsStarsRevenueStatsTypeID = 0xc92bb73b

// Ensuring interfaces in compile-time for PaymentsStarsRevenueStats.
var (
	_ bin.Encoder     = &PaymentsStarsRevenueStats{}
	_ bin.Decoder     = &PaymentsStarsRevenueStats{}
	_ bin.BareEncoder = &PaymentsStarsRevenueStats{}
	_ bin.BareDecoder = &PaymentsStarsRevenueStats{}
)

func (s *PaymentsStarsRevenueStats) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.RevenueGraph == nil) {
		return false
	}
	if !(s.Status.Zero()) {
		return false
	}
	if !(s.UsdRate == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *PaymentsStarsRevenueStats) String() string {
	if s == nil {
		return "PaymentsStarsRevenueStats(nil)"
	}
	type Alias PaymentsStarsRevenueStats
	return fmt.Sprintf("PaymentsStarsRevenueStats%+v", Alias(*s))
}

// FillFrom fills PaymentsStarsRevenueStats from given interface.
func (s *PaymentsStarsRevenueStats) FillFrom(from interface {
	GetRevenueGraph() (value StatsGraphClass)
	GetStatus() (value StarsRevenueStatus)
	GetUsdRate() (value float64)
}) {
	s.RevenueGraph = from.GetRevenueGraph()
	s.Status = from.GetStatus()
	s.UsdRate = from.GetUsdRate()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PaymentsStarsRevenueStats) TypeID() uint32 {
	return PaymentsStarsRevenueStatsTypeID
}

// TypeName returns name of type in TL schema.
func (*PaymentsStarsRevenueStats) TypeName() string {
	return "payments.starsRevenueStats"
}

// TypeInfo returns info about TL type.
func (s *PaymentsStarsRevenueStats) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "payments.starsRevenueStats",
		ID:   PaymentsStarsRevenueStatsTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "RevenueGraph",
			SchemaName: "revenue_graph",
		},
		{
			Name:       "Status",
			SchemaName: "status",
		},
		{
			Name:       "UsdRate",
			SchemaName: "usd_rate",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *PaymentsStarsRevenueStats) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode payments.starsRevenueStats#c92bb73b as nil")
	}
	b.PutID(PaymentsStarsRevenueStatsTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *PaymentsStarsRevenueStats) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode payments.starsRevenueStats#c92bb73b as nil")
	}
	if s.RevenueGraph == nil {
		return fmt.Errorf("unable to encode payments.starsRevenueStats#c92bb73b: field revenue_graph is nil")
	}
	if err := s.RevenueGraph.Encode(b); err != nil {
		return fmt.Errorf("unable to encode payments.starsRevenueStats#c92bb73b: field revenue_graph: %w", err)
	}
	if err := s.Status.Encode(b); err != nil {
		return fmt.Errorf("unable to encode payments.starsRevenueStats#c92bb73b: field status: %w", err)
	}
	b.PutDouble(s.UsdRate)
	return nil
}

// Decode implements bin.Decoder.
func (s *PaymentsStarsRevenueStats) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode payments.starsRevenueStats#c92bb73b to nil")
	}
	if err := b.ConsumeID(PaymentsStarsRevenueStatsTypeID); err != nil {
		return fmt.Errorf("unable to decode payments.starsRevenueStats#c92bb73b: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *PaymentsStarsRevenueStats) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode payments.starsRevenueStats#c92bb73b to nil")
	}
	{
		value, err := DecodeStatsGraph(b)
		if err != nil {
			return fmt.Errorf("unable to decode payments.starsRevenueStats#c92bb73b: field revenue_graph: %w", err)
		}
		s.RevenueGraph = value
	}
	{
		if err := s.Status.Decode(b); err != nil {
			return fmt.Errorf("unable to decode payments.starsRevenueStats#c92bb73b: field status: %w", err)
		}
	}
	{
		value, err := b.Double()
		if err != nil {
			return fmt.Errorf("unable to decode payments.starsRevenueStats#c92bb73b: field usd_rate: %w", err)
		}
		s.UsdRate = value
	}
	return nil
}

// GetRevenueGraph returns value of RevenueGraph field.
func (s *PaymentsStarsRevenueStats) GetRevenueGraph() (value StatsGraphClass) {
	if s == nil {
		return
	}
	return s.RevenueGraph
}

// GetStatus returns value of Status field.
func (s *PaymentsStarsRevenueStats) GetStatus() (value StarsRevenueStatus) {
	if s == nil {
		return
	}
	return s.Status
}

// GetUsdRate returns value of UsdRate field.
func (s *PaymentsStarsRevenueStats) GetUsdRate() (value float64) {
	if s == nil {
		return
	}
	return s.UsdRate
}