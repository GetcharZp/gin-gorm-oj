<template>
<div class="login-page">
 <el-form
    ref="ruleFormRef"
    :model="ruleForm"
    :rules="rules"
    label-width="60px"
    class="login-form"
    :size="formSize"
  >
    <el-form-item label="账户" prop="name">
      <el-input v-model="ruleForm.name" />
    </el-form-item>
     
    <el-form-item label="密码" prop="password">
      <el-input v-model="ruleForm.password" />
    </el-form-item>
     <div style="text-align:center;">
         <el-button type="primary" @click="submitForm(ruleFormRef)"
        >登录</el-button
      >
     </div>
     
  </el-form>
</div>
 
</template>

<script lang="ts" setup>
import { reactive, ref } from 'vue'
import {useStore} from 'vuex'
import {useRouter} from 'vue-router'
import type { FormInstance } from 'element-plus'
const emits=defineEmits(['loginSucc'])
const formSize = ref('default')
const store=useStore()
const ruleFormRef = ref<FormInstance>()
const ruleForm = reactive({
  name: 'Hello',
 password:''
})

const rules = reactive({
  name: [
    { required: true, message: 'Please input Activity name', trigger: 'blur' },
    { min: 3, max: 5, message: 'Length should be 3 to 5', trigger: 'blur' },
  ],
  password: [
    { required: true, message: 'Please input Activity name', trigger: 'blur' },
    { min: 3, max: 5, message: 'Length should be 3 to 5', trigger: 'blur' },
  ],
   
})
const router=useRouter()
console.log(router)
const submitForm = async (formEl: FormInstance | undefined) => {
  if (!formEl) return
  await formEl.validate((valid, fields) => {
    if (valid) {
      localStorage.setItem('login',1)
      store.commit('loginSucc')
      emits('loginSucc')
      // router.push('/index')
    } else {
      console.log('error submit!', fields)
    }
  })
}

const resetForm = (formEl: FormInstance | undefined) => {
  if (!formEl) return
  formEl.resetFields()
}
</script>
<style lang="scss" scoped>
  .login-page{
    width: 100%;
    height: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
  }
  .login-form{
    // widows: 300px;
    border: 1px solid #eee;
    padding: 40px 80px 20px 80px;
    border-radius: 10px;
  }
</style>