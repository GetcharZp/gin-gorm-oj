import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  base:'./',
  server: {	
    port:3555,			// ← ← ← ← ← ←
    // host: '0.0.0.0'	// ← 新增内容 ←
  },
   // 强制预构建插件包
   optimizeDeps: {
    include: [
      `monaco-editor/esm/vs/language/json/json.worker`,
      `monaco-editor/esm/vs/language/css/css.worker`,
      `monaco-editor/esm/vs/language/html/html.worker`,
      `monaco-editor/esm/vs/language/typescript/ts.worker`,
      `monaco-editor/esm/vs/editor/editor.worker`
    ], 
  },
  plugins: [vue()]
})
