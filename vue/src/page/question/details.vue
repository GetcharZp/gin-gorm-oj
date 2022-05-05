<template>
    
        <div class="ques-cont">
            <div class="left">
                    <h3>{{detail.title}}</h3>
                    <p>{{detail.content}}</p>
                    <div class="msg">
                       分类：<span v-for="mi in detail.problem_categories" :key="mi.id">
                                {{mi.category_basic.name}}
                            </span>
                    </div>
                    <div class="msg">
                        <span>创建时间{{detail.created_at}}</span>
                        <span>提交次数：{{detail.submit_num}}</span>
                        <span>通过次数：{{detail.pass_num}}</span>
                        <span>最少用时：{{detail.pass_num}}ms</span>
                        
                    </div>
            </div>
            <div class="right">

            </div>
        </div>
   
</template>
<script lang="ts" setup>
 import { reactive,ref } from '@vue/reactivity'
 import {useRoute} from 'vue-router'
 import api from '../../api/api.js'
 import {
  Edit
} from '@element-plus/icons-vue'
const route=useRoute()
const detail=ref({

})
const getDetail=()=>{
    api.getProblemDetail({
        identity:route.query.identity
    }).then(res=>{
            if(res.data.data){
                detail.value=res.data.data
            }
    })
}
getDetail()
</script>
<style scoped lang="scss">
.ques-cont{
    display: flex;
    height: 100%;
    .left{
        width: 50%;
        line-height: 2;
        border-right: 10px solid #eee;
        padding: 10px;
        .msg{
            font-size: 12px;
            
            span{
                margin-right: 20px;
                color: #999;
            }
        }
    }
}
</style>