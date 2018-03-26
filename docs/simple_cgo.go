package main

/* 												//Declare C-code
typedef struct{ int y; int x; }TPoint; // HL
TPoint NewTPointC(){ // HL
	TPoint out = {1,2};
	return out;
}
*/
import "C"
import (
	"fmt"
	"reflect"
)

func main() {
	A := C.TPoint{1, 2} // Equivelent to C99 `TPoint A = {1,2};`
	B := C.NewTPointC() // Calls C function `NewTPointC`
	if reflect.DeepEqual(A, B) {
		fmt.Println("A and B are Equal")
	}
}
