# README

## About EPG TV Viewer

This program allows you to download and parse EPG (Electronic Program Guide) data, search for TV shows, and stream live TV with a schedule list and rewind functionality. It interacts with the Stream TV API to fetch live TV shows and allows users to view their favorite shows with an easy-to-navigate interface.

Features
- Download EPG: Download and parse EPG data in XML format.
- Search TV Shows: Search for TV shows from the EPG based on titles or genres.
- Show and Download TV Shows: Access and download TV shows using your Stream TV API key.
- Live TV Streaming: View live TV with an interactive schedule list and the ability to rewind to a specific program.

## About EPG

EPG (Electronic Program Guide) XML is a standard format used to represent TV program schedules in XML (Extensible Markup Language). It provides information about upcoming television shows, including the program's title, description, start and end times, channel, and other metadata like genre and language.

EPG XML files are typically used by TV receivers, set-top boxes, and applications to display TV schedules to viewers. The data helps users browse through available programs and plan what to watch, similar to a digital guide you might find on TV networks or streaming services.

These XML files can vary in structure but generally include elements like:

- `<channel>`: Information about the TV channel (e.g., name, ID).
- `<programme>`: A specific TV program with details such as title, start and end times, and description.
- `<title>`: The program's title.
- `<desc>`: A description or summary of the program.
- `<start>` and `<stop>`: Timestamps for when the program starts and ends.

Many online sources or TV networks distribute EPG data in this XML format to integrate with various devices or software.
