#!/usr/bin/env python3

import requests


class Sonarr:
    def __init__(self, api_key, base_url):
        self.api_key = api_key
        self.base_url = base_url

    def get_ids(self, titles):
        ids = []
        sonarr_url = f"{self.base_url}/api/series?apikey={self.api_key}"
        headers = {"User-Agent": "Housekeeper"}

        response = requests.get(sonarr_url, headers=headers)
        if response.status_code != 200:
            raise ValueError("Failed to retrieve Sonarr series.")

        sonarr_model = response.json()

        for series in sonarr_model:
            if series["title"] in titles:
                ids.append(series["id"])

        return ids

    def delete_shows(self, ids):
        for series_id in ids:
            print(f"Deleting series with ID: {series_id}")

            sonarr_url = f"{self.base_url}/api/series/{series_id}?deleteFiles=true&apikey={self.api_key}"
            headers = {"User-Agent": "Housekeeper"}

            response = requests.delete(sonarr_url, headers=headers)
            if response.status_code != 200:
                raise ValueError(f"Failed to delete series with ID: {series_id}")
