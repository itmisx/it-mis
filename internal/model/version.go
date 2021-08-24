package model

type Version struct {
	ID      int    `json:"id"`
	Version string `json:"version"`
}

func (Version) TableName() string {
	return "version"
}
