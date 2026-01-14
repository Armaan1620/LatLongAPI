// DemoPreview Component - Vue Options API
const DemoPreview = {
    name: 'DemoPreview',
    data() {
        return {
            map: null,
            marker: null,
            coordinates: [
                { location: 'Bangalore, India', lat: 12.9716, lng: 77.5946, latDir: 'N', lngDir: 'E' },
                { location: 'New York, USA', lat: 40.7128, lng: 74.0060, latDir: 'N', lngDir: 'W' },
                { location: 'Tokyo, Japan', lat: 35.6762, lng: 139.6503, latDir: 'N', lngDir: 'E' }
            ]
        };
    },
    mounted() {
        this.$nextTick(() => {
            const mapContainer = this.$refs.mapContainer;
            if (!mapContainer || typeof L === 'undefined') {
                return;
            }

            const bangaloreCoords = [12.9716, 77.5946];
            this.map = L.map(mapContainer, {
                zoomControl: true,
                scrollWheelZoom: false,
            }).setView(bangaloreCoords, 11);

            L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
                maxZoom: 19,
                attribution: '© OpenStreetMap contributors',
            }).addTo(this.map);

            this.marker = L.marker(bangaloreCoords).addTo(this.map);
            this.marker.bindPopup('<strong>Bangalore, India</strong><br>12.9716° N, 77.5946° E').openPopup();

            window.addEventListener('resize', () => {
                if (this.map) {
                    this.map.invalidateSize();
                }
            });
        });
    },
    beforeUnmount() {
        if (this.map) {
            this.map.remove();
        }
    },
    template: `
        <section class="demo-preview" aria-label="Demo preview section">
            <div class="container">
                <div class="demo-preview-grid">
                    <div class="demo-preview-content">
                        <h2 class="demo-preview-title">See It In Action</h2>
                        <p class="demo-preview-subtitle">
                            Try searching for Bangalore coordinates and see the magic happen!
                        </p>
                        <div class="demo-preview-coordinates">
                            <div v-for="coord in coordinates" :key="coord.location" class="demo-preview-coordinate">
                                <strong>{{ coord.location }}:</strong>
                                <span>{{ coord.lat }}° {{ coord.latDir }}, {{ coord.lng }}° {{ coord.lngDir }}</span>
                            </div>
                        </div>
                        <a href="/demo" class="btn btn-primary demo-preview-button">
                            Try Full Demo
                        </a>
                    </div>
                    <div class="demo-preview-map-wrapper">
                        <div ref="mapContainer" class="demo-preview-map"></div>
                    </div>
                </div>
            </div>
        </section>
    `
};

