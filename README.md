# ForeFlight Track View

This project aims to be a one-off Go CLI tool that shrinks and cleans ForeFlight tracklog (`.kml`) files for interactive use in apps like Google Earth. The goal is to make the collection and use of the files as easy as possible, including adding to the collection as more tracklogs become available.


## Project Overview
The standard ForeFlight tracklog file includes a lot of "noisy" data that bloats without contributing in a meaningful way to a basic visual representation of the tracks. This includes things like:

- Timestamps for every coordinate captured
- Pitch/bank
- Course
- Speed
- Horizontal/vertical acceleration
- Start/end placemarks

At a minimum, this software strips away everything except altitude. This can result in significantly shrinking the files: a sample file from my testing went from 1.9MB to 488KB - **a 75% reduction in size.** This is before any .kmz compression, and economizes the storage and usage of track files.

The software also renames processed files in a way that helps users organize them by date or registration/callsign. The name will reflect discovered metadata:
```
2026-01-01-N123AB.kml
```

If desired, location data can also be added:
```
2026-01-01-N123AB-KABC,KDEF.kml
```
> **NOTE:** *This may capture some odd things like GPS coordinates or nearby airports if ForeFlight had to make a guess during tracklog generation. The naming can be handled case-by-case during conversion if you prefer.*