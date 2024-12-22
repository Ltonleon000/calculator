package calculator

import "testing"

func TestEvaluate(t *testing.T) {
	tests := []struct {
		name        string
		expression  string
		want        float64
		wantErr     bool
		errorString string
	}{
		{
			name:       "simple addition",
			expression: "2+2",
			want:       4,
			wantErr:    false,
		},
		{
			name:       "addition and subtraction",
			expression: "10-5+3",
			want:       8,
			wantErr:    false,
		},
		{
			name:        "empty expression",
			expression:  "",
			want:        0,
			wantErr:     true,
			errorString: "empty expression",
		},
		{
			name:        "invalid character",
			expression:  "2+a",
			want:        0,
			wantErr:     true,
			errorString: "invalid character in expression: a",
		},
		{
			name:       "expression with spaces",
			expression: "2 + 2",
			want:       4,
			wantErr:    false,
		},
		{
			name:       "negative numbers",
			expression: "-5+3",
			want:       -2,
			wantErr:    false,
		},
		{
			name:        "invalid format consecutive operators",
			expression:  "2++2",
			want:        0,
			wantErr:     true,
			errorString: "invalid format: consecutive operators",
		},
		{
			name:        "invalid format ends with operator",
			expression:  "2+",
			want:        0,
			wantErr:     true,
			errorString: "invalid format: expression ends with operator",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Evaluate(tt.expression)
			if (err != nil) != tt.wantErr {
				t.Errorf("Evaluate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && err != nil && err.Error() != tt.errorString {
				t.Errorf("Evaluate() error = %v, wantErr %v", err.Error(), tt.errorString)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("Evaluate() = %v, want %v", got, tt.want)
			}
		})
	}
}
