// Code generated by "stringer -output=zsublayerflags_strings.go -type=SublayerFlags -trimprefix=SublayerFlags"; DO NOT EDIT.

package wf

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[SublayerFlagsPersistent-1]
}

const _SublayerFlags_name = "Persistent"

var _SublayerFlags_index = [...]uint8{0, 10}

func (i SublayerFlags) String() string {
	i -= 1
	if i >= SublayerFlags(len(_SublayerFlags_index)-1) {
		return "SublayerFlags(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _SublayerFlags_name[_SublayerFlags_index[i]:_SublayerFlags_index[i+1]]
}
