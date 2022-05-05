import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  base:'./',
  server: {	
    port:3555,			// ← ← ← ← ← ←
    // host: '0.0.0.0'	// ← 新增内容 ←
  },
  plugins: [vue()]
})
