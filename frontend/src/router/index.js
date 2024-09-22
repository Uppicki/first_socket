import AuthComponent from "@/components/Auth.vue";
import LoginComponent from "@/components/Login.vue";
import SignupComponent from "@/components/Signup.vue";
import { createRouter, createWebHistory } from "vue-router";


const routes = [
    {
        path: '/',
        name: 'AuthComponent',
        component: AuthComponent
    },
    {
        path: '/auth/login',
        name: 'LoginComponent',
        component: LoginComponent
    },
    {
        path: '/auth/signup',
        name: 'SignupComponent',
        component: SignupComponent
    }
];


const router = createRouter({
    history: createWebHistory(),
    routes
});

export default router;

