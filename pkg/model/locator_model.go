package model

import (
	"encoding/xml"
	"time"
)

// PlexPyLibraryInfo will gather all the library related info. We really just need the count from this...
type PlexPyLibraryInfo struct {
	Response struct {
		Message interface{} `json:"message"`
		Data    struct {
			Count           int         `json:"count"`
			SectionID       int         `json:"section_id"`
			SectionName     string      `json:"section_name"`
			LibraryArt      string      `json:"library_art"`
			ParentCount     interface{} `json:"parent_count"`
			SectionType     string      `json:"section_type"`
			DoNotifyCreated int         `json:"do_notify_created"`
			KeepHistory     int         `json:"keep_history"`
			ChildCount      interface{} `json:"child_count"`
			LibraryThumb    string      `json:"library_thumb"`
			DoNotify        int         `json:"do_notify"`
		} `json:"data"`
		Result string `json:"result"`
	} `json:"response"`
}

// PlexPyMediaInfo - This is the information for the Media Library and related media.
type PlexPyMediaInfo struct {
	Response struct {
		Message interface{} `json:"message"`
		Data    struct {
			Draw             int    `json:"draw"`
			RecordsTotal     string `json:"recordsTotal"`
			TotalFileSize    int64  `json:"total_file_size"`
			RecordsFiltered  int    `json:"recordsFiltered"`
			FilteredFileSize int64  `json:"filtered_file_size"`
			Data             []struct {
				Year                 string `json:"year"`
				SortTitle            string `json:"sort_title"`
				ParentRatingKey      string `json:"parent_rating_key"`
				AudioCodec           string `json:"audio_codec"`
				FileSize             string `json:"file_size"`
				RatingKey            string `json:"rating_key"`
				Container            string `json:"container"`
				Thumb                string `json:"thumb"`
				Title                string `json:"title"`
				SectionType          string `json:"section_type"`
				MediaType            string `json:"media_type"`
				VideoResolution      string `json:"video_resolution"`
				GrandparentRatingKey string `json:"grandparent_rating_key"`
				AudioChannels        string `json:"audio_channels"`
				LastPlayed           int64  `json:"last_played,omitempty"`
				SectionID            int    `json:"section_id"`
				PlayCount            int    `json:"play_count"`
				Bitrate              string `json:"bitrate"`
				VideoFramerate       string `json:"video_framerate"`
				MediaIndex           string `json:"media_index"`
				AddedAt              int64  `json:"added_at,string"`
				VideoCodec           string `json:"video_codec"`
				ParentMediaIndex     string `json:"parent_media_index"`
			} `json:"data"`
		} `json:"data"`
		Result string `json:"result"`
	} `json:"response"`
}

type XMLPlexLibraryType struct {
	XMLName          xml.Name `xml:"MediaContainer"`
	Text             string   `xml:",chardata"`
	Size             string   `xml:"size,attr"`
	AllowSync        string   `xml:"allowSync,attr"`
	Art              string   `xml:"art,attr"`
	Content          string   `xml:"content,attr"`
	Identifier       string   `xml:"identifier,attr"`
	LibrarySectionID string   `xml:"librarySectionID,attr"`
	MediaTagPrefix   string   `xml:"mediaTagPrefix,attr"`
	MediaTagVersion  string   `xml:"mediaTagVersion,attr"`
	Nocache          string   `xml:"nocache,attr"`
	Thumb            string   `xml:"thumb,attr"`
	Title1           string   `xml:"title1,attr"`
	ViewGroup        string   `xml:"viewGroup,attr"`
	ViewMode         string   `xml:"viewMode,attr"`
	Directory        []struct {
		Text      string `xml:",chardata"`
		Key       string `xml:"key,attr"`
		Title     string `xml:"title,attr"`
		Secondary string `xml:"secondary,attr"`
		Prompt    string `xml:"prompt,attr"`
		Search    string `xml:"search,attr"`
	} `xml:"Directory"`
}

// SonarrSeries type takes all the data from Sonarr and places it in a struct
type SonarrSeries []struct {
	Title           string `json:"title"`
	AlternateTitles []struct {
		Title        string `json:"title"`
		SeasonNumber int    `json:"seasonNumber"`
	} `json:"alternateTitles"`
	SortTitle         string    `json:"sortTitle"`
	SeasonCount       int       `json:"seasonCount"`
	TotalEpisodeCount int       `json:"totalEpisodeCount"`
	EpisodeCount      int       `json:"episodeCount"`
	EpisodeFileCount  int       `json:"episodeFileCount"`
	SizeOnDisk        int64     `json:"sizeOnDisk"`
	Status            string    `json:"status"`
	Overview          string    `json:"overview"`
	PreviousAiring    time.Time `json:"previousAiring"`
	Network           string    `json:"network"`
	AirTime           string    `json:"airTime,omitempty"`
	Images            []struct {
		CoverType string `json:"coverType"`
		URL       string `json:"url"`
	} `json:"images"`
	Seasons []struct {
		SeasonNumber int  `json:"seasonNumber"`
		Monitored    bool `json:"monitored"`
		Statistics   struct {
			PreviousAiring    time.Time `json:"previousAiring"`
			EpisodeFileCount  int       `json:"episodeFileCount"`
			EpisodeCount      int       `json:"episodeCount"`
			TotalEpisodeCount int       `json:"totalEpisodeCount"`
			SizeOnDisk        int64     `json:"sizeOnDisk"`
			PercentOfEpisodes float64   `json:"percentOfEpisodes"`
		} `json:"statistics"`
	} `json:"seasons"`
	Year              int           `json:"year"`
	Path              string        `json:"path"`
	ProfileID         int           `json:"profileId"`
	SeasonFolder      bool          `json:"seasonFolder"`
	Monitored         bool          `json:"monitored"`
	UseSceneNumbering bool          `json:"useSceneNumbering"`
	Runtime           int           `json:"runtime"`
	TvdbID            int           `json:"tvdbId"`
	TvRageID          int           `json:"tvRageId"`
	TvMazeID          int           `json:"tvMazeId"`
	FirstAired        time.Time     `json:"firstAired"`
	LastInfoSync      time.Time     `json:"lastInfoSync"`
	SeriesType        string        `json:"seriesType"`
	CleanTitle        string        `json:"cleanTitle"`
	ImdbID            string        `json:"imdbId,omitempty"`
	TitleSlug         string        `json:"titleSlug"`
	Certification     string        `json:"certification,omitempty"`
	Genres            []string      `json:"genres"`
	Tags              []interface{} `json:"tags"`
	Added             time.Time     `json:"added"`
	Ratings           struct {
		Votes int     `json:"votes"`
		Value float64 `json:"value"`
	} `json:"ratings"`
	QualityProfileID int       `json:"qualityProfileId"`
	ID               int       `json:"id"`
	NextAiring       time.Time `json:"nextAiring,omitempty"`
}
