<template>
    <div class="app-shell">
        <NuxtRouteAnnouncer />
        <NuxtPage />
    </div>
</template>

<style>
:root {
    color-scheme: light;
    --bg: #f5efe6;
    --bg-accent: #ded1c0;
    --panel: rgba(255, 255, 255, 0.82);
    --panel-border: rgba(76, 54, 41, 0.14);
    --text: #1f1a17;
    --muted: #6d5f57;
    --accent: #c65f2b;
    --accent-deep: #8d3f1f;
    --shadow: 0 24px 70px rgba(50, 32, 20, 0.14);
}

html,
body,
#__nuxt {
    min-height: 100%;
}

body {
    margin: 0;
    font-family:
        Inter,
        ui-sans-serif,
        system-ui,
        -apple-system,
        BlinkMacSystemFont,
        "Segoe UI",
        sans-serif;
    background: radial-gradient(circle at top left, rgba(198, 95, 43, 0.18), transparent 35%), radial-gradient(circle at bottom right, rgba(141, 63, 31, 0.16), transparent 30%), linear-gradient(135deg, var(--bg), var(--bg-accent));
    color: var(--text);
}

a {
    color: inherit;
}

button,
input,
select {
    font: inherit;
}

.app-shell {
    min-height: 100vh;
}
</style>
