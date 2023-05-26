#!/usr/bin/env python3

import configparser
from lib.sonarr import Sonarr
from lib.radarr import Radarr
from lib.tautulli import Tautulli
from lib.plex import Plex


def read_config():
    config = configparser.ConfigParser()
    config.read("config.ini")
    return config


def main():

    # Read config
    config = read_config()

    # Initialize Housekeeper
    dry_run = config.getboolean("Housekeeper", "dry_run", fallback=True)
    inactive_days = int(config.get("Housekeeper", "inactive_days"))
    exclude_list = config.get("Housekeeper", "exclude_list")

    # Initialize Sonarr
    sonarr_api_key = config.get("Sonarr", "api_key")
    sonarr_base_url = config.get("Sonarr", "base_url")
    sonarr = Sonarr(sonarr_api_key, sonarr_base_url)

    # Initialize Radarr
    radarr_api_key = config.get("Radarr", "api_key")
    radarr_base_url = config.get("Radarr", "base_url")
    radarr = Radarr(radarr_api_key, radarr_base_url)

    # Initialize Tautulli
    tautulli_api_key = config.get("Tautulli", "api_key")
    tautulli_base_url = config.get("Tautulli", "base_url")
    tautulli = Tautulli(tautulli_api_key, tautulli_base_url)

    # Initialize Plex
    plex_api_key = config.get("Plex", "token")
    plex_base_url = config.get("Plex", "base_url")
    plex = Plex(plex_api_key, plex_base_url)

    # Get section IDs
    section_ids = config.get("Plex", "section_ids").split(",")

    # Generate an empty list to sort movies and show titles.
    titles_full_list = []

    for section_id in section_ids:
        # Get library type
        library_type = plex.get_library_type(section_id)

        # Get IDs and titles for section.
        titles = tautulli.get_titles(section_id, inactive_days, exclude_list)

        # Add titles to full list.
        for title in titles:
            titles_full_list.append(title)

        # If dry run, continue
        if dry_run:
            continue

        if library_type == "movie":
            # Get Radarr IDs for titles.
            radarr_ids = radarr.get_ids(titles)
            # Delete movies from Radarr.
            radarr.delete_movies(radarr_ids)
        elif library_type == "show":
            # Get Sonarr IDs for titles.
            sonarr_ids = sonarr.get_ids(titles)
            # Delete shows from Sonarr.
            sonarr.delete_shows(sonarr_ids)

    # Sort and print full list if this is a dry_run
    if dry_run:
        titles_full_list.sort()
        print(f"{len(titles_full_list)} items would be deleted due to {inactive_days} days of inactivity:")
        for title in titles_full_list:
            print(title)


if __name__ == "__main__":
    main()
