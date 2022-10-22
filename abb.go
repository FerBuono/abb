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
	nodo := &nodoAbb[K, V]{clave: clave, dato: dato}
	if a.raiz == nil {
		a.raiz = nodo
		a.cant++
	} else {
		a.agregarNodo(a.raiz, nodo)
	}
	a.verArbol()

}

func (a *abb[K, V]) verArbol() {
	for iter := a.Iterador(); iter.HaySiguiente(); {
		clave, dato := iter.VerActual()
		fmt.Print("- {", clave, ":", dato, "} -")
		iter.Siguiente()
	}
	fmt.Println("\n-----------")
}

func (a *abb[K, V]) agregarNodo(nodoPadre, nodo *nodoAbb[K, V]) {
	if a.cmp(nodo.clave, nodoPadre.clave) < 0 {
		if nodoPadre.izq == nil {
			nodoPadre.izq = nodo
			a.cant++
		} else {
			a.agregarNodo(nodoPadre.izq, nodo)
		}
	}

	if a.cmp(nodo.clave, nodoPadre.clave) > 0 {
		if nodoPadre.der == nil {
			nodoPadre.der = nodo
			a.cant++
		} else {
			a.agregarNodo(nodoPadre.der, nodo)
		}
	}

	if a.cmp(nodo.clave, nodoPadre.clave) == 0 {
		nodoPadre.dato = nodo.dato
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
	if puntero != nil {
		dato = puntero.dato
		if a.cant == 1 {
			a.raiz = nil
			return dato
		}
		if a.cantidadDeHijos(puntero) == 0 {
			fmt.Println(puntero, "Entro aca: 0 hijos")
			puntero = nil
		} else if a.cantidadDeHijos(puntero) == 1 {
			fmt.Println(puntero, "Entro aca: 1 hijo")
			nuevaClave := a.obtenerHijo(puntero).clave
			nuevoDato := a.Borrar(nuevaClave)
			puntero.clave = nuevaClave
			puntero.dato = nuevoDato
		} else {
			fmt.Println(puntero, "Entro aca: 2 hijos")
			nuevaClave := a.buscarReemplazo(puntero.izq).clave
			nuevoDato := a.Borrar(nuevaClave)
			puntero.clave = nuevaClave
			puntero.dato = nuevoDato
		}
		a.cant--
	} else {
		panic("La clave no pertenece al diccionario")
	}
	a.verArbol()

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
	if nodo.der != nil {
		i.pila.Apilar(nodo.der)
		i.abb.apilarHijosIzq(nodo.der, i.pila)
	}

	return nodo.clave
}

// Primitivas del DiccionarioOrdenado

func (a *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	a.iterarPorRango(a.raiz, visitar, desde, hasta)
}

func (a *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	iter := new(iterAbb[K, V])
	iter.pila = TDAPila.CrearPilaDinamica[*nodoAbb[K, V]]()
	iter.apilarRango(iter.abb.raiz, desde, hasta)
	return iter
}

// Funciones y métodos auxiliares

func (a *abb[K, V]) buscarPuntero(clave K, nodo *nodoAbb[K, V]) *nodoAbb[K, V] {
	if nodo == nil {
		return nodo
	}
	if a.cmp(clave, nodo.clave) < 0 {
		if nodo.izq == nil || a.cmp(clave, nodo.izq.clave) == 0 {
			return nodo.izq
		}
		return a.buscarPuntero(clave, nodo.izq)
	} else if a.cmp(clave, nodo.clave) > 0 {
		if nodo.der == nil || a.cmp(clave, nodo.der.clave) == 0 {
			return nodo.der
		}
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
	} else if nodo.izq == nil && nodo.der == nil {
		return 0
	} else {
		return 1
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

func (a *abb[K, V]) iterarPorRango(actual *nodoAbb[K, V], f func(K, V) bool, desde *K, hasta *K) {
	if actual == nil {
		return
	}
	if a.cmp(actual.clave, *desde) > 0 {
		a.iterarPorRango(actual.izq, f, desde, hasta)
	}
	if a.cmp(actual.clave, *desde) >= 0 && a.cmp(actual.clave, *hasta) <= 0 {
		if !f(actual.clave, actual.dato) {
			return
		}
	}
	if a.cmp(actual.clave, *hasta) < 0 {
		a.iterarPorRango(actual.der, f, desde, hasta)
	}
}

func (i *iterAbb[K, V]) apilarRango(actual *nodoAbb[K, V], desde, hasta *K) {
	if actual == nil {
		return
	}
	if i.abb.cmp(actual.clave, *desde) > 0 {
		i.apilarRango(actual.izq, desde, hasta)
	}
	if i.abb.cmp(actual.clave, *desde) >= 0 && i.abb.cmp(actual.clave, *hasta) <= 0 {
		i.pila.Apilar(actual)
	}
	if i.abb.cmp(actual.clave, *hasta) < 0 {
		i.apilarRango(actual.der, desde, hasta)
	}
}
