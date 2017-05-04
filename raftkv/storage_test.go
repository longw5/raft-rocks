package raftkv

import (
	"testing"

	"os"

	"github.com/stretchr/testify/assert"
)

func TestRocksDBStorage(t *testing.T) {
	A := assert.New(t)
	const dir = "./testrocks"
	db, err := MakeRocksDBStore(dir)
	A.NoError(err)
	defer db.Close()
	defer os.RemoveAll(dir)

	// first get should miss
	value, ok := db.Get("hello")
	A.Empty(value)
	A.False(ok)

	// then put a kv
	db.Put("hello", "world")

	// this get should hit
	value, ok = db.Get("hello")
	A.Equal("world", value)
	A.True(ok)
}
