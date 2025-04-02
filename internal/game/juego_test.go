package ahorcado

import "testing"

func TestNuevaPartidaYGanadoPerdido(t *testing.T) {
	juego := NuevaPartida("gol", 3) // palabra de 3 letras, 3 intentos permitidos
	// Al iniciar, no debe estar ganado ni perdido
	if juego.Ganado() {
		t.Errorf("El juego no debería empezar en estado 'ganado'")
	}
	if juego.Perdido() {
		t.Errorf("El juego no debería empezar en estado 'perdido'")
	}
	// El progreso inicial debe ser tres guiones bajos
	esperado := "___"
	if string(juego.Progreso) != esperado {
		t.Errorf("Progreso inicial incorrecto, got=%s, want=%s", string(juego.Progreso), esperado)
	}
}

func TestIntentarLetra(t *testing.T) {
	juego := NuevaPartida("gol", 3)
	// Intentar una letra correcta
	errMsg := juego.IntentarLetra("o")
	if errMsg != "" {
		t.Errorf("No debería haber error al ingresar 'o' por primera vez, pero se obtuvo: %s", errMsg)
	}
	// 'o' está en "gol" en la posición 2 (índice 1), debe revelarse
	if string(juego.Progreso) != "_O_" {
		t.Errorf("La letra 'o' no se reveló correctamente en el progreso. got=%s, want=_O_", string(juego.Progreso))
	}
	// Intentos no debe haber disminuido (seguirá en 3)
	if juego.IntentosRestantes != 3 {
		t.Errorf("IntentosRestantes debería seguir en 3 tras acierto, got=%d", juego.IntentosRestantes)
	}

	// Intentar una letra incorrecta
	errMsg = juego.IntentarLetra("z")
	if errMsg != "" {
		t.Errorf("No debería haber error al ingresar 'z' (letra no usada), pero se obtuvo: %s", errMsg)
	}
	// 'z' no está en "gol", los intentos deben disminuir a 2
	if juego.IntentosRestantes != 2 {
		t.Errorf("IntentosRestantes no se decrementó correctamente tras error. got=%d, want=%d", juego.IntentosRestantes, 2)
	}

	// Intentar una letra repetida
	errMsg = juego.IntentarLetra("o")
	if errMsg == "" {
		t.Errorf("Debería haber un mensaje de error al repetir la letra 'o', pero errMsg fue vacío")
	}
	// Verificar que no cambió IntentosRestantes ni progreso
	if juego.IntentosRestantes != 2 {
		t.Errorf("IntentosRestantes cambió indebidamente al repetir letra. got=%d, want=2", juego.IntentosRestantes)
	}
	if string(juego.Progreso) != "_O_" {
		t.Errorf("Progreso cambió indebidamente al repetir letra. got=%s, want=_O_", string(juego.Progreso))
	}
}
