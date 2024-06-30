# Groupie Tracker Geolocalization

## Description
Groupie Tracker is a web application designed to display information about various artists, their locations, concert dates, and relations. Implemented in Go, the server provides a simple and interactive graphical user interface (GUI) to browse and view detailed information about different artists. Utilizing HTML templates and CSS, the application ensures a seamless and visually appealing user experience. Whether viewing artist details or navigating through the site, Groupie Trackers offers users a platform where information meets elegance. Additionally, it offers mapping the concert locations of various artists and bands, by ensuring robustness and efficiency in handling and displaying geospatial data.

## Authors
- Ali (aabdulhu)
- Sayed Ali (sayedalawi)
- Yusuf (yrahma)

## Features
- Displays a list of artists with details such as image, members, creation date, and first album.
- Allows viewing detailed information about each artist.
- Implements a search functionality to find artists based on name, members, locations, first album date, and creation date.
- Provides typing suggestions in the search bar for a better user experience.
- Implements a filter functionality to find artists based on number of members, locations, first album date, and creation date.
- Maps the concert locations of artists/bands using geographic coordinates derived from addresses.
- Utilizes a Maps API to convert addresses into coordinates and place markers on the map.

## Usage: How to Run
1. Clone the repository:
   git clone (repository_url)
2. Navigate to the repository
   cd groupie-tracker
3. Run the code
   go run main.go
   The website will run right away

## Implementation Details: Algorithm
The implementation of server-side logic for displaying artist information is implemented using Go. Below is a broad overview of the algorithm employed:

1. HTTP Endpoints:
   * GET /:
      - Renders the main page containing a list of artists.on.
   * GET /view?id={id}:
      - Displays detailed information about a specific artist based on the provided ID.
      - Converts concert addresses into geographic coordinates and updates the map with markers.
   * POST /search
     - Handles search queries for artists by name, members, locations, first album date, and creation date.
   * POST /filtered
     - Handles feltering for artists by number of members, locations, first album date, and creation date.


2. Data Extraction:
   - Extracts data from the Groupie Trackers API and populates the structs for artists, locations, dates, and relations.

3. Template Rendering:
   - Uses HTML templates to render the main page, artist details, and error pages.
   - Handles error scenarios gracefully by displaying appropriate messages (404, 400, 500, 405).

4. Error Handling:
   - Provides suitable HTTP status codes for various scenarios (OK, Not Found, Bad Request, Internal Server Error, Method Not Allowed).

5. Utility Functions:
   - OpenBrowser: Automatically opens the default web browser to the application's URL.
   - ExtractToStruct: Fetches and unmarshals JSON data from the provided API URLs into the corresponding Go structs.

## Directory Structure
   - main.go: The main Go file containing the server implementation.
   - Templates/: Directory containing HTML templates for the main page, artist details page, search page, and error pages.
   - Style/: Directory containing CSS files for the main page, artist details page, search page, and error pages.
   - Functions/: Includes Go files with functions used in the application logic.


## Dependencies
   - Go programming language
   - Maps API for geocoding and mapping concert locations
   