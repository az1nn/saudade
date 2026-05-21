/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./**/*.html", "./**/*.templ", "./**/*.go"],
  safelist: [],
  theme: {
    extend: {
      colors: {
        parchment: "#F4F1EA",
        "aged-ink": "#3E362E",
        terracotta: "#C28B7E",
        sage: "#8A9A86",
      },
      fontFamily: {
        display: ['"Playfair Display"', "Georgia", "serif"],
        body: ["Lato", "system-ui", "sans-serif"],
      },
    },
  },
};
