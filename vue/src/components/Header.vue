<script setup lang="ts">
import { computed, ref } from 'vue'
import {useStore} from 'vuex'
import LoginPage from '../page/login.vue'
import {ElMessage} from 'element-plus'
import {useRouter} from 'vue-router'
import {
  Fold,Expand,Setting
} from '@element-plus/icons-vue'
 const  circleUrl='https://cube.elemecdn.com/3/7c/3ea6beec64369c2642b92c6726f1epng.png'
const store=useStore()
const router=useRouter()
const collapse=computed(()=>store.state.collapse)
const isLogin=computed(()=>store.state.isLogin)
const is_admin=computed(()=>store.state.is_admin)
const username=computed(()=>store.state.username)
function changeMenu(){
  store.commit('changeCollapse',!collapse.value)
}
const handleCommand = (command: string | number | object) => {
  if(command=='a'){
    localStorage.clear()
    // location.reload()
    store.commit('logout')
  }else if(command=='b'){//分类
    router.push('/questionManage')
  }else if(command=='c'){//问题
    ElMessage('正在开发中')

  }
}
 const showLogin=ref(false)
 const loginSucc=()=>{
   showLogin.value=false
 }
</script>

<template>
 
 
  <div>
      <el-dropdown @command="handleCommand" v-if="isLogin">

        <div class="el-dropdown-link">
          <el-avatar :size="25" :src="circleUrl" />
          <span>您好！{{username}}</span>
           
           <el-icon><setting /></el-icon>
        </div>
        <template #dropdown>
          <el-dropdown-menu>
            <el-dropdown-item :icon="Plus" command="b" v-if="is_admin">进入管理</el-dropdown-item>
            

            <el-dropdown-item :icon="Plus" command="a">退出登录</el-dropdown-item>
            
          </el-dropdown-menu>
        </template>
      </el-dropdown>
      <span class="login" @click="showLogin=true" v-else>登录</span>
    <el-dialog
    v-model="showLogin"
    title="用户登录/注册"
    width="500px"
    :before-close="handleClose"
  >
    <LoginPage @loginSucc="loginSucc" /> 
    <!-- <template #footer>
      <span class="dialog-footer">
        <el-button @click="dialogVisible = false">Cancel</el-button>
        <el-button type="primary" @click="dialogVisible = false"
          >Confirm</el-button
        >
      </span>
    </template> -->
  </el-dialog>
  </div>

  
   <!-- 管理员
   <div class="setting">
     <el-icon><setting /></el-icon>
   </div> -->
 
</template>

<style scoped lang="scss">
.fold{
  padding-left: 10px;
}
 /* .user-msg{
   display: flex;
   justify-content: space-between;
   align-items: center;
   width: 20%;
   max-width: 250px;
   padding: 0 10px;
 } */
 .el-dropdown-link{
   display: flex;
   justify-content: space-between;
   align-items: center;
    padding-right: 20px;
    span{
      padding: 0 10px;
    }
 }
 .login{
   color: white;
   padding: 0 20px;
   cursor: pointer;
 }
</style>
