package formaterror

import (
	"errors"
	"strings"
)

func FormatError(err string) error {

	if strings.Contains(err, "username") {
		return errors.New("username Already Taken")
	}

	if strings.Contains(err, "nama_lengkap") {
		return errors.New("nama lengkap Already Taken")
	}

	return errors.New("Incorrect Details")
}
