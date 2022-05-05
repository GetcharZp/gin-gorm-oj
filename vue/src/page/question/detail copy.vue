<template>
    <div>
        <div id="monaco" ref="mymonaco"></div>
    </div>
</template>
<script lang="ts" setup>
 import { onMounted } from '@vue/runtime-core'
import * as monaco from 'monaco-editor/esm/vs/editor/editor.api.js'
// import * as monaco from 'monaco-editor/esm/vs/editor/editor.main.js'
//  import * as monaco from 'monaco-editor'
import 'monaco-editor/esm/vs/basic-languages/javascript/javascript.contribution'
// import jsonWorker from 'monaco-editor/esm/vs/language/json/json.worker?worker'
// import editorWorker from 'monaco-editor/esm/vs/editor/editor.worker?worker'
// import 'monaco-editor/esm/vs/editor/contrib/find/findController.js'

import { ref } from 'vue'

const mymonaco=ref<HTMLElement|null>()
onMounted(
  () => {
      // @ts-ignore
// self.MonacoEnvironment = {
//   getWorker: function (_: any, label: string) {
//     console.log(label)
//     if (label === 'json') {
//       return new jsonWorker()
//     }
//     new editorWorker()
//   },
// }

     if(mymonaco.value){
         const monacoInstance=monaco.editor.create(mymonaco.value,{
             value:'console.log("aaa")',
             language:'javascript'
         })
        //  monacoInstance.dispose()
     }
      
        
       
    
  }
)
 
</script>
<style scoped lang="scss">
#monaco{
    width: 500px;
    height: 500px;
}
</style>