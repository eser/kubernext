package shared

import "github.com/oklog/ulid/v2"

func GenerateUniqueId() string {
	return ulid.MustNew(ulid.Now(), nil).String()
}
