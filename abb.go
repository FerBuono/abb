package diccionario

import (
	TDAPila "pila"
)

type nodoAbb[K comparable, V any] struct {
	izq   *nodoAbb[K, V]
	der   *nodoAbb[K, V]
	clave K
	dato  V
}

type abb[K comparable, V any] struct {
	raiz *nodoAbb[K, V]
	cant int
	cmp  func(K, K) int
}

type iterAbb[K comparable, V any] struct {
	abb  *abb[K, V]
	pila TDAPila.Pila[*nodoAbb[K, V]]
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	a := new(abb[K, V])
	a.cmp = funcion_cmp
	return a
}

// Primitivas del Diccionario

func (a *abb[K, V]) Guardar(clave K, dato V) {
	puntero := a.buscarPuntero(clave, a.raiz)
	if puntero != nil {
		puntero.dato = dato
	} else {
		puntero = a.crearNodo(clave, dato)
		a.cant++
	}
}

func (a *abb[K, V]) Pertenece(clave K) bool {
	return a.buscarPuntero(clave, a.raiz) != nil
}

func (a *abb[K, V]) Obtener(clave K) V {
	puntero := a.buscarPuntero(clave, a.raiz)
	if puntero != nil {
		return puntero.dato
	} else {
		panic("La clave no pertenece al diccionario")
	}
}

func (a *abb[K, V]) Borrar(clave K) V {
	puntero := a.buscarPuntero(clave, a.raiz)
	var dato V
	if puntero == nil {
		dato = puntero.dato
		if a.cantidadDeHijos(puntero) == 0 {
			puntero = nil
		} else if a.cantidadDeHijos(puntero) == 1 {
			nuevaClave := a.obtenerHijo(puntero).clave
			nuevoDato := a.Borrar(nuevaClave)
			puntero.clave = nuevaClave
			puntero.dato = nuevoDato
		} else {
			nuevaClave := a.buscarReemplazo(puntero.izq).clave
			nuevoDato := a.Borrar(nuevaClave)
			puntero.clave = nuevaClave
			puntero.dato = nuevoDato
		}
		a.cant--
	} else {
		panic("La clave no pertenece al diccionario")
	}

	return dato
}

func (a *abb[K, V]) Cantidad() int {
	return a.cant
}

func (a *abb[K, V]) Iterar(funcion func(K, V) bool) {
	a.iterar(a.raiz, funcion)
}

func (a *abb[K, V]) Iterador() IterDiccionario[K, V] {
	iter := new(iterAbb[K, V])
	iter.pila = TDAPila.CrearPilaDinamica[*nodoAbb[K, V]]()
	iter.pila.Apilar(a.raiz)
	a.apilarHijosIzq(a.raiz, iter.pila)
	return iter
}

// Primitivas del IterDiccionario

func (i *iterAbb[K, V]) HaySiguiente() bool {
	return !i.pila.EstaVacia()
}

func (i *iterAbb[K, V]) VerActual() (K, V) {
	return i.pila.VerTope().clave, i.pila.VerTope().dato
}

func (i *iterAbb[K, V]) Siguiente() K {
	nodo := i.pila.Desapilar()
	if nodo.der != nil {
		i.pila.Apilar(nodo.der)
		i.abb.apilarHijosIzq(nodo.der, i.pila)
	}

	return nodo.clave
}

// Primitivas del DiccionarioOrdenado

// Funciones y métodos auxiliares

func (a *abb[K, V]) crearNodo(clave K, dato V) *nodoAbb[K, V] {
	n := new(nodoAbb[K, V])
	n.clave = clave
	n.dato = dato
	return n
}

func (a *abb[K, V]) buscarPuntero(clave K, nodo *nodoAbb[K, V]) *nodoAbb[K, V] {
	if nodo == nil {
		return nodo
	}

	if nodo.izq != nil && a.cmp(clave, nodo.izq.clave) == 0 {
		return nodo.izq
	}
	if nodo.der != nil && a.cmp(clave, nodo.der.clave) == 0 {
		return nodo.der
	}

	if a.cmp(clave, nodo.clave) < 0 {
		return a.buscarPuntero(clave, nodo.izq)
	} else if a.cmp(clave, nodo.clave) > 0 {
		return a.buscarPuntero(clave, nodo.der)
	} else {
		return nodo
	}
}

func (a *abb[K, V]) buscarReemplazo(nodo *nodoAbb[K, V]) *nodoAbb[K, V] {
	if nodo.der == nil {
		return nodo
	} else {
		return a.buscarReemplazo(nodo.der)
	}
}

func (a *abb[K, V]) cantidadDeHijos(nodo *nodoAbb[K, V]) int {
	if nodo.izq != nil && nodo.der != nil {
		return 2
	} else if nodo.izq == nil || nodo.der == nil {
		return 1
	} else {
		return 0
	}
}

func (a *abb[K, V]) obtenerHijo(nodo *nodoAbb[K, V]) *nodoAbb[K, V] {
	if nodo.izq != nil {
		return nodo.izq
	} else {
		return nodo.der
	}
}

func (a *abb[K, V]) iterar(nodo *nodoAbb[K, V], f func(K, V) bool) {
	if nodo == nil {
		return
	}
	a.iterar(nodo.izq, f)
	f(nodo.clave, nodo.dato)
	a.iterar(nodo.der, f)
}

func (a *abb[K, V]) apilarHijosIzq(nodo *nodoAbb[K, V], pila TDAPila.Pila[*nodoAbb[K, V]]) {
	if nodo.izq == nil {
		return
	} else {
		pila.Apilar(nodo.izq)
		a.apilarHijosIzq(nodo.izq, pila)
	}
}
