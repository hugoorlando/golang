package main

import "fmt"

type ByteSize float64

const (
	_           = iota
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func (b ByteSize) String() string {
	// Iota uso de constantes de golang

	// En la definición de constante, iota puede iterar fácilmente un valor de 0 en incrementos de 1, 0,1,2,3,4,5 ...
	// En este ejemplo, el formato de tamaño de archivo 2 se redondea a la décima potencia, y KB es 1 a la izquierda por 10 bits y MB a la izquierda por 20 bits
	// Sprintf ("% f", x) en este artículo no causará errores de bucle sin fin porque está definido en el método String, porque% f no intentará llamar a String ()
	
	switch {
	case b >= YB:
		return fmt.Sprintf("%.2fYB", b/YB)
	case b >= ZB:
		return fmt.Sprintf("%.2fZB", b/ZB)
	case b >= EB:
		return fmt.Sprintf("%.2fEB", b/EB)
	case b >= PB:
		return fmt.Sprintf("%.2fPB", b/PB)
	case b >= TB:
		return fmt.Sprintf("%.2fTB", b/TB)
	case b >= GB:
		return fmt.Sprintf("%.2fGB", b/GB)
	case b >= MB:
		return fmt.Sprintf("%.2fMB", b/MB)
	case b >= KB:
		return fmt.Sprintf("%.2fKB", b/KB)

	}
	return fmt.Sprintf("%.2fB", b)

}
func main() {
	fmt.Println(ByteSize(1e10))
}


