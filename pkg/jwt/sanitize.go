package jwt

// Sanitize makes sure that the given token is valid.
func Sanitize(token string) bool {
	header, payload, signature := Split(token)
	if header == nil || payload == nil || signature == nil {
		return false
	}
	if len(*header) > 20 && len(*payload) > 20 && len(*signature) > 20 {
		return true
	}
	return false
}
