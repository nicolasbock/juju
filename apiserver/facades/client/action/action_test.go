// Copyright 2014 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package action_test

import (
	"github.com/juju/testing"
	gc "gopkg.in/check.v1"
)

type baseSuite struct {
	testing.IsolationSuite
}

type actionSuite struct {
	baseSuite
}

var _ = gc.Suite(&actionSuite{})

func (s *actionSuite) TestStub(c *gc.C) {
	c.Skip(`This suite is missing tests for the following scenarios:
- Enqueueing an action against multiple units, verifying persisted receivers/params/status etc.
- Enqueueing actions and cancelling a subset, verifying actions in the cancelled state.
- Beginning an action and then cancelling, verifying the aborting status.
- Ensuring the ApplicationsCharmsActions returns the correct result for input app tags.
- Enqueueing and watching an action, observing its action messages as it progresses. 
`)
}
