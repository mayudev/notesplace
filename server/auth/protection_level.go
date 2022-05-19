package auth

type ProtectionLevel uint8

// WriteProtected checks if notebook is protected against writes,
// i.e. is read only by unauthenticated users.
func (p ProtectionLevel) WriteProtected() bool {
	return p > 0
}

// Protected checks if notebook is protected,
// meaning authentication is required to see and motify its contents.
func (p ProtectionLevel) Protected() bool {
	return p > 1
}
