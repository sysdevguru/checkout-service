package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	filePath := "../config.yml"

	global, err := loadConfig(filePath)
	assert.Nil(t, err)
	assert.Equal(t, "€", global.Server.Currency)
	assert.Equal(t, ":8080", global.Server.GRPCPort)
	assert.Equal(t, 5.00, global.Server.PenPrice)
	assert.Equal(t, 20.00, global.Server.TshirtPrice)
	assert.Equal(t, 7.50, global.Server.MugPrice)
}