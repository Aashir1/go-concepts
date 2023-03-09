package main

func main() {

	smell, _ := userJustFarted() //its not a problematic, because it will keep there for very short amount of time
	if err := userJustShitInTheirPant(); err != nil {
		changeDiapers()
	} //its a problematic and we need to handle this

}

// func Foo() (int, error) {
// 	return 666, nil
// }
// func Bar() (int, error) {
// 	return 666, nil
// }
// func Baz() (int, error) {
// 	return 666, nil
// }

func Foo() error {
	return nil
}

func bar() error {
	if err := Foo(); err != nil {
		return err
	}

	return nil
}
func changeDiapers() {

}

func userJustFarted() (int, error) {
	return 10, nil
}

func userJustShitInTheirPant() error {
	return nil
}
