package kit

import (
	"time"

	"github.com/mattgonewild/kit/internal/help"
)

func BoolToInt(true bool) int { return help.BoolToInt(true) }
func UnixNano() int64         { return time.Now().UnixNano() }
