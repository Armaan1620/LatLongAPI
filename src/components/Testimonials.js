// Testimonials Component - Vue Options API
const Testimonials = {
    name: 'Testimonials',
    data() {
        return {
            testimonials: [
                {
                    quote: 'Perfect for logistics companies needing precise addresses.',
                    author: 'Logistics Pro',
                    meta: 'Delivery Company'
                },
                {
                    quote: 'Essential tool for our research team\'s location analysis.',
                    author: 'Research Team',
                    meta: 'University'
                },
                {
                    quote: 'Simple API integration saved us weeks of development time.',
                    author: 'Dev Team',
                    meta: 'Tech Startup'
                }
            ]
        };
    },
    template: `
        <section class="testimonials" aria-label="Customer testimonials">
            <div class="container">
                <div class="testimonials-grid">
                    <article v-for="testimonial in testimonials" :key="testimonial.author" class="testimonial-card">
                        <p class="testimonial-quote">"{{ testimonial.quote }}"</p>
                        <div class="testimonial-author">
                            <span class="testimonial-author-name">{{ testimonial.author }}</span>
                            <span class="testimonial-author-meta">— {{ testimonial.meta }}</span>
                        </div>
                    </article>
                </div>
            </div>
        </section>
    `
};

