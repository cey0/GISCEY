import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  server: {
    proxy: {
      '/types': {
        target: 'http://49.13.157.100:8080',
        changeOrigin: true,
        secure: false,
      },
    },
  },
})
