// Copyright 2024 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package state

import (
	"testing"

	gc "gopkg.in/check.v1"

	"github.com/juju/juju/domain/application"
	"github.com/juju/juju/domain/application/architecture"
	"github.com/juju/juju/domain/application/charm"
	schematesting "github.com/juju/juju/domain/schema/testing"
)

func TestPackage(t *testing.T) {
	gc.TestingT(t)
}

type baseSuite struct {
	schematesting.ModelSuite
}

func (s *baseSuite) minimalMetadata(c *gc.C, name string) charm.Metadata {
	return charm.Metadata{
		Name: name,
	}
}

func (s *baseSuite) minimalManifest(c *gc.C) charm.Manifest {
	return charm.Manifest{
		Bases: []charm.Base{
			{
				Name: "ubuntu",
				Channel: charm.Channel{
					Risk: charm.RiskStable,
				},
				Architectures: []string{"amd64"},
			},
		},
	}
}

func (s *baseSuite) addApplicationArgForResources(c *gc.C,
	name string,
	charmResources map[string]charm.Resource,
	addResourcesArgs []application.AddApplicationResourceArg) application.AddApplicationArg {
	platform := application.Platform{
		Channel:      "666",
		OSType:       application.Ubuntu,
		Architecture: architecture.ARM64,
	}
	channel := &application.Channel{
		Track:  "track",
		Risk:   "risk",
		Branch: "branch",
	}

	metadata := s.minimalMetadata(c, name)
	metadata.Resources = charmResources
	return application.AddApplicationArg{
		Platform: platform,
		Charm: charm.Charm{
			Metadata:      metadata,
			Manifest:      s.minimalManifest(c),
			Source:        charm.CharmHubSource,
			ReferenceName: name,
			Revision:      42,
			Architecture:  architecture.ARM64,
		},
		CharmDownloadInfo: &charm.DownloadInfo{
			Provenance:         charm.ProvenanceDownload,
			CharmhubIdentifier: "ident-1",
			DownloadURL:        "http://example.com/charm",
			DownloadSize:       666,
		},
		Scale:     1,
		Channel:   channel,
		Resources: addResourcesArgs,
	}
}
