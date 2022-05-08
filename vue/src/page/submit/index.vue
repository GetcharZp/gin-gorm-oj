<template>
    <div class="submit-list">
        
        <div class="list" v-loading="loading">
            <div class="msg title">
                <span>问题</span><span>用户</span><span>提交时间</span>
                <span style="display:flex;white-space:nowrap;align-items:center">状态：

                    <el-select v-model="mystatus" clearable @change="getSubmitList">
                        <el-option :value="i" v-for="(mi,i) in status" :key="mi" :label="mi">{{mi}}</el-option>
                    </el-select>
                </span>
            </div>
            <div class="msg" v-for="item in submitList" :key="item.id">
               
                    
                   
                        <span @click="toProblem(item.problem_basic)" class="name">{{item.problem_basic.title}}</span>
                        <span v-if="item.user_basic">{{item.user_basic.name}}</span>
                        <span v-if="item.user_basic">{{item.created_at}}</span>
                        <span>{{status[item.status]}}</span>
                   
                 
                
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
import {useRouter,useRoute} from 'vue-router'
const loading=ref(false)
const router=useRouter()
const mystatus=ref('')
 const submitList=ref([])
const status=ref(['未知','答案正确','答案错误','超时','运行超内存','编译错误'])

 const route=useRoute()
 const pageSize=ref(10)
const currentPage=ref(1)
const total=ref(0)
 
const handleSizeChange = (val: number) => {
   getSubmitList()

}
const handleCurrentChange = (val: number) => {
  getSubmitList()
}

 const getSubmitList=()=>{
     loading.value=true
      api.getSubmitList({
          problem_identity:route.query.identity,
           size:pageSize.value,
      page:currentPage.value,
     status: mystatus.value
 }).then(res=>{
     loading.value=false
     if(res&&res.data){
         submitList.value=res.data.data.list
          total.value=res.data.data.count
     }
     console.log(res)
 })
 }
getSubmitList()
 const toProblem=(detail:any)=>{
      router.push({
         path:'/questionDetail',
         query:detail
     })
 }
 


</script>
<style scoped lang="scss">
.submit-list{
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
            width:50%;
        }
        span:nth-child(2){
            width: 10%;
        }
        span:nth-child(3){
            width:20%;
        }
        span:nth-child(4){
            width: 20%;
        }
        &.title{
           border-bottom: 1px solid #0087bf;
border-top: 1px solid #0087bf;
            // color: white;
            align-items: center;
        }
        .name{
            color: skyblue;
            cursor: pointer;
        }
    }
</style>