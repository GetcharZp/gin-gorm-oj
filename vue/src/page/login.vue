<template>
  <div class="login-page">
    <el-tabs
      v-model="activeName"
      type="card"
      class="demo-tabs"
      @tab-click="handleClick"
    >
      <el-tab-pane label="登录" name="first">
        <el-form
          ref="ruleFormRef"
          :model="ruleForm"
          :rules="rules"
          label-width="60px"
          class="login-form"
          :size="formSize"
        >
          <el-form-item label="账户" prop="username">
            <el-input v-model="ruleForm.username" />
          </el-form-item>

          <el-form-item label="密码" prop="password">
            <el-input v-model="ruleForm.password" />
          </el-form-item>
          <div style="text-align: center">
            <el-button type="primary" @click="submitForm(ruleFormRef)"
              >登录</el-button
            >
          </div>
        </el-form>
      </el-tab-pane>
      <el-tab-pane label="注册" name="second">
         
        <el-form
          ref="registFormRef"
          :model="registForm"
          :rules="rules"
          label-width="60px"
          class="login-form"
          :size="formSize"
        >
          <el-form-item label="账户" prop="name">
            <el-input v-model="registForm.name" />
          </el-form-item>

          <el-form-item label="密码" prop="password">
            <el-input v-model="registForm.password" />
          </el-form-item>
           <el-form-item label="邮箱" prop="email">
            <el-input v-model="registForm.email" />
          </el-form-item>
          <el-form-item label="验证码" prop="code">
            <el-input v-model="registForm.code" />
            <el-button @click="sendCode">发送验证码</el-button>
            <el-button  disabled>{{remainTime}}秒</el-button>
          </el-form-item>

          <div style="text-align: center">
            <el-button type="primary" @click="subRegister(registFormRef)"
              >注册</el-button
            >
          </div>
        </el-form>
      </el-tab-pane>
    
    </el-tabs>
  </div>
</template>

<script lang="ts" setup>
import { reactive, ref } from "vue";
import { useStore } from "vuex";
import { useRouter } from "vue-router";
import type { FormInstance } from "element-plus";
import type { TabsPaneContext } from "element-plus";
import {ElMessage } from "element-plus"
import api from '../api/api.js'
const activeName = ref("first");

const handleClick = (tab: TabsPaneContext, event: Event) => {
  console.log(tab, event);
};
const emits = defineEmits(["loginSucc"]);
const formSize = ref("default");
const store = useStore();
const ruleFormRef = ref<FormInstance>();
const registFormRef=ref<FormInstance>()
const remainTime=ref(60)
const ruleForm = reactive({
  username: "Hello",
  password: "",
});
const registForm=reactive({
   name: "Hello",
  password: "",
  email:'',
  code:''
})
const rules = reactive({
  username: [
    { required: true, message: "Please input Activity name", trigger: "blur" },
    
  ],
  password: [
    { required: true, message: "Please input Activity name", trigger: "blur" },
     
  ],
});
const router = useRouter();
console.log(router);
const submitForm = async (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  await formEl.validate((valid, fields) => {
    if (valid) {
      api.login(ruleForm).then(res=>{
        if(res.data&&res.data.data){
           localStorage.setItem("token", res.data.data.token);
      store.commit("loginSucc",res.data.data.token);
      emits("loginSucc");
        }
        
      })
     
      // router.push('/index')
    } else {
      console.log("error submit!", fields);
    }
  });
};
const subRegister = async (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  await formEl.validate((valid, fields) => {
    if (valid) {
      localStorage.setItem("login", 1);
      store.commit("loginSucc");
      emits("loginSucc");
      // router.push('/index')
    } else {
      console.log("error submit!", fields);
    }
  });
};
const resetForm = (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  formEl.resetFields();
};
const startRemain=()=>{
  if(remainTime.value>0){
    remainTime.value--
      setTimeout(startRemain,1000)
  }else{
    remainTime.value=60
  }
}
const sendCode=()=>{
  const re=/^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$/
  if(re.test(registForm.email)){
      api.sendCode({
        email:registForm.email
      }).then(res=>{
        startRemain()
      })
  }else{
    ElMessage('请输入正确的邮箱')
  }
}
</script>
<style lang="scss" scoped>
.login-page {
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
}
.login-form {
  // widows: 300px;
  border: 1px solid #eee;
  padding: 40px 80px 20px 80px;
  border-radius: 10px;
}
</style>