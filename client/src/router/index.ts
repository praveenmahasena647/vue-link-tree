import { createRouter, createWebHistory } from 'vue-router'

//@ts-ignore
function loggedInRedirectDashboard(to, from, next) {
    if (localStorage.getItem('mhm')) {
        next()
        return
    } else {
        next('/login')
    }
}

//@ts-ignore
function isLoggedIn(to, from, next) {
    if (localStorage.getItem('mhm')) {
        next('/DashBoard')
        return
    }
    next()
}

//@ts-ignore
function adminLoggedIn(to, form, next) {
    if (localStorage.getItem('adminCum')) {
        next('/cuminMouth')
        return
    }
    next()
}

//@ts-ignore
function adminIsLogged(to, form, next) {
    if (localStorage.getItem('adminCum')) {
        next()
        return
    }
    next('/cum')
}

// @ts-ignore
const routes = [
    {
        path: '/',
        // @ts-ignore
        component: () => import('../views/AllUsers.vue'),
    },
    {
        path: '/:id',
        // @ts-ignore
        component: () => import('../views/ViewOne.vue'),
    },
    {
        path: '/create',
        beforeEnter: isLoggedIn,
        // @ts-ignore
        component: () => import('../views/CreateUser.vue'),
    },
    {
        path: '/login',
        beforeEnter: isLoggedIn,
        // @ts-ignore
        component: () => import('../views/Login.vue'),
    },
    {
        path: '/paths',
        // @ts-ignore
        component: () => import('../views/Paths.vue'),
    },
    {
        path: '/DashBoard',
        beforeEnter: loggedInRedirectDashboard,
        // @ts-ignore
        component: () => import('../views/DashBoard.vue'),
    },
    {
        path: '/cum',
        beforeEnter: adminLoggedIn,
        //@ts-ignore
        component: () => import('../views/AdminLog.vue'),
    },
    {
        path: '/cuminMouth',
        beforeEnter: adminIsLogged,
        //@ts-ignore
        component: () => import('../views/AdminCum.vue'),
    },
]

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    // @ts-ignore
    routes,
})

export default router
