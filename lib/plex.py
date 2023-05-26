#!/usr/bin/env python3

import requests


class Plex:
    def __init__(self, api_key, base_url):
        self.api_key = api_key
        self.base_url = base_url

    def get_library_type(self, section_id):
        url = f"{self.base_url}/library/sections/{section_id}?X-Plex-Token={self.api_key}"
        headers = {"User-Agent": "Housekeeper", "Accept": "application/json"}

        response = requests.get(url, headers=headers)
        if response.status_code != 200:
            raise ValueError("Failed to retrieve library type from the Plex API.")

        library_info = response.json()["MediaContainer"]
        thumb = library_info.get("thumb", "")
        if thumb == "/:/resources/movie.png":
            return "movie"
        elif thumb == "/:/resources/show.png":
            return "show"
        else:
            raise ValueError("Unsupported library type found. This app only supports movies and shows.")
