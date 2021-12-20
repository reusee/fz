package fz

import (
	"time"

	"github.com/google/uuid"
)

type CreatedAt string

func (_ Def) CreatedAt() CreatedAt {
	return CreatedAt(time.Now().Format(time.RFC3339))
}

func (_ Def) CreatedAtConfigItem(
	t CreatedAt,
) ConfigItems {
	return ConfigItems{t}
}

func (_ Def) UUID() uuid.UUID {
	return uuid.New()
}

func (_ Def) UUIDConfigItem(
	id uuid.UUID,
) ConfigItems {
	return ConfigItems{id}
}
