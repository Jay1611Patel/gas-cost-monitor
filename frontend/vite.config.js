import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    proxy: {
      // Proxy requests from /api to your backend server
      '/api': {
        target: 'http://localhost:3000', // Your Express API's port
        changeOrigin: true,
        //rewrite: (path) => path.replace(/^\/api/, ''),
      },
    }
  }
})