/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./templates/**/*.{html,js,tmpl}"],
  theme: {
    extend: {},
  },
  plugins: [require("daisyui")],
  darkMode: "class",
  jit: true,
};
