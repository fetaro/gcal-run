package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsCheckTiming1MinuteAgo(t *testing.T) {
	r := NewRunTimingCalculator(1)
	assert.False(t, r.IsRunTiming(12))
	assert.False(t, r.IsRunTiming(13))
	assert.True(t, r.IsRunTiming(14))
	assert.False(t, r.IsRunTiming(15))

	assert.False(t, r.IsRunTiming(27))
	assert.False(t, r.IsRunTiming(28))
	assert.True(t, r.IsRunTiming(29))
	assert.False(t, r.IsRunTiming(30))

	assert.False(t, r.IsRunTiming(42))
	assert.False(t, r.IsRunTiming(43))
	assert.True(t, r.IsRunTiming(44))
	assert.False(t, r.IsRunTiming(45))

	assert.False(t, r.IsRunTiming(57))
	assert.False(t, r.IsRunTiming(58))
	assert.True(t, r.IsRunTiming(59))
	assert.False(t, r.IsRunTiming(0))
}

func TestIsCheckTiming2MinuteAgo(t *testing.T) {
	r := NewRunTimingCalculator(2)
	assert.False(t, r.IsRunTiming(12))
	assert.True(t, r.IsRunTiming(13))
	assert.True(t, r.IsRunTiming(14))
	assert.False(t, r.IsRunTiming(15))

	assert.False(t, r.IsRunTiming(27))
	assert.True(t, r.IsRunTiming(28))
	assert.True(t, r.IsRunTiming(29))
	assert.False(t, r.IsRunTiming(30))

	assert.False(t, r.IsRunTiming(42))
	assert.True(t, r.IsRunTiming(43))
	assert.True(t, r.IsRunTiming(44))
	assert.False(t, r.IsRunTiming(45))

	assert.False(t, r.IsRunTiming(57))
	assert.True(t, r.IsRunTiming(58))
	assert.True(t, r.IsRunTiming(59))
	assert.False(t, r.IsRunTiming(0))
}

func TestIsCheckTiming3MinuteAgo(t *testing.T) {
	r := NewRunTimingCalculator(3)
	assert.True(t, r.IsRunTiming(12))
	assert.True(t, r.IsRunTiming(13))
	assert.True(t, r.IsRunTiming(14))
	assert.False(t, r.IsRunTiming(15))

	assert.True(t, r.IsRunTiming(27))
	assert.True(t, r.IsRunTiming(28))
	assert.True(t, r.IsRunTiming(29))
	assert.False(t, r.IsRunTiming(30))

	assert.True(t, r.IsRunTiming(42))
	assert.True(t, r.IsRunTiming(43))
	assert.True(t, r.IsRunTiming(44))
	assert.False(t, r.IsRunTiming(45))

	assert.True(t, r.IsRunTiming(57))
	assert.True(t, r.IsRunTiming(58))
	assert.True(t, r.IsRunTiming(59))
	assert.False(t, r.IsRunTiming(0))
}
