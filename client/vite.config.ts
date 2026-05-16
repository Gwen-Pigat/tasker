import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';
import { SvelteKitPWA } from '@vite-pwa/sveltekit';

export default defineConfig({
	plugins: [
		sveltekit(),
		SvelteKitPWA({
			registerType: 'autoUpdate',
			manifest: {
				name: 'Tasker',
				short_name: 'Tasker',
				description: 'A premium task management application',
				theme_color: '#0f172a',
				background_color: '#0f172a',
				display: 'standalone',
				icons: [
					{
						src: '/images/favicon-96x96.png',
						sizes: '96x96',
						type: 'image/png'
					},
					{
						src: '/images/apple-touch-icon.png',
						sizes: '180x180',
						type: 'image/png'
					},
					{
						src: '/images/web-app-manifest-192x192.png',
						sizes: '192x192',
						type: 'image/png'
					},
					{
						src: '/images/web-app-manifest-512x512.png',
						sizes: '512x512',
						type: 'image/png'
					}
				]
			},
			workbox: {
				maximumFileSizeToCacheInBytes: 6000000 // 6MB
			}
		})
	]
});
