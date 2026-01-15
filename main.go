package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"latlongapi/backend/handlers"
	"latlongapi/backend/middleware"
	"latlongapi/backend/store"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

// tmplCache holds parsed templates.
var tmplCache *template.Template

// loadTemplates parses all templates in the templates directory.
func loadTemplates() {
	var err error
	tmplCache, err = template.ParseGlob("frontend/templates/*.html")
	if err != nil {
		log.Fatalf("error parsing templates: %v", err)
	}
}

// renderTemplate renders a named template wrapped in the base layout.
func renderTemplate(w http.ResponseWriter, name string, data any) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// Execute the page template which includes layout.html via blocks
	err := tmplCache.ExecuteTemplate(w, name, data)
	if err != nil {
		log.Printf("error rendering template %s: %v", name, err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// homeHandler serves the landing page.
func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		notFoundHandler(w, r)
		return
	}
	renderTemplate(w, "index.html", map[string]any{
		"Title": "LatLongAPI - Simple Latitude & Longitude API in Go",
	})
}

// docsHandler serves a simple API documentation page.
func docsHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "docs.html", map[string]any{
		"Title": "LatLongAPI Docs",
	})
}

// demoHandler serves the interactive demo page.
func demoHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "demo.html", map[string]any{
		"Title": "LatLongAPI Demo",
	})
}

// pricingHandler serves a simple pricing page.
func pricingHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "pricing.html", map[string]any{
		"Title": "LatLongAPI Pricing",
	})
}

// whyLatLongGPSHandler serves the WhyLatLongGPS page.
func whyLatLongGPSHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "why-latlong-gps.html", map[string]any{
		"Title": "Why LatLongGPS",
	})
}

// personalHandler serves the Personal page.
func personalHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "personal.html", map[string]any{
		"Title": "Personal - LatLongAPI",
	})
}

// businessHandler serves the Business page.
func businessHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "business.html", map[string]any{
		"Title": "Business - LatLongAPI",
	})
}

// loginHandler serves the login page.
func loginHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "login.html", map[string]any{
		"Title": "Login - LatLongAPI",
	})
}

// notFoundHandler renders a custom 404 page.
func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	renderTemplate(w, "404.html", map[string]any{
		"Title": "Page not found",
		"Path":  r.URL.Path,
	})
}

// reverseGeocode uses OpenStreetMap Nominatim API to convert coordinates to address.
func reverseGeocode(lat, lng float64) (map[string]interface{}, error) {
	// Parse coordinates
	latStr := strconv.FormatFloat(lat, 'f', 6, 64)
	lngStr := strconv.FormatFloat(lng, 'f', 6, 64)

	// Build Nominatim API URL
	apiURL := fmt.Sprintf("https://nominatim.openstreetmap.org/reverse?format=json&lat=%s&lon=%s&zoom=18&addressdetails=1",
		url.QueryEscape(latStr), url.QueryEscape(lngStr))

	// Create HTTP client with proper User-Agent (required by Nominatim)
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "LatLongAPI-Go/1.0")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("nominatim API returned status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return data, nil
}

