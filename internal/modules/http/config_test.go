package http

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	tests := []struct {
		name         string
		env          map[string]string
		expectedPort string
	}{
		{
			name:         "Given PORT set to 8082, expect 8082",
			env:          map[string]string{"PORT": "8082"},
			expectedPort: "8082",
		},
		{
			name:         "Given no PORT set, expect default port 8080",
			env:          map[string]string{"PORT": "8080"},
			expectedPort: "8080",
		},
		{
			name:         "Given PORT set to non-numeric, expect default port 8080",
			env:          map[string]string{"PORT": "abc"},
			expectedPort: "8080",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d: %s", i+1, tt.name), func(t *testing.T) {
			for key, value := range tt.env {
				os.Setenv(key, value)
				defer os.Unsetenv(key)
			}

			c := Config()

			assert.Equal(t, tt.expectedPort, c.Port, "The port configuration does not match the expected value.")
		})
	}
}
