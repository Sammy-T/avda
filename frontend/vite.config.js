import { defineConfig } from 'vite'
import { svelte } from '@sveltejs/vite-plugin-svelte'
import { resolve } from 'path'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [svelte()],
  resolve: {
    alias: {
        $assets: resolve('./src/assets'),
        $lib: resolve('./src/lib'),
        $stores: resolve('./src/lib/stores'),
        $util: resolve('./src/lib/util'),
        $wails: resolve('./wailsjs')
    }
  }
})
