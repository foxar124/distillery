package install

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/glamorousis/distillery/pkg/config"
	"github.com/glamorousis/distillery/pkg/provider"
	"github.com/glamorousis/distillery/pkg/source"
)

func Test_NewSource(t *testing.T) {
	t.Parallel()

	cases := []struct {
		source        string
		defaultSource string
		error         bool
		want          provider.ISource
	}{
		{
			source: "ekristen/aws-nuke",
			want: &source.GitHub{
				Owner:   "ekristen",
				Repo:    "aws-nuke",
				Version: "latest",
			},
		},
		{
			source: "github/ekristen/aws-nuke",
			want: &source.GitHub{
				Owner:   "ekristen",
				Repo:    "aws-nuke",
				Version: "latest",
			},
		},
		{
			source: "github.com/ekristen/aws-nuke",
			want: &source.GitHub{
				Owner:   "ekristen",
				Repo:    "aws-nuke",
				Version: "latest",
			},
		},
		{
			source: "ekristen/aws-nuke@3.1.1",
			want: &source.GitHub{
				Owner:   "ekristen",
				Repo:    "aws-nuke",
				Version: "3.1.1",
			},
		},
		{
			source: "github/ekristen/aws-nuke@3.1.1",
			want: &source.GitHub{
				Owner:   "ekristen",
				Repo:    "aws-nuke",
				Version: "3.1.1",
			},
		},
		{
			source: "github.com/ekristen/aws-nuke@3.1.1",
			want: &source.GitHub{
				Owner:   "ekristen",
				Repo:    "aws-nuke",
				Version: "3.1.1",
			},
		},
		{
			source: "homebrew/aws-nuke",
			want: &source.Homebrew{
				Formula: "aws-nuke",
				Version: "latest",
			},
		},
		{
			source: "hashicorp/terraform",
			want: &source.Hashicorp{
				Owner: "hashicorp",
				Repo:  "terraform",
			},
		},
		{
			source:        "opentufu",
			defaultSource: "homebrew",
			want: &source.Homebrew{
				Formula: "opentufu",
				Version: "latest",
			},
		},
		{
			source:        "terraform",
			defaultSource: "hashicorp",
			want: &source.Hashicorp{
				Owner: "hashicorp",
				Repo:  "terraform",
			},
		},
		{
			source:        "gitlab-org/gitlab-runner",
			defaultSource: "gitlab",
			want: &source.GitLab{
				Owner: "gitlab-org",
				Repo:  "gitlab-runner",
			},
		},
		{
			source:        "terraform",
			defaultSource: "unknown",
			error:         true,
			want:          nil,
		},
		{
			source: "github/hashicorp/terraform",
			want: &source.Hashicorp{
				Owner: "hashicorp",
				Repo:  "terraform",
			},
		},
		{
			source: "gitlab/gitlab-org/gitlab-runner",
			want: &source.GitLab{
				Owner: "gitlab-org",
				Repo:  "gitlab-runner",
			},
		},
		{
			source:        "unknown/unknown",
			defaultSource: "unknown",
			error:         true,
		},
		{
			source: "unknown/some-owner/some-repo",
			error:  true,
		},
		{
			source: "unknown/some-owner/some-repo/extra@3.1.1",
			error:  true,
		},
	}

	for _, tt := range cases {
		t.Run(tt.source, func(t *testing.T) {
			cfg, err := config.New("/tmp/test/path")
			assert.NoError(t, err)
			if tt.defaultSource != "" {
				cfg.DefaultSource = tt.defaultSource
			}

			got, err := NewSource(tt.source, &provider.Options{Config: cfg})
			if tt.error {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.want.GetSource(), got.GetSource())
		})
	}
}
