import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

const devHost = 'localhost:3050'

// https://vitejs.dev/config/
export default defineConfig({
  resolve: {
    alias: {
      '@': './src'
    }
  },
  server: {
    proxy: {
      '/api': {
        target: `http://${devHost}`,
        changeOrigin: true
      }
    }
  },
  plugins: [vue()]
})
