// Code generated by "stringer -type=Names"; DO NOT EDIT.

package worker

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Unknown-0]
	_ = x[Bashful-1]
	_ = x[Doc-2]
	_ = x[Dopey-3]
	_ = x[Grumpy-4]
	_ = x[Happy-5]
	_ = x[Sleepy-6]
	_ = x[Sneezy-7]
}

const _Names_name = "UnknownBashfulDocDopeyGrumpyHappySleepySneezy"

var _Names_index = [...]uint8{0, 7, 14, 17, 22, 28, 33, 39, 45}

func (i Names) String() string {
	if i < 0 || i >= Names(len(_Names_index)-1) {
		return "Names(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Names_name[_Names_index[i]:_Names_index[i+1]]
}
