# Housekeeper

Housekeeper is a cleanup tool for Plex environements. 

It requires two things - Plex and PlexPy (now known as Tautulli). It has a option for communicating with Telegram for cleanup purposes.

Please use the example configuration file for configuration.

The command line args that can be passed: 

```Usage of ./housekeeper:
  -c string
        Configuration to load
  -check
        Perform only a check. Do not delete. (default true)
  -days int
        days to poll
  -delete
        Perform the delete task.
  -sectionid int
        pick a section ID```