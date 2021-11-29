package rand

import (
	"github.com/mazen160/go-random"
)

func GenerateCode(val int) (string, error) {
	data, err := random.String(val)

	return data, err
}
