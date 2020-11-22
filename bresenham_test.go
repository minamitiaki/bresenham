/*
 * @Author: chenjingwei
 * @Date: 2020-05-11 10:40:36
 * @Last Modified by: chenjingwei
 * @Last Modified time: 2020-11-22 19:36:39
 */

package util

import (
	"testing"

	"github.com/bmizerany/assert"
)

//base
func Test1(t *testing.T) {
	rlt := Path(0, 2, 0, 1)
	assert.Equal(t, 3, len(rlt))
	assert.Equal(t, int16(0), rlt[0].X)
	assert.Equal(t, int16(0), rlt[0].Y)
	assert.Equal(t, int16(1), rlt[1].X)
	assert.Equal(t, int16(0), rlt[1].Y)
	assert.Equal(t, int16(2), rlt[2].X)
	assert.Equal(t, int16(1), rlt[2].Y)
}

//long
func Test2(t *testing.T) {
	rlt := Path(0, 6, 0, 4)
	assert.Equal(t, 7, len(rlt))
	assert.Equal(t, int16(0), rlt[0].X)
	assert.Equal(t, int16(0), rlt[0].Y)
	assert.Equal(t, int16(1), rlt[1].X)
	assert.Equal(t, int16(1), rlt[1].Y)
	assert.Equal(t, int16(2), rlt[2].X)
	assert.Equal(t, int16(1), rlt[2].Y)
	assert.Equal(t, int16(3), rlt[3].X)
	assert.Equal(t, int16(2), rlt[3].Y)
	assert.Equal(t, int16(4), rlt[4].X)
	assert.Equal(t, int16(3), rlt[4].Y)
	assert.Equal(t, int16(5), rlt[5].X)
	assert.Equal(t, int16(3), rlt[5].Y)
	assert.Equal(t, int16(6), rlt[6].X)
	assert.Equal(t, int16(4), rlt[6].Y)
}

//negative
func Test3(t *testing.T) {
	rlt := Path(0, 6, 4, 0)
	assert.Equal(t, 7, len(rlt))
	assert.Equal(t, int16(0), rlt[0].X)
	assert.Equal(t, int16(4), rlt[0].Y)
	assert.Equal(t, int16(1), rlt[1].X)
	assert.Equal(t, int16(3), rlt[1].Y)
	assert.Equal(t, int16(2), rlt[2].X)
	assert.Equal(t, int16(3), rlt[2].Y)
	assert.Equal(t, int16(3), rlt[3].X)
	assert.Equal(t, int16(2), rlt[3].Y)
	assert.Equal(t, int16(4), rlt[4].X)
	assert.Equal(t, int16(1), rlt[4].Y)
	assert.Equal(t, int16(5), rlt[5].X)
	assert.Equal(t, int16(1), rlt[5].Y)
	assert.Equal(t, int16(6), rlt[6].X)
	assert.Equal(t, int16(0), rlt[6].Y)
}

//steep
func Test4(t *testing.T) {
	rlt := Path(0, 4, 0, 6)
	assert.Equal(t, 7, len(rlt))
	assert.Equal(t, int16(0), rlt[0].X)
	assert.Equal(t, int16(0), rlt[0].Y)
	assert.Equal(t, int16(1), rlt[1].X)
	assert.Equal(t, int16(1), rlt[1].Y)
	assert.Equal(t, int16(1), rlt[2].X)
	assert.Equal(t, int16(2), rlt[2].Y)
	assert.Equal(t, int16(2), rlt[3].X)
	assert.Equal(t, int16(3), rlt[3].Y)
	assert.Equal(t, int16(3), rlt[4].X)
	assert.Equal(t, int16(4), rlt[4].Y)
	assert.Equal(t, int16(3), rlt[5].X)
	assert.Equal(t, int16(5), rlt[5].Y)
	assert.Equal(t, int16(4), rlt[6].X)
	assert.Equal(t, int16(6), rlt[6].Y)
}

//negative & steep
func Test5(t *testing.T) {
	rlt := Path(0, 4, 6, 0)
	assert.Equal(t, 7, len(rlt))
	assert.Equal(t, int16(4), rlt[0].X)
	assert.Equal(t, int16(0), rlt[0].Y)
	assert.Equal(t, int16(3), rlt[1].X)
	assert.Equal(t, int16(1), rlt[1].Y)
	assert.Equal(t, int16(3), rlt[2].X)
	assert.Equal(t, int16(2), rlt[2].Y)
	assert.Equal(t, int16(2), rlt[3].X)
	assert.Equal(t, int16(3), rlt[3].Y)
	assert.Equal(t, int16(1), rlt[4].X)
	assert.Equal(t, int16(4), rlt[4].Y)
	assert.Equal(t, int16(1), rlt[5].X)
	assert.Equal(t, int16(5), rlt[5].Y)
	assert.Equal(t, int16(0), rlt[6].X)
	assert.Equal(t, int16(6), rlt[6].Y)
}

//horizon
func Test6(t *testing.T) {
	rlt := Path(0, 6, 4, 4)
	assert.Equal(t, 7, len(rlt))
	assert.Equal(t, int16(0), rlt[0].X)
	assert.Equal(t, int16(4), rlt[0].Y)
	assert.Equal(t, int16(1), rlt[1].X)
	assert.Equal(t, int16(4), rlt[1].Y)
	assert.Equal(t, int16(2), rlt[2].X)
	assert.Equal(t, int16(4), rlt[2].Y)
	assert.Equal(t, int16(3), rlt[3].X)
	assert.Equal(t, int16(4), rlt[3].Y)
	assert.Equal(t, int16(4), rlt[4].X)
	assert.Equal(t, int16(4), rlt[4].Y)
	assert.Equal(t, int16(5), rlt[5].X)
	assert.Equal(t, int16(4), rlt[5].Y)
	assert.Equal(t, int16(6), rlt[6].X)
	assert.Equal(t, int16(4), rlt[6].Y)
}

//vertical
func Test7(t *testing.T) {
	rlt := Path(4, 4, 0, 6)
	assert.Equal(t, 7, len(rlt))
	assert.Equal(t, int16(4), rlt[0].X)
	assert.Equal(t, int16(0), rlt[0].Y)
	assert.Equal(t, int16(4), rlt[1].X)
	assert.Equal(t, int16(1), rlt[1].Y)
	assert.Equal(t, int16(4), rlt[2].X)
	assert.Equal(t, int16(2), rlt[2].Y)
	assert.Equal(t, int16(4), rlt[3].X)
	assert.Equal(t, int16(3), rlt[3].Y)
	assert.Equal(t, int16(4), rlt[4].X)
	assert.Equal(t, int16(4), rlt[4].Y)
	assert.Equal(t, int16(4), rlt[5].X)
	assert.Equal(t, int16(5), rlt[5].Y)
	assert.Equal(t, int16(4), rlt[6].X)
	assert.Equal(t, int16(6), rlt[6].Y)
}
