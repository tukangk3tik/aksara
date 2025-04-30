package utils

import "fmt"

func LogErrorMessageBuilder(Error string, TraceID string) string {
	return fmt.Sprintf("{trace_id: %s, error: %s}", TraceID, Error)
}
