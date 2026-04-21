package main

import "context"

type contextKey string

const requestIDKey contextKey = "request_id"

// ✅ ADD THIS
func getRequestID(ctx context.Context) string {
	if v, ok := ctx.Value(requestIDKey).(string); ok {
		return v
	}
	return ""
}
