package main

import "testing"

func TestDivide(t *testing.T) {
	_, err := divide(10.0, 1.0)
	if err != nil {
		t.Error("Got an error when we should not have")
	}
}

func TestDivideByZero(t *testing.T) {
	_, err := divide(10.0, 0.0)
	if err == nil {
		t.Error("Expected an error when dividing by zero")
	}
}

// テーブルテスト
var tests = []struct {
	name     string
	dividend float32
	divisor  float32
	expected float32
	isErr    bool
}{
	{"10/2", 10.0, 2.0, 5.0, false},
	{"10/0", 10.0, 0.0, 0.0, true},
	{"10/-2", 10.0, -2.0, -5.0, false},
	{"-10/-2", -10.0, -2.0, 5.0, false},
	{"-10/2", -10.0, 2.0, -5.0, false},
	{"-10/0", -10.0, 0.0, 0.0, true},
}

func TestTableDivide(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			d, err := divide(test.dividend, test.divisor)
			if test.isErr {
				if err == nil {
					t.Errorf("Expected an error when dividing %f by %f", test.dividend, test.divisor)
				}
				return
			}
			if err != nil {
				t.Errorf("Got an error when we should not have: %v", err)
			}
			if d != test.expected {
				t.Errorf("Expected %f but got %f", test.expected, d)
			}
		})
	}
}

// testの実行
// go test

// testカバレッジを図る
// go test --cover

// testカバレッジを出力
// go test --coverprofile=coverage.out && go tool cover -html=coverage.out
