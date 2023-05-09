package shark

import (
	"testdoubles/prey"
	"testdoubles/simulator"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockStorage struct{}

func (m *mockStorage) GetValue(key string) interface{} {
	switch key {
	case "max time to catch":
		return 10.0
	default:
		return nil
	}
}

func TestWhiteShark_Hunt_Success(t *testing.T) {
	storage := &mockStorage{}
	catchSimulator := simulator.NewCatchSimulator(storage)
	prey := prey.CreateTuna()
	whiteShark := shark.CreateWhiteShark(catchSimulator)
	err := whiteShark.Hunt(prey)

	assert.NoError(t, err)
}

func TestWhiteShark_Hunt_TooSlow(t *testing.T) {
	storage := &mockStorage{}
	catchSimulator := simulator.NewCatchSimulator(storage)
	prey := &mockPrey{speed: 10}
	whiteShark := shark.CreateWhiteShark(catchSimulator)

	err := whiteShark.Hunt(prey)
	assert.Error(t, err)
}

func TestWhiteShark_Hunt_TooFar(t *testing.T) {
	storage := &mockStorage{}
	catchSimulator := simulator.NewCatchSimulator(storage)
	prey := &mockPrey{speed: 100}
	whiteShark := shark.CreateWhiteShark(catchSimulator)

	whiteShark.SetPosition([2]float64{1000, 1000})

	err := whiteShark.Hunt(prey)
	assert.Error(t, err)
	assert.True(t, catchSimulator.GetLinearDistanceCalled)
}

type mockPrey struct {
	speed float64
}

func (m *mockPrey) GetSpeed() float64 {
	return m.speed
}
