package forms

type errors map[string][]string

// Add adds an error message to the errors map for a given key.
func (e errors) Add(key, message string) {

	e[key] = append(e[key], message)
}

// Get retrieves the first error message for a given key from the errors map.
func (e errors) Get(key string) string {
	if len(e[key]) == 0 {
		return ""
	}
	return e[key][0]
}
