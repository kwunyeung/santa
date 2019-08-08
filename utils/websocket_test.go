package utils

import (
	"os"
	"bytes"
	"log"
	"strings"
	"time"

	"testing"
	"github.com/stretchr/testify/require"
)
func TestWebsocketListen(t *testing.T) {
	app := setupWithPlentyBalanceAccount(t)

	// DO NOT TOUCH - wait next block
	time.Sleep(time.Second * 10)

	var logBuf bytes.Buffer
	log.SetOutput(&logBuf)
	defer log.SetOutput(os.Stderr)

	app.TriggerInterval = "1"
	app.ListenNewBLock(true)
	
	result := logBuf.String()
	require.True(t, strings.Contains(result, "[Success]"))
}