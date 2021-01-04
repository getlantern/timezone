package timezone

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIANANameForTime(t *testing.T) {
	shanghai, err := time.LoadLocation("Asia/Shanghai")
	require.NoError(t, err)

	now := time.Now()
	nowInChina := now.In(shanghai)
	tzName, err := IANANameForTime(nowInChina)
	require.NoError(t, err)

	t.Logf("tzName: %v", tzName)

	tz, err := time.LoadLocation(tzName)
	require.NoError(t, err)
	nowInTZ := now.In(tz)
	assert.Equal(t, nowInChina.String(), nowInTZ.String())
}
