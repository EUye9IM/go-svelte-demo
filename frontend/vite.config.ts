import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';

export default defineConfig(({ command }) => {
	let cfg: any = { plugins: [sveltekit()] }
	if (command == "serve") {
		cfg.server = {
			proxy: { "/api": "http://localhost:8001" } // 与go后端端口一致，用于在 npm run dev 下开发。
		}
	}
	return cfg
});
