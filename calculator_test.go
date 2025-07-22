package main

import(
	"testing"
)

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5.0
	if result != expected {
		t.Errorf("Add(2, 3) = %.2f; want %.2f", result, expected)
	}
}

func TestDivide(t *testing.T) {
	t.Run("Divide by non-zero", func(t *testing.T) {
		result, err := Divide(6, 3)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		if result != 2.0 {
			t.Errorf("Divide(6, 3) = %.2f; want 2.0", result)
		}
	})

	t.Run("Divide by zero", func(t *testing.T) {
		_, err := Divide(6, 0)
		if err == nil {
			t.Fatal("Expected error, got nil")
		}
	})
}

func TestCalculate(t *testing.T) {
	tests := []struct {
		name     string
		a, b     float64
		op       string
		expected float64
		wantErr  bool
	}{
		{"Add", 2, 3, "+", 5.0, false},
		{"Subtract", 5, 3, "-", 2.0, false},
		{"Multiply", 2, 3, "*", 6.0, false},
		{"Divide", 6, 3, "/", 2.0, false},
		{"Divide by zero", 6, 0, "/", 0, true},
		{"Unknown op", 6, 3, "?", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Calculate(tt.a, tt.b, tt.op)
			if (err != nil) != tt.wantErr {
				t.Errorf("Calculate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && result != tt.expected {
				t.Errorf("Calculate() = %.2f, want %.2f", result, tt.expected)
			}
		})
	}
}
