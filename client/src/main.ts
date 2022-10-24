import { createApp } from 'vue'
import './style.scss'
import App from './App.vue'

import ElementPlus from 'element-plus'
import ja from 'element-plus/lib/locale/lang/ja.js'

import router from '@/router'

const app = createApp(App)
app.use(router)
app.use(ElementPlus, { locale: ja })
app.mount('#app')
