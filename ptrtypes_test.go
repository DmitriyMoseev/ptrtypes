package ptrtypes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructors(t *testing.T) {
	assert.Equal(t, true, *NewBool(true))
	assert.Equal(t, false, *NewBool(false))
	assert.Equal(t, byte(1), *NewByte(1))
	assert.Equal(t, float32(3.2), *NewFloat32(3.2))
	assert.Equal(t, float64(6.4), *NewFloat64(6.4))
	assert.Equal(t, 7, *NewInt(7))
	assert.Equal(t, int8(8), *NewInt8(8))
	assert.Equal(t, int16(16), *NewInt16(16))
	assert.Equal(t, int32(32), *NewInt32(32))
	assert.Equal(t, int64(64), *NewInt64(64))
	assert.Equal(t, "hello world!", *NewString("hello world!"))
	assert.Equal(t, uint(7), *NewUInt(7))
	assert.Equal(t, uint8(8), *NewUInt8(8))
	assert.Equal(t, uint16(16), *NewUInt16(16))
	assert.Equal(t, uint32(32), *NewUInt32(32))
	assert.Equal(t, uint64(64), *NewUInt64(64))
}
