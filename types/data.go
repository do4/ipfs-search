package types

import (
	"github.com/ipfs-search/ipfs-search/types/references"
	"time"
)

// TODO: Remove/reorganise me!

type common struct {
	*Resource
	Size       int                     `json:"size"`
	FirstSeen  time.Time               `json:"first_seen"`
	LastSeen   time.Time               `json:"last_seen"`
	References []references.References `json:"references"`
}

type Metadata map[string]interface{}

type Language struct {
	Language   string  `json:"language"`
	Confidence string  `json:"confidence"`
	RawScore   float64 `json:"rawScore"`
}

type ExtractedData struct {
	Metadata        Metadata   `json:"metadata"`
	Content         string     `json:"content"`
	URLs            []string   `json:"urls"`
	Languages       []Language `json:"language"`
	IpfsTikaVersion string     `json:"ipfs_tika_version"`
}

type File struct {
	common
	ExtractedData
}

type Link struct {
	Hash string
	Name string
	Size uint64
	Type string
}

type Directory struct {
	common
	Links []Link `json:"links"`
}
