package diccionario_test

import (
	TDA_ABB "diccionario"
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

var TAMS_VOLUMEN = []int{1000, 2000, 4000}

func TestAbb(t *testing.T) {
	dic := TDA_ABB.CrearABB[int, int](func(a, b int) int { return a - b })

	desde := 2
	hasta := 5
	dic.Guardar(1, 10)
	dic.Guardar(2, 10)
	dic.Guardar(3, 10)
	dic.Guardar(4, 10)
	dic.Guardar(5, 10)
	dic.Guardar(6, 10)
	dic.Guardar(7, 10)
	for iter := dic.IteradorRango(&desde, &hasta); iter.HaySiguiente(); {
		fmt.Println(iter.VerActual())
		iter.Siguiente()
	}

}

func TestDiccionarioVacio(t *testing.T) {
	t.Log("Comprueba que Diccionario vacio no tiene claves")
	dic := TDA_ABB.CrearABB[int, int](func(a, b int) int { return a - b })
	require.NotNil(t, dic)
	require.Equal(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece(1))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(1) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(1) })
}

func TestUnElemento(t *testing.T) {
	t.Log("Comprueba que Diccionario con un elemento tiene esa Clave, unicamente")
	dic := TDA_ABB.CrearABB[string, int](func(a, b string) int { return strings.Compare(a, b) })
	dic.Guardar("A", 10)
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece("A"))
	require.False(t, dic.Pertenece("B"))
	require.EqualValues(t, 10, dic.Obtener("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("B") })
}

func TestDiccionarioGuardar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se comprueba que en todo momento funciona acorde")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}

	dic := TDA_ABB.CrearABB[string, string](func(a, b string) int { return strings.Compare(a, b) })
	require.False(t, dic.Pertenece(claves[0]))
	require.False(t, dic.Pertenece(claves[0]))
	dic.Guardar(claves[0], valores[0])
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))

	require.False(t, dic.Pertenece(claves[1]))
	require.False(t, dic.Pertenece(claves[2]))
	dic.Guardar(claves[1], valores[1])
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[1]))
	require.EqualValues(t, 2, dic.Cantidad())
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[1], dic.Obtener(claves[1]))

	require.False(t, dic.Pertenece(claves[2]))
	dic.Guardar(claves[2], valores[2])
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[1]))
	require.True(t, dic.Pertenece(claves[2]))
	require.EqualValues(t, 3, dic.Cantidad())
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[1], dic.Obtener(claves[1]))
	require.EqualValues(t, valores[2], dic.Obtener(claves[2]))
}

func TestReemplazoDato(t *testing.T) {
	t.Log("Guarda un par de claves, y luego vuelve a guardar, buscando que el dato se haya reemplazado")
	clave := "Gato"
	clave2 := "Perro"
	dic := TDA_ABB.CrearABB[string, string](func(a, b string) int { return strings.Compare(a, b) })
	dic.Guardar(clave, "miau")
	dic.Guardar(clave2, "guau")
	require.True(t, dic.Pertenece(clave))
	require.True(t, dic.Pertenece(clave2))
	require.EqualValues(t, "miau", dic.Obtener(clave))
	require.EqualValues(t, "guau", dic.Obtener(clave2))
	require.EqualValues(t, 2, dic.Cantidad())

	dic.Guardar(clave, "miu")
	dic.Guardar(clave2, "baubau")
	require.True(t, dic.Pertenece(clave))
	require.True(t, dic.Pertenece(clave2))
	require.EqualValues(t, 2, dic.Cantidad())
	require.EqualValues(t, "miu", dic.Obtener(clave))
	require.EqualValues(t, "baubau", dic.Obtener(clave2))
}

func TestDiccionarioBorrar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se los borra, revisando que en todo momento " +
		"el diccionario se comporte de manera adecuada")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	dic := TDA_ABB.CrearABB[string, string](func(a, b string) int { return strings.Compare(a, b) })

	require.False(t, dic.Pertenece(claves[0]))
	require.False(t, dic.Pertenece(claves[0]))
	dic.Guardar(claves[0], valores[0])
	dic.Guardar(claves[1], valores[1])
	dic.Guardar(claves[2], valores[2])

	require.True(t, dic.Pertenece(claves[2]))
	require.EqualValues(t, valores[2], dic.Borrar(claves[2]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[2]) })
	require.EqualValues(t, 2, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[2]))

	require.True(t, dic.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], dic.Borrar(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[0]) })
	require.EqualValues(t, 1, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(claves[0]) })

	require.True(t, dic.Pertenece(claves[1]))
	require.EqualValues(t, valores[1], dic.Borrar(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[1]) })
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(claves[1]) })
}

