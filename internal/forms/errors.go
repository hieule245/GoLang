package forms

type errors map[string][]string

func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

// Get return the first error
func (e errors) Get(field string) string {
	es := e[field]
	if len(es) == 0 {
		return "" 
	}
	return es[0]
}
