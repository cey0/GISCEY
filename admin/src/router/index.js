import { createRouter, createWebHistory } from 'vue-router';
import LoginComponent from '../components/LoginComponent.vue';
import AdminComponent from '../components/AdminComponent.vue';

const routes = [
  {
    path: '/',
    name: 'login',
    component: LoginComponent
  },
  {
    path: '/dash',
    name: 'dash',
    component: AdminComponent
  }
];

const router = createRouter({
  history: createWebHistory(),
  routes
});

export default router;
