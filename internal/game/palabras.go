package ahorcado

import (
	"math/rand"
	"time"
)

// Lista de posibles palabras secretas para el modo de un jugador.
var palabras = []string{
	"murcielago", "golang", "programacion", "ahorcado", "gol", "consola", "terminal",
	"monitor", "freelance", "soplete", "Malaga", "gas", "coche", "paquete", "circuito",
	"programador", "desarrollador", "software", "hardware", "teclado", "raton",
	"pantalla", "disco", "memoria", "red", "internet", "nube", "servidor", "cliente",
	"calcetin", "datos", "algoritmo", "estructura", "dato", "programa", "codigo",
	"funcion", "variable", "constante", "bucle", "condicional", "array", "lista",
	"elefante", "raton", "gato", "perro", "pajaro", "pez", "tortuga", "caballo",
	"vaca", "oveja", "cerdo", "gallo", "pato", "pavo", "gallina", "pajaro",
	"pez", "tortuga", "caballo", "vaca", "oveja", "cerdo", "gallo", "pato",
	"balon", "pelota", "raqueta", "pala", "paleta", "bici", "bicicleta", "patin",
	"patines", "skate", "skateboard", "surf", "surfista", "vela", "vela", "viento",
	"mar", "playa", "arena", "sol", "luna", "estrella", "nube", "cielo", "tierra",
	"agua", "fuego", "aire", "madera", "metal", "plastica", "papel", "carton",
	"medico", "enfermero", "dentista", "farmaceutico", "psicologo", "fisioterapeuta",
	"suma", "resta", "multiplicacion", "division", "matematicas", "geometria",
	"calculo", "bueno", "malo", "bonito", "feo", "grande", "pequeño", "alto", "bajo",
	"largo", "corto", "ancho", "estrecho", "delgado", "gordo", "fino", "grueso",
	"caliente", "frio", "templado", "humedo", "seco", "mojado", "sucio", "limpio",
	"nuevo", "viejo", "joven", "adulto", "anciano", "sabio", "inteligente",
	"tonto", "listo", "astuto", "lento", "rapido", "tranquilo", "nervioso",
	"feliz", "triste", "enojado", "contento", "satisfecho", "aburrido", "divertido",
	"interesante", "aburrido", "divertido", "entretenido", "monotonía", "rutina",
	"diversion", "entretenimiento", "ocio", "pasatiempo", "hobby", "aficion",
}

// init se ejecuta al cargar el paquete, y se usa para inicializar la semilla aleatoria.

func init() {
	rand.Seed(time.Now().UnixNano())
}

// EscogerPalabraAleatoria devuelve una palabra aleatoria de la lista de palabras.

func EscogerPalabraAleatoria() string {
	n := len(palabras)
	if n == 0 {
		return "" // si la lista estuviera vacía, devolvemos cadena vacía (caso extremo manejado)
	}
	indice := rand.Intn(n)  // índice aleatorio entre 0 y n-1
	return palabras[indice] // devolvemos la palabra en esa posición
}