// apiConvertHandler is a reverse geocoding API endpoint.
// It accepts lat and lng query parameters and returns address information.
func apiConvertHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	latStr := r.URL.Query().Get("lat")
	lngStr := r.URL.Query().Get("lng")

	if latStr == "" || lngStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Missing required parameters: lat and lng",
		})
		return
	}

	lat, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": fmt.Sprintf("Invalid latitude: %s", latStr),
		})
		return
	}

	lng, err := strconv.ParseFloat(lngStr, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": fmt.Sprintf("Invalid longitude: %s", lngStr),
		})
		return
	}

	// Validate coordinate ranges
	if lat < -90 || lat > 90 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Latitude must be between -90 and 90",
		})
		return
	}

	if lng < -180 || lng > 180 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Longitude must be between -180 and 180",
		})
		return
	}

	// Perform reverse geocoding
	geocodeData, err := reverseGeocode(lat, lng)
	if err != nil {
		log.Printf("Reverse geocoding error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Failed to geocode coordinates",
		})
		return
	}

	// Extract address information
	response := map[string]interface{}{
		"latitude":  latStr,
		"longitude": lngStr,
	}

	// Extract display_name (full address)
	if displayName, ok := geocodeData["display_name"].(string); ok {
		response["address"] = displayName
	}

	// Extract address components
	if addr, ok := geocodeData["address"].(map[string]interface{}); ok {
		if country, ok := addr["country"].(string); ok {
			response["country"] = country
		}
		if city, ok := addr["city"].(string); ok {
			response["city"] = city
		} else if town, ok := addr["town"].(string); ok {
			response["city"] = town
		} else if village, ok := addr["village"].(string); ok {
			response["city"] = village
		}
		if state, ok := addr["state"].(string); ok {
			response["state"] = state
		}
		if postcode, ok := addr["postcode"].(string); ok {
			response["postcode"] = postcode
		}
		if road, ok := addr["road"].(string); ok {
			response["road"] = road
		}
		if houseNumber, ok := addr["house_number"].(string); ok {
			response["house_number"] = houseNumber
		}
	}

	// Build a formatted address string if components exist
	var addressParts []string
	if addr, ok := geocodeData["address"].(map[string]interface{}); ok {
		if houseNum, ok := addr["house_number"].(string); ok {
			addressParts = append(addressParts, houseNum)
		}
		if road, ok := addr["road"].(string); ok {
			addressParts = append(addressParts, road)
		}
		if city, ok := addr["city"].(string); ok {
			addressParts = append(addressParts, city)
		} else if town, ok := addr["town"].(string); ok {
			addressParts = append(addressParts, town)
		} else if village, ok := addr["village"].(string); ok {
			addressParts = append(addressParts, village)
		}
		if state, ok := addr["state"].(string); ok {
			addressParts = append(addressParts, state)
		}
		if country, ok := addr["country"].(string); ok {
			addressParts = append(addressParts, country)
		}
		if len(addressParts) > 0 && response["address"] == nil {
			response["address"] = strings.Join(addressParts, ", ")
		}
	}

	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	_ = enc.Encode(response)
}

// healthHandler is a basic health check endpoint.
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write([]byte(`{"status":"ok"}`))
}

func main() {
	loadTemplates()

	// Initialize user store
	userStore := store.NewMemoryStore()

	// Initialize auth handler
	authHandler := handlers.NewAuthHandler(userStore)

	mux := http.NewServeMux()

	// Page routes.
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/docs", docsHandler)
	mux.HandleFunc("/demo", demoHandler)
	mux.HandleFunc("/pricing", pricingHandler)
	mux.HandleFunc("/why-latlong-gps", whyLatLongGPSHandler)
	mux.HandleFunc("/personal", personalHandler)
	mux.HandleFunc("/business", businessHandler)
	mux.HandleFunc("/login", loginHandler)

	// Auth API routes.
	mux.HandleFunc("/api/auth/register", authHandler.Register)
	mux.HandleFunc("/api/auth/login", authHandler.Login)
	mux.HandleFunc("/api/auth/logout", authHandler.Logout)
	
	// Protected routes (require authentication)
	authMiddleware := middleware.AuthMiddleware(userStore)
	mux.Handle("/api/auth/me", authMiddleware(http.HandlerFunc(authHandler.Me)))

	// API routes.
	mux.HandleFunc("/api/v1/convert", apiConvertHandler)
	mux.HandleFunc("/healthz", healthHandler)

	// Static files.
	staticDir := http.Dir("frontend/static")
	fileServer := http.FileServer(staticDir)
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	// Custom 404.
	muxWithNotFound := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Let the mux try to serve the route first.
		rr := &responseRecorder{ResponseWriter: w, status: http.StatusOK}
		mux.ServeHTTP(rr, r)
		if rr.status == http.StatusNotFound {
			notFoundHandler(w, r)
		}
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	addr := ":" + port
	log.Printf("LatLongAPI Go server listening on %s", addr)
	if err := http.ListenAndServe(addr, muxWithNotFound); err != nil {
		log.Fatalf("server error: %v", err)
	}
}

// responseRecorder wraps an http.ResponseWriter to capture the status code.
type responseRecorder struct {
	http.ResponseWriter
	status int
}

func (rr *responseRecorder) WriteHeader(code int) {
	rr.status = code
	rr.ResponseWriter.WriteHeader(code)
}
