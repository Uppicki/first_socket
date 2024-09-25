import AuthComponent from "@/components/Auth.vue";
import LoginComponent from "@/components/Login.vue";
import SignupComponent from "@/components/Signup.vue";
import AnyComponent from "@/components/Any.vue";
import store from "@/store/store";
import { createRouter, createWebHistory } from "vue-router";


const routes = [
    {
        path: '/',
        name: 'AnyComponent',
        component: AnyComponent,
        meta: {requiresAuth: true}
    },
    {
        path: '/auth',
        name: 'AuthComponent',
        component: AuthComponent,
        meta: {requiresAuth: false}
    },
    {
        path: '/auth/login',
        name: 'LoginComponent',
        component: LoginComponent,
        meta: {requiresAuth: false}
    },
    {
        path: '/auth/signup',
        name: 'SignupComponent',
        component: SignupComponent,
        meta: {requiresAuth: false}
    }
];


const router = createRouter({
    history: createWebHistory(),
    routes
});

router.beforeEach((to, from, next) => {
    console.log(to.matched.some(record => record.meta.requiresAuth))
    console.log(store.getters.isAuthenticated())
    if (to.matched.some(record => record.meta.requiresAuth)) {
        if (store.getters.isAuthenticated()) {
            next()
        } else {
            next({
                path: '/auth',
            })
        }
    } else {
        if (store.getters.isAuthenticated()) {
            next(next({
                path: '/',
            }))
        } else {
            next()
        }
       
    }
});

export default router;