func TestReutlizacionDeBorrados(t *testing.T) {
	t.Log("Prueba de caja blanca: revisa que no haya problema reinsertando un elemento borrado")
	dic := TDA_ABB.CrearABB[string, string](func(a, b string) int { return strings.Compare(a, b) })
	clave := "hola"
	dic.Guardar(clave, "mundo!")
	dic.Borrar(clave)
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece(clave))
	dic.Guardar(clave, "mundooo!")
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, "mundooo!", dic.Obtener(clave))
}

func TestConClavesNumericas(t *testing.T) {
	t.Log("Valida que no solo funcione con strings")
	dic := TDA_ABB.CrearABB[int, string](func(a, b int) int { return a - b })
	clave := 10
	valor := "Gatito"

	dic.Guardar(clave, valor)
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, valor, dic.Obtener(clave))
	require.EqualValues(t, valor, dic.Borrar(clave))
	require.False(t, dic.Pertenece(clave))
}

func TestConClavesStructs(t *testing.T) {
	t.Log("Valida que tambien funcione con estructuras mas complejas")
	type basico struct {
		cadena string
		num    int
	}

	dic := TDA_ABB.CrearABB[basico, int](func(a, b basico) int { return strings.Compare(a.cadena, b.cadena) })

	b2 := basico{cadena: "odnum", num: 14}
	b1 := basico{cadena: "mundo", num: 8}
	b3 := basico{cadena: "world", num: 8}

	dic.Guardar(b1, 0)
	dic.Guardar(b2, 1)
	dic.Guardar(b3, 2)

	require.True(t, dic.Pertenece(b1))
	require.True(t, dic.Pertenece(b2))
	require.True(t, dic.Pertenece(b3))
	require.EqualValues(t, 0, dic.Obtener(b1))
	require.EqualValues(t, 1, dic.Obtener(b2))
	require.EqualValues(t, 2, dic.Obtener(b3))
	dic.Guardar(b1, 5)
	require.EqualValues(t, 5, dic.Obtener(b1))
	require.EqualValues(t, 2, dic.Obtener(b3))
	require.EqualValues(t, 5, dic.Borrar(b1))
	require.False(t, dic.Pertenece(b1))
	require.EqualValues(t, 2, dic.Obtener(b3))
}

func TestClaveVacia(t *testing.T) {
	t.Log("Guardamos una clave vacía (i.e. \"\") y deberia funcionar sin problemas")
	dic := TDA_ABB.CrearABB[string, string](func(a, b string) int { return strings.Compare(a, b) })
	clave := ""
	dic.Guardar(clave, clave)
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, clave, dic.Obtener(clave))
}

func TestValorNulo(t *testing.T) {
	t.Log("Probamos que el valor puede ser nil sin problemas")
	dic := TDA_ABB.CrearABB[string, *int](func(a, b string) int { return strings.Compare(a, b) })
	clave := "Pez"
	dic.Guardar(clave, nil)
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, (*int)(nil), dic.Obtener(clave))
	require.EqualValues(t, (*int)(nil), dic.Borrar(clave))
	require.False(t, dic.Pertenece(clave))
}

func buscar(clave string, claves []string) int {
	for i, c := range claves {
		if c == clave {
			return i
		}
	}
	return -1
}

func TestIteradorInternoClaves(t *testing.T) {
	t.Log("Valida que todas las claves sean recorridas (y una única vez) con el iterador interno")
	clave1 := "Perro"
	clave2 := "Gato"
	clave3 := "Vaca"
	claves := []string{clave1, clave2, clave3}
	dic := TDA_ABB.CrearABB[string, *int](func(a, b string) int { return strings.Compare(a, b) })
	dic.Guardar(claves[0], nil)
	dic.Guardar(claves[1], nil)
	dic.Guardar(claves[2], nil)

	cs := []string{"", "", ""}
	cantidad := 0
	cantPtr := &cantidad

	dic.Iterar(func(clave string, dato *int) bool {
		cs[cantidad] = clave
		*cantPtr = *cantPtr + 1
		return true
	})

	require.EqualValues(t, 3, cantidad)
	require.NotEqualValues(t, -1, buscar(cs[0], claves))
	require.NotEqualValues(t, -1, buscar(cs[1], claves))
	require.NotEqualValues(t, -1, buscar(cs[2], claves))
	require.NotEqualValues(t, cs[0], cs[1])
	require.NotEqualValues(t, cs[0], cs[2])
	require.NotEqualValues(t, cs[2], cs[1])
}

