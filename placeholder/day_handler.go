package placeholder

import (
	"fmt"
	"time"

	"platform/logging"
)

type DayHandler struct {
	logging.Logger
}

func (dh DayHandler) GetDay() string {
	return fmt.Sprintf("Day: %v", time.Now().Day())
}
