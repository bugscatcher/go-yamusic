package yamusic

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type (
	// PlaylistsService is a service to deal with accounts.
	PlaylistsService struct {
		client *Client
	}

	// PlaylistsList describes get user's playlists response
	PlaylistsList struct {
		InvocationInfo struct {
			Hostname           string `json:"hostname"`
			ReqID              string `json:"req-id"`
			ExecDurationMillis string `json:"exec-duration-millis"`
		} `json:"invocationInfo"`
		Result []struct {
			UID        int           `json:"uid"`
			Kind       int           `json:"kind"`
			DurationMs int           `json:"durationMs"`
			Revision   int           `json:"revision"`
			TrackCount int           `json:"trackCount"`
			Collective bool          `json:"collective"`
			Available  bool          `json:"available"`
			IsBanner   bool          `json:"isBanner"`
			IsPremiere bool          `json:"isPremiere"`
			Title      string        `json:"title"`
			Visibility string        `json:"visibility"`
			OgImage    string        `json:"ogImage"`
			Created    time.Time     `json:"created"`
			Modified   time.Time     `json:"modified"`
			Tags       []interface{} `json:"tags"`
			Owner      struct {
				UID      int    `json:"uid"`
				Login    string `json:"login"`
				Name     string `json:"name"`
				Verified bool   `json:"verified"`
			} `json:"owner"`
			Cover struct {
				Type     string   `json:"type"`
				ItemsURI []string `json:"itemsUri"`
				Custom   bool     `json:"custom"`
			} `json:"cover"`
		} `json:"result"`
	}
	// PlaylistsGet describes get user's playlist by kind response
	PlaylistsGet struct {
		InvocationInfo struct {
			Hostname           string `json:"hostname"`
			ReqID              string `json:"req-id"`
			ExecDurationMillis string `json:"exec-duration-millis"`
		} `json:"invocationInfo"`
		Result struct {
			UID        int           `json:"uid"`
			Kind       int           `json:"kind"`
			Revision   int           `json:"revision"`
			TrackCount int           `json:"trackCount"`
			DurationMs int           `json:"durationMs"`
			LikesCount int           `json:"likesCount"`
			Collective bool          `json:"collective"`
			Available  bool          `json:"available"`
			IsBanner   bool          `json:"isBanner"`
			IsPremiere bool          `json:"isPremiere"`
			Title      string        `json:"title"`
			OgImage    string        `json:"ogImage"`
			Visibility string        `json:"visibility"`
			Created    time.Time     `json:"created"`
			Modified   time.Time     `json:"modified"`
			Tags       []interface{} `json:"tags"`
			Owner      struct {
				UID      int    `json:"uid"`
				Login    string `json:"login"`
				Name     string `json:"name"`
				Verified bool   `json:"verified"`
			} `json:"owner"`
			Cover struct {
				Type    string `json:"type"`
				Dir     string `json:"dir"`
				Version string `json:"version"`
				URI     string `json:"uri"`
				Custom  bool   `json:"custom"`
			} `json:"cover"`
			Tracks []struct {
				ID        int       `json:"id"`
				Timestamp time.Time `json:"timestamp"`
				Recent    bool      `json:"recent"`
				Track     struct {
					DurationMs               int    `json:"durationMs"`
					FileSize                 int    `json:"fileSize"`
					Available                bool   `json:"available"`
					AvailableForPremiumUsers bool   `json:"availableForPremiumUsers"`
					LyricsAvailable          bool   `json:"lyricsAvailable"`
					ID                       string `json:"id"`
					RealID                   string `json:"realId"`
					Title                    string `json:"title"`
					StorageDir               string `json:"storageDir"`
					CoverURI                 string `json:"coverUri"`
					OgImage                  string `json:"ogImage"`
					Major                    struct {
						ID   int    `json:"id"`
						Name string `json:"name"`
					} `json:"major"`
					Normalization struct {
						Gain float64 `json:"gain"`
						Peak int     `json:"peak"`
					} `json:"normalization"`
					Artists []struct {
						ID       int    `json:"id"`
						Name     string `json:"name"`
						Various  bool   `json:"various"`
						Composer bool   `json:"composer"`
						Cover    struct {
							Type   string `json:"type"`
							Prefix string `json:"prefix"`
							URI    string `json:"uri"`
						} `json:"cover"`
						Genres []interface{} `json:"genres"`
					} `json:"artists"`
					Albums []struct {
						ID                       int           `json:"id"`
						Title                    string        `json:"title"`
						Year                     int           `json:"year"`
						ReleaseDate              time.Time     `json:"releaseDate"`
						CoverURI                 string        `json:"coverUri"`
						OgImage                  string        `json:"ogImage"`
						Genre                    string        `json:"genre"`
						Buy                      []interface{} `json:"buy"`
						TrackCount               int           `json:"trackCount"`
						Recent                   bool          `json:"recent"`
						VeryImportant            bool          `json:"veryImportant"`
						Available                bool          `json:"available"`
						AvailableForPremiumUsers bool          `json:"availableForPremiumUsers"`
						AvailableForMobile       bool          `json:"availableForMobile"`
						AvailablePartially       bool          `json:"availablePartially"`
						Bests                    []int         `json:"bests"`
						Artists                  []struct {
							ID       int    `json:"id"`
							Name     string `json:"name"`
							Various  bool   `json:"various"`
							Composer bool   `json:"composer"`
							Cover    struct {
								Type   string `json:"type"`
								Prefix string `json:"prefix"`
								URI    string `json:"uri"`
							} `json:"cover"`
							Genres []interface{} `json:"genres"`
						} `json:"artists"`
						Labels []struct {
							ID          int    `json:"id"`
							Name        string `json:"name"`
							Description string `json:"description"`
							Image       string `json:"image"`
						} `json:"labels"`
						TrackPosition struct {
							Volume int `json:"volume"`
							Index  int `json:"index"`
						} `json:"trackPosition"`
					} `json:"albums"`
				} `json:"track"`
			} `json:"tracks"`
		} `json:"result"`
	}
	// PlaylistsGetByKinds describes get user's playlists by kinds response
	PlaylistsGetByKinds struct {
		InvocationInfo struct {
			Hostname           string `json:"hostname"`
			ReqID              string `json:"req-id"`
			ExecDurationMillis string `json:"exec-duration-millis"`
		} `json:"invocationInfo"`
		Result []struct {
			UID        int           `json:"uid"`
			Kind       int           `json:"kind"`
			Revision   int           `json:"revision"`
			TrackCount int           `json:"trackCount"`
			DurationMs int           `json:"durationMs"`
			Collective bool          `json:"collective"`
			Available  bool          `json:"available"`
			IsBanner   bool          `json:"isBanner"`
			IsPremiere bool          `json:"isPremiere"`
			Visibility string        `json:"visibility"`
			Title      string        `json:"title"`
			OgImage    string        `json:"ogImage"`
			Created    time.Time     `json:"created"`
			Modified   time.Time     `json:"modified"`
			Tags       []interface{} `json:"tags"`
			Owner      struct {
				UID      int    `json:"uid"`
				Login    string `json:"login"`
				Name     string `json:"name"`
				Verified bool   `json:"verified"`
			} `json:"owner"`
			Cover struct {
				Type    string `json:"type"`
				Dir     string `json:"dir"`
				Version string `json:"version"`
				URI     string `json:"uri"`
				Custom  bool   `json:"custom"`
			} `json:"cover"`
			Tracks []struct {
				ID        int       `json:"id"`
				AlbumID   int       `json:"albumId"`
				Timestamp time.Time `json:"timestamp"`
			} `json:"tracks"`
		} `json:"result"`
	}

	// PlaylistsRename describes method rename playlist response
	PlaylistsRename struct {
		InvocationInfo struct {
			Hostname           string `json:"hostname"`
			ReqID              string `json:"req-id"`
			ExecDurationMillis string `json:"exec-duration-millis"`
		} `json:"invocationInfo"`
		Result struct {
			UID        int       `json:"uid"`
			Kind       int       `json:"kind"`
			Revision   int       `json:"revision"`
			TrackCount int       `json:"trackCount"`
			DurationMs int       `json:"durationMs"`
			Collective bool      `json:"collective"`
			Available  bool      `json:"available"`
			IsBanner   bool      `json:"isBanner"`
			IsPremiere bool      `json:"isPremiere"`
			OgImage    string    `json:"ogImage"`
			Title      string    `json:"title"`
			Visibility string    `json:"visibility"`
			Created    time.Time `json:"created"`
			Modified   time.Time `json:"modified"`
			Owner      struct {
				UID      int    `json:"uid"`
				Login    string `json:"login"`
				Name     string `json:"name"`
				Verified bool   `json:"verified"`
			} `json:"owner"`
			Cover struct {
				Error string `json:"error"`
			} `json:"cover"`
		} `json:"result"`
	}
	// PlaylistsCreate describes method create playlist response
	PlaylistsCreate struct {
		InvocationInfo struct {
			Hostname           string `json:"hostname"`
			ReqID              string `json:"req-id"`
			ExecDurationMillis string `json:"exec-duration-millis"`
		} `json:"invocationInfo"`
		Result struct {
			UID        int           `json:"uid"`
			Kind       int           `json:"kind"`
			DurationMs int           `json:"durationMs"`
			Revision   int           `json:"revision"`
			TrackCount int           `json:"trackCount"`
			Collective bool          `json:"collective"`
			Available  bool          `json:"available"`
			IsBanner   bool          `json:"isBanner"`
			IsPremiere bool          `json:"isPremiere"`
			Title      string        `json:"title"`
			Visibility string        `json:"visibility"`
			OgImage    string        `json:"ogImage"`
			Tags       []interface{} `json:"tags"`
			Created    time.Time     `json:"created"`
			Modified   time.Time     `json:"modified"`
			Owner      struct {
				UID      int    `json:"uid"`
				Login    string `json:"login"`
				Name     string `json:"name"`
				Verified bool   `json:"verified"`
			} `json:"owner"`
			Cover struct {
				Error string `json:"error"`
			} `json:"cover"`
		} `json:"result"`
	}
	// PlaylistsDelete describes method delete playlist response
	PlaylistsDelete struct {
		InvocationInfo struct {
			Hostname           string `json:"hostname"`
			ReqID              string `json:"req-id"`
			ExecDurationMillis string `json:"exec-duration-millis"`
		} `json:"invocationInfo"`
		Result string `json:"result"`
	}
)

