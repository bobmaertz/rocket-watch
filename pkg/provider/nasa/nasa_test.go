package nasa

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetNasaLaunches(t *testing.T) {

	if os.Getenv("TEST-LIVE") == "" {
		t.Skip()
	}

	tc := []struct {
		name          string
		expectedError bool
		from          int
		size          int
	}{
		{"testGetNext10Launches", false, 0, 10}, // test default getLaunch
		{"testGetNextLaunch", false, 0, 1},      // test size getLaunch

	}

	for _, test := range tc {
		t.Run(test.name, func(t *testing.T) {
			prov, err := NewProvider()
			require.NoError(t, err)

			resp, err := prov.GetLaunches(test.from, test.size)

			if test.expectedError {
				require.Error(t, err)
			} else {
				require.Equal(t, len(resp.Hits.Hits), test.size)
				require.NoError(t, err)
			}
		})
	}

}
