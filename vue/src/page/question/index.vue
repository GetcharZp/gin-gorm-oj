<template>
    <div>
        <div class="sort-list">
            问题分类：
            <div>
                <span @click="getProblem(null)" :class="!actSort?'act':''">全部</span>
                <span v-for="item in sortList" :key="item.id" @click="getProblem(item.identity)" :class="actSort==item.identity?'act':''">{{item.name}}</span>
            </div>
            
        </div>
        <div class="list">
            <div class="item" v-for="item in quesList" :key="item.id">
                <div>
                    <div class="title">
                        {{item.title}}
                        <div class="sort">
                            <span v-for="mi in item.problem_categories" :key="mi.id">
                                {{mi.category_basic.name}}
                            </span>
                        </div>
                    </div>
                    <div class="msg">
                        <span>创建时间：{{item.created_at}}</span>
                        <span>提交人数：{{item.submit_num}}</span>
                        <span>通过人数：{{item.pass_num}}</span>
                    </div>
                </div>
                <div class="edit">
                   <el-icon @click="toDetail(item)" title="详情"><edit /></el-icon>
                    <el-icon title="排行"><histogram /></el-icon>
                    <el-icon  title="提交列表"><list /></el-icon>
                </div>
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
 const quesList=ref([])
 const sortList=ref([])
 const actSort=ref<null|number>(null)
 const getProblem=(sortId:number|null)=>{
     actSort.value=sortId
      api.getProblemList({
          category_identity:sortId
 }).then(res=>{
     if(res&&res.data){
         quesList.value=res.data.data.list
     }
     console.log(res)
 })
 }
getProblem(null)
 api.getSortList({}).then(res=>{
      if(res&&res.data){
         sortList.value=res.data.data.list
     }
 })
 const toDetail=(item:any)=>{
     router.push({
         path:'/questionDetail',
         query:item
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
    .edit{
        width: 80px;
        display: flex;
        justify-content: space-between;
        color: #13b6f9;
        cursor: pointer;
    }
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
            margin-right: 20px;
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