// List returns playlists of the user
func (s *PlaylistsService) List(
	ctx context.Context,
	userID int,
) (*PlaylistsList, *http.Response, error) {

	uri := fmt.Sprintf("users/%v/playlists/list", userID)
	req, err := s.client.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	playlists := new(PlaylistsList)
	resp, err := s.client.Do(ctx, req, playlists)
	return playlists, resp, err
}

// Get returns playlist of the user by kind
func (s *PlaylistsService) Get(
	ctx context.Context,
	userID int,
	kind int,
) (*PlaylistsGet, *http.Response, error) {

	uri := fmt.Sprintf("users/%v/playlists/%v", userID, kind)
	req, err := s.client.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	playlist := new(PlaylistsGet)
	resp, err := s.client.Do(ctx, req, playlist)
	return playlist, resp, err
}

type (
	// PlaylistsGetByKindOptions options for GetByKinds method
	PlaylistsGetByKindOptions struct {
		Kinds      []int
		Mixed      bool
		RichTracks bool
	}
)

// GetByKinds returns several playlists by kinds with track ids
func (s *PlaylistsService) GetByKinds(
	ctx context.Context,
	userID int,
	opts *PlaylistsGetByKindOptions,
) (*PlaylistsGetByKinds, *http.Response, error) {

	if opts == nil {
		opts = &PlaylistsGetByKindOptions{}
	}

	queryParams := url.Values{}
	queryParams.Set("kinds", func() string {
		stringKinds := make([]string, 0, len(opts.Kinds))
		for _, kind := range opts.Kinds {
			stringKinds = append(stringKinds, strconv.Itoa(kind))
		}
		return strings.Join(stringKinds, ",")
	}())
	queryParams.Set("mixed", strconv.FormatBool(opts.Mixed))
	queryParams.Set("rich-tracks", strconv.FormatBool(opts.RichTracks))

	uri := fmt.Sprintf("users/%v/playlists?%v", userID, queryParams.Encode())
	req, err := s.client.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	playlists := new(PlaylistsGetByKinds)
	resp, err := s.client.Do(ctx, req, playlists)
	return playlists, resp, err
}

