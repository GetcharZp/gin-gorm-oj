<template>
    <div class="top-list">
        
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
        <div class="pagi">
      <el-pagination
        v-model:currentPage="currentPage"
        v-model:page-size="pageSize"
        :page-sizes="[10, 20, 50, 100]"
        
        layout="total,sizes, prev, pager, next"
        :total="total"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
      />
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
 const pageSize=ref(10)
const currentPage=ref(1)
const total=ref(0)
 
const handleSizeChange = (val: number) => {
   getRankList()

}
const handleCurrentChange = (val: number) => {
  getRankList()
}
 const getRankList=( )=>{
     
      api.getRankList({
             size:pageSize.value,
      page:currentPage.value
 }).then(res=>{
     if(res&&res.data){
         rankList.value=res.data.data.list
           total.value=res.data.data.count
     }
     console.log(res)
 })
 }
getRankList( )
 
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
.top-list{
   height: 100%;
   display: flex;
   justify-content: space-between;
   flex-direction: column;
}
.list{
    flex: 1;
    overflow: auto;
}
.pagi{
    text-align: center;
    display: flex;
    justify-content: center;
    padding: 10px 0;
    border-top: 1px solid #eee;
}
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