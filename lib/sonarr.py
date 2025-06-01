#!/usr/bin/env python3

import requests
import asyncio
import aiohttp


class Sonarr:
    def __init__(self, api_key, base_url):
        self.api_key = api_key
        self.base_url = base_url

    def get_ids(self, titles):
        ids_with_titles = []
        sonarr_url = f"{self.base_url}/api/v3/series?apikey={self.api_key}"
        headers = {"User-Agent": "Housekeeper"}

        response = requests.get(sonarr_url, headers=headers)
        if response.status_code != 200:
            raise ValueError("Failed to retrieve Sonarr series.")

        sonarr_model = response.json()
        
        # Convert titles to lowercase for case-insensitive matching
        lowercase_titles = [title.lower() for title in titles]
        
        # Track matched titles for potential duplicate detection
        matched_titles = set()
        unmatched_titles = set(titles)  # Start with all titles as unmatched
        
        for series in sonarr_model:
            series_title = series["title"]
            series_title_lower = series_title.lower()
            
            # Check for exact match (case-insensitive)
            exact_match = False
            for title in titles:
                if title.lower() == series_title_lower:
                    ids_with_titles.append({"id": series["id"], "title": series_title})
                    matched_titles.add(series_title_lower)
                    unmatched_titles.discard(title)  # Remove from unmatched
                    exact_match = True
                    break
                    
            if exact_match:
                continue
                
            # Check for titles containing each other
            for title in list(unmatched_titles):  # Use a list copy to safely modify set during iteration
                title_lower = title.lower()
                if (title_lower in series_title_lower or series_title_lower in title_lower) and title_lower not in matched_titles:
                    print(f"Fuzzy match: Tautulli title '{title}' matched with Sonarr title '{series_title}'")
                    ids_with_titles.append({"id": series["id"], "title": series_title})
                    matched_titles.add(title_lower)
                    unmatched_titles.discard(title)  # Remove from unmatched
                    break

        # Report statistics
        if len(ids_with_titles) < len(titles):
            print(f"Notice: Only {len(ids_with_titles)} of {len(titles)} shows were found in Sonarr")
            if unmatched_titles:
                print("Unmatched show titles from Tautulli:")
                for title in sorted(unmatched_titles):
                    print(f"  - {title}")
            
        return ids_with_titles, list(unmatched_titles)

    def delete_shows(self, items):
        for item in items:
            series_id = item["id"]
            title = item["title"]
            print(f"Deleting series: {title} (ID: {series_id})")

            sonarr_url = f"{self.base_url}/api/v3/series/{series_id}?deleteFiles=true&apikey={self.api_key}"
            headers = {"User-Agent": "Housekeeper"}

            response = requests.delete(sonarr_url, headers=headers)
            if response.status_code != 200:
                raise ValueError(f"Failed to delete series: {title} (ID: {series_id})")
                
    async def delete_show(self, session, item, max_retries=3):
        series_id = item["id"]
        title = item["title"]
        print(f"Deleting series: {title} (ID: {series_id})")
        
        sonarr_url = f"{self.base_url}/api/v3/series/{series_id}?deleteFiles=true&apikey={self.api_key}"
        headers = {"User-Agent": "Housekeeper"}
        
        retries = 0
        while retries < max_retries:
            try:
                async with session.delete(sonarr_url, headers=headers) as response:
                    if response.status == 200:
                        return {"id": series_id, "title": title, "status": "deleted"}
                    else:
                        print(f"Attempt {retries+1}/{max_retries}: Failed to delete series: {title} (ID: {series_id}) - Status: {response.status}")
            except Exception as e:
                print(f"Attempt {retries+1}/{max_retries}: Exception when deleting series: {title} (ID: {series_id}) - {str(e)}")
            
            retries += 1
            if retries < max_retries:
                # Wait longer between each retry attempt
                await asyncio.sleep(2 * retries)
        
        raise ValueError(f"Failed to delete series after {max_retries} attempts: {title} (ID: {series_id})")
        
    
    async def delete_shows_async(self, items):
        print(f"Starting deletion of {len(items)} shows sequentially")
        results = []
        
        async with aiohttp.ClientSession() as session:
            for item in items:
                try:
                    result = await self.delete_show(session, item)
                    results.append(result)
                except Exception as e:
                    print(f"Error deleting show: {str(e)}")
                    results.append(e)
            
            # Check for exceptions and print summary
            successes = sum(1 for r in results if not isinstance(r, Exception))
            failures = sum(1 for r in results if isinstance(r, Exception))
            
            print(f"Shows deletion completed: {successes} succeeded, {failures} failed")
            return results
