// CallToAction Component - Vue Options API
const CallToAction = {
    name: 'CallToAction',
    template: `
        <section class="call-to-action" aria-label="Call to action">
            <div class="container">
                <div class="call-to-action-inner">
                    <div class="call-to-action-card">
                        <h3 class="call-to-action-title">Ready to Get Started?</h3>
                        <p class="call-to-action-text">
                            Try our demo now and see how easy it is to work with GPS coordinates.
                        </p>
                        <a href="/demo" class="call-to-action-button">
                            <svg xmlns="http://www.w3.org/2000/svg" class="call-to-action-icon" width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round" aria-hidden="true">
                                <path stroke="none" d="M0 0h24v24H0z" fill="none"></path>
                                <path d="M4 13a8 8 0 0 1 7 7a6 6 0 0 0 3 -5a9 9 0 0 0 6 -8a3 3 0 0 0 -3 -3a9 9 0 0 0 -8 6a6 6 0 0 0 -5 3"></path>
                                <path d="M7 14a6 6 0 0 0 -3 6a6 6 0 0 0 6 -3"></path>
                                <path d="M15 9m-1 0a1 1 0 1 0 2 0a1 1 0 1 0 -2 0"></path>
                            </svg>
                            Start Using LatLongAPI
                        </a>
                    </div>
                </div>
            </div>
        </section>
    `
};

