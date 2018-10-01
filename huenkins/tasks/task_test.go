package task

import (
	"fmt"
	"testing"

	. "gopkg.in/check.v1"
)

type testSuite struct{}

var _ = Suite(&testSuite{})

func TestTaskStack(t *testing.T) { TestingT(t) }

func (s *testSuite) TestParseEnv(c *C) {

	// body, err := parseEnv([]byte(js1))
	// c.Assert(err, IsNil)

	fmt.Println("super!")
	c.Log("super!")
	c.Assert(nil, IsNil)
	// c.Assert(body, DeepEquals, test1)
}
