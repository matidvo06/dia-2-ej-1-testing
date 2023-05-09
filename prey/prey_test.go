/*Se crea un mock para la interfaz Prey llamado mockPrey que devuelve valores
aleatorios para el método GetSpeed. En el primer caso de prueba, se utiliza
la implementación dreal de la interfaz Prey, mientras que en el segundo se
utiliza el mock para verificar que GetSpeed devuelve valores dentro del
rango esperado.*/

package prey

import (
	"math/rand"
	"testing"

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
