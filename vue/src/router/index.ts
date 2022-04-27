import { createRouter,createWebHashHistory,createWebHistory } from "vue-router";
import {useStore} from 'vuex'
import store from "../store";
const routes=[
    {
        path:"/",
        redirect:"/index",
        

    },
    {
        path:"/index",
        component:()=>import('../page/index.vue'),
        children:[
            {
                path:'',
                name:'Home',
                component:()=>import('../page/home.vue')
            },
            {
                path:'/questionList',
                name:'questionList',

                component:()=>import('../page/question/index.vue')
            },
            {
                path:'/questionDetail',
                name:'questionDetail',

                component:()=>import('../page/question/detail.vue')
            },
            {
                path:'/topList',
                name:'topList',

                component:()=>import('../page/top/index.vue')
            },
            {
                path:'/submitList',
                name:'submitList',

                component:()=>import('../page/submit/index.vue')
            },
            
        ]
    },
    {
        path:'/login',
        component:()=>import('../page/login.vue')
    }
]
const router=createRouter({
    history:createWebHashHistory(),
    routes
})
router.beforeEach((to,from,next)=>{
    const role = localStorage.getItem('login');
    if (!role && to.path !== '/login') {
        // next('/login');
        if(store.state.isLogin){
            store.commit('logout')
       }
       next();
    } else if (to.meta.permission) {
        // 如果是管理员权限则可进入，这里只是简单的模拟管理员权限而已
        role === 'admin'
            ? next()
            : next('/403');
    } else {
        if(!store.state.isLogin){
             store.commit('loginSucc')
        }
       
        next();
       

    }
    // next()
})
export default router