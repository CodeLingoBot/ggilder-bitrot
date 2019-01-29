package main

type ManifestStorage struct {
	Path string
}

type ManifestStorageEntry struct {
	Path      string
	Id        string
	Manifests []*ManifestFileEntry
}

type ManifestFileEntry struct {
	SourcePath string
}

func NewManifestStorage(path string) *ManifestStorage {
	return &ManifestStorage{Path: path}
}

func (m *ManifestStorage) List() []*ManifestStorageEntry {
	return []*ManifestStorageEntry{}
}
