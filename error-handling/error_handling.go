package erratum

// Use creates a resource and handles any errors or otherwise calls `Frob` on the resource
func Use(o ResourceOpener, input string) (err error) {
	r, err := o()
	for err != nil {
		if _, ok := err.(TransientError); !ok {
			return
		}
		r, err = o()
	}

	defer r.Close()
	defer func() {
		if re := recover(); re != nil {
			if frob, ok := re.(FrobError); ok {
				r.Defrob(frob.defrobTag)
			}
			err = re.(error)
		}
	}()
	r.Frob(input)
	return
}

