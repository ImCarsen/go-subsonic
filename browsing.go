package subsonic

func (s *SubsonicClient) GetMusicFolders() ([]*MusicFolder, error) {
	resp, err := s.Get("getMusicFolders", nil)
	if err != nil {
		return nil, err
	}
	return resp.MusicFolders.Folders, nil
}

// GetIndexes returns the index of entries by letter/number.
// Optional Parameters:
// * musicFolderId:    Only return songs in the music folder with the given ID. See getMusicFolders.
// * ifModifiedSince:  If specified, only return a result if the artist collection has changed since the given time (in milliseconds since 1 Jan 1970).
func (s *SubsonicClient) GetIndexes(parameters map[string]string) (*IndexContainer, error) {
	resp, err := s.Get("getIndexes", parameters)
	if err != nil {
		return nil, err
	}
	return resp.Indexes, nil
}

// GetMusicDirectory returns the context around an object ID from the database.
func (s *SubsonicClient) GetMusicDirectory(id string) (*Directory, error) {
	resp, err := s.Get("getMusicDirectory", map[string]string{"id": id})
	if err != nil {
		return nil, err
	}
	return resp.Directory, nil
}

// GetGenres returns all genres in the server.
func (s *SubsonicClient) GetGenres() ([]*Genre, error) {
	resp, err := s.Get("getGenres", nil)
	if err != nil {
		return nil, err
	}
	return resp.Genres.Genre, nil
}

// GetArtists returns all artists in the server.
// Optional Parameters:
// * musicFolderId:  Only return songs in the music folder with the given ID. See getMusicFolders.
func (s *SubsonicClient) GetArtists(parameters map[string]string) (*ArtistsContainer, error) {
	resp, err := s.Get("getArtists", parameters)
	if err != nil {
		return nil, err
	}
	return resp.Artists, nil
}

// GetAlbum returns an Artist by ID.
func (s *SubsonicClient) GetArtist(id string) (*Artist, error) {
	resp, err := s.Get("getArtist", map[string]string{"id": id})
	if err != nil {
		return nil, err
	}
	return resp.Artist, nil
}

// GetAlbum returns an Album by ID.
func (s *SubsonicClient) GetAlbum(id string) (*Album, error) {
	resp, err := s.Get("getAlbum", map[string]string{"id": id})
	if err != nil {
		return nil, err
	}
	return resp.Album, nil
}

// GetSong returns a Song by ID.
func (s *SubsonicClient) GetSong(id string) (*Song, error) {
	resp, err := s.Get("getSong", map[string]string{"id": id})
	if err != nil {
		return nil, err
	}
	return resp.Song, nil
}

// GetArtistInfo returns biography, image links, and similar artists from last.fm.
// Optional Parameters:
// * count:             Max number of similar artists to return.
// * includeNotPresent: Whether to return artists that are not present in the media library.
func (s *SubsonicClient) GetArtistInfo(id string, parameters map[string]string) (*ArtistInfo, error) {
	params := make(map[string]string)
	params["id"] = id
	for k, v := range parameters {
		params[k] = v
	}
	resp, err := s.Get("getArtistInfo", params)
	if err != nil {
		return nil, err
	}
	return resp.ArtistInfo, nil
}

// GetArtistInfo2 returns biography, image links, and similar artists like GetArtistInfo, but using id3 tags.
// Optional Parameters:
// * count:             Max number of similar artists to return.
// * includeNotPresent: Whether to return artists that are not present in the media library.
func (s *SubsonicClient) GetArtistInfo2(id string, parameters map[string]string) (*ArtistInfo, error) {
	params := make(map[string]string)
	params["id"] = id
	for k, v := range parameters {
		params[k] = v
	}
	resp, err := s.Get("getArtistInfo2", params)
	if err != nil {
		return nil, err
	}
	return resp.ArtistInfo2, nil
}

// GetAlbumInfo returns album notes, image data, etc using data from last.fm.
// This accepts both album and song IDs.
func (s *SubsonicClient) GetAlbumInfo(id string) (*AlbumInfo, error) {
	resp, err := s.Get("getAlbumInfo", map[string]string{"id": id})
	if err != nil {
		return nil, err
	}
	return resp.AlbumInfo, nil
}

// GetAlbumInfo2 returns the same data as GetAlbumInfo, but organized by id3 tag.
// It only accepts album IDs.
func (s *SubsonicClient) GetAlbumInfo2(id string) (*AlbumInfo, error) {
	resp, err := s.Get("getAlbumInfo2", map[string]string{"id": id})
	if err != nil {
		return nil, err
	}
	return resp.AlbumInfo, nil
}

// GetSimilarSongs finds similar songs to an album, track, or artist.
// This is mostly used for radio features. This accepts artist, album, or song IDs.
// Optional Parameters:
// * count: Number of songs to return
func (s *SubsonicClient) GetSimilarSongs(id string, parameters map[string]string) ([]*Song, error) {
	params := make(map[string]string)
	params["id"] = id
	for k, v := range parameters {
		params[k] = v
	}
	resp, err := s.Get("getSimilarSongs", params)
	if err != nil {
		return nil, err
	}
	return resp.SimilarSongs.Songs, nil
}

// GetSimilarSongs2 finds similar songs like GetSimilarSongs, but using id3 tags.
// Optional Parameters:
// * count: Number of songs to return
func (s *SubsonicClient) GetSimilarSongs2(id string, parameters map[string]string) ([]*Song, error) {
	params := make(map[string]string)
	params["id"] = id
	for k, v := range parameters {
		params[k] = v
	}
	resp, err := s.Get("getSimilarSongs2", params)
	if err != nil {
		return nil, err
	}
	return resp.SimilarSongs2.Songs, nil
}

// GetTopSongs returns the top songs for a given artist by name.
// Optional Parameters:
// * count: Number of songs to return
func (s *SubsonicClient) GetTopSongs(name string, parameters map[string]string) ([]*Song, error) {
	params := make(map[string]string)
	params["artist"] = name
	for k, v := range parameters {
		params[k] = v
	}
	resp, err := s.Get("getTopSongs", params)
	if err != nil {
		return nil, err
	}
	return resp.TopSongs.Songs, nil
}
