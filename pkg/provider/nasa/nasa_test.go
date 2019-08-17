package nasa

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetNasaLaunches(t *testing.T) {
	
	tc := []struct{
		name string
		expectedError bool
	}{
		{"testDefault", false}, // test default getLaunch
	}
	
	for _, test := range tc {
		t.Run(test.name, func(t *testing.T) {
			prov, err := NewProvider()
			require.NoError(t,err)

			_, err = prov.GetLaunches()

			if test.expectedError {
				require.Error(t, err)
			}else {
				require.NoError(t,err)
			}
		})
	}

}
