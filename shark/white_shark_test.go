package shark

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"testdoubles/mocks"
	"testdoubles/prey"
)

func TestWhiteShark_Hunt_Success(t *testing.T) {
	// Crear el mock del simulador
	sim := mocks.CatchSimulatorMock{}
	sim.On("CanCatch", 100.0, 144.0, 252.0).Return(true)
	sim.On("GetLinearDistance", [2]float64{0, 0}).Return(100.0)

	// Crear el stub del atún
	tuna := &prey.TunaStub{MaxSpeed: 252.0}

	// Crear el tiburón blanco con el simulador y cazar al atún
	shark := CreateWhiteShark(&sim)
	err := shark.Hunt(tuna)

	// Afirmar que no hubo error y que se llamó al método GetLinearDistance
	assert.NoError(t, err)
	sim.AssertCalled(t, "GetLinearDistance", [2]float64{0, 0})
}

func TestWhiteShark_Hunt_TooSlow(t *testing.T) {
	// Crear el mock del simulador
	sim := mocks.CatchSimulatorMock{}
	sim.On("CanCatch", 100.0, 144.0, 100.0).Return(false)

	// Crear el stub del atún
	tuna := &prey.TunaStub{MaxSpeed: 100.0}

	// Crear el tiburón blanco con el simulador y tratar de cazar al atún
	shark := CreateWhiteShark(&sim)
	err := shark.Hunt(tuna)

	// Afirmar que hubo un error y que el método GetLinearDistance no fue llamado
	assert.Error(t, err)
	sim.AssertNotCalled(t, "GetLinearDistance")
}

func TestWhiteShark_Hunt_TooFar(t *testing.T) {
	// Crear el mock del simulador
	sim := mocks.CatchSimulatorMock{}
	sim.On("CanCatch", 10000.0, 144.0, 252.0).Return(false)

	// Crear el stub del atún
	tuna := &prey.TunaStub{MaxSpeed: 252.0}

	// Crear el tiburón blanco con el simulador y tratar de cazar al atún
	shark := CreateWhiteShark(&sim)
	err := shark.Hunt(tuna)

	// Afirmar que hubo un error y que el método GetLinearDistance fue llamado
	assert.Error(t, err)
	sim.AssertCalled(t, "GetLinearDistance", [2]float64{0, 0})
}
