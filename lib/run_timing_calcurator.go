package lib

type RunTimingCalculator struct {
	minutesAgo int
}

func NewRunTimingCalculator(minutesAgo int) *RunTimingCalculator {
	return &RunTimingCalculator{minutesAgo: minutesAgo}
}

func (r *RunTimingCalculator) IsRunTiming(minute int) bool {
	if minute%15 > 14-r.minutesAgo {
		return true
	}
	return false
}
