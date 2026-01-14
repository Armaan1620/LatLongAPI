# LatLongAPI - Go Implementation

A Go-based recreation of [latlongapi.com](https://latlongapi.com), providing reverse geocoding services to convert GPS coordinates (latitude/longitude) to human-readable addresses.

## Features

- **Reverse Geocoding API**: Convert GPS coordinates to addresses using OpenStreetMap's Nominatim service
- **Interactive Demo**: Try the API with an interactive map powered by Leaflet.js
- **Clean UI**: Modern, responsive design with dark theme
- **RESTful API**: Simple JSON API endpoints
- **Documentation**: Complete API documentation
- **Pricing Page**: Transparent pricing tiers

## Getting Started

### Prerequisites

- Go 1.19 or later
- Internet connection (for OpenStreetMap Nominatim API)

### Installation

1. Clone or navigate to the project directory:
```bash
cd latlongapi
```

2. Run the server:
```bash
go run .
```

The server will start on port 8080 by default (or the port specified in the `PORT` environment variable).

### Usage

Visit `http://localhost:8080` in your browser to access the website.

#### Pages

- **Homepage** (`/`): Overview of the service
- **Demo** (`/demo`): Interactive tool to test coordinates and see results on a map
- **Docs** (`/docs`): API documentation
- **Pricing** (`/pricing`): Pricing plans

#### API Endpoint

**Reverse Geocoding**
```
GET /api/v1/convert?lat={latitude}&lng={longitude}
```

**Example Request:**
```bash
curl "http://localhost:8080/api/v1/convert?lat=51.5074&lng=-0.1278"
```

**Example Response:**
```json
{
  "latitude": "51.5074",
  "longitude": "-0.1278",
  "address": "Westminster, London, Greater London, England, SW1A 1AA, United Kingdom",
  "city": "London",
  "country": "United Kingdom",
  "state": "England",
  "postcode": "SW1A 1AA",
  "road": "Parliament Square"
}
```

## Project Structure

```
latlongapi/
├── main.go              # Main server application
├── go.mod               # Go module file
├── templates/           # HTML templates
│   ├── layout.html      # Base layout template
│   ├── index.html       # Homepage
│   ├── demo.html        # Demo page with interactive map
│   ├── docs.html        # API documentation
│   ├── pricing.html     # Pricing page
│   └── 404.html         # 404 error page
└── static/
    └── css/
        └── styles.css   # Stylesheet
```

## Technology Stack

- **Backend**: Go standard library (`net/http`, `html/template`)
- **Geocoding**: OpenStreetMap Nominatim API
- **Maps**: Leaflet.js (via CDN)
- **Styling**: Custom CSS with modern design

## API Rate Limits

This implementation uses OpenStreetMap's Nominatim service, which has usage policies:
- Maximum 1 request per second
- Requires proper User-Agent header (already implemented)
- For production use, consider using a commercial geocoding service or self-hosting Nominatim

## License

This is a recreation/clone project for educational purposes.



