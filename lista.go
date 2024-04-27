package lista

type Lista[T any] interface {

	// EstaVacia devuelve verdadero si la lista no tiene elementos, false en caso contrario.
	EstaVacia() bool

	// InsertarPrimero agrega un elemento en la primera posicion de la lista.
	InsertarPrimero(T)

	// InsertarUltimoagrega un elemento en la ultimo posicion de la lista.
	InsertarUltimo(T)

	// BorrarPrimero elimina el primer elemento de la lista. Si está vacía, entra en pánico con un mensaje
	// "La lista esta vacia".
	BorrarPrimero() T

	// VerPrimero obtiene el valor del primero de la lista. Si está vacía, entra en pánico con un mensaje
	// "La lista esta vacia".
	VerPrimero() T

	// VerUltimoobtiene el valor del ultimo de la lista. Si está vacía, entra en pánico con un mensaje
	// "La lista esta vacia".
	VerUltimo() T

	// Largo devuelve la longitud de la lista.
	Largo() int

	// Iterar recorre la lista de manera externa.
	Iterar(visitar func(T) bool)

	//Iterador devuelve el tipo IteradorLista[T] para iterar de forma interna con las primitivas del mismo.
	Iterador() IteradorLista[T]
}

type IteradorLista[T any] interface {

	// Ver actual devuelve el elemento en el que se encuentra el Iterador.
	VerActual() T

	// HaySiguiente devuelve si existe el elemento en el que sigue en el Iterador.
	HaySiguiente() bool

	// Siguiente avanza al siguiente elemento
	Siguiente()

	// Insertar
	Insertar(T)

	// Borrar
	Borrar() T
}
