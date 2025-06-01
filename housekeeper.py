#!/usr/bin/env python3

import configparser
import asyncio
from lib.sonarr import Sonarr
from lib.radarr import Radarr
from lib.tautulli import Tautulli
from lib.plex import Plex


def read_config():
    config = configparser.ConfigParser()
    config.read("config.ini")
    return config


async def collect_titles(section_id, library_type, titles, items_to_delete):
    """Collect titles from a section and add them to the appropriate list"""
    if library_type == "movie":
        for title in titles:
            items_to_delete["movies"].append(title)
    elif library_type == "show":
        for title in titles:
            items_to_delete["shows"].append(title)
            
    return library_type

async def main_async():
    # Read config
    config = read_config()

    # Initialize Housekeeper
    dry_run = config.getboolean("Housekeeper", "dry_run", fallback=True)
    inactive_days = int(config.get("Housekeeper", "inactive_days"))
    exclude_list = config.get("Housekeeper", "exclude_list", fallback="")

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

    # Dictionary to hold items that should be deleted
    items_to_delete = {
        "movies": [],
        "shows": []
    }
    
    # Collection tasks to gather all titles by type
    collection_tasks = []

    # First pass: collect all titles from all sections
    for section_id in section_ids:
        # Get library type
        library_type = plex.get_library_type(section_id)

        # Get IDs and titles for section.
        titles = tautulli.get_titles(section_id, inactive_days, exclude_list)
        
        # Create task to collect titles
        collection_tasks.append(collect_titles(section_id, library_type, titles, items_to_delete))
    
    # Wait for all collection tasks to complete
    await asyncio.gather(*collection_tasks)
    
    # Calculate total items
    total_movies = len(items_to_delete["movies"])
    total_shows = len(items_to_delete["shows"])
    all_titles = items_to_delete["movies"] + items_to_delete["shows"]
    
    # Sort and print full list if this is a dry_run
    if dry_run:
        all_titles.sort()
        print(f"{len(all_titles)} items would be deleted due to {inactive_days} days of inactivity:")
        for title in all_titles:
            print(title)
    # Process movies and shows in parallel (but each item sequentially within its type)
    else:
        tasks = []
        movie_results = None
        show_results = None
        unmatched_movies = []
        unmatched_shows = []
        
        # Process movies if any
        if total_movies > 0:
            print(f"Found {total_movies} movies to delete")
            movie_ids, unmatched_movies = radarr.get_ids(items_to_delete["movies"])
            if movie_ids:
                tasks.append(radarr.delete_movies_async(movie_ids))
        
        # Process shows if any
        if total_shows > 0:
            print(f"Found {total_shows} shows to delete")
            show_ids, unmatched_shows = sonarr.get_ids(items_to_delete["shows"])
            if show_ids:
                tasks.append(sonarr.delete_shows_async(show_ids))
        
        # Run movie and show deletion concurrently (but items within each type sequentially)
        if tasks:
            print(f"Starting deletion in parallel: {total_movies} movies and {total_shows} shows")
            results = await asyncio.gather(*tasks)
            
            # Print final summary statistics
            print("\n--- FINAL DELETION SUMMARY ---")
            print(f"Total inactive items from Tautulli: {total_movies + total_shows}")
            print(f"  - Movies: {total_movies}")
            print(f"  - Shows: {total_shows}")
            
            movies_found = total_movies - len(unmatched_movies)
            shows_found = total_shows - len(unmatched_shows)
            print(f"\nItems found in *arr applications: {movies_found + shows_found}")
            print(f"  - Movies found in Radarr: {movies_found}")
            print(f"  - Shows found in Sonarr: {shows_found}")
            
            total_unmatched = len(unmatched_movies) + len(unmatched_shows)
            if total_unmatched > 0:
                print(f"\nTotal unmatched items: {total_unmatched}")
                print(f"  - Unmatched movies: {len(unmatched_movies)}")
                print(f"  - Unmatched shows: {len(unmatched_shows)}")
                
                if len(unmatched_movies) > 0 or len(unmatched_shows) > 0:
                    print("\nThese titles might:")
                    print("  - Be spelled differently in Tautulli vs *arr applications")
                    print("  - Exist in Plex but not be managed by Radarr/Sonarr")
                    print("  - Have been renamed in Radarr/Sonarr but not in Plex")
            
            print("\nAll deletion tasks completed")

def main():
    # Run async main and handle dry run printing inside main_async
    asyncio.run(main_async())


if __name__ == "__main__":
    try:
        main()
    except KeyboardInterrupt:
        print("\nOperation cancelled by user")
    except Exception as e:
        print(f"Error: {str(e)}")
