<template>
  <el-form
    ref="ruleFormRef"
    :model="ruleForm"
    :rules="rules"
    label-width="120px"
    class="demo-ruleForm"
    :size="formSize"
  >
    <el-form-item label="问题名称" prop="title">
      <el-input v-model="ruleForm.title" />
    </el-form-item>
    <el-form-item label="内容" prop="content ">
      <el-input v-model="ruleForm.content" type="textarea" />
    </el-form-item>
    <el-form-item label="分类" prop="category_ids">
      <el-select
        v-model="ruleForm.category_ids"
        placeholder="请选择"
        multiple
        style="width: 100%"
      >
        <el-option
          :label="mi.name"
          :value="mi.identity"
          v-for="mi in sortList"
          :key="mi.id"
        />
      </el-select>
    </el-form-item>
    <el-form-item label="最大运行时间" prop="max_runtime">
      <el-input v-model="ruleForm.max_runtime" />
    </el-form-item>
    <el-form-item label="最大占用内存" prop="max_mem">
      <el-input v-model="ruleForm.max_mem" />
    </el-form-item>

    <el-form-item
      label="测试案例"
      prop="test_cases"
      v-for="(mi, i) in ruleForm.test_cases"
      :key="i"
    >
      <el-row :gutter="20">
        <el-col :span="4" style="text-align:right;">输入：</el-col>
        <el-col :span="6"><el-input v-model="mi.input" /></el-col>
        <el-col :span="4" style="text-align:right;">输出：</el-col>
        <el-col :span="6"> <el-input v-model="mi.output" /></el-col>
        <el-col :span="4"> 
          <el-icon @click="addCase"><circle-plus /></el-icon> 
           <el-icon @click="removeCase(i)"><remove /></el-icon
        ></el-col>
      </el-row>
    </el-form-item>
    <el-form-item  >
      <el-button type="primary" @click="submitForm(ruleFormRef)"
        >创建</el-button
      >
      <el-button @click="resetForm(ruleFormRef)">取消</el-button>
    </el-form-item>
  </el-form>
</template>

<script lang="ts" setup>
import { reactive, ref } from "vue";
import type { FormInstance, FormRules } from "element-plus";
import { CirclePlus, Remove } from "@element-plus/icons-vue";
import api from "../../api/api.js";
const formSize = ref("default");
const ruleFormRef = ref<FormInstance>();
// interface formBase{
//    title :string,
//   content : '',
//   max_runtime : '',
//   max_mem : '',
//   test_cases : [],
//   category_ids: [],
// }
const ruleForm = ref({
  title: "",
  content: "",
  max_runtime: "",
  max_mem: "",
  test_cases: [{ input: "", output: "" }],
  category_ids: [],
});
defineProps(["sortList"]);
const rules = reactive<FormRules>({
  title: [{ required: true, message: "请输入", trigger: "blur" }],
  content: [{ required: true, message: "请输入", trigger: "blur" }],
  max_runtime: [{ required: true, message: "请输入", trigger: "blur" }],
  max_mem: [{ required: true, message: "请输入", trigger: "blur" }],
  test_cases: [{ required: true, message: "请输入", trigger: "blur" }],
  category_ids: [
    {
      type: "array",
      required: true,
      message: "Please select at least one activity type",
      trigger: "change",
    },
  ],
});
const addCase = () => {
  ruleForm.value.test_cases.push({
    input: "",
    output: "",
  });
};
const removeCase = (i:number) => {
  if (ruleForm.value.test_cases.length > 1) {
    ruleForm.value.test_cases.splice(i,1);
  }
};
const submitForm = async (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  await formEl.validate((valid, fields) => {
    if (valid) {
      api.addProblem(ruleForm.value);
    } else {
      console.log("error submit!", fields);
    }
  });
};

const resetForm = (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  formEl.resetFields();
};
</script>