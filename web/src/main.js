import {createApp} from 'vue'
import App from './App.vue'
import router from './router'
import 'element-plus/dist/index.css'
import {createPinia} from 'pinia'
import { createPersistedState } from 'pinia-plugin-persistedstate'

const app = createApp(App)
const pinia = createPinia();
pinia.use(createPersistedState({
    auto: true, storage: localStorage
}));

app.use(router).use(pinia).mount('#app')

