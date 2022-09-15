package models

import (
	"fmt"
	"getcharzp.cn/define"
	"time"
)

type MyTime time.Time

func (t *MyTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%v\"", time.Time(*t).Format(define.DateLayout))), nil
}
