package gotime

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	t1 := NewTimeByStr("1:1:1")
	tim := Now()
	fmt.Println(tim.String())
	fmt.Println(tim.SecondsOfDay())
	fmt.Println(tim.SecondsOfHour())

	fmt.Println(tim.Sub(t1).Complete())
}
