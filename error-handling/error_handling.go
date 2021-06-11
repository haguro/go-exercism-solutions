package erratum

// Use creates a resource and handles any errors or otherwise calls `Frob` on the resource
func Use(o ResourceOpener, input string) (err error) {
	r, err := o()
	if _, ok := err.(TransientError); ok {
		return Use(o, input)
	} else if err != nil {
		return err
	}

	defer func() error { return r.Close() }()
	defer func() {
		re := recover()
		if e, ok := re.(FrobError); ok {
			r.Defrob(e.defrobTag)
			err = e
		}
		if re != nil {
			err = re.(error)
		}

	}()
	r.Frob(input)
	return err
}
