package diccionario_test

import (
	TDA_abb "diccionario"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestArbolVacio(t *testing.T) {
	arbol := TDA_abb.CrearABB[int, int]()
	require.NotNil(t, arbol)
}
