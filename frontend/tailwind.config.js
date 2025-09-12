/** @type {import('tailwindcss').Config} */
export default {
    content: ["./src/**/*.{js,ts,jsx,tsx}"],
    theme: {
        extend: {
            colors: {
                mainBg: "var(--color-main-bg)",
                mainBgTransparent: "var(--color-main-bg-transparent)",
                firstBg: "var(--color-first-bg)",
                secondBg: "var(--color-second-bg)",
                thirdBg: "var(--color-third-bg)",
                textPrimary: "var(--color-text-primary)",
                textSecondary: "var(--color-text-secondary)",
                border: "var(--color-border)",
                gradient: "var(--color-gradient)",
            },
        },
    },
    plugins: [],
};
