// HowItWorks Component - Vue Options API
const HowItWorks = {
    name: 'HowItWorks',
    data() {
        return {
            steps: [
                {
                    number: 1,
                    title: 'Enter Coordinates',
                    text: 'Input latitude and longitude values in decimal format.'
                },
                {
                    number: 2,
                    title: 'Visualize Location',
                    text: 'See your coordinates plotted on an interactive map.'
                },
                {
                    number: 3,
                    title: 'Get Address',
                    text: 'Receive detailed address information for your coordinates.'
                }
            ]
        };
    },
    template: `
        <section class="how-it-works" aria-label="How it works">
            <div class="container">
                <div class="how-it-works-inner">
                    <h2 class="how-it-works-title">How It Works</h2>
                    <div class="how-it-works-steps">
                        <article v-for="step in steps" :key="step.number" class="how-it-works-step">
                            <div class="how-it-works-badge">{{ step.number }}</div>
                            <h3 class="how-it-works-step-title">{{ step.title }}</h3>
                            <p class="how-it-works-step-text">{{ step.text }}</p>
                        </article>
                    </div>
                </div>
            </div>
        </section>
    `
};

