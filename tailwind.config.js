/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./templates/**/*.{html,js,tmpl}"],
  theme: {
    extend: {},
    container: {
      center: true,
      padding: "2rem",
    },
  },
  plugins: [require("daisyui")],
  purge: {
    enabled: true,
    content: ["./src/**/*.html"],
  },
};
