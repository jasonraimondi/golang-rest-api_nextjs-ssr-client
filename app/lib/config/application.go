package config

type JWTSecureKey string

type S3Cred struct {
	Host         string
	Region       string
	Identifier   string
	Secret       string
	SessionToken string // @todo what is session token?
}

type DBCred struct {
	Driver     string
	Connection string
}
