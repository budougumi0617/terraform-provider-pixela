package pixela

import (
	"context"
	"os"
	"testing"

	pixela "github.com/ebc-2in2crc/pixela4go"
	"github.com/google/go-cmp/cmp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

func Test_dataSourceGraphsRead(t *testing.T) {

	tests := []struct {
		name string
		want diag.Diagnostics
	}{
		{
			name: "confirmClientResponse",
			want: nil,
		},
	}
	for _, tt := range tests {
		usename := os.Getenv("PIXELA_USERNAME")
		token := os.Getenv("PIXELA_TOKEN")
		if usename == "" || token == "" {
			t.SkipNow()
		}
		t.Run(tt.name, func(t *testing.T) {

			m := pixela.New(usename, token)
			d := dataSourceGraphs().TestResourceData()
			got := dataSourceGraphsRead(context.TODO(), d, m)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("dataSourceGraphsRead: (-got +want)\n%s", diff)
			}
		})
	}
}
