import { createApp } from 'vue'
import App from './App.vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import zhCn from 'element-plus/es/locale/lang/zh-cn'

 import store from './store'
import router from './router'
createApp(App).use(router).use(ElementPlus,{
    locale: zhCn,
  }).use(store).mount('#app')

