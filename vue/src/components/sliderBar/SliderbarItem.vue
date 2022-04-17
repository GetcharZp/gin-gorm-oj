<template>
 
      <!-- <template>
        <template> -->
        <el-sub-menu :index="item.path" v-if="item.children">
          <template #title>
             
             
               <el-icon><icon-menu /></el-icon>
                
               <span>{{item.name}}</span>
              
                 
           
            <!-- <el-icon><location /></el-icon>
            
            <span>{{item.name}}</span> -->
          </template>
          <!-- <el-menu-item-group title="Group One">
            <el-menu-item index="1-1">item one</el-menu-item>
            <el-menu-item index="1-2">item one</el-menu-item>
          </el-menu-item-group>
          <el-menu-item-group title="Group Two">
            <el-menu-item index="1-3">item three</el-menu-item>
          </el-menu-item-group> -->
          <sliderbar-item v-for="(mi,i) in item.children" :index="item.path" :key="mi.id" :item="mi"></sliderbar-item>
          
        </el-sub-menu>
           

      
          
             <el-menu-item  :index="item.path" v-else>
               <el-icon><icon-menu /></el-icon>
               <template #title>{{item.name}}</template>
               <!-- <span></span> -->
                 </el-menu-item>
          
      
        
        
     
</template>
 

<script lang="ts" setup>

import {
  Location,
  Document,
  Menu as IconMenu,
  Setting,
} from '@element-plus/icons-vue'
const props=defineProps({
    item:{
        type:Object,
        default(){
            return {}
        }
    },
    index:{}
})
 
 
</script>
<style scoped>
/* .link-tag{
  display: block;
  width: 100%;
} */
</style>