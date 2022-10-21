package pila_test

import (
	TDAPila "pila"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPilaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.True(t, pila.EstaVacia())
}

func TestApilarDesapilarElementos(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[string]()
	pila.Apilar("hola")
	require.False(t, pila.EstaVacia())
	pila.Apilar("mundo,")
	require.False(t, pila.EstaVacia())
	pila.Apilar("esto es")
	require.False(t, pila.EstaVacia())
	pila.Apilar("una prueba")
	require.False(t, pila.EstaVacia())
	require.Equal(t, "una prueba", pila.VerTope())
	require.False(t, pila.EstaVacia())
	require.Equal(t, "una prueba", pila.Desapilar())
	require.False(t, pila.EstaVacia())
	require.Equal(t, "esto es", pila.VerTope())
	require.False(t, pila.EstaVacia())
	require.Equal(t, "esto es", pila.Desapilar())
	require.False(t, pila.EstaVacia())
	require.Equal(t, "mundo,", pila.VerTope())
	require.False(t, pila.EstaVacia())
	require.Equal(t, "mundo,", pila.Desapilar())
	require.False(t, pila.EstaVacia())
	require.Equal(t, "hola", pila.VerTope())
	require.False(t, pila.EstaVacia())
	require.Equal(t, "hola", pila.Desapilar())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.True(t, pila.EstaVacia())
}

func TestVolumen(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	for i := 0; i <= 1000; i++ {
		pila.Apilar(i)
		require.Equal(t, i, pila.VerTope())
		require.False(t, pila.EstaVacia())
	}
	require.False(t, pila.EstaVacia())
	for j := 1000; j >= 0; j-- {
		require.Equal(t, j, pila.VerTope())
		require.False(t, pila.EstaVacia())
		pila.Desapilar()
	}
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.True(t, pila.EstaVacia())
}
