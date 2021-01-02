package helper

import (
	"errors"
	"fmt"
	"time"

	"github.com/eaciit/toolkit"
	"gorm.io/gorm"
)

// MakeID MakeID
func MakeID(prefix string, l int) string {
	if l < 20 {
		l = 20
	}
	p1 := toolkit.Date2String(time.Now(), "YYMMddHHmmss")
	lp2 := l - len(prefix) - len(p1)
	if lp2 <= 0 {
		if prefix == "" {
			return p1
		}
		return fmt.Sprintf("%s%s", prefix, p1)
	}
	p2 := toolkit.GenerateRandomString("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789", lp2)
	if prefix == "" {
		return fmt.Sprintf("%s%s", p1, p2)
	}
	return fmt.Sprintf("%s%s%s", prefix, p1, p2)
}

// IsNotFound IsNotFound
func IsNotFound(res *gorm.DB) bool {
	return res.Error != nil && errors.Is(res.Error, gorm.ErrRecordNotFound)
}
