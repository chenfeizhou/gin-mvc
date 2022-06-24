package helpers

func handlerErr(err error) {
	if err != nil {
		panic(err)
	}
}

func InitConfig() {
	LoadConfig()
}
