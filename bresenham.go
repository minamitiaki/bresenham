/*
 * @Author: chenjingwei
 * @Date: 2020-05-11 09:37:26
 * @Last Modified by: chenjingwei
 * @Last Modified time: 2020-11-22 19:36:30
 */

package util

import (
	"game/def"
)

func Abs(v int16) int16 {
	if v < 0 {
		return -v
	}
	return v
}

//[]
func Path(x0, x1, y0, y1 int16) []def.Pos {
	rlt := make([]def.Pos, 0)
	steep := Abs(y1-y0) > Abs(x1-x0)
	if steep {
		x0, y0 = y0, x0
		x1, y1 = y1, x1
	}

	if x0 > x1 {
		x0, x1 = x1, x0
		y0, y1 = y1, y0
	}

	dx := x1 - x0
	dy := Abs(y1 - y0)
	err := dx >> 1

	var ystep int16
	if y0 < y1 {
		ystep = 1
	} else {
		ystep = -1
	}

	y := y0
	for x := x0; x <= x1; x++ {
		if steep {
			rlt = append(rlt, def.Pos{X: y, Y: x})
		} else {
			rlt = append(rlt, def.Pos{X: x, Y: y})
		}
		err -= dy
		if err < 0 {
			y += ystep
			err += dx
		}
	}
	return rlt
}
