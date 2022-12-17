package car

import "fmt"

/* CarPublic car con acceso publico */
type CarPublic struct {
	Brand string
	Year  int
}

/* CarPrivate car con acceso privado */
type carPrivate struct {
	brand string
	year  int
}

/* Function public */
func PrintMessage1(text string) {
	fmt.Println(text)
}

/* Function private */
func PrintMessage2(text string) {
	fmt.Println(text)
}
