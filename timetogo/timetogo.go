package timetogo

import "time"

// Durationer An object that can be converted to a duration
type Durationer interface {
	Duration() time.Duration
}

func times(v interface{}, d time.Duration) (res time.Duration) {
	// we check if v is a number and do the correct cast depending on results
	// TODO : this should be far more complete, maybe using a loop on all possible numerical type
	i, isInt := v.(int)
	if isInt {
		res = time.Duration(i) * d
		return
	}
	f64, isFloat64 := v.(float64)
	if isFloat64 {
		// multiply float by the duration
		product := f64 * float64(d)
		// cast product to an int (= number of nanosecond), then to a duration
		res = time.Duration(int(product))
		return
	}
	f32, isFloat32 := v.(float32)
	if isFloat32 {
		// multiply float by the duration
		product := f32 * float32(d)
		res = time.Duration(int(product))
		return
	}
	// TODO add panic return
	return
}

// Duration Converts an int to a duration (number of nanoseconds)
func Duration(v interface{}) (d time.Duration) {
	i, isInt := v.(int)
	if isInt {
		d = time.Duration(i)
		return
	}
	f64, isFloat64 := v.(float64)
	if isFloat64 {
		d = time.Duration(int(f64))
		return
	}
	f32, isFloat32 := v.(float32)
	if isFloat32 {
		d = time.Duration(int(f32))
		return
	}
	return
}

// Seconds Converts an value value in minutes
func Seconds(v interface{}) (d time.Duration) {
	d = times(v, time.Second)
	return
}

// Minutes Converts an value value in minutes
func Minutes(v interface{}) (d time.Duration) {
	d = times(v, time.Minute)
	return
}

// Hours Converts an value value in minutes
func Hours(v interface{}) (d time.Duration) {
	d = times(v, time.Hour)
	return
}

// Days Converts an value value in minutes
func Days(v interface{}) (d time.Duration) {
	// there are no unified definition of a day, so this is an approx
	d = times(v, time.Hour) * 24
	return
}
