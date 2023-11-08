package utils

func ThrowPanic(err error) {
	if err != nil {
		panic(err)
	}
}
