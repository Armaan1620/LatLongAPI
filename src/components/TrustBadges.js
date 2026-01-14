// TrustBadges Component - Vue Options API
const TrustBadges = {
    name: 'TrustBadges',
    template: `
        <section class="trust-badges" aria-label="Trust badges">
            <div class="container">
                <div class="trust-badges-inner">
                    <div class="trust-badge">
                        <div class="trust-badge-icon">
                            <img src="/static/img/openstreetmap-logo.svg" alt="OpenStreetMap logo">
                        </div>
                        <p class="trust-badge-label">Powered by OpenStreetMap</p>
                    </div>
                    <div class="trust-badge">
                        <div class="trust-badge-icon trust-badge-icon-leaflet">
                            <img src="/static/img/leaflet-logo.svg" alt="Leaflet.js logo">
                        </div>
                        <p class="trust-badge-label">Built with Leaflet.js</p>
                    </div>
                    <div class="trust-badge">
                        <div class="trust-badge-icon trust-badge-icon-shield" aria-hidden="true">
                            <svg viewBox="0 0 24 24" class="trust-badge-shield-svg" role="presentation">
                                <path d="M12 2.25L5 4.75v6.42c0 4.03 2.93 7.76 7 8.83 4.07-1.07 7-4.8 7-8.83V4.75L12 2.25z" fill="currentColor" opacity="0.12" />
                                <path d="M12 2.25L5 4.75v6.42c0 4.03 2.93 7.76 7 8.83 4.07-1.07 7-4.8 7-8.83V4.75L12 2.25zm0 1.71l5 1.77v5.44c0 3.19-2.17 6.21-5 7.14-2.83-.93-5-3.95-5-7.14V5.73l5-1.77z" fill="currentColor" />
                                <path d="M10.75 13.69l-1.69-1.69-.96.96 2.65 2.65 4.46-4.46-.96-.96-3.5 3.5z" fill="currentColor" />
                            </svg>
                        </div>
                        <p class="trust-badge-label">SSL Secured</p>
                    </div>
                </div>
            </div>
        </section>
    `
};

