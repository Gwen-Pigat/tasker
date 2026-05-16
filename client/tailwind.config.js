/** @type {import('tailwindcss').Config} */
export default {
  content: ['./src/**/*.{html,js,svelte,ts}'],
  theme: {
    extend: {
      colors: {
        primary: '#6366f1',
        'primary-hover': '#4f46e5',
        'bg-dark': '#0f172a',
        'card-bg': 'rgba(30, 41, 59, 0.7)',
      }
    },
  },
  plugins: [],
}
