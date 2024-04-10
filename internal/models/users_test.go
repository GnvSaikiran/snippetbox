package models

import (
	"testing"

	"github.com/GnvSaikiran/snippetbox/internal/assert"
)

func TestUserModelExists(t *testing.T) {
	if testing.Short() {
		t.Skip("models: skipping integration test")
	}

	tests := []struct {
		name string
		ID   int
		want bool
	}{
		{
			name: "Valid ID",
			ID:   1,
			want: true,
		},
		{
			name: "Zero ID",
			ID:   0,
			want: false,
		},
		{
			name: "Non-existent ID",
			ID:   2,
			want: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			db := newTestDB(t)
			user := UserModel{DB: db}
			got, err := user.Exists(test.ID)

			assert.Equal(t, test.want, got)
			assert.NilError(t, err)
		})
	}
}
