package homebrew

type Formula struct {
	Name                    string                `json:"name"`
	FullName                string                `json:"full_name"`
	Tap                     string                `json:"tap"`
	Oldnames                []any                 `json:"oldnames"`
	Aliases                 []string              `json:"aliases"`
	VersionedFormulae       []string              `json:"versioned_formulae"`
	Desc                    string                `json:"desc"`
	License                 string                `json:"license"`
	Homepage                string                `json:"homepage"`
	Versions                Versions              `json:"versions"`
	Urls                    Urls                  `json:"urls"`
	Revision                int                   `json:"revision"`
	VersionScheme           int                   `json:"version_scheme"`
	Bottle                  Bottle                `json:"bottle"`
	PourBottleOnlyIf        any                   `json:"pour_bottle_only_if"`
	KegOnly                 bool                  `json:"keg_only"`
	KegOnlyReason           any                   `json:"keg_only_reason"`
	Options                 []any                 `json:"options"`
	BuildDependencies       []string              `json:"build_dependencies"`
	Dependencies            []string              `json:"dependencies"`
	TestDependencies        []any                 `json:"test_dependencies"`
	RecommendedDependencies []any                 `json:"recommended_dependencies"`
	OptionalDependencies    []any                 `json:"optional_dependencies"`
	UsesFromMacos           []string              `json:"uses_from_macos"`
	UsesFromMacosBounds     []UsesFromMacosBounds `json:"uses_from_macos_bounds"`
	Requirements            []any                 `json:"requirements"`
	ConflictsWith           []any                 `json:"conflicts_with"`
	ConflictsWithReasons    []any                 `json:"conflicts_with_reasons"`
	LinkOverwrite           []any                 `json:"link_overwrite"`
	Caveats                 any                   `json:"caveats"`
	Installed               []any                 `json:"installed"`
	LinkedKeg               any                   `json:"linked_keg"`
	Pinned                  bool                  `json:"pinned"`
	Outdated                bool                  `json:"outdated"`
	Deprecated              bool                  `json:"deprecated"`
	DeprecationDate         any                   `json:"deprecation_date"`
	DeprecationReason       any                   `json:"deprecation_reason"`
	Disabled                bool                  `json:"disabled"`
	DisableDate             any                   `json:"disable_date"`
	DisableReason           any                   `json:"disable_reason"`
	PostInstallDefined      bool                  `json:"post_install_defined"`
	Service                 any                   `json:"service"`
	TapGitHead              string                `json:"tap_git_head"`
	RubySourcePath          string                `json:"ruby_source_path"`
	RubySourceChecksum      RubySourceChecksum    `json:"ruby_source_checksum"`
	Variations              Variations            `json:"variations"`
	Analytics               Analytics             `json:"analytics"`
	GeneratedDate           string                `json:"generated_date"`
}

type Versions struct {
	Stable string `json:"stable"`
	Head   string `json:"head"`
	Bottle bool   `json:"bottle"`
}

type URLStable struct {
	URL      string `json:"url"`
	Tag      any    `json:"tag"`
	Revision any    `json:"revision"`
	Using    any    `json:"using"`
	Checksum string `json:"checksum"`
}
type URLHead struct {
	URL    string `json:"url"`
	Branch string `json:"branch"`
	Using  any    `json:"using"`
}
type Urls struct {
	Stable URLStable `json:"stable"`
	Head   URLHead   `json:"head"`
}

type FileVariant struct {
	Cellar string `json:"cellar"`
	URL    string `json:"url"`
	Sha256 string `json:"sha256"`
}

type BottleStable struct {
	Rebuild int                    `json:"rebuild"`
	RootURL string                 `json:"root_url"`
	Files   map[string]FileVariant `json:"files"`
}
type Bottle struct {
	Stable BottleStable `json:"stable"`
}

type UsesFromMacosBounds struct {
}
type RubySourceChecksum struct {
	Sha256 string `json:"sha256"`
}

type Variation struct {
	BuildDependencies []string `json:"build_dependencies"`
	Dependencies      []string `json:"dependencies"`
}

type Variations struct {
	Sequoia    Variation `json:"sequoia"`
	Sonoma     Variation `json:"sonoma"`
	Ventura    Variation `json:"ventura"`
	Monterey   Variation `json:"monterey"`
	BigSur     Variation `json:"big_sur"`
	Catalina   Variation `json:"catalina"`
	Mojave     Variation `json:"mojave"`
	HighSierra Variation `json:"high_sierra"`
	Sierra     Variation `json:"sierra"`
	ElCapitan  Variation `json:"el_capitan"`
	X8664Linux Variation `json:"x86_64_linux"`
}

type Three0D struct {
	Ffmpeg     int `json:"ffmpeg"`
	FfmpegHEAD int `json:"ffmpeg --HEAD"`
}
type Nine0D struct {
	Ffmpeg     int `json:"ffmpeg"`
	FfmpegHEAD int `json:"ffmpeg --HEAD"`
}
type Three65D struct {
	Ffmpeg     int `json:"ffmpeg"`
	FfmpegHEAD int `json:"ffmpeg --HEAD"`
}
type Install struct {
	Three0D  Three0D  `json:"30d"`
	Nine0D   Nine0D   `json:"90d"`
	Three65D Three65D `json:"365d"`
}
type InstallOnRequest struct {
	Three0D  Three0D  `json:"30d"`
	Nine0D   Nine0D   `json:"90d"`
	Three65D Three65D `json:"365d"`
}
type BuildError struct {
	Three0D Three0D `json:"30d"`
}
type Analytics struct {
	Install          Install          `json:"install"`
	InstallOnRequest InstallOnRequest `json:"install_on_request"`
	BuildError       BuildError       `json:"build_error"`
}
