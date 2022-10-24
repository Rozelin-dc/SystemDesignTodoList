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
    port: 3000,
    proxy: {
      '/api': {
        target: `http://${devHost}`,
        changeOrigin: true,
        rewrite:
          process.env.NODE_ENV === 'development'
            ? path => path.replace(/^\/api/, '')
            : null
      }
    }
  },
  plugins: [vue()]
})
