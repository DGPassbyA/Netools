import { createApp } from 'vue'
import App from './App.vue'
import './index.css'
import axios from 'axios'

axios.defaults.baseURL = "http://127.0.0.1:8081/api/"
createApp(App).mount('#app')


