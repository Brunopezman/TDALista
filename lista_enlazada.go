package lista

type nodoLista[T any] struct {
	dato      T
	siguiente *nodoLista[T]
}

type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}

type iterListaEnlazada[T any] struct {
	anterior *nodoLista[T]
	actual   *nodoLista[T]
	lista    *listaEnlazada[T]
}

func crearNodo[T any](dato T) *nodoLista[T] {
	return &nodoLista[T]{dato: dato, siguiente: nil}
}

func CrearListaEnlazada[T any]() Lista[T] {
	return &listaEnlazada[T]{primero: nil, ultimo: nil, largo: 0}
}

func (l *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return &iterListaEnlazada[T]{anterior: nil, actual: nil, lista: l}
}

func (l *listaEnlazada[T]) EstaVacia() bool {
	return l.largo == 0
}

func (l *listaEnlazada[T]) InsertarPrimero(dato T) {
	nodo := crearNodo[T](dato)
	if l.EstaVacia() {
		l.primero = nodo
	} else {
		l.ultimo.siguiente = nodo
	}
	l.ultimo = nodo
}

func (l *listaEnlazada[T]) InsertarUltimo(dato T) {
	nodo := crearNodo[T](dato)
	if l.EstaVacia() {
		l.primero = nodo
	} else {
		l.ultimo.siguiente = nodo
	}
	l.ultimo = nodo
	l.largo++

}

func (l *listaEnlazada[T]) BorrarPrimero() T {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}
	elem := l.VerPrimero()
	l.primero = l.primero.siguiente

	if l.primero == nil {
		l.ultimo = nil
	}
	l.largo--

	return elem
}

func (l *listaEnlazada[T]) VerPrimero() T {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}
	return l.primero.dato
}

func (l *listaEnlazada[T]) VerUltimo() T {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}
	return l.ultimo.dato
}

func (l *listaEnlazada[T]) Largo() int {
	return l.largo
}

func (l *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	actual := l.primero
	for actual != nil {
		visitar(actual.dato)
		actual = actual.siguiente
	}
	
}

func (iter *iterListaEnlazada[T]) VerActual() T {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return iter.actual.dato
}

func (iter *iterListaEnlazada[T]) HaySiguiente() bool {
	return iter.actual != nil
}

func (iter *iterListaEnlazada[T]) Siguiente() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	iter.anterior = iter.actual
	iter.actual = iter.actual.siguiente
}

func (iter *iterListaEnlazada[T]) Insertar(dato T) {
	if iter == nil || iter.lista == nil {
		return
	}

	//Iterador al principio de la lista
	if iter.anterior == nil {
		nodo := crearNodo(dato)
		nodo.siguiente = iter.actual
		iter.lista.primero = nodo //actualizo la referencia del primer elemento en la lista

		if iter.lista.ultimo == nil {
			iter.lista.ultimo = nodo
		}

		iter.lista.largo++
		return
	}

	nodo := crearNodo(dato)
	//Iterador al final de la lista
	if !iter.HaySiguiente() {
		iter.lista.ultimo = nodo
	}
	//Iterador en el medio de la lista
	iter.anterior.siguiente = nodo
	nodo.siguiente = iter.actual
	iter.lista.largo++

}

func (iter *iterListaEnlazada[T]) Borrar() T {
	if iter == nil || iter.lista == nil || iter.lista.primero == nil {
		panic("El iterador termino de iterar")
	}

	//Iterador al principio de la lista
	if iter.anterior == nil {
		dato := iter.lista.primero.dato
		iter.lista.primero = iter.lista.primero.siguiente //actualizo la referencia del primer elemento en la lista

		if iter.lista.primero == nil {
			iter.lista.ultimo = nil
		}

		iter.lista.largo--
		return dato
	}

	//Iterador al medio de la lista
	dato := iter.actual.dato
	iter.anterior.siguiente = iter.actual.siguiente

	//Iterador al final de la lista
	if !iter.HaySiguiente() {
		iter.lista.ultimo = iter.anterior
	}
	iter.lista.largo--
	return dato
}
