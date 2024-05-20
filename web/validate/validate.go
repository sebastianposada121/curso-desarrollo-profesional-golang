package validate

import "regexp"

func Password(password string) bool {
	// Expresión regular para validar la contraseña:
	// Al menos 8 caracteres
	// Al menos una letra mayúscula
	// Al menos una letra minúscula
	// Al menos un dígito
	// Verificar la longitud
	if len(password) < 8 {
		return false
	}

	// Verificar la presencia de al menos una letra minúscula
	minuscula := regexp.MustCompile(`[a-z]`)
	if !minuscula.MatchString(password) {
		return false
	}

	// Verificar la presencia de al menos una letra mayúscula
	mayuscula := regexp.MustCompile(`[A-Z]`)
	if !mayuscula.MatchString(password) {
		return false
	}

	// Verificar la presencia de al menos un dígito
	digito := regexp.MustCompile(`\d`)
	return digito.MatchString(password)
}

func Email(email string) bool {
	if len(email) == 0 {
		return false
	}
	regex, error := regexp.Compile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if error != nil {
		return false
	}
	return regex.MatchString(email)
}
