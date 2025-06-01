#!/usr/bin/env python3

import requests
import asyncio
import aiohttp


class Radarr:
    def __init__(self, api_key, base_url):
        self.api_key = api_key
        self.base_url = base_url

    def get_ids(self, titles):
        ids_with_titles = []
        radarr_url = f"{self.base_url}/api/v3/movie?apikey={self.api_key}"
        headers = {"User-Agent": "Housekeeper"}

        response = requests.get(radarr_url, headers=headers)
        if response.status_code != 200:
            raise ValueError("Failed to retrieve Radarr movies.")

        radarr_model = response.json()
        
        # Convert titles to lowercase for case-insensitive matching
        lowercase_titles = [title.lower() for title in titles]
        
        # Track matched titles for potential duplicate detection
        matched_titles = set()
        unmatched_titles = set(titles)  # Start with all titles as unmatched
        
        for movie in radarr_model:
            movie_title = movie["title"]
            movie_title_lower = movie_title.lower()
            
            # Check for exact match (case-insensitive)
            exact_match = False
            for title in titles:
                if title.lower() == movie_title_lower:
                    ids_with_titles.append({"id": movie["id"], "title": movie_title})
                    matched_titles.add(movie_title_lower)
                    unmatched_titles.discard(title)  # Remove from unmatched
                    exact_match = True
                    break
                    
            if exact_match:
                continue
                
            # Check for titles containing each other
            for title in list(unmatched_titles):  # Use a list copy to safely modify set during iteration
                title_lower = title.lower()
                if (title_lower in movie_title_lower or movie_title_lower in title_lower) and title_lower not in matched_titles:
                    print(f"Fuzzy match: Tautulli title '{title}' matched with Radarr title '{movie_title}'")
                    ids_with_titles.append({"id": movie["id"], "title": movie_title})
                    matched_titles.add(title_lower)
                    unmatched_titles.discard(title)  # Remove from unmatched
                    break

        # Report statistics
        if len(ids_with_titles) < len(titles):
            print(f"Notice: Only {len(ids_with_titles)} of {len(titles)} movies were found in Radarr")
            if unmatched_titles:
                print("Unmatched movie titles from Tautulli:")
                for title in sorted(unmatched_titles):
                    print(f"  - {title}")
            
        return ids_with_titles, list(unmatched_titles)

    def delete_movies(self, items):
        for item in items:
            movie_id = item["id"]
            title = item["title"]
            print(f"Deleting movie: {title} (ID: {movie_id})")

            radarr_url = f"{self.base_url}/api/v3/movie/{movie_id}?deleteFiles=true&apikey={self.api_key}"
            headers = {"User-Agent": "Housekeeper"}

            response = requests.delete(radarr_url, headers=headers)
            if response.status_code != 200:
                raise ValueError(f"Failed to delete movie: {title} (ID: {movie_id})")
                
    async def delete_movie(self, session, item, max_retries=3):
        movie_id = item["id"]
        title = item["title"]
        print(f"Deleting movie: {title} (ID: {movie_id})")
        
        radarr_url = f"{self.base_url}/api/v3/movie/{movie_id}?deleteFiles=true&apikey={self.api_key}"
        headers = {"User-Agent": "Housekeeper"}
        
        retries = 0
        while retries < max_retries:
            try:
                async with session.delete(radarr_url, headers=headers) as response:
                    if response.status == 200:
                        return {"id": movie_id, "title": title, "status": "deleted"}
                    else:
                        print(f"Attempt {retries+1}/{max_retries}: Failed to delete movie: {title} (ID: {movie_id}) - Status: {response.status}")
            except Exception as e:
                print(f"Attempt {retries+1}/{max_retries}: Exception when deleting movie: {title} (ID: {movie_id}) - {str(e)}")
            
            retries += 1
            if retries < max_retries:
                # Wait longer between each retry attempt
                await asyncio.sleep(2 * retries)
        
        raise ValueError(f"Failed to delete movie after {max_retries} attempts: {title} (ID: {movie_id})")
        
    
    async def delete_movies_async(self, items):
        print(f"Starting deletion of {len(items)} movies sequentially")
        results = []
        
        async with aiohttp.ClientSession() as session:
            for item in items:
                try:
                    result = await self.delete_movie(session, item)
                    results.append(result)
                except Exception as e:
                    print(f"Error deleting movie: {str(e)}")
                    results.append(e)
            
            # Check for exceptions and print summary
            successes = sum(1 for r in results if not isinstance(r, Exception))
            failures = sum(1 for r in results if isinstance(r, Exception))
            
            print(f"Movies deletion completed: {successes} succeeded, {failures} failed")
            return results
