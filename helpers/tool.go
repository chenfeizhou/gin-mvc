package helpers

func HandlerErr(err error) {
	if err != nil {
		panic(err)
	}
}
