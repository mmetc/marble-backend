package infra

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLicenseConfigurationDefaults(t *testing.T) {
	config := NewLicenseConfiguration()

	assert.Equal(t, "", config.LicenseKey)
	assert.False(t, config.KillIfReadLicenseError)
}

func TestNewLicenseConfigurationReadsEnv(t *testing.T) {
	t.Setenv("LICENSE_KEY", "license-key")
	t.Setenv("KILL_IF_READ_LICENSE_ERROR", "true")

	config := NewLicenseConfiguration()

	assert.Equal(t, "license-key", config.LicenseKey)
	assert.True(t, config.KillIfReadLicenseError)
}
