package log

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFunctionNameAtRuntime(t *testing.T) {
	tests := []struct {
		name string
		want Caller
	}{
		{
			name: "test 1",
			want: Caller{
				FunctionName: "TestGetFunctionNameAtRuntime.func1",
				FilePath:     "pkg/log/runtime_test.go",
				Line:         25,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetFunctionNameAtRuntime(1)
			assert.Equal(t, tt.want, got)
		})
	}
}
