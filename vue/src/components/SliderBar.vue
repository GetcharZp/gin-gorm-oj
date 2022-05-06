<template>
<div :class="['slide-bar',collapse?'collapse':'']">
    <div class="logo" v-if="!collapse">
       <router-link to="/index">在线练题系统</router-link>
     </div>
     <div class="logo" v-else>
       <router-link to="/index">logo</router-link>
     </div>
     <div class="headers">
 <el-menu
        active-text-color="#dfff7d"
        background-color="#0087bf"
        class="el-menu-vertical-demo"
         :default-active="onRoutes" 
        text-color="#fff"
        style="border:none"
        :collapse="collapse"
        router
         mode="horizontal"
       
      >
       <sliderbar-item v-for="(item,index) in menuList" :index="index" :key="item.id" :item="item"></sliderbar-item>  
        
      </el-menu>
       <v-header></v-header>
     </div>
     
</div>
  
</template>

<script lang="ts" setup>
import {
  Location,
  Document,
  Menu as IconMenu,
  Setting,
} from '@element-plus/icons-vue'
import {computed, ref} from 'vue'
import {useStore} from 'vuex'
import {useRoute} from 'vue-router'
import SliderbarItem from './sliderBar/SliderbarItem.vue'
 import vHeader from "../components/Header.vue";
const store=useStore()
const collapse=computed(()=>store.state.collapse)
const route=useRoute()
  const onRoutes = computed(() => {
            return route.path;
        });
// const handleOpen = (key: string, keyPath: string[]) => {
//   console.log(key, keyPath)
// }
// const handleClose = (key: string, keyPath: string[]) => {
//   console.log(key, keyPath)
// }
const menuList=[
  {name:'问题列表',id:1,path:'/questionList' },
  {name:'提交列表',id:2,path:'/submitList'},
  {name:'排行榜',id:3,path:'/topList'},
]
</script>
<style>

</style>
<style scoped lang="scss">
.logo{
  height: 50px;
  display: none;
  background-color: #0087bf;
  border-bottom: 1px solid #dfff7d;
  // display: flex;
  align-items: center;
  justify-content: center;
  a{
    color: #fff;
    font-weight: 600;
  }
}
 .slide-bar{
    //  width: 30%;
     transition: 0.3s;
     min-width: 250px;
    
     background-color: #0087bf;
     &.collapse{
       min-width: 0;
     }
   }
   .el-menu-vertical-demo{
     flex: 1;
   }
   .headers{
     display: flex;
     align-items: center;
   }
</style>