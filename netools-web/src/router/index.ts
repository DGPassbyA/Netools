import { createRouter,createWebHashHistory} from "vue-router";
import TextArea from '../components/TextArea.vue';
import FileArea from '../components/FileArea.vue';
import ListArea from '../components/ListArea.vue';

const routes = [
    { path: '/', component: TextArea },
    { path: '/text', component: TextArea },
    { path: '/file', component: FileArea },
    { path: '/list', component: ListArea },
];
export const router = createRouter({
    history: createWebHashHistory(),
    routes: routes
})