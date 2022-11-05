package diccionario

import (
	TDAPila "diccionario/pila"
	"fmt"
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
	abb   *abb[K, V]
	pila  TDAPila.Pila[*nodoAbb[K, V]]
	desde *K
	hasta *K
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	a := new(abb[K, V])
	a.cmp = funcion_cmp
	return a
}

// Primitivas del Diccionario

func (a *abb[K, V]) Guardar(clave K, dato V) {
	puntero := a.buscarPuntero(clave, &a.raiz)
	if *puntero == nil {
		*puntero = &nodoAbb[K, V]{clave: clave, dato: dato}
		a.cant++
	} else {
		(*puntero).dato = dato
	}
}

func (a *abb[K, V]) verArbol() {
	if a.cant == 0 {
		return
	}
	for iter := a.Iterador(); iter.HaySiguiente(); {
		clave, dato := iter.VerActual()
		fmt.Print("- {", clave, ":", dato, "} -")
		iter.Siguiente()
	}
	fmt.Println("\n-----------")
}

func (a *abb[K, V]) Pertenece(clave K) bool {
	return *(a.buscarPuntero(clave, &a.raiz)) != nil
}

func (a *abb[K, V]) Obtener(clave K) V {
	puntero := a.buscarPuntero(clave, &a.raiz)
	if *puntero == nil {
		panic("La clave no pertenece al diccionario")
	}
	return (*puntero).dato
}

func (a *abb[K, V]) Borrar(clave K) V {
	puntero := a.buscarPuntero(clave, &a.raiz)
	return a.borrar(puntero)
}

func (a *abb[K, V]) Cantidad() int {
	return a.cant
}

func (a *abb[K, V]) Iterar(funcion func(K, V) bool) {
	a.iterarPorRango(a.raiz, funcion, nil, nil)
}

func (a *abb[K, V]) Iterador() IterDiccionario[K, V] {
	return a.IteradorRango(nil, nil)
}

// Primitivas del IterDiccionario

func (i *iterAbb[K, V]) HaySiguiente() bool {
	return !i.pila.EstaVacia()
}

func (i *iterAbb[K, V]) VerActual() (K, V) {
	if !i.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return i.pila.VerTope().clave, i.pila.VerTope().dato
}

func (i *iterAbb[K, V]) Siguiente() K {
	if !i.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	nodo := i.pila.Desapilar()
	i.apilarHijosIzq(nodo.der)
	return nodo.clave
}

// Primitivas del DiccionarioOrdenado

func (a *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	a.iterarPorRango(a.raiz, visitar, desde, hasta)
}

func (a *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	iter := new(iterAbb[K, V])
	iter.abb = a
	iter.pila = TDAPila.CrearPilaDinamica[*nodoAbb[K, V]]()
	iter.desde = desde
	iter.hasta = hasta
	primero := iter.buscarPrimero(a.raiz, desde)
	if primero != nil {
		iter.pila.Apilar(primero)
		iter.apilarHijosPrimero(primero)
	}
	return iter
}

// Funciones y métodos auxiliares

func (a *abb[K, V]) buscarPuntero(clave K, nodo **nodoAbb[K, V]) **nodoAbb[K, V] {
	if *nodo == nil {
		return nodo
	}
	if a.cmp(clave, (*nodo).clave) < 0 {
		if (*nodo).izq == nil || a.cmp(clave, (*nodo).izq.clave) == 0 {
			return &(*nodo).izq
		}
		return a.buscarPuntero(clave, &(*nodo).izq)
	} else if a.cmp(clave, (*nodo).clave) > 0 {
		if (*nodo).der == nil || a.cmp(clave, (*nodo).der.clave) == 0 {
			return &(*nodo).der
		}
		return a.buscarPuntero(clave, &(*nodo).der)
	} else {
		return nodo
	}
}

