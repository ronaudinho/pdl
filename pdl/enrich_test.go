package pdl_test

import (
	"os"
	"testing"

	"github.com/ronaudinho/pdl/pdl"
)

func TestPerson_Enrich(t *testing.T) {
	tests := []struct {
		name   string
		params map[string]string
		// want   map[string]interface{} // not sure what expecations to define in tests right now
		want    bool // simply checks if results is nil
		wantErr error
	}{
		{
			name: "200",
			params: map[string]string{
				"profile": "https://linkedin.com/in/seanthorne",
			},
			want: true,
		},
		{
			name:    "400",
			wantErr: pdl.ErrMissingParams,
		},
	}

	// these are shared, could probably instantiate in test.Main
	key := os.Getenv("PDL_API_KEY")
	cli := pdl.NewV5(key)
	person := pdl.NewPerson(cli)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := person.Enrich(tt.params)
			if err != nil && tt.wantErr != err {
				t.Errorf("want %v, got %v", tt.wantErr, err)
				return
			}
			if res == nil && tt.want {
				t.Error("want response, got nil")
				return
			}
		})
	}
}
