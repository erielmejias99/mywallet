package util

import (
	"crypto/sha256"
	"encoding/hex"
)

func Login( password string ) bool{
	hashed_password := sha256.Sum256( []byte( password ) )
	my_password := "33B19169C8C9A19B41DFDCD9E9F3630C4C033FF86B13507DCFEB6FD02B767566"

	return hex.EncodeToString( hashed_password[:] ) == my_password
}