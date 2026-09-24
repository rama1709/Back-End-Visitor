package firebase

import (
	"context"
	"fmt"
)

func (f *Firebase) CreateCustomToken(
	uid string,
) (string, error) {
	if f == nil {
		return "", fmt.Errorf(
			"firebase client nil",
		)
	}

	if f.Auth == nil {
		return "", fmt.Errorf(
			"firebase auth client nil",
		)
	}

	if uid == "" {
		return "", fmt.Errorf(
			"firebase uid kosong",
		)
	}

	ctx := context.Background()

	token, err :=
		f.Auth.CustomToken(ctx, uid)

	if err != nil {
		return "", fmt.Errorf(
			"gagal membuat Firebase custom token: %w",
			err,
		)
	}

	return token, nil
}