package model

import "encoding/xml"

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
