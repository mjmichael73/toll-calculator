package main

import "github.com/mjmichael73/toll-calculator/types"

type LogMiddleware struct {
	next DataProducer
}
func NewLogMiddleware(next DataProducer) *LogMiddleware {
	return &LogMiddleware{
		next: next,
	}
}
func (l *LogMiddleware) produceData(data types.OBUData) error {
	return l.next.produceData(data)
}