// Copyright 2026 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.
package status

import (
	"bytes"
	"fmt"
	"io"

	jc "github.com/juju/testing/checkers"
	gc "gopkg.in/check.v1"

	jujutesting "github.com/juju/juju/testing"
)

var _ = gc.Suite(&outputTabularSuite{})

type outputTabularSuite struct {
	jujutesting.BaseSuite
}

func (s *outputTabularSuite) TestFormatOnelinePortsGroupedNumerically(c *gc.C) {
	fs := formattedStatus{
		Applications: map[string]applicationStatus{
			"app": {
				Units: map[string]unitStatus{
					"app/0": {
						OpenedPorts: []string{
							"9998/tcp",
							"9999/tcp",
							"10000/tcp",
							"10002/tcp",
							"10003/tcp",
							"10004/tcp",
						},
					},
				},
			},
		},
	}

	buff := &bytes.Buffer{}
	err := formatOneline(buff, false, fs, func(out io.Writer, format, uName string, u unitStatus, level int) {
		fmt.Fprintf(out, format, uName, "running", level)
	})
	c.Assert(err, jc.ErrorIsNil)

	c.Assert(buff.String(), jc.Contains, "9998-10000,10002-10004/tcp")
}