// Rename renames playlist of current user
func (s *PlaylistsService) Rename(
	ctx context.Context,
	kind int,
	newName string,
) (*PlaylistsRename, *http.Response, error) {

	uri := fmt.Sprintf("users/%v/playlists/%v/name", s.client.userID, kind)

	form := url.Values{}
	form.Set("value", newName)

	req, err := s.client.NewRequest(http.MethodPost, uri, form)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	renamedPlaylist := new(PlaylistsRename)
	resp, err := s.client.Do(ctx, req, renamedPlaylist)
	return renamedPlaylist, resp, err
}

// Create creates playlist for current user
func (s *PlaylistsService) Create(
	ctx context.Context,
	title string,
	isPublic bool,
) (*PlaylistsCreate, *http.Response, error) {

	var visibility string
	if isPublic {
		visibility = "public"
	} else {
		visibility = "private"
	}

	form := url.Values{}
	form.Set("title", title)
	form.Set("visibility", visibility)

	uri := fmt.Sprintf("users/%v/playlists/create", s.client.userID)

	req, err := s.client.NewRequest(http.MethodPost, uri, form)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	createdPlaylist := new(PlaylistsCreate)
	resp, err := s.client.Do(ctx, req, createdPlaylist)
	return createdPlaylist, resp, err
}

// Delete deletes playlist for current user
func (s *PlaylistsService) Delete(
	ctx context.Context,
	kind int,
) (*PlaylistsDelete, *http.Response, error) {

	uri := fmt.Sprintf("users/%v/playlists/%v/delete", s.client.userID, kind)
	req, err := s.client.NewRequest(http.MethodPost, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	deletedPlaylist := new(PlaylistsDelete)
	resp, err := s.client.Do(ctx, req, deletedPlaylist)
	return deletedPlaylist, resp, err
}