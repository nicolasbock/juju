// Copyright 2018 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package common

import (
	"github.com/juju/errors"
	"gopkg.in/juju/names.v3"

	"github.com/juju/juju/apiserver/facade"
	"github.com/juju/juju/apiserver/params"
	"github.com/juju/juju/core/leadership"
	"github.com/juju/juju/permission"
	"github.com/juju/juju/state"
)

//go:generate mockgen -package mocks -destination mocks/leadership.go github.com/juju/juju/apiserver/common LeadershipPinningBackend,LeadershipMachine

// LeadershipMachine is an indirection for state.machine.
type LeadershipMachine interface {
	ApplicationNames() ([]string, error)
}

type leadershipMachine struct {
	*state.Machine
}

// LeadershipPinningBacked describes state method wrappers used by this API.
type LeadershipPinningBackend interface {
	Machine(string) (LeadershipMachine, error)
}

type leadershipPinningBackend struct {
	*state.State
}

// Machine wraps state.Machine to return an implementation
// of the LeadershipMachine indirection.
func (s leadershipPinningBackend) Machine(name string) (LeadershipMachine, error) {
	m, err := s.State.Machine(name)
	if err != nil {
		return nil, err
	}
	return leadershipMachine{m}, nil
}

// LeadershipPinningAPI exposes leadership pinning and unpinning functionality
// for remote use.
type LeadershipPinningAPI interface {
	PinMachineApplications() (params.PinApplicationsResults, error)
	UnpinMachineApplications() (params.PinApplicationsResults, error)
	PinnedLeadership() (params.PinnedLeadershipResult, error)
}

// NewLeadershipPinningFromContext creates and returns a new leadership from
// a facade context.
// This signature is suitable for facade registration.
func NewLeadershipPinningFromContext(ctx facade.Context) (*LeadershipPinning, error) {
	st := ctx.State()
	model, err := st.Model()
	if err != nil {
		return nil, errors.Trace(err)
	}
	pinner, err := ctx.LeadershipPinner(model.UUID())
	if err != nil {
		return nil, errors.Trace(err)
	}
	return NewLeadershipPinning(leadershipPinningBackend{st}, model.ModelTag(), pinner, ctx.Auth())
}

// NewLeadershipPinning creates and returns a new leadership API from the
// input tag, Pinner implementation and facade Authorizer.
func NewLeadershipPinning(
	st LeadershipPinningBackend, modelTag names.ModelTag, pinner leadership.Pinner, authorizer facade.Authorizer,
) (*LeadershipPinning, error) {
	return &LeadershipPinning{
		st:         st,
		modelTag:   modelTag,
		pinner:     pinner,
		authorizer: authorizer,
	}, nil
}

// LeadershipPinning defines a type for pinning and unpinning application
// leaders.
type LeadershipPinning struct {
	st         LeadershipPinningBackend
	modelTag   names.ModelTag
	pinner     leadership.Pinner
	authorizer facade.Authorizer
}

// PinnedLeadership returns all pinned applications and the entities that
// require their pinned behaviour, for leadership in the current model.
func (a *LeadershipPinning) PinnedLeadership() (params.PinnedLeadershipResult, error) {
	result := params.PinnedLeadershipResult{}

	canAccess, err := a.authorizer.HasPermission(permission.ReadAccess, a.modelTag)
	if err != nil {
		return result, errors.Trace(err)
	}
	if !canAccess {
		return result, ErrPerm
	}

	result.Result = a.pinner.PinnedLeadership()
	return result, nil
}

// PinApplicationLeaders pins leadership for applications based on the auth
// tag provided.
func (a *LeadershipPinning) PinApplicationLeaders() (params.PinApplicationsResults, error) {
	if !a.authorizer.AuthMachineAgent() {
		return params.PinApplicationsResults{}, ErrPerm
	}

	tag := a.authorizer.GetAuthTag()
	switch tag.Kind() {
	case names.MachineTagKind:
		return a.pinMachineApplications()
	default:
		return params.PinApplicationsResults{}, ErrPerm
	}
}

// UnpinApplicationLeaders unpins leadership for applications based on the auth
// tag provided.
func (a *LeadershipPinning) UnpinApplicationLeaders() (params.PinApplicationsResults, error) {
	if !a.authorizer.AuthMachineAgent() {
		return params.PinApplicationsResults{}, ErrPerm
	}

	tag := a.authorizer.GetAuthTag()
	switch tag.Kind() {
	case names.MachineTagKind:
		return a.unpinMachineApplications()
	default:
		return params.PinApplicationsResults{}, ErrPerm
	}
}

// GetMachineApplicationNames returns the applications associated with a
// machine.
func (a *LeadershipPinning) GetMachineApplicationNames() ([]string, error) {
	if !a.authorizer.AuthMachineAgent() {
		return nil, ErrPerm
	}

	tag := a.authorizer.GetAuthTag()
	m, err := a.st.Machine(tag.Id())
	if err != nil {
		return nil, errors.Trace(err)
	}
	apps, err := m.ApplicationNames()
	if err != nil {
		return nil, errors.Trace(err)
	}
	return apps, nil
}

// PinMachineApplicationsByName takes a slice of application names and attempts
// to pin them accordingly.
func (a *LeadershipPinning) PinMachineApplicationsByName(appNames []string) (params.PinApplicationsResults, error) {
	if !a.authorizer.AuthMachineAgent() {
		return params.PinApplicationsResults{}, ErrPerm
	}

	return a.pinMachineAppsOps(appNames, a.pinner.PinLeadership)
}

// UnpinMachineApplicationsByName takes a slice of application names and
// attempts to pin them accordingly.
func (a *LeadershipPinning) UnpinMachineApplicationsByName(appNames []string) (params.PinApplicationsResults, error) {
	if !a.authorizer.AuthMachineAgent() {
		return params.PinApplicationsResults{}, ErrPerm
	}

	return a.pinMachineAppsOps(appNames, a.pinner.UnpinLeadership)
}

// pinMachineApplications pins leadership for applications represented by units
// running on the auth'd machine.
func (a *LeadershipPinning) pinMachineApplications() (params.PinApplicationsResults, error) {
	appNames, err := a.GetMachineApplicationNames()
	if err != nil {
		return params.PinApplicationsResults{}, ErrPerm
	}
	return a.pinMachineAppsOps(appNames, a.pinner.PinLeadership)
}

// unpinMachineApplications unpins leadership for applications represented by
// units running on the auth'd machine.
func (a *LeadershipPinning) unpinMachineApplications() (params.PinApplicationsResults, error) {
	appNames, err := a.GetMachineApplicationNames()
	if err != nil {
		return params.PinApplicationsResults{}, ErrPerm
	}
	return a.pinMachineAppsOps(appNames, a.pinner.UnpinLeadership)
}

// pinMachineAppsOps runs the input pin/unpin operation against all
// applications represented by units on the authorised machine.
// An assumption is made that the validity of the auth tag has been verified
// by the caller.
func (a *LeadershipPinning) pinMachineAppsOps(appNames []string, op func(string, string) error) (params.PinApplicationsResults, error) {
	var result params.PinApplicationsResults

	tag := a.authorizer.GetAuthTag()

	results := make([]params.PinApplicationResult, len(appNames))
	for i, app := range appNames {
		results[i] = params.PinApplicationResult{
			ApplicationName: app,
		}
		if err := op(app, tag.String()); err != nil {
			results[i].Error = ServerError(err)
		}
	}
	result.Results = results
	return result, nil
}