func (a *abb[K, V]) borrar(puntero **nodoAbb[K, V]) V {
	if *puntero == nil {
		panic("La clave no pertenece al diccionario")
	}
	dato := (*puntero).dato
	if a.cantidadDeHijos(puntero) == 0 {
		*puntero = nil
	} else if a.cantidadDeHijos(puntero) == 1 {
		reemplazo := a.obtenerHijo(puntero)
		*puntero = *reemplazo
	} else {
		reemplazo := a.buscarReemplazo(&(*puntero).izq)
		nuevaClave, nuevoDato := (*reemplazo).clave, (*reemplazo).dato
		a.borrar(reemplazo)
		(*puntero).clave = nuevaClave
		(*puntero).dato = nuevoDato
	}
	a.cant--
	return dato
}

func (a *abb[K, V]) buscarReemplazo(nodo **nodoAbb[K, V]) **nodoAbb[K, V] {
	if (*nodo).der == nil {
		return nodo
	} else {
		return a.buscarReemplazo(&(*nodo).der)
	}
}

func (a *abb[K, V]) cantidadDeHijos(nodo **nodoAbb[K, V]) int {
	if (*nodo).izq != nil && (*nodo).der != nil {
		return 2
	} else if (*nodo).izq == nil && (*nodo).der == nil {
		return 0
	} else {
		return 1
	}
}

func (a *abb[K, V]) obtenerHijo(nodo **nodoAbb[K, V]) **nodoAbb[K, V] {
	if (*nodo).izq != nil {
		return &(*nodo).izq
	} else {
		return &(*nodo).der
	}
}

func (i *iterAbb[K, V]) apilarHijosPrimero(nodo *nodoAbb[K, V]) {
	if nodo == nil {
		return
	}
	if nodo.izq != nil {
		if i.desde == nil || i.abb.cmp(*i.desde, nodo.izq.clave) <= 0 {
			i.pila.Apilar(nodo.izq)
			i.apilarHijosPrimero(nodo.izq)
			return
		}
		if i.desde != nil && nodo.izq.der != nil && i.abb.cmp(*i.desde, nodo.izq.der.clave) <= 0 {
			i.pila.Apilar(nodo.izq.der)
			i.apilarHijosPrimero(nodo.izq.der)
			return
		}
	}
}

func (i *iterAbb[K, V]) apilarHijosIzq(nodo *nodoAbb[K, V]) {
	if nodo == nil {
		return
	}
	if (i.desde == nil || i.abb.cmp(*i.desde, nodo.clave) <= 0) && (i.hasta == nil || i.abb.cmp(*i.hasta, nodo.clave) >= 0) {
		i.pila.Apilar(nodo)
	}
	i.apilarHijosIzq(nodo.izq)
}

func (i *iterAbb[K, V]) buscarPrimero(nodo *nodoAbb[K, V], desde *K) *nodoAbb[K, V] {
	if nodo == nil {
		return nil
	}

	if i.desde == nil && i.hasta == nil {
		return nodo
	}

	if i.desde != nil && i.abb.cmp(*i.desde, nodo.clave) > 0 {
		return i.buscarPrimero(nodo.der, desde)
	}
	if i.hasta != nil && i.abb.cmp(*i.hasta, nodo.clave) < 0 {
		return i.buscarPrimero(nodo.izq, desde)
	}
	return nodo
}

func (a *abb[K, V]) iterarPorRango(actual *nodoAbb[K, V], f func(K, V) bool, desde *K, hasta *K) bool {
	if actual == nil {
		return true
	}
	anterior := true
	if desde == nil || a.cmp(actual.clave, *desde) > 0 {
		anterior = a.iterarPorRango(actual.izq, f, desde, hasta)
	}
	if anterior && (desde == nil || a.cmp(actual.clave, *desde) >= 0) && (hasta == nil || a.cmp(actual.clave, *hasta) <= 0) {
		anterior = f(actual.clave, actual.dato)
	}
	if anterior && (hasta == nil || a.cmp(actual.clave, *hasta) < 0) {
		return a.iterarPorRango(actual.der, f, desde, hasta)
	}
	return false
}
