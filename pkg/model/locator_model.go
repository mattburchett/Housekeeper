package model

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
				LastPlayed           int64  `json:"last_played"`
				SectionID            int    `json:"section_id"`
				PlayCount            int    `json:"play_count"`
				Bitrate              string `json:"bitrate"`
				VideoFramerate       string `json:"video_framerate"`
				MediaIndex           string `json:"media_index"`
				AddedAt              string `json:"added_at"`
				VideoCodec           string `json:"video_codec"`
				ParentMediaIndex     string `json:"parent_media_index"`
			} `json:"data"`
		} `json:"data"`
		Result string `json:"result"`
	} `json:"response"`
}
