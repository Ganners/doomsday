// generated by stringer -type=monthOfYear; DO NOT EDIT

package doomsday

import "fmt"

const _monthOfYear_name = "JanuaryFebruaryMarchAprilMayJuneJulyAugustSeptemberOctoberNovemberDecember"

var _monthOfYear_index = [...]uint8{0, 7, 15, 20, 25, 28, 32, 36, 42, 51, 58, 66, 74}

func (i monthOfYear) String() string {
	i -= 1
	if i < 0 || i+1 >= monthOfYear(len(_monthOfYear_index)) {
		return fmt.Sprintf("monthOfYear(%d)", i+1)
	}
	return _monthOfYear_name[_monthOfYear_index[i]:_monthOfYear_index[i+1]]
}
