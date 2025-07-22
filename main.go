import (
	"fmt"
	"flag"
	"os"
)

func main() {
	// 
	a := flag.Float64("a", 0, "first operand")
	b := flag.Float64("b", 0, "second operand")
	op := flag.String("op", "+", "operation (+, -, *, /)")
	flag.Parse()

	// Выполняем расчет
	result, err := Calculate(*a, *b, *op)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}


	// Выводим результат
	fmt.Printf("Result: %.2f\n", result)
}
