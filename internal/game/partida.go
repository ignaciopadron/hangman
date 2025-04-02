package ahorcado

import (
	"strings"

	"unicode"
)

// Partida representa el estado de una partida del ahorcado.
type Partida struct {
	PalabraSecreta    string // palabra a adivinar (en mayúsculas para facilitar comparación)
	Progreso          []rune // slice de runas con el progreso (letras adivinadas y guiones bajos '_' para no reveladas)
	LetrasProbadas    []rune // letras que el jugador ha intentado (correctas o incorrectas)
	IntentosRestantes int    // intentos que le quedan al jugador antes de perder
}

// NuevaPartida crea una nueva partida con una palabra secreta dada y un número fijo de intentos.
func NuevaPartida(palabra string, maxIntentos int) *Partida {
	palabra = strings.ToUpper(palabra) // convertimos a mayúsculas para no distinguir entre 'a' y 'A'
	longitud := len(palabra)
	progreso := make([]rune, longitud)
	for i := 0; i < longitud; i++ {
		// Inicia el progreso con '_' por cada letra (no revelada).
		if palabra[i] == ' ' {
			progreso[i] = ' ' // si la palabra contiene espacios (ej. frases), opcional: mostramos espacio
		} else {
			progreso[i] = '_' // letra no adivinada aún
		}
	}
	return &Partida{
		PalabraSecreta:    palabra,
		Progreso:          progreso,
		LetrasProbadas:    []rune{},
		IntentosRestantes: maxIntentos,
	}
}

// IntentarLetra procesa la letra ingresada por el jugador.
// Devuelve un mensaje de error si la entrada es inválida o ya fue probada, o una cadena vacía ("") si fue aceptada.
func (p *Partida) IntentarLetra(input string) string {
	if len(input) == 0 {
		return "No ingresaste ninguna letra"
	}
	// Tomar solo la primera letra si el jugador introdujo más de una
	// (alternativamente, podríamos rechazar entradas de más de una letra).
	runeInput := []rune(input)
	letra := runeInput[0]
	// Normalizar a mayúscula para comparar con PalabraSecreta
	letra = unicode.ToUpper(letra)

	// Validar que sea una letra del alfabeto
	if !unicode.IsLetter(letra) {
		return "Entrada no válida: introduce una letra"
	}
	// Verificar si ya fue intentada
	for _, l := range p.LetrasProbadas {
		if l == letra {
			return "Ya probaste la letra " + string(letra)
		}
	}

	// Registrar la letra en las probadas
	p.LetrasProbadas = append(p.LetrasProbadas, letra)

	// Comprobar si la letra está en la palabra secreta
	var acierto bool
	for i, l := range p.PalabraSecreta {
		if rune(l) == letra { // si coincide la letra (en mayúsculas)
			p.Progreso[i] = letra // revela la letra en el progreso
			acierto = true
		}
	}
	if !acierto {
		p.IntentosRestantes-- // letra incorrecta, resta un intento
	}
	return "" // sin error, intento procesado correctamente
}

// Ganado verifica si el jugador descubrió toda la palabra.
func (p *Partida) Ganado() bool {
	// Si en el progreso no quedan '_' significa que todas las letras fueron reveladas.
	for _, r := range p.Progreso {
		if r == '_' {
			return false
		}
	}
	return true
}

// Perdido verifica si al jugador no le quedan intentos.
func (p *Partida) Perdido() bool {
	return p.IntentosRestantes <= 0
}
