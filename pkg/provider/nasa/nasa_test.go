package nasa

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestGetNasaLaunches(t *testing.T) {

	if os.Getenv("TEST-LIVE") == "" {
		t.Skip()
	}

	tc := []struct {
		name          string
		expectedError bool
		size          int
		from          time.Time
		to            time.Time
	}{
		{"testGetNext10Launches", false, 10, time.Now(), time.Now().AddDate(10, 0, 0)}, // test default getLaunch
		{"testGetNextLaunch", false, 1, time.Now(), time.Now().AddDate(10, 0, 0)},      // test size getLaunch

	}

	for _, test := range tc {
		t.Run(test.name, func(t *testing.T) {
			prov, err := NewProvider()
			require.NoError(t, err)

			resp, err := prov.GetLaunches(test.from, test.to, test.size)

			if test.expectedError {
				require.Error(t, err)
			} else {
				require.Equal(t, len(resp.Hits.Hits), test.size)
				require.NoError(t, err)
			}
		})
	}

}
