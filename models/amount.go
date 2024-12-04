package models

import (
	"fmt"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"strconv"
)

// Amount is currency amount data type
type Amount int64

func (a Amount) PersianString() string {
	return message.NewPrinter(language.Persian).Sprintf("%d", a)
}

func (a Amount) String() string {
	return fmt.Sprintf("%d", a)
}

func ParseAmount(in string) Amount {
	v, err := strconv.ParseInt(in, 10, 64)
	if err != nil {
		return 0
	}

	return Amount(v)
}
