/** @type {import('tailwindcss').Config} */
module.exports = {
	content: ["./public/views/**/*.{html,js}"],
	theme: {
		extend: {
			colors: {
				clifford: '#da373d',
			}
		}
	},
	plugins: [
		require('daisyui'),
	],
}
