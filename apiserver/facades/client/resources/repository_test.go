// Copyright 2020 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package resources

import (
	"time"

	"github.com/golang/mock/gomock"
	"github.com/juju/charm/v9"
	charmresource "github.com/juju/charm/v9/resource"
	jc "github.com/juju/testing/checkers"
	"github.com/juju/utils/v2/hash"
	gc "gopkg.in/check.v1"

	"github.com/juju/juju/apiserver/facades/client/resources/mocks"
	"github.com/juju/juju/charmhub/transport"
	corecharm "github.com/juju/juju/core/charm"
)

var _ = gc.Suite(&CharmHubClientSuite{})

type CharmHubClientSuite struct {
	client *mocks.MockCharmHub
	logger *mocks.MockLogger
}

func (s *CharmHubClientSuite) TestResolveResources(c *gc.C) {
	defer s.setupMocks(c).Finish()
	s.expectRefresh()
	s.expectListResourceRevisions(2)

	fp, err := charmresource.ParseFingerprint("38b060a751ac96384cd9327eb1b1e36a21fdb71114be07434c0cc7bf63f6e1da274edebfe76f65fbd51ad2f14898b95b")
	c.Assert(err, jc.ErrorIsNil)
	result, err := s.newClient().ResolveResources([]charmresource.Resource{{
		Meta:        charmresource.Meta{Name: "wal-e", Type: 1, Path: "wal-e.snap", Description: "WAL-E Snap Package"},
		Origin:      charmresource.OriginUpload,
		Revision:    1,
		Fingerprint: fp,
		Size:        0,
	}, {
		Meta:        charmresource.Meta{Name: "wal-e", Type: 1, Path: "wal-e.snap", Description: "WAL-E Snap Package"},
		Origin:      charmresource.OriginStore,
		Revision:    2,
		Fingerprint: fp,
		Size:        0,
	}})
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(result, gc.DeepEquals, []charmresource.Resource{{
		Meta:        charmresource.Meta{Name: "wal-e", Type: 1, Path: "wal-e.snap", Description: "WAL-E Snap Package"},
		Origin:      charmresource.OriginUpload,
		Revision:    1,
		Fingerprint: fp,
		Size:        0,
	}, {
		Meta:        charmresource.Meta{Name: "wal-e", Type: 1, Path: "wal-e.snap", Description: "WAL-E Snap Package"},
		Origin:      charmresource.OriginStore,
		Revision:    2,
		Fingerprint: fp,
		Size:        0,
	}})
}

func (s *CharmHubClientSuite) TestResolveResourcesFromStore(c *gc.C) {
	defer s.setupMocks(c).Finish()
	s.expectRefresh()
	s.expectListResourceRevisions(1)

	fp, err := charmresource.ParseFingerprint("38b060a751ac96384cd9327eb1b1e36a21fdb71114be07434c0cc7bf63f6e1da274edebfe76f65fbd51ad2f14898b95b")
	c.Assert(err, jc.ErrorIsNil)
	result, err := s.newClient().ResolveResources([]charmresource.Resource{{
		Meta:     charmresource.Meta{Name: "wal-e", Type: 1, Path: "wal-e.snap", Description: "WAL-E Snap Package"},
		Origin:   charmresource.OriginStore,
		Revision: 1,
		Size:     0,
	}})
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(result, gc.DeepEquals, []charmresource.Resource{{
		Meta:        charmresource.Meta{Name: "wal-e", Type: 1, Path: "wal-e.snap", Description: "WAL-E Snap Package"},
		Origin:      charmresource.OriginStore,
		Revision:    1,
		Fingerprint: fp,
		Size:        0,
	}})
}

func (s *CharmHubClientSuite) TestResolveResourcesFromStoreNoRevision(c *gc.C) {
	defer s.setupMocks(c).Finish()
	s.expectRefreshWithRevision(1)

	fp, err := charmresource.ParseFingerprint("38b060a751ac96384cd9327eb1b1e36a21fdb71114be07434c0cc7bf63f6e1da274edebfe76f65fbd51ad2f14898b95b")
	c.Assert(err, jc.ErrorIsNil)
	result, err := s.newClient().ResolveResources([]charmresource.Resource{{
		Meta:     charmresource.Meta{Name: "wal-e", Type: 1, Path: "wal-e.snap", Description: "WAL-E Snap Package"},
		Origin:   charmresource.OriginStore,
		Revision: -1,
		Size:     0,
	}})
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(result, gc.DeepEquals, []charmresource.Resource{{
		Meta:        charmresource.Meta{Name: "wal-e", Type: 1, Path: "wal-e.snap", Description: "WAL-E Snap Package"},
		Origin:      charmresource.OriginStore,
		Revision:    1,
		Fingerprint: fp,
		Size:        0,
	}})
}

