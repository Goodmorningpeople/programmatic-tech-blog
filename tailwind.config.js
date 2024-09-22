/** @type {import('tailwindcss').Config} */
module.exports = {
  daisyui: {
    themes: ["light", "dark", "night"],
  },
  content: [
    "./templates/**/*.html",
    "./static/**/*.{js, ts}",
    "./node_modules/daisyui/dist/**/*.js"
  ],
  theme: {
    extend: {
      fontFamily: {
        'comfortaa': ['Comfortaa', 'sans-serif'],
        'lora': ['Lora', 'serif'],
        'opensans': ['Open Sans', 'sans-serif']
      }
    },
  },
  plugins: [
    require('daisyui'),
  ],
}

