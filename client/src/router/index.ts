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
         beforeEnter:isLoggedIn,  
        // @ts-ignore
        component: () => import('../views/CreateUser.vue'),
    },
    {
        path: '/login',
         beforeEnter:isLoggedIn, 
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
        path: '/changeDetails',
        /* beforeEnter:loggedIn, */
        // @ts-ignore
        components: () => import('../views/Change.vue'),
    },
    {
        path:'/addingAdmin',
        // @ts-ignore
        component:()=>import('../views/Admin.vue')
    }
]

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    // @ts-ignore
    routes,
})

export default router