<template>
    <div>
        
        <div class="list">
            <div class="item" v-for="item in rankList" :key="item.id">
                
                    <div class="title">
                        {{item.id}}.
                        {{item.name}}
                        <div class="sort">
                            <span>{{item.mail}}</span>
                        </div>
                    </div>
                    <div class="msg">
                        
                        <span>提交数：{{item.submit_num}}</span>
                        <span>通过数：{{item.pass_num}}</span>
                    </div>
                
                <!-- <div class="edit">
                   <el-icon @click="toDetail(item)" title="详情"><edit /></el-icon>
                    <el-icon title="排行" @click="toRank(item)"><histogram /></el-icon>
                    <el-icon  title="提交列表" @click="toSubList(item)"><list /></el-icon>
                </div> -->
            </div>
        </div>
    </div>
</template>
<script lang="ts" setup>
 import { reactive,ref } from '@vue/reactivity'
 import {
  Edit,Histogram,List
} from '@element-plus/icons-vue'
import api from '../../api/api.js'
import {useRouter} from 'vue-router'
const router=useRouter()
 const rankList=ref([])
 const sortList=ref([])
 const actSort=ref<null|number>(null)
 const getRankList=(sortId:number|null)=>{
     actSort.value=sortId
      api.getRankList({
          
 }).then(res=>{
     if(res&&res.data){
         rankList.value=res.data.data.list
     }
     console.log(res)
 })
 }
getRankList(null)
 
 const toDetail=(item:any)=>{
     router.push({
         path:'/questionDetail',
         query:item
     })
 }
  const toRank=(item:any)=>{
     router.push({
         path:'/topList',
         query:{identity:item.identity}
     })
 }
  const toSubList=(item:any)=>{
     router.push({
         path:'/submitList',
        query:{identity:item.identity}

     })
 }
</script>
<style scoped lang="scss">
.item{
    padding: 20px;
    border-bottom: 1px solid #eee;
    display: flex;
    justify-content: space-between;
    align-items: center;
 
    .title{
        margin-bottom: 20px;
        display: flex;
        align-items: center;
        font-size: 18px;
        span{
            background-color: #d3ebff;
            font-size: 12px;
            margin: 0 10px;
            padding: 4px;
            border-radius: 5px;
            border: 1px solid #62a9ff;
            color: #62a9ff;
        }
    }
    .msg{
        font-size: 14px;
       
        color: #999;
        span{
            margin-right: 40px;
        }
    }
}
.sort-list{
    display: flex;
    background-color: #d6fff2;
    padding: 20px;
    span{
        color: #13b6f9;
        margin-right: 20px;
        cursor: pointer;
        &.act{
            color: black;
        }
    }
}
</style>