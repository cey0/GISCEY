import { createRouter, createWebHistory } from 'vue-router'
import MapsView from '../views/HomeView.vue';
import HomeView from '../components/HomeComponent.vue'

const routes = [
  {
    path: '/maps',
    name: 'maps',
    component: MapsView,
  },
  {
    path: '/',
    name: 'home',
    component: HomeView
  }
];

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
});

export default router;
