import { sveltekit } from '@sveltejs/kit/vite';

/** @type {import('vite').UserConfig} */
const config = {
	plugins: [sveltekit()],
	css: {
		preprocessorOptions: {
			scss: {
				additionalData: '@use "src/variables.scss" as *;'
			}
		},
	},
	optimizeDeps: {
		include: ["highlight.js", "highlight.js/lib/core"],
	}
};

export default config;
