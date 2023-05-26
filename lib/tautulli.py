#!/usr/bin/env python3

import requests
import time


class Tautulli:
    def __init__(self, api_key, base_url):
        self.api_key = api_key
        self.base_url = base_url

    def get_count(self, section_id):
        count_url = f"{self.base_url}/api/v2?apikey={self.api_key}&cmd=get_library&section_id={section_id}"
        headers = {"User-Agent": "Housekeeper"}

        response = requests.get(count_url, headers=headers)
        if response.status_code != 200:
            raise ValueError(f"Failed to retrieve count for section ID: {section_id}")

        count_model = response.json()
        count = count_model["response"]["data"]["count"]
        return count

    def get_titles(self, section_id, days, exclude_list):
        count = self.get_count(section_id)

        titles_url = f"{self.base_url}/api/v2?apikey={self.api_key}&cmd=get_library_media_info&section_id={section_id}&order_column=last_played&refresh=true&order_dir=asc&length={count}"
        headers = {"User-Agent": "Housekeeper"}

        response = requests.get(titles_url, headers=headers)
        if response.status_code != 200:
            raise ValueError(f"Failed to retrieve titles for section ID: {section_id}")

        title_model = response.json()
        data = title_model["response"]["data"]["data"]

        titles = []

        epoch = int(time.time()) - (days * 24 * 60 * 60)

        exclude = [ex.strip() for ex in exclude_list.split(",")]

        for item in data:
            exclude_match = False
            for ex in exclude:
                if ex in item["title"]:
                    exclude_match = True
                    break
            if exclude_match:
                continue

            last_played = item.get("last_played")
            if last_played is not None and int(last_played) <= epoch:
                titles.append(item["title"])
            elif last_played is None and int(item["added_at"]) <= epoch:
                titles.append(item["title"])

        titles.sort()

        return titles
