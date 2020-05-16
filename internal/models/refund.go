package models

import (
	"gitlab.com/thorchain/midgard/internal/clients/thorchain/types"
)

type EventRefund struct {
	Event
	Code   uint32 `json:"code,string"`
	Reason string `json:"reason"`
}

func NewRefundEvent(refund types.EventRefund, event types.Event) EventRefund {
	return EventRefund{
		Code:   uint32(refund.Code),
		Reason: refund.Reason,
		Event:  newEvent(event),
	}
}
