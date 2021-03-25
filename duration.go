package gotime

// 时间间隔
type Duration struct {
	Seconds int
}

func NewDuration(hour int, minute int, second int) *Duration{
	return &Duration{Seconds: hour_ * hour + minute_ * minute + second}
}

func (d *Duration)Hour() int {
	return d.Seconds / hour_
}

func (d *Duration)Minute() int {
	return d.Seconds / minute_
}

func (d *Duration)Second() int {
	return d.Seconds
}

func (d *Duration) Complete() (hour int, minute int, second int){
	hour = d.Hour()
	minute = (d.Seconds - hour * hour_) / minute_
	second = d.Seconds - hour * hour_ - minute * minute_

	return
}

func (d *Duration) MinuteSecond() (minute int, second int) {
	minute = d.Seconds / minute_
	second = d.Seconds % minute_
	return
}