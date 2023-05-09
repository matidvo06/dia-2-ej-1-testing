package simulator

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockCatchSimulator struct {
	maxTimeToCatch float64
	distance       float64
	speed          float64
	catchSpeed     float64
	isSpyCalled    bool
}

func (m *mockCatchSimulator) CanCatch(distance float64, speed float64, catchSpeed float64) bool {
	return distance/m.speed <= m.maxTimeToCatch
}

func (m *mockCatchSimulator) GetLinearDistance(position [2]float64) float64 {
	m.isSpyCalled = true
	dx := math.Abs(position[0] - 1.0)
	dy := math.Abs(position[1] - 1.0)
	return math.Sqrt(dx*dx + dy*dy)
}

func TestCatchSimulator_CanCatch(t *testing.T) {
	maxTimeToCatch := 10.0
	distance := 500.0
	speed := 20.0
	catchSpeed := 5.0

	mock := &mockCatchSimulator{
		maxTimeToCatch: maxTimeToCatch,
		distance:       distance,
		speed:          speed,
		catchSpeed:     catchSpeed,
	}

	canCatch := mock.CanCatch(distance, speed, catchSpeed)
	assert.True(t, canCatch)

	distance = 50.0
	canCatch = mock.CanCatch(distance, speed, catchSpeed)
	assert.False(t, canCatch)
}

func TestCatchSimulator_GetLinearDistance(t *testing.T) {
	mock := &mockCatchSimulator{}
	position := [2]float64{3.0, 4.0}

	distance := mock.GetLinearDistance(position)
	assert.True(t, mock.isSpyCalled)
	assert.InDelta(t, distance, 3.60555127, 0.0001)
}