func TestIteradorInternoValores(t *testing.T) {
	t.Log("Valida que los datos sean recorridas correctamente (y una única vez) con el iterador interno")
	clave1 := "Hamster"
	clave2 := "Vaca"
	clave3 := "Perro"
	clave4 := "Burrito"
	clave5 := "Gato"

	dic := TDA_ABB.CrearABB[string, int](func(a, b string) int { return strings.Compare(a, b) })
	dic.Guardar(clave1, 6)
	dic.Guardar(clave2, 2)
	dic.Guardar(clave3, 3)
	dic.Guardar(clave4, 4)
	dic.Guardar(clave5, 5)

	factorial := 1
	ptrFactorial := &factorial
	dic.Iterar(func(_ string, dato int) bool {
		*ptrFactorial *= dato
		return true
	})

	require.EqualValues(t, 720, factorial)
}

func ejecutarPruebaVolumen(b *testing.B, n int) {
	dic := TDA_ABB.CrearABB[string, int](func(a, b string) int { return strings.Compare(a, b) })

	claves := make([]string, n)
	valores := make([]int, n)

	for i := 0; i < n; i++ {
		claves[i] = fmt.Sprintf("%08d", i)
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(claves), func(i, j int) { claves[i], claves[j] = claves[j], claves[i] })

	for i := 0; i < n; i++ {
		valores[i] = i
		dic.Guardar(claves[i], valores[i])
	}

	require.EqualValues(b, n, dic.Cantidad(), "La cantidad de elementos es incorrecta")

	// Verifica que devuelva los valores correctos
	ok := true
	for i := 0; i < n; i++ {
		ok = dic.Pertenece(claves[i])
		if !ok {
			break
		}
		ok = dic.Obtener(claves[i]) == valores[i]
		if !ok {
			break
		}
	}

	require.True(b, ok, "Pertenece y Obtener con muchos elementos no funciona correctamente")
	require.EqualValues(b, n, dic.Cantidad(), "La cantidad de elementos es incorrecta")

	// Verifica que borre y devuelva los valores correctos
	for i := 0; i < n; i++ {
		ok = dic.Borrar(claves[i]) == valores[i]
		if !ok {
			break
		}
	}

	require.True(b, ok, "Borrar muchos elementos no funciona correctamente")
	require.EqualValues(b, 0, dic.Cantidad())
}

func BenchmarkDiccionario(b *testing.B) {
	b.Log("Prueba de stress del Diccionario. Prueba guardando distinta cantidad de elementos (muy grandes), " +
		"ejecutando muchas veces las pruebas para generar un benchmark. Valida que la cantidad " +
		"sea la adecuada. Luego validamos que podemos obtener y ver si pertenece cada una de las claves geeneradas, " +
		"y que luego podemos borrar sin problemas")
	for _, n := range TAMS_VOLUMEN {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebaVolumen(b, n)
			}
		})
	}
}

