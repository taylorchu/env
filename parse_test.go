package env

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	assert.Equal(t, "fallback1", String("test1", "fallback1"))

	os.Setenv("test2", "test2")
	assert.Equal(t, "test2", String("test2", "fallback2"))
}

func TestDuration(t *testing.T) {
	assert.Equal(t, 10*time.Minute, Duration("test1", "10m"))

	os.Setenv("test2", "1m")
	assert.Equal(t, time.Minute, Duration("test2", "10m"))
}

func TestBool(t *testing.T) {
	assert.Equal(t, false, Bool("test1", "false"))

	os.Setenv("test2", "1")
	assert.Equal(t, true, Bool("test2", "false"))
}

func TestFloat64(t *testing.T) {
	assert.Equal(t, 1.1, Float64("test1", "1.1"))

	os.Setenv("test2", "1.3")
	assert.Equal(t, 1.3, Float64("test2", "1.1"))
}

func TestInt64(t *testing.T) {
	assert.Equal(t, int64(3), Int64("test1", "3"))

	os.Setenv("test2", "12")
	assert.Equal(t, int64(12), Int64("test2", "3"))
}

func TestUint64(t *testing.T) {
	assert.Equal(t, uint64(3), Uint64("test1", "3"))

	os.Setenv("test2", "12")
	assert.Equal(t, uint64(12), Uint64("test2", "3"))
}
