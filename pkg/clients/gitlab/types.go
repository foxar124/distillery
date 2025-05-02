package gitlab

import "time"

type Release struct {
	Name            string      `json:"name"`
	TagName         string      `json:"tag_name"`
	Description     string      `json:"description"`
	CreatedAt       time.Time   `json:"created_at"`
	ReleasedAt      time.Time   `json:"released_at"`
	UpcomingRelease bool        `json:"upcoming_release"`
	Author          Author      `json:"author"`
	Commit          Commit      `json:"commit"`
	CommitPath      string      `json:"commit_path"`
	TagPath         string      `json:"tag_path"`
	Assets          *Assets     `json:"assets"`
	Evidences       []Evidences `json:"evidences"`
}
type Author struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Name      string `json:"name"`
	State     string `json:"state"`
	Locked    bool   `json:"locked"`
	AvatarURL string `json:"avatar_url"`
	WebURL    string `json:"web_url"`
}
type Trailers struct {
}
type ExtendedTrailers struct {
}
type Commit struct {
	ID               string           `json:"id"`
	ShortID          string           `json:"short_id"`
	CreatedAt        time.Time        `json:"created_at"`
	ParentIds        []string         `json:"parent_ids"`
	Title            string           `json:"title"`
	Message          string           `json:"message"`
	AuthorName       string           `json:"author_name"`
	AuthorEmail      string           `json:"author_email"`
	AuthoredDate     time.Time        `json:"authored_date"`
	CommitterName    string           `json:"committer_name"`
	CommitterEmail   string           `json:"committer_email"`
	CommittedDate    time.Time        `json:"committed_date"`
	Trailers         Trailers         `json:"trailers"`
	ExtendedTrailers ExtendedTrailers `json:"extended_trailers"`
	WebURL           string           `json:"web_url"`
}
type Sources struct {
	Format string `json:"format"`
	URL    string `json:"url"`
}
type Links struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	URL            string `json:"url"`
	DirectAssetURL string `json:"direct_asset_url"`
	LinkType       string `json:"link_type"`
}
type Assets struct {
	Count   int        `json:"count"`
	Sources []*Sources `json:"sources"`
	Links   []*Links   `json:"links"`
}
type Evidences struct {
	Sha         string    `json:"sha"`
	Filepath    string    `json:"filepath"`
	CollectedAt time.Time `json:"collected_at"`
}
