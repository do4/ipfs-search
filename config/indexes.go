package config

type Indexes struct {
	Files       string `yaml:"files"`
	Directories string `yaml:"directories"`
	Invalids    string `yaml:"invalids"`
}

func IndexesDefaults() Indexes {
	return Indexes{
		Files:       "ipfs_files",
		Directories: "ipfs_directories",
		Invalids:    "ipfs_invalids",
	}
}
