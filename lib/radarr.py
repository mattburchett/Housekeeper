#!/usr/bin/env python3

import requests


class Radarr:
    def __init__(self, api_key, base_url):
        self.api_key = api_key
        self.base_url = base_url

    def get_ids(self, titles):
        ids = []
        radarr_url = f"{self.base_url}/api/v3/movie?apikey={self.api_key}"
        headers = {"User-Agent": "Housekeeper"}

        response = requests.get(radarr_url, headers=headers)
        if response.status_code != 200:
            raise ValueError("Failed to retrieve Radarr movies.")

        radarr_model = response.json()

        for movie in radarr_model:
            if movie["title"] in titles:
                ids.append(movie["id"])

        return ids


def delete_movies(self, ids):
    for movie_id in ids:
        print(f"Deleting movie with ID: {movie_id}")

        radarr_url = f"{self.base_url}/api/v3/movie/{movie_id}?deleteFiles=true&apikey={self.api_key}"
        headers = {"User-Agent": "Housekeeper"}

        response = requests.delete(radarr_url, headers=headers)
        if response.status_code != 200:
            raise ValueError(f"Failed to delete movie with ID: {movie_id}")
