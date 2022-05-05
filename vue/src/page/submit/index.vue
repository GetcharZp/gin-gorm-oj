<template>
    <div>
        
        <div class="list">
            <div class="msg title">
                <span>问题</span><span>用户</span><span>状态</span>
            </div>
            <div class="msg" v-for="item in submitList" :key="item.id">
               
                    
                   
                        <span @click="toProblem(item.problem_basic)" class="name">{{item.problem_basic.title}}</span>
                        <span v-if="item.user_basic">{{item.user_basic.name}}</span>
                        <span>{{status[item.status]}}</span>
                   
                 
                
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
import {useRouter,useRoute} from 'vue-router'
const router=useRouter()
 const submitList=ref([])
const status=ref(['未知','答案正确','答案错误','超时','运行超内存','编译错误'])
 const actSort=ref<null|number>(null)
 const route=useRoute()
 const getSubmitList=(sortId:number|null)=>{
     actSort.value=sortId
      api.getSubmitList({
          problem_identity:route.query.identity
 }).then(res=>{
     if(res&&res.data){
         submitList.value=res.data.data.list
     }
     console.log(res)
 })
 }
getSubmitList(null)
 const toProblem=(detail:any)=>{
      router.push({
         path:'/questionDetail',
         query:detail
     })
 }
 


</script>
<style scoped lang="scss">
.item{
    // padding: 20px;
  
    display: flex;
    justify-content: space-between;
    align-items: center;
   

   
}
 .msg{
        font-size: 14px;
    padding: 20px;
        display: flex;
        color: #999;
          border-bottom: 1px solid #eee;
        span:nth-child(1){
            width: 80%;
        }
        span:nth-child(2){
            width: 10%;
        }
        span:nth-child(3){
            width: 10%;
        }
        &.title{
            background-color: #eee;
        }
        .name{
            color: skyblue;
            cursor: pointer;
        }
    }
</style>