func (s *CharmHubClientSuite) TestResolveResourcesNoMatchingRevision(c *gc.C) {
	defer s.setupMocks(c).Finish()
	s.expectRefresh()
	s.expectRefreshWithRevision(99)
	s.expectListResourceRevisions(3)

	_, err := s.newClient().ResolveResources([]charmresource.Resource{{
		Meta:     charmresource.Meta{Name: "wal-e", Type: 1, Path: "wal-e.snap", Description: "WAL-E Snap Package"},
		Origin:   charmresource.OriginStore,
		Revision: 1,
		Size:     0,
	}})
	c.Assert(err, gc.ErrorMatches, `charm resource "wal-e" at revision 1 not found`)
}

func (s *CharmHubClientSuite) TestResolveResourcesUpload(c *gc.C) {
	defer s.setupMocks(c).Finish()
	s.expectRefresh()

	result, err := s.newClient().ResolveResources([]charmresource.Resource{{
		Meta:     charmresource.Meta{Name: "wal-e", Type: 1, Path: "wal-e.snap", Description: "WAL-E Snap Package"},
		Origin:   charmresource.OriginUpload,
		Revision: 3,
		Fingerprint: charmresource.Fingerprint{
			Fingerprint: hash.Fingerprint{}},
		Size: 0,
	}})
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(result, gc.DeepEquals, []charmresource.Resource{{
		Meta:     charmresource.Meta{Name: "wal-e", Type: 1, Path: "wal-e.snap", Description: "WAL-E Snap Package"},
		Origin:   charmresource.OriginUpload,
		Revision: 3,
		Fingerprint: charmresource.Fingerprint{
			Fingerprint: hash.Fingerprint{}},
		Size: 0,
	}})
}

func (s *CharmHubClientSuite) setupMocks(c *gc.C) *gomock.Controller {
	ctrl := gomock.NewController(c)
	s.client = mocks.NewMockCharmHub(ctrl)
	s.logger = mocks.NewMockLogger(ctrl)
	s.logger.EXPECT().Tracef(gomock.Any(), gomock.Any()).AnyTimes().Do(
		func(msg string, args ...interface{}) {
			c.Logf("Trace: "+msg, args...)
		},
	)
	s.logger.EXPECT().Debugf(gomock.Any(), gomock.Any()).AnyTimes().Do(
		func(msg string, args ...interface{}) {
			c.Logf("Debug: "+msg, args...)
		},
	)
	return ctrl
}

func (s *CharmHubClientSuite) newClient() *charmHubClient {
	curl := charm.MustParseURL("ubuntu")
	channel, _ := charm.ParseChannel("stable")
	c := &charmHubClient{
		client: s.client,
	}
	c.resourceClient = resourceClient{
		client: c,
		id: CharmID{
			URL: curl,
			Origin: corecharm.Origin{
				ID:      "meshuggah",
				Source:  corecharm.CharmHub,
				Channel: &channel,
				Platform: corecharm.Platform{
					OS:           "ubuntu",
					Series:       "focal",
					Architecture: "amd64",
				},
			},
		},
		logger: s.logger,
	}
	return c
}

func (s *CharmHubClientSuite) expectRefresh() {
	s.expectRefreshWithRevision(0)
}

func resourceRevision(rev int) transport.ResourceRevision {
	return transport.ResourceRevision{
		Download: transport.Download{
			HashSHA384: "38b060a751ac96384cd9327eb1b1e36a21fdb71114be07434c0cc7bf63f6e1da274edebfe76f65fbd51ad2f14898b95b",
			Size:       0,
			URL:        "https://api.staging.charmhub.io/api/v1/resources/download/charm_jmeJLrjWpJX9OglKSeUHCwgyaCNuoQjD.wal-e_0",
		},
		Name:        "wal-e",
		Revision:    rev,
		Type:        "file",
		Filename:    "wal-e.snap",
		Description: "WAL-E Snap Package",
	}
}

func (s *CharmHubClientSuite) expectRefreshWithRevision(rev int) {
	resp := []transport.RefreshResponse{
		{
			Entity: transport.RefreshEntity{
				CreatedAt: time.Date(2020, 7, 7, 9, 39, 44, 132000000, time.UTC),
				Download:  transport.Download{HashSHA256: "c97e1efc5367d2fdcfdf29f4a2243b13765cc9cbdfad19627a29ac903c01ae63", Size: 5487460, URL: "https://api.staging.charmhub.io/api/v1/charms/download/jmeJLrjWpJX9OglKSeUHCwgyaCNuoQjD_208.charm"},
				ID:        "jmeJLrjWpJX9OglKSeUHCwgyaCNuoQjD",
				Name:      "ubuntu",
				Resources: []transport.ResourceRevision{
					resourceRevision(rev),
				},
				Summary: "PostgreSQL object-relational SQL database (supported version)",
				Version: "208",
			},
			EffectiveChannel: "latest/stable",
			Error:            (*transport.APIError)(nil),
			Name:             "postgresql",
			Result:           "download",
		},
	}
	s.client.EXPECT().Refresh(gomock.Any(), gomock.Any()).Return(resp, nil)
}

func (s *CharmHubClientSuite) expectListResourceRevisions(rev int) {
	resp := []transport.ResourceRevision{
		resourceRevision(rev),
	}
	s.client.EXPECT().ListResourceRevisions(gomock.Any(), gomock.Any(), gomock.Any()).Return(resp, nil)
}
