package formatting_string_review

import (
	"fmt"
)

func getSMSErrorString(cost float64, recipient string) string {
	return fmt.Sprintf("SMS that costs $%.2f to be sent to '%v' can not be sent", cost, recipient)
}

// SMS that costs $COST to be sent to 'RECIPIENT' can not be sent
