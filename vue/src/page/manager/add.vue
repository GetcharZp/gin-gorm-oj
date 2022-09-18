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
      <el-input v-model="ruleForm.title"/>
    </el-form-item>
    <el-form-item label="内容" prop="content ">
      <el-input v-model="ruleForm.content" type="textarea"/>
    </el-form-item>
    <el-form-item label="分类" prop="problem_categories">
      <el-select
          v-model="ruleForm.problem_categories"
          placeholder="请选择"
          multiple
          style="width: 100%"
      >
        <el-option
            :label="mi.name"
            :value="mi.id"
            v-for="mi in sortList"
            :key="mi.id"
        />
      </el-select>
    </el-form-item>
    <el-form-item label="最大运行时间" prop="max_runtime">
      <el-input v-model.number="ruleForm.max_runtime"/>
    </el-form-item>
    <el-form-item label="最大占用内存" prop="max_mem">
      <el-input v-model.number="ruleForm.max_mem"/>
    </el-form-item>

    <el-form-item label="测试案例" prop="test_cases">
      <el-form-item
          label-width="0"
          v-for="(mi, i) in ruleForm.test_cases"
          :key="i"
      >
        <el-row :gutter="20">
          <el-col :span="4" style="text-align: right">输入：</el-col>
          <el-col :span="6">
            <el-input type="textarea" v-model="mi.input" rows="6" cols="40"/>
          </el-col>
          <el-col :span="4" style="text-align: right">输出：</el-col>
          <el-col :span="6">
            <el-input type="textarea" v-model="mi.output" rows="6" cols="40"/>
          </el-col>
          <el-col :span="4">
            <el-icon @click="addCase">
              <circle-plus/>
            </el-icon>
            <el-icon @click="removeCase(i)">
              <remove/>
            </el-icon
            >
          </el-col>
        </el-row>
      </el-form-item>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="submitForm(ruleFormRef)"
      >创建
      </el-button
      >
      <el-button @click="closeBox">取消</el-button>
    </el-form-item>
  </el-form>
</template>

<script lang="ts" setup>
import {reactive, ref} from "vue";
import type {FormInstance, FormRules} from "element-plus";
import {ElMessage} from "element-plus";
import {CirclePlus, Remove} from "@element-plus/icons-vue";
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
  max_runtime: 0,
  max_mem: 0,
  test_cases: [{input: "", output: ""}],
  problem_categories: [],
});
const emits = defineEmits(["cancel", "submit"]);
const props = defineProps(["sortList", "selectQues"]);
if (props.selectQues) {
  let se = props.selectQues;
  if (se.problem_categories instanceof Array) {
    se.problem_categories = se.problem_categories.map((it: any) => {
      return it.category_id;
    });
  }
  if (!Array.isArray(se.test_cases) || se.test_cases.length == 0) {
    se.test_cases = [{input: "", output: ""}];
  }
  ruleForm.value = se;
}
const closeBox = () => {
  emits("cancel");
};
const rules = reactive<FormRules>({
  title: [{required: true, message: "请输入", trigger: "blur"}],
  content: [{required: true, message: "请输入", trigger: "blur"}],
  max_runtime: [
    {required: true, message: "请输入", trigger: "blur"},
    {type: "number", message: "请输入数字", trigger: "blur"},
  ],
  max_mem: [
    {required: true, message: "请输入", trigger: "blur"},
    {type: "number", message: "请输入数字", trigger: "blur"},
  ],
  test_cases: [{required: true, message: "请输入", trigger: "blur"}],
  problem_categories: [
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
const removeCase = (i: number) => {
  if (ruleForm.value.test_cases.length > 1) {
    ruleForm.value.test_cases.splice(i, 1);
  }
};
const submitForm = async (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  await formEl.validate((valid, fields) => {
    if (valid) {
      if (ruleForm.value.identity) {
        api.editProblem(ruleForm.value).then((res) => {
          if (res.data.code == 200) {
            ElMessage.success("成功");
            emits("submit");
          } else {
            ElMessage(res.data.msg);
          }
        });
      } else {
        api.addProblem(ruleForm.value).then((res: any) => {
          if (res.data.code == 200) {
            ElMessage.success("成功");
            emits("submit");
          } else {
            ElMessage(res.data.msg);
          }
        });
      }
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