func TestIterarDiccionarioVacio(t *testing.T) {
	t.Log("Iterar sobre diccionario vacio es simplemente tenerlo al final")
	dic := TDA_ABB.CrearABB[string, int](func(a, b string) int { return strings.Compare(a, b) })
	iter := dic.Iterador()
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestDiccionarioIterar(t *testing.T) {
	t.Log("Guardamos 3 valores en un Diccionario, e iteramos validando que las claves sean todas diferentes " +
		"pero pertenecientes al diccionario. Además los valores de VerActual y Siguiente van siendo correctos entre sí")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	dic := TDA_ABB.CrearABB[string, string](func(a, b string) int { return strings.Compare(a, b) })
	dic.Guardar(claves[0], valores[0])
	dic.Guardar(claves[1], valores[1])
	dic.Guardar(claves[2], valores[2])
	iter := dic.Iterador()

	require.True(t, iter.HaySiguiente())
	primero, _ := iter.VerActual()
	require.NotEqualValues(t, -1, buscar(primero, claves))

	require.EqualValues(t, primero, iter.Siguiente())
	segundo, segundo_valor := iter.VerActual()
	require.NotEqualValues(t, -1, buscar(segundo, claves))
	require.EqualValues(t, valores[buscar(segundo, claves)], segundo_valor)
	require.NotEqualValues(t, primero, segundo)
	require.True(t, iter.HaySiguiente())

	require.EqualValues(t, segundo, iter.Siguiente())
	require.True(t, iter.HaySiguiente())
	tercero, _ := iter.VerActual()
	require.NotEqualValues(t, -1, buscar(tercero, claves))
	require.NotEqualValues(t, primero, tercero)
	require.NotEqualValues(t, segundo, tercero)
	require.EqualValues(t, tercero, iter.Siguiente())

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestIteradorNoLlegaAlFinal(t *testing.T) {
	t.Log("Crea un iterador y no lo avanza. Luego crea otro iterador y lo avanza.")
	dic := TDA_ABB.CrearABB[string, string](func(a, b string) int { return strings.Compare(a, b) })
	claves := []string{"A", "B", "C"}
	dic.Guardar(claves[0], "")
	dic.Guardar(claves[1], "")
	dic.Guardar(claves[2], "")

	dic.Iterador()
	iter2 := dic.Iterador()
	iter2.Siguiente()
	iter3 := dic.Iterador()
	primero := iter3.Siguiente()
	segundo := iter3.Siguiente()
	tercero := iter3.Siguiente()
	require.False(t, iter3.HaySiguiente())
	require.NotEqualValues(t, primero, segundo)
	require.NotEqualValues(t, tercero, segundo)
	require.NotEqualValues(t, primero, tercero)
	require.NotEqualValues(t, -1, buscar(primero, claves))
	require.NotEqualValues(t, -1, buscar(segundo, claves))
	require.NotEqualValues(t, -1, buscar(tercero, claves))
}

func TestPruebaIterarTrasBorrados(t *testing.T) {
	t.Log("Prueba de caja blanca: Esta prueba intenta verificar el comportamiento del hash abierto cuando " +
		"queda con listas vacías en su tabla. El iterador debería ignorar las listas vacías, avanzando hasta " +
		"encontrar un elemento real.")

	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"

	dic := TDA_ABB.CrearABB[string, string](func(a, b string) int { return strings.Compare(a, b) })
	dic.Guardar(clave1, "")
	dic.Guardar(clave2, "")
	dic.Guardar(clave3, "")
	dic.Borrar(clave1)
	dic.Borrar(clave2)
	dic.Borrar(clave3)
	iter := dic.Iterador()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	dic.Guardar(clave1, "A")
	iter = dic.Iterador()

	require.True(t, iter.HaySiguiente())
	c1, v1 := iter.VerActual()
	require.EqualValues(t, clave1, c1)
	require.EqualValues(t, "A", v1)
	require.EqualValues(t, clave1, iter.Siguiente())
	require.False(t, iter.HaySiguiente())
}

func ejecutarPruebasVolumenIterador(b *testing.B, n int) {
	dic := TDA_ABB.CrearABB[string, *int](func(a, b string) int { return strings.Compare(a, b) })

	claves := make([]string, n)
	valores := make([]int, n)

	for i := 0; i < n; i++ {
		claves[i] = fmt.Sprintf("%08d", i)
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(claves), func(i, j int) { claves[i], claves[j] = claves[j], claves[i] })

	for i := 0; i < n; i++ {
		valores[i] = i
		dic.Guardar(claves[i], &valores[i])
	}

	// Prueba de iteración sobre las claves almacenadas.
	iter := dic.Iterador()
	require.True(b, iter.HaySiguiente())

	ok := true
	var i int
	var clave string
	var valor *int

	for i = 0; i < n; i++ {
		if !iter.HaySiguiente() {
			ok = false
			break
		}
		c1, v1 := iter.VerActual()
		clave = c1
		if clave == "" {
			ok = false
			break
		}
		valor = v1
		if valor == nil {
			ok = false
			break
		}
		*valor = n
		iter.Siguiente()
	}
	require.True(b, ok, "Iteracion en volumen no funciona correctamente")
	require.EqualValues(b, n, i, "No se recorrió todo el largo")
	require.False(b, iter.HaySiguiente(), "El iterador debe estar al final luego de recorrer")

	ok = true
	for i = 0; i < n; i++ {
		if valores[i] != n {
			ok = false
			break
		}
	}
	require.True(b, ok, "No se cambiaron todos los elementos")
}

func BenchmarkIterador(b *testing.B) {
	b.Log("Prueba de stress del Iterador del Diccionario. Prueba guardando distinta cantidad de elementos " +
		"(muy grandes) b.N elementos, iterarlos todos sin problemas. Se ejecuta cada prueba b.N veces para generar " +
		"un benchmark")
	for _, n := range TAMS_VOLUMEN {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebasVolumenIterador(b, n)
			}
		})
	}
}

