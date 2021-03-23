package gotime

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	tim := Now()
	fmt.Println(tim.String())
	fmt.Println(tim.SecondsOfDay())
	fmt.Println(tim.SecondsOfHour())
}
