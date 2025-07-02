package math

import "testing"

func TestAbs(t *testing.T) {
        tests := []struct {
                input    int
                expected int
        }{
                {5, 5},
                {-5, 5},
                {0, 0},
                {-100, 100},
        }

        for _, test := range tests {
                result := Abs(test.input)
                if result != test.expected {
                        t.Errorf("Abs(%d) = %d; expected %d", test.input, result, test.expected)
                }
        }
}

func TestClamp(t *testing.T) {
        tests := []struct {
                value    int
                min      int
                max      int
                expected int
        }{
                {5, 0, 10, 5},
                {-5, 0, 10, 0},
                {15, 0, 10, 10},
                {0, 0, 10, 0},
                {10, 0, 10, 10},
        }

        for _, test := range tests {
                result := Clamp(test.value, test.min, test.max)
                if result != test.expected {
                        t.Errorf("Clamp(%d, %d, %d) = %d; expected %d", test.value, test.min, test.max, result, test.expected)
                }
        }
}

func TestGCD(t *testing.T) {
        tests := []struct {
                a        int
                b        int
                expected int
        }{
                {12, 18, 6},
                {48, 18, 6},
                {7, 5, 1},
                {0, 5, 5},
        }

        for _, test := range tests {
                result := GCD(test.a, test.b)
                if result != test.expected {
                        t.Errorf("GCD(%d, %d) = %d; expected %d", test.a, test.b, result, test.expected)
                }
        }
}

func TestLCM(t *testing.T) {
        tests := []struct {
                a        int
                b        int
                expected int
        }{
                {4, 6, 12},
                {12, 18, 36},
                {7, 5, 35},
        }

        for _, test := range tests {
                result := LCM(test.a, test.b)
                if result != test.expected {
                        t.Errorf("LCM(%d, %d) = %d; expected %d", test.a, test.b, result, test.expected)
                }
        }
}

func TestIsPrime(t *testing.T) {
        tests := []struct {
                input    int
                expected bool
        }{
                {2, true},
                {3, true},
                {4, false},
                {17, true},
                {25, false},
                {1, false},
                {0, false},
                {-5, false},
        }

        for _, test := range tests {
                result := IsPrime(test.input)
                if result != test.expected {
                        t.Errorf("IsPrime(%d) = %v; expected %v", test.input, result, test.expected)
                }
        }
}

func TestFactorial(t *testing.T) {
        tests := []struct {
                input    int
                expected int
        }{
                {0, 1},
                {1, 1},
                {5, 120},
                {6, 720},
                {-1, 0},
        }

        for _, test := range tests {
                result := Factorial(test.input)
                if result != test.expected {
                        t.Errorf("Factorial(%d) = %d; expected %d", test.input, result, test.expected)
                }
        }
}

func TestFibonacci(t *testing.T) {
        tests := []struct {
                input    int
                expected int
        }{
                {0, 0},
                {1, 1},
                {2, 1},
                {3, 2},
                {10, 55},
        }

        for _, test := range tests {
                result := Fibonacci(test.input)
                if result != test.expected {
                        t.Errorf("Fibonacci(%d) = %d; expected %d", test.input, result, test.expected)
                }
        }
}

func TestRandomInt(t *testing.T) {
        min, max := 10, 20
        
        // Test multiple times to ensure it's within range
        for i := 0; i < 100; i++ {
                result := RandomInt(min, max)
                if result < min || result > max {
                        t.Errorf("RandomInt(%d, %d) = %d; expected value between %d and %d", min, max, result, min, max)
                }
        }
        
        // Test edge case where min == max
        singleValue := RandomInt(5, 5)
        if singleValue != 5 {
                t.Errorf("RandomInt(5, 5) = %d; expected 5", singleValue)
        }
        
        // Test that different calls can produce different results (probabilistic test)
        results := make(map[int]bool)
        for i := 0; i < 50; i++ {
                results[RandomInt(1, 100)] = true
        }
        if len(results) < 10 { // Should have at least 10 different values in 50 calls
                t.Errorf("RandomInt not producing diverse results: got %d unique values in 50 calls", len(results))
        }
}

func TestPower(t *testing.T) {
        tests := []struct {
                x        int
                y        int
                expected int
        }{
                {2, 3, 8},
                {5, 2, 25},
                {10, 0, 1},
                {3, 1, 3},
        }

        for _, test := range tests {
                result := Power(test.x, test.y)
                if result != test.expected {
                        t.Errorf("Power(%d, %d) = %d; expected %d", test.x, test.y, result, test.expected)
                }
        }
}

func TestIsEven(t *testing.T) {
        tests := []struct {
                input    int
                expected bool
        }{
                {2, true},
                {3, false},
                {0, true},
                {-2, true},
                {-3, false},
        }

        for _, test := range tests {
                result := IsEven(test.input)
                if result != test.expected {
                        t.Errorf("IsEven(%d) = %v; expected %v", test.input, result, test.expected)
                }
        }
}

func TestIsOdd(t *testing.T) {
        tests := []struct {
                input    int
                expected bool
        }{
                {2, false},
                {3, true},
                {0, false},
                {-2, false},
                {-3, true},
        }

        for _, test := range tests {
                result := IsOdd(test.input)
                if result != test.expected {
                        t.Errorf("IsOdd(%d) = %v; expected %v", test.input, result, test.expected)
                }
        }
}

func TestRound(t *testing.T) {
        tests := []struct {
                input    float64
                expected int
        }{
                {3.14, 3},
                {3.5, 4},
                {3.7, 4},
                {-3.14, -3},
                {-3.5, -4},
        }

        for _, test := range tests {
                result := Round(test.input)
                if result != test.expected {
                        t.Errorf("Round(%f) = %d; expected %d", test.input, result, test.expected)
                }
        }
}
