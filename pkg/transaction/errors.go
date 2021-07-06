package transaction

import "github.com/solace-iot-team/agent-sdk/pkg/util/errors"

// Transaction errors
var (
	ErrInRedactions = errors.Newf(1550, "error when redacting %v: %v")
)
