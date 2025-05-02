package inventory

import "path/filepath"

type Bin struct {
	Name     string
	Versions []*Version
	Source   string
	Owner    string
	Repo     string
}

func (b *Bin) ListVersions() []string {
	versionMap := make(map[string]struct{})
	for _, v := range b.Versions {
		versionMap[v.Version] = struct{}{}
	}

	var versions []string
	for version := range versionMap {
		versions = append(versions, version)
	}

	return versions
}

func (b *Bin) GetInstallPath(base string) string {
	return filepath.Join(base, b.Source, b.Owner, b.Repo)
}

type Version struct {
	Version string
	Path    string
	Latest  bool
	Target  string
}
