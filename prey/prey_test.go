package prey

import (
	"math/rand"
	"testing"

	"github.com/matidvo06/dia-2-ej-1-testing/prey"
	"github.com/stretchr/testify/assert"
)

type mockPrey struct {
	maxSpeed float64
}

func (m *mockPrey) GetSpeed() float64 {
	return m.maxSpeed * rand.Float64()
}

func TestTuna_GetSpeed(t *testing.T) {
	tuna := prey.CreateTuna()
	assert.NotNil(t, tuna)

	// Caso 1
	expectedSpeed := tuna.GetSpeed()
	assert.GreaterOrEqual(t, expectedSpeed, 0.0, "Speed should be greater or equal than 0")
	assert.LessOrEqual(t, expectedSpeed, 252.0, "Speed should be less or equal than 252")

	// Caso 2
	mock := &mockPrey{maxSpeed: 150.0}
	tuna = mock
	expectedSpeed = tuna.GetSpeed()
	assert.GreaterOrEqual(t, expectedSpeed, 0.0, "Speed should be greater or equal than 0")
	assert.LessOrEqual(t, expectedSpeed, 150.0, "Speed should be less or equal than 252")
}
