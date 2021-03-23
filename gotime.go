package gotime

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	day_    = hour_ * 24
	hour_   = minute_ * 60
	minute_ = second_ * 60
	second_ = 1
)

type Time struct {
	Hour   int // 0 - 23
	Minute int
	Second int
	//Millisecond int
}

func NewTime(hour int, minute int, second int) *Time {
	return &Time{
		Hour:   hour,
		Minute: minute,
		Second: second,
	}
}

func NewTimeByStr(timeStr string) *Time {
	timeList := strings.Split(timeStr, ":")
	hour, _ := strconv.Atoi(timeList[0])
	minute, _ := strconv.Atoi(timeList[1])
	second, _ := strconv.Atoi(timeList[2])
	return NewTime(hour, minute, second)
}

func (t *Time) String(layout ...string) string {
	if len(layout) > 0 {
		return fmt.Sprintf(layout[0], t.Hour, t.Minute, t.Second)
	}
	return fmt.Sprintf("%02d:%02d:%02d", t.Hour, t.Minute, t.Second)
}

func (t *Time) Clone() *Time {
	return NewTime(t.Hour, t.Minute, t.Second)
}

func (t *Time) Copy(time *Time) {
	t.Hour = time.Hour
	t.Minute = time.Minute
	t.Second = time.Second
}

// 返回当前时间是这天的第几秒
func (t *Time) SecondsOfDay() int {
	return t.Hour*60*60 + t.Minute*60 + t.Second
}

func (t *Time) SecondsOfHour() int {
	return t.Minute*60 + t.Second
}

// 从秒反推到时间
func turnSecondsToTime(x int) *Time {
	if x < 0 {
		left := day_ + x
		for left < 0 {
			left = day_ + left
		}
		return turnSecondsToTime(left)
	}

	if x >= day_ {
		x = x % day_
	}

	h := x / hour_
	x = x % hour_
	m := x / minute_
	x = x % minute_

	return &Time{
		Hour:   h,
		Minute: m,
		Second: x,
	}
}

func (t *Time) AddSeconds(x int) {
	seconds := t.SecondsOfDay() + x
	t.Copy(turnSecondsToTime(seconds))
}

func (t *Time) SubSeconds(x int) {
	t.AddSeconds(-x)
}

func (t *Time) AddMinutes(x int) {
	t.AddSeconds(x * minute_)
}

func (t *Time) SubMinutes(x int) {
	t.AddMinutes(-x)
}

func (t *Time) AddHour(x int) {
	t.AddSeconds(x * hour_)
}

func (t *Time) SubHour(x int) {
	t.AddHour(-x)
}

func (t *Time) Early(time *Time) bool {
	return t.SecondsOfDay() < time.SecondsOfDay()
}

func (t *Time) Later(time *Time) bool {
	return t.SecondsOfDay() > time.SecondsOfDay()
}

func (t *Time) Equal(time *Time) bool {
	return t.SecondsOfDay() == time.SecondsOfDay()
}

func (t *Time) EarlyEqual(time *Time) bool {
	return t.SecondsOfDay() <= time.SecondsOfDay()
}

func (t *Time) LaterEqual(time *Time) bool {
	return t.SecondsOfDay() >= time.SecondsOfDay()
}

func (t *Time) Increase() {
	t.AddSeconds(1)
}

func (t *Time) IncreaseMinute() {
	t.AddMinutes(1)
}

func (t *Time) IncreaseHour() {
	t.AddHour(1)
}

func (t *Time) Reduce() {
	t.SubSeconds(1)
}

func (t *Time) ReduceMinute() {
	t.SubMinutes(1)
}

func (t *Time) ReduceHour() {
	t.SubHour(1)
}

func (t *Time) Sub(time *Time) int {
	return t.SecondsOfDay() - time.SecondsOfDay()
}

func (t *Time) Accurate(x int) {
	t.Copy(turnSecondsToTime(x))
}

func Now() *Time {
	s := strings.Split(fmt.Sprintf("%v", time.Now()), " ")[1]
	tim := strings.Split(s, ".")[0]
	return NewTimeByStr(tim)
}
