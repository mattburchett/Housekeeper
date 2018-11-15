# Housekeeper

Housekeeper is a cleanup tool for Plex environments. 

It requires two things - Plex and PlexPy (now known as Tautulli). It has a option for communicating with Telegram for cleanup purposes.

Please use the example configuration file for configuration.

The command line args that can be passed: 

```
Usage of ./housekeeper:
  -c string
        Configuration to load
  -check
        Perform only a check. This will send the message out to Telegram with what can be removed. Does not delete. (default true)
  -days int
        How many days of inactivity to look for on Plex.
  -delete
        Perform the delete task.
  -sectionid int
        Plex Section ID
```