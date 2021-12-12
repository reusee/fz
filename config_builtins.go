package fz

import (
	"time"

	"github.com/google/uuid"
)

type CreatedTime string

func (_ ConfigScope) CreateTime() CreatedTime {
	return CreatedTime(time.Now().Format(time.RFC3339))
}

func (_ ConfigScope) CreatedTimeConfig(
	t CreatedTime,
) ConfigItems {
	return ConfigItems{t}
}

func (_ ConfigScope) UUID() uuid.UUID {
	return uuid.New()
}

func (_ ConfigScope) UUIDConfig(
	id uuid.UUID,
) ConfigItems {
	return ConfigItems{id}
}
