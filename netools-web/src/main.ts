import { createApp } from 'vue'
import App from './App.vue'
import './index.css'
import axios from 'axios'
import {router} from './router'
import VueClipboard from 'vue-clipboard2'

axios.defaults.baseURL = "http://127.0.0.1:8081/api/"

const app = createApp(App)
app.use(VueClipboard)
app.use(router)
app.mount('#app')


