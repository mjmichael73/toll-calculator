package main

import (
	"time"

	"github.com/mjmichael73/toll-calculator/types"
	"github.com/sirupsen/logrus"
)

type LogMiddleware struct {
	next Aggregator
}

func NewLogMiddleware(next Aggregator) Aggregator {
	return &LogMiddleware{
		next: next,
	}
}

func (m *LogMiddleware) AggregateDistance(distance types.Distance) (err error) {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"took": time.Since(start),
			"err":  err,
		}).Info("Aggregate distance")
	}(time.Now())
	err = m.next.AggregateDistance(distance)
	return
}

func (m *LogMiddleware) CalculateInvoice(obuId int) (inv *types.Invoice, err error) {
	defer func(start time.Time) {
		var (
			distance float64
			amount   float64
		)
		if inv != nil {
			distance = inv.TotalDistance
			amount = inv.TotalAmount
		}
		logrus.WithFields(logrus.Fields{
			"took":        time.Since(start),
			"err":         err,
			"obuId":       obuId,
			"totalAmount": amount,
			"distance":    distance,
		}).Info("Aggregate distance")
	}(time.Now())
	inv, err = m.next.CalculateInvoice(obuId)
	return
}
