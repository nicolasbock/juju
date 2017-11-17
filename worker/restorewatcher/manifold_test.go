// Copyright 2017 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package restorewatcher_test

import (
	"github.com/juju/errors"
	"github.com/juju/testing"
	jc "github.com/juju/testing/checkers"
	gc "gopkg.in/check.v1"
	worker "gopkg.in/juju/worker.v1"

	"github.com/juju/juju/state"
	"github.com/juju/juju/worker/dependency"
	dt "github.com/juju/juju/worker/dependency/testing"
	"github.com/juju/juju/worker/restorewatcher"
	"github.com/juju/juju/worker/workertest"
)

type ManifoldSuite struct {
	testing.IsolationSuite

	manifold     dependency.Manifold
	context      dependency.Context
	st           *state.State
	pool         *state.StatePool
	stateTracker stubStateTracker

	stub testing.Stub
}

var _ = gc.Suite(&ManifoldSuite{})

func (s *ManifoldSuite) SetUpTest(c *gc.C) {
	s.IsolationSuite.SetUpTest(c)

	s.st = new(state.State)
	s.pool = state.NewStatePool(s.st)
	s.stateTracker = stubStateTracker{pool: s.pool}
	s.stub.ResetCalls()

	s.context = s.newContext(nil)
	s.manifold = restorewatcher.Manifold(restorewatcher.ManifoldConfig{
		StateName:            "state",
		NewWorker:            s.newWorker,
		RestoreStatusChanged: s.restoreStatusChanged,
	})
}

func (s *ManifoldSuite) newContext(overlay map[string]interface{}) dependency.Context {
	resources := map[string]interface{}{"state": &s.stateTracker}
	for k, v := range overlay {
		resources[k] = v
	}
	return dt.StubContext(nil, resources)
}

func (s *ManifoldSuite) newWorker(config restorewatcher.Config) (worker.Worker, error) {
	s.stub.MethodCall(s, "NewWorker", config)
	if err := s.stub.NextErr(); err != nil {
		return nil, err
	}
	w := worker.NewRunner(worker.RunnerParams{})
	s.AddCleanup(func(c *gc.C) { workertest.DirtyKill(c, w) })
	return w, nil
}

func (s *ManifoldSuite) restoreStatusChanged(status state.RestoreStatus) error {
	s.stub.MethodCall(s, "RestoreStatusChanged", status)
	return s.stub.NextErr()
}

var expectedInputs = []string{"state"}

func (s *ManifoldSuite) TestInputs(c *gc.C) {
	c.Assert(s.manifold.Inputs, jc.SameContents, expectedInputs)
}

func (s *ManifoldSuite) TestMissingInputs(c *gc.C) {
	for _, input := range expectedInputs {
		context := s.newContext(map[string]interface{}{
			input: dependency.ErrMissing,
		})
		_, err := s.manifold.Start(context)
		c.Assert(errors.Cause(err), gc.Equals, dependency.ErrMissing)
	}
}

func (s *ManifoldSuite) TestStart(c *gc.C) {
	w := s.startWorkerClean(c)
	workertest.CleanKill(c, w)

	s.stub.CheckCallNames(c, "NewWorker")
	args := s.stub.Calls()[0].Args
	c.Assert(args, gc.HasLen, 1)
	c.Assert(args[0], gc.FitsTypeOf, restorewatcher.Config{})
	config := args[0].(restorewatcher.Config)

	c.Assert(config.RestoreStatusChanged, gc.NotNil)
	err := config.RestoreStatusChanged(state.RestorePending)
	c.Assert(err, jc.ErrorIsNil)
	s.stub.CheckCall(c, 1, "RestoreStatusChanged", state.RestorePending)
	config.RestoreStatusChanged = nil

	c.Assert(config, jc.DeepEquals, restorewatcher.Config{
		RestoreInfoWatcher: restorewatcher.RestoreInfoWatcherShim{s.st},
	})
}

func (s *ManifoldSuite) TestStopWorkerClosesState(c *gc.C) {
	w := s.startWorkerClean(c)
	defer workertest.CleanKill(c, w)

	s.stateTracker.CheckCallNames(c, "Use")

	workertest.CleanKill(c, w)
	s.stateTracker.CheckCallNames(c, "Use", "Done")
}

func (s *ManifoldSuite) startWorkerClean(c *gc.C) worker.Worker {
	w, err := s.manifold.Start(s.context)
	c.Assert(err, jc.ErrorIsNil)
	workertest.CheckAlive(c, w)
	return w
}

type stubStateTracker struct {
	testing.Stub
	pool *state.StatePool
}

func (s *stubStateTracker) Use() (*state.StatePool, error) {
	s.MethodCall(s, "Use")
	return s.pool, s.NextErr()
}

func (s *stubStateTracker) Done() error {
	s.MethodCall(s, "Done")
	return s.NextErr()
}
