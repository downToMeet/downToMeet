// Code generated by "stringer -type=SessionKey"; DO NOT EDIT.

package impl

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[UserID-0]
	_ = x[FacebookState-1]
	_ = x[GoogleState-2]
}

const _SessionKey_name = "UserIDFacebookStateGoogleState"

var _SessionKey_index = [...]uint8{0, 6, 19, 30}

func (i SessionKey) String() string {
	if i < 0 || i >= SessionKey(len(_SessionKey_index)-1) {
		return "SessionKey(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _SessionKey_name[_SessionKey_index[i]:_SessionKey_index[i+1]]
}
