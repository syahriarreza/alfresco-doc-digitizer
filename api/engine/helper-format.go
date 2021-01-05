package engine

import "github.com/eaciit/toolkit"

// FormatNumberDigit format number into string with certain length
func FormatNumberDigit(number, length int) string {
	nStr := toolkit.ToString(number)
	if len(nStr) > length {
		return nStr
	}

	reqLen := length - len(nStr)
	for i := 0; i < reqLen; i++ {
		nStr = "0" + nStr
	}

	return nStr
}