func TestIterarRangoVacio(t *testing.T) {
	dic := TDA_ABB.CrearABB[string, *int](func(a, b string) int { return strings.Compare(a, b) })
	suma := 0
	visitar := func(clave string, dato *int) bool {
		suma += *dato
		return true
	}
	desde := "Hola"
	hasta := "Chau"
	dic.IterarRango(&desde, &hasta, visitar)
	require.Zero(t, suma)
}

func TestIterarElementosFueraRango(t *testing.T) {
	cmp := func(a, b int) int {
		return a - b
	}
	dic := TDA_ABB.CrearABB[int, *int](cmp)
	//Guardo 10 como raiz para que el abb tenga al menos 2 ramas
	raiz := 10
	dic.Guardar(raiz, &raiz)
	for i := 0; i > 20; i++ {
		dic.Guardar(i, &i)
	}
	suma := 0
	visitar := func(clave int, dato *int) bool {
		suma += *dato
		return true
	}
	desde := 30
	hasta := 50
	dic.IterarRango(&desde, &hasta, visitar)
	require.Zero(t, suma)
}

func TestIterarRangoVolumen(t *testing.T) {
	cmp := func(a, b int) int {
		return a - b
	}
	dic := TDA_ABB.CrearABB[int, *int](cmp)
	//Guardo 500 como raiz para que el abb tenga al menos 2 ramas
	raiz := 500
	dic.Guardar(raiz, &raiz)

	//Se guardan elementos aleatorios
	for i := 0; i < 2500; i++ {
		random := rand.Int()
		dic.Guardar(rand.Intn(1000), &random)
	}

	desde := 500
	hasta := 750
	visitar := func(clave int, dato *int) bool {
		require.True(t, clave >= desde && clave <= hasta)
		return true
	}
	dic.IterarRango(&desde, &hasta, visitar)

}

func TestIterarCortePorFucion(t *testing.T) {
	cmp := func(a, b int) int {
		return a - b
	}
	dic := TDA_ABB.CrearABB[int, *int](cmp)
	//Guardo 500 como raiz para que el abb tenga al menos 2 ramas
	raiz := 500
	dic.Guardar(raiz, &raiz)

	//Se guardan numeros positivos aleatorios
	for i := 0; i < 2500; i++ {
		random := rand.Int()
		dic.Guardar(rand.Intn(1000), &random)
	}

	//Con estas condiciones suma debe si o si superar 1000 en el caso de que itere todo
	desde := 500
	hasta := 1750
	suma := 0

	visitar := func(clave int, dato *int) bool {
		suma += *dato
		return suma < 1000
	}
	dic.IterarRango(&desde, &hasta, visitar)
	require.Less(t, 1000, suma)
}

func TestIteradorRangoVacio(t *testing.T) {
	dic := TDA_ABB.CrearABB[string, *int](func(a, b string) int { return strings.Compare(a, b) })
	desde := "Hola"
	hasta := "Chau"
	iterador := dic.IteradorRango(&desde, &hasta)
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.VerActual() })
	require.False(t, iterador.HaySiguiente())
}

func TestIterardorElementosFueraRango(t *testing.T) {
	cmp := func(a, b int) int {
		return a - b
	}
	dic := TDA_ABB.CrearABB[int, *int](cmp)
	//Guardo 10 como raiz para que el abb tenga al menos 2 ramas
	raiz := 10
	dic.Guardar(raiz, &raiz)
	for i := 0; i > 20; i++ {
		dic.Guardar(i, &i)
	}
	desde := 30
	hasta := 50
	iterador := dic.IteradorRango(&desde, &hasta)
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.VerActual() })
	require.False(t, iterador.HaySiguiente())
}
func TestIteradorRangoVolumen(t *testing.T) {
	cmp := func(a, b int) int {
		return a - b
	}
	dic := TDA_ABB.CrearABB[int, *int](cmp)
	//Guardo 500 como raiz para que el abb tenga al menos 2 ramas
	raiz := 500
	dic.Guardar(raiz, &raiz)

	rand.Seed(time.Now().UnixNano())
	//Se guardan elementos aleatorios
	for i := 0; i < 2500; i++ {
		random := rand.Int()
		dic.Guardar(rand.Intn(1000), &random)
	}

	desde := 500
	hasta := 750

	for iter := dic.IteradorRango(&desde, &hasta); iter.HaySiguiente(); {
		clave, _ := iter.VerActual()
		require.True(t, clave >= desde && clave <= hasta)
		iter.Siguiente()
	}
}
