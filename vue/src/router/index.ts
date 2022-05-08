import { createRouter,createWebHashHistory,createWebHistory } from "vue-router";
import {useStore} from 'vuex'
import store from "../store";
const routes=[
    {
        path:"/",
        redirect:"/questionList",
        

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
                path:'/questionManage',
                name:'questionManage',

                component:()=>import('../page/manager/index.vue')
            },
            {
                path:'/questionDetail',
                name:'questionDetail',

                component:()=>import('../page/question/details.vue')
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
    const role = localStorage.getItem('token');
   
    if (!role && to.path !== '/login') {
        // next('/login');
    
       next();
    } else if (to.meta.permission) {
        // 如果是管理员权限则可进入，这里只是简单的模拟管理员权限而已
        role === 'admin'
            ? next()
            : next('/403');
    } else {
        if(role&&!store.state.isLogin){
            const username = localStorage.getItem('username');
            const is_admin = localStorage.getItem('is_admin');
            store.commit('loginSucc',role)
            store.commit('setUser',{username,is_admin})
        }
        // if(){
             
        // }
       
        next();
       

    }
    // next()
})
export default router