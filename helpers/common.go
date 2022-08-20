package helpers

// PanicErr panics if error not nil
func PanicErr(err error) {
	if err != nil {
		panic(err)
	}
}
