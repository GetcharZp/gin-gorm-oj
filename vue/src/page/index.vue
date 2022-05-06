<script lang="ts" setup>

import vSlideBar from "../components/SliderBar.vue";
import vHistoryTag from "../components/historyTag.vue";
</script>

<template>
  <div class="page">
    <v-slide-bar></v-slide-bar>
    <div class="content">
      <div class="top-bar">
       
        <v-history-tag></v-history-tag>
      </div>

      <div class="page-box">
        <div class="page-cont">
            <router-view v-slot="{ Component }">
                   
                          <transition name="component-fade" mode="out-in">
                            <component :is="Component" />
                             </transition>
                        </router-view>
         
          
            
         
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped lang="scss">
.component-fade-enter-active,
.component-fade-leave-active {
  transition: opacity 0.3s ease;
}

.component-fade-enter-from,
.component-fade-leave-to {
  opacity: 0;
}

.page {
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;

  .content {
    flex: 1;
    overflow: hidden;
    .top-bar {
      // height: 80px;
      box-shadow: 0 10px 10px #eee;
    }
    .header {
      height: 40px;
      border-bottom: 1px solid #eee;
      display: flex;
      justify-content: space-between;
      align-items: center;
    }
    .page-box {
      // height: calc(100% - 80px);
      height: 100%;
      padding: 10px;
      box-sizing: border-box;
      background-color: #f5f5f5;
      .page-cont {
        box-shadow: 0 0 10px #eee;
        width: 100%;
        height: 100%;
        overflow-y: auto;
        overflow-x: hidden;
        background-color: white;
        .comp-box {
          padding: 10px;
           
          height: 100%;
          box-sizing: border-box;
        }
      }
    }
  }
}
</style>
