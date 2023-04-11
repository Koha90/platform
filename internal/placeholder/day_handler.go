// Package placeholder ...
package placeholder

import (
	"fmt"
	"time"

	"github.com/koha90/platform/pkg/logging"
)

type DayHandler struct {
	logging.Logger
}

func (dh DayHandler) GetDay() string {
	return fmt.Sprintf("Day: %v", time.Now().Day())
}
