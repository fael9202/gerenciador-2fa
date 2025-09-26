package config

import (
	"time"
)

const (
	TokenExpiration = 24 * time.Hour
	TOTPIssuer     = "GerenciadorApp2FA"
	MongoTimeout   = 10 * time.Second
) 