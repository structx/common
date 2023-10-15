package kv_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/structx/common/kv"
	"github.com/syndtr/goleveldb/leveldb"
)

func init() {
	_ = os.Setenv("LEVELDB_PATH", "./testfiles/leveldb")
}

type LevelDBSuite struct {
	suite.Suite
	db *leveldb.DB
}

func (s *LevelDBSuite) SetupTest() {

	assert := s.Assert()

	db, err := kv.NewLevelDB()
	assert.NoError(err)

	s.db = db
}

func TestLevelDBSuitee(t *testing.T) {
	suite.Run(t, new(LevelDBSuite))
}
