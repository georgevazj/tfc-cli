package models

type Data struct {
	Data []Workspace `json:"data"`
}

type Workspace struct {
	ID         string    `json:"id"`
	Type       string    `json:"type"`
	Attributes Attribute `json:"attributes"`
}

type Attribute struct {
	Name          string   `json:"name"`
	AutoApply     bool     `json:"auto-apply"`
	ResourceCount int      `json:"resource-count"`
	Repo          string   `json:"vcs-repo-identifier"`
	Description   string   `json:"description"`
	TagNames      []string `json:"tag-names"`
}
