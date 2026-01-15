# LatLongAPI - Project Summary

## Overview

LatLongAPI is a Go-based web application that provides reverse geocoding services to convert GPS coordinates (latitude/longitude) into human-readable addresses. It is a recreation of [latlongapi.com](https://latlongapi.com) built with Go, featuring a modern web interface, RESTful API, and user authentication system.

## Purpose

The application enables developers and users to:
- Convert GPS coordinates to physical addresses using OpenStreetMap's Nominatim service
- Access reverse geocoding through a simple RESTful API
- Test and visualize geocoding results with an interactive map interface
- Manage user accounts with secure authentication

## Technology Stack

### Backend
- **Language**: Go 1.22+
- **Web Framework**: Go standard library (`net/http`, `html/template`)
- **Authentication**: JWT (JSON Web Tokens) using `github.com/golang-jwt/jwt/v5`
- **Password Hashing**: bcrypt via `golang.org/x/crypto`
- **Geocoding Service**: OpenStreetMap Nominatim API

### Frontend
- **Templating**: Go `html/template` package
- **Styling**: Custom CSS with modern dark theme design
- **Mapping**: Leaflet.js (via CDN)
- **Interactive Components**: Vanilla JavaScript

## Architecture

### Project Structure

```
LatLongAPI/
├── main.go                  # Main application entry point, HTTP server, and core handlers
├── go.mod                   # Go module dependencies
├── go.sum                   # Dependency checksums
├── Makefile                 # Build and development automation
├── .env                     # Environment variables (not tracked in git)
├── backend/                 # Backend business logic
│   ├── auth/               # Authentication logic
│   │   ├── jwt.go          # JWT token generation and validation
│   │   └── password.go     # Password hashing and verification
│   ├── handlers/           # HTTP request handlers
│   │   └── auth.go         # Authentication endpoints (register, login, logout, me)
│   ├── middleware/         # HTTP middleware
│   │   └── auth.go         # JWT authentication middleware
│   ├── models/             # Data models
│   │   └── user.go         # User model and UserStore interface
│   └── store/              # Data storage implementations
│       └── memory.go       # In-memory user storage
├── templates/              # HTML templates
│   ├── layout.html         # Base layout template
│   ├── index.html          # Homepage
│   ├── demo.html           # Interactive demo with map
│   ├── docs.html           # API documentation page
│   ├── pricing.html        # Pricing information
│   ├── login.html          # Login/registration page
│   ├── personal.html       # Personal plan page
│   ├── business.html       # Business plan page
│   ├── why-latlong-gps.html # Information page
│   └── 404.html            # 404 error page
├── static/                 # Static assets
│   └── css/
│       └── styles.css      # Application styles
└── docs/                   # Project documentation
    └── PROJECT_SUMMARY.md  # This file
```

### Backend Architecture

#### Core Components

1. **HTTP Server** (`main.go`)
   - Handles HTTP routing and request dispatching
   - Serves HTML templates and static files
   - Implements custom 404 error handling
   - Configurable port via environment variable

2. **Authentication System**
   - **JWT-based authentication**: Tokens valid for 24 hours
   - **Password security**: bcrypt hashing for secure password storage
   - **Middleware protection**: Routes can be protected with JWT validation
   - **Token delivery**: Supports both Authorization header and cookie-based tokens

3. **User Management**
   - In-memory user storage (MemoryStore)
   - User registration with email validation
   - Secure login with password verification
   - User profile retrieval for authenticated users

4. **Geocoding Service**
   - Reverse geocoding via OpenStreetMap Nominatim
   - Extracts comprehensive address components (country, city, state, postcode, road)
   - Rate-limited to comply with Nominatim usage policies (1 req/sec)
   - Proper User-Agent header for API compliance

## API Endpoints

### Public Endpoints

#### Reverse Geocoding
```
GET /api/v1/convert?lat={latitude}&lng={longitude}
```

**Description**: Converts GPS coordinates to a human-readable address.

**Parameters**:
- `lat` (required): Latitude (-90 to 90)
- `lng` (required): Longitude (-180 to 180)

**Example Request**:
```bash
curl "http://localhost:8080/api/v1/convert?lat=51.5074&lng=-0.1278"
```

**Example Response**:
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

#### Health Check
```
GET /healthz
```

**Description**: Basic health check endpoint.

**Response**:
```json
{
  "status": "ok"
}
```

### Authentication Endpoints

#### User Registration
```
POST /api/auth/register
```

**Request Body**:
```json
{
  "email": "user@example.com",
  "password": "password123"
}
```

**Response** (201 Created):
```json
{
  "token": "eyJhbGciOiJIUzI1NiIs...",
  "user": {
    "id": 1,
    "email": "user@example.com",
    "created_at": "2026-01-15T13:00:00Z"
  }
}
```

#### User Login
```
POST /api/auth/login
```

**Request Body**:
```json
{
  "email": "user@example.com",
  "password": "password123"
}
```

**Response** (200 OK):
```json
{
  "token": "eyJhbGciOiJIUzI1NiIs...",
  "user": {
    "id": 1,
    "email": "user@example.com",
    "created_at": "2026-01-15T13:00:00Z"
  }
}
```

#### Get Current User
```
GET /api/auth/me
Authorization: Bearer {token}
```

**Response** (200 OK):
```json
{
  "id": 1,
  "email": "user@example.com",
  "created_at": "2026-01-15T13:00:00Z"
}
```

#### User Logout
```
POST /api/auth/logout
```

**Response** (200 OK):
```json
{
  "message": "Logged out successfully"
}
```

### Web Pages

- `GET /` - Homepage with service overview
- `GET /demo` - Interactive demo with map interface
- `GET /docs` - API documentation
- `GET /pricing` - Pricing plans
- `GET /login` - User login and registration
- `GET /personal` - Personal plan information
- `GET /business` - Business plan information
- `GET /why-latlong-gps` - Educational content about GPS coordinates

## Key Features

### 1. Reverse Geocoding
- Converts latitude/longitude coordinates to addresses
- Returns detailed address components
- Integrates with OpenStreetMap Nominatim service
- Includes coordinate validation

### 2. Interactive Demo
- Live map interface powered by Leaflet.js
- Click-to-geocode functionality
- Visual coordinate display
- Real-time address lookup

### 3. User Authentication
- Secure JWT-based authentication
- Password hashing with bcrypt
- Token expiration (24-hour validity)
- Protected API endpoints
- Support for both header and cookie-based authentication

### 4. Modern Web Interface
- Responsive design
- Dark theme aesthetic
- Clean, intuitive navigation
- Mobile-friendly layout

### 5. API Documentation
- Comprehensive endpoint documentation
- Example requests and responses
- Interactive testing tools
- Clear usage guidelines

## Development

### Prerequisites
- Go 1.22 or later
- Internet connection (for OpenStreetMap API)

### Installation & Setup

1. Clone the repository:
```bash
git clone https://github.com/Armaan1620/LatLongAPI.git
cd LatLongAPI
```

2. (Optional) Create a `.env` file for environment variables:
```bash
PORT=8080
JWT_SECRET=your-secret-key-change-in-production
```

3. Install dependencies:
```bash
go mod download
```

### Running the Application

#### Option 1: Using Make (Recommended)
```bash
# Development mode (auto-loads .env if present)
make dev

# Production build and run
make run

# Build only
make build

# Clean build artifacts
make clean

# Stop running server
make stop
```

#### Option 2: Direct Go Commands
```bash
# Run directly
go run .

# Build and run
go build -o latlongapi .
./latlongapi
```

The server starts on port 8080 by default (or the port specified in the `PORT` environment variable).

### Configuration

Environment variables:
- `PORT`: HTTP server port (default: 8080)
- `JWT_SECRET`: Secret key for JWT token signing (default: "your-secret-key-change-in-production")

### API Rate Limits

The application uses OpenStreetMap's Nominatim service, which has usage policies:
- Maximum 1 request per second
- Requires proper User-Agent header (already implemented)
- For production use, consider:
  - Using a commercial geocoding service
  - Self-hosting Nominatim
  - Implementing caching

## Security Considerations

### Authentication
- JWT tokens expire after 24 hours
- Passwords hashed with bcrypt (cost factor 10)
- Tokens can be invalidated client-side
- Protected routes require valid JWT

### Password Requirements
- Minimum 6 characters
- Stored as bcrypt hashes
- Never exposed in API responses

### API Security
- CORS headers enabled for API endpoints
- Input validation on all endpoints
- Coordinate range validation
- Proper error handling without information leakage

## Data Storage

Currently uses **in-memory storage** for user data:
- Users stored in memory (MemoryStore)
- Data persists only during application runtime
- Not suitable for production use
- Should be replaced with persistent storage (PostgreSQL, MySQL, etc.) for production

## Future Enhancements

Potential improvements for production deployment:
1. **Database Integration**: Replace in-memory storage with PostgreSQL/MySQL
2. **Caching**: Add Redis for geocoding result caching
3. **Rate Limiting**: Implement API rate limiting per user
4. **API Keys**: Add API key management for users
5. **Usage Tracking**: Monitor API usage per user
6. **Enhanced Security**: Add refresh tokens, password reset, email verification
7. **Scalability**: Add horizontal scaling support
8. **Monitoring**: Implement logging, metrics, and error tracking
9. **Testing**: Add comprehensive unit and integration tests
10. **CI/CD**: Set up automated testing and deployment pipelines

## License

This is a recreation/clone project for educational purposes.

## Support & Contact

For issues, questions, or contributions, please visit the GitHub repository:
https://github.com/Armaan1620/LatLongAPI
