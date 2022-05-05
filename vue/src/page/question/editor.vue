<template>
  <div>
   
    <codemirror
      ref="cm"
      v-model="code"
      :options="cmOptions"
      @input="inputChange"
    ></codemirror>
  </div>
</template>
<script lang="ts" setup>
//
import {codemirror} from 'vue-codemirror';
// 全局引入vue-codemirror
// 引入主题 可以从 codemirror/theme/ 下引入多个
 // 引入css文件
  import 'codemirror/lib/codemirror.css'
import 'codemirror/theme/idea.css'
// 引入语言模式 可以从 codemirror/mode/ 下引入多个
import 'codemirror/mode/sql/sql.js';
import { ref } from 'vue';
const cm=ref()
const code= ref('select a from table1 where b = 1')
const cmOptions= ref({
          // 语言及语法模式
          mode: 'text/x-sql',
          // 主题
          theme: 'idea',
          // 显示函数
          line: true,
          lineNumbers: true,
          // 软换行
          lineWrapping: true,
          // tab宽度
          tabSize: 4,
        })
const inputChange=(content:any)=>{
        
          console.log("code:" + code);
          console.log("content:" + content)
        
      }
</script>
<style scoped lang="scss">
#monaco {
  width: 500px;
  height: 500px;
}
</style>