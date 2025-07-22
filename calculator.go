import(
	"fmt"
	"errors"
)

// Add возвращает сумму двух чисел
func Add(a, b float64) float64 {
	return a + b
}

// Subtract возвращает разность a - b
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply возвращает произведение a * b
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide возвращает результат деления a / b или ошибку
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero is not allowed")
	}
	return a / b, nil
}

// Calculate выполняет операцию в зависимости от переданного оператора
func Calculate(a, b float64, operation string) (float64, error) {
	switch operation {
	case "+":
		return Add(a, b), nil
	case "-":
		return Subtract(a, b), nil
	case "*":
		return Multiply(a, b), nil
	case "/":
		return Divide(a, b)
	default:
		return 0, fmt.Errorf("unknown operation: %s", operation)
	}
}
