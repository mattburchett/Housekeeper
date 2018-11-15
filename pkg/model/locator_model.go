package model

import "encoding/xml"

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

// XMLPlexAPI - This is the XML version of the struct below it.
type XMLPlexAPI struct {
	XMLName             xml.Name `xml:"MediaContainer"`
	Text                string   `xml:",chardata"`
	Size                string   `xml:"size,attr"`
	AllowSync           string   `xml:"allowSync,attr"`
	Identifier          string   `xml:"identifier,attr"`
	LibrarySectionID    string   `xml:"librarySectionID,attr"`
	LibrarySectionTitle string   `xml:"librarySectionTitle,attr"`
	LibrarySectionUUID  string   `xml:"librarySectionUUID,attr"`
	MediaTagPrefix      string   `xml:"mediaTagPrefix,attr"`
	MediaTagVersion     string   `xml:"mediaTagVersion,attr"`
	Video               struct {
		Text                  string `xml:",chardata"`
		RatingKey             string `xml:"ratingKey,attr"`
		Key                   string `xml:"key,attr"`
		GUID                  string `xml:"guid,attr"`
		LibrarySectionTitle   string `xml:"librarySectionTitle,attr"`
		LibrarySectionID      string `xml:"librarySectionID,attr"`
		LibrarySectionKey     string `xml:"librarySectionKey,attr"`
		Studio                string `xml:"studio,attr"`
		Type                  string `xml:"type,attr"`
		Title                 string `xml:"title,attr"`
		ContentRating         string `xml:"contentRating,attr"`
		Summary               string `xml:"summary,attr"`
		Rating                string `xml:"rating,attr"`
		AudienceRating        string `xml:"audienceRating,attr"`
		Year                  string `xml:"year,attr"`
		Tagline               string `xml:"tagline,attr"`
		Thumb                 string `xml:"thumb,attr"`
		Art                   string `xml:"art,attr"`
		Duration              string `xml:"duration,attr"`
		OriginallyAvailableAt string `xml:"originallyAvailableAt,attr"`
		AddedAt               string `xml:"addedAt,attr"`
		UpdatedAt             string `xml:"updatedAt,attr"`
		AudienceRatingImage   string `xml:"audienceRatingImage,attr"`
		ChapterSource         string `xml:"chapterSource,attr"`
		PrimaryExtraKey       string `xml:"primaryExtraKey,attr"`
		RatingImage           string `xml:"ratingImage,attr"`
		Media                 struct {
			Text            string `xml:",chardata"`
			VideoResolution string `xml:"videoResolution,attr"`
			ID              string `xml:"id,attr"`
			Duration        string `xml:"duration,attr"`
			Bitrate         string `xml:"bitrate,attr"`
			Width           string `xml:"width,attr"`
			Height          string `xml:"height,attr"`
			AspectRatio     string `xml:"aspectRatio,attr"`
			AudioChannels   string `xml:"audioChannels,attr"`
			AudioCodec      string `xml:"audioCodec,attr"`
			VideoCodec      string `xml:"videoCodec,attr"`
			Container       string `xml:"container,attr"`
			VideoFrameRate  string `xml:"videoFrameRate,attr"`
			AudioProfile    string `xml:"audioProfile,attr"`
			VideoProfile    string `xml:"videoProfile,attr"`
			Part            struct {
				Text         string `xml:",chardata"`
				ID           string `xml:"id,attr"`
				Key          string `xml:"key,attr"`
				Duration     string `xml:"duration,attr"`
				File         string `xml:"file,attr"`
				Size         string `xml:"size,attr"`
				AudioProfile string `xml:"audioProfile,attr"`
				Container    string `xml:"container,attr"`
				VideoProfile string `xml:"videoProfile,attr"`
				Stream       []struct {
					Text               string `xml:",chardata"`
					ID                 string `xml:"id,attr"`
					StreamType         string `xml:"streamType,attr"`
					Default            string `xml:"default,attr"`
					Codec              string `xml:"codec,attr"`
					Index              string `xml:"index,attr"`
					Bitrate            string `xml:"bitrate,attr"`
					Language           string `xml:"language,attr"`
					LanguageCode       string `xml:"languageCode,attr"`
					BitDepth           string `xml:"bitDepth,attr"`
					ChromaLocation     string `xml:"chromaLocation,attr"`
					ChromaSubsampling  string `xml:"chromaSubsampling,attr"`
					FrameRate          string `xml:"frameRate,attr"`
					HasScalingMatrix   string `xml:"hasScalingMatrix,attr"`
					Height             string `xml:"height,attr"`
					Level              string `xml:"level,attr"`
					Profile            string `xml:"profile,attr"`
					RefFrames          string `xml:"refFrames,attr"`
					ScanType           string `xml:"scanType,attr"`
					Title              string `xml:"title,attr"`
					Width              string `xml:"width,attr"`
					DisplayTitle       string `xml:"displayTitle,attr"`
					Selected           string `xml:"selected,attr"`
					Channels           string `xml:"channels,attr"`
					AudioChannelLayout string `xml:"audioChannelLayout,attr"`
					SamplingRate       string `xml:"samplingRate,attr"`
					Key                string `xml:"key,attr"`
				} `xml:"Stream"`
			} `xml:"Part"`
		} `xml:"Media"`
		Genre []struct {
			Text   string `xml:",chardata"`
			ID     string `xml:"id,attr"`
			Filter string `xml:"filter,attr"`
			Tag    string `xml:"tag,attr"`
		} `xml:"Genre"`
		Director struct {
			Text   string `xml:",chardata"`
			ID     string `xml:"id,attr"`
			Filter string `xml:"filter,attr"`
			Tag    string `xml:"tag,attr"`
		} `xml:"Director"`
		Writer []struct {
			Text   string `xml:",chardata"`
			ID     string `xml:"id,attr"`
			Filter string `xml:"filter,attr"`
			Tag    string `xml:"tag,attr"`
		} `xml:"Writer"`
		Producer []struct {
			Text   string `xml:",chardata"`
			ID     string `xml:"id,attr"`
			Filter string `xml:"filter,attr"`
			Tag    string `xml:"tag,attr"`
		} `xml:"Producer"`
		Country struct {
			Text   string `xml:",chardata"`
			ID     string `xml:"id,attr"`
			Filter string `xml:"filter,attr"`
			Tag    string `xml:"tag,attr"`
		} `xml:"Country"`
		Role []struct {
			Text   string `xml:",chardata"`
			ID     string `xml:"id,attr"`
			Filter string `xml:"filter,attr"`
			Tag    string `xml:"tag,attr"`
			Role   string `xml:"role,attr"`
			Thumb  string `xml:"thumb,attr"`
		} `xml:"Role"`
	} `xml:"Video"`
}
