<template>
  <div class="ques-list">
    <div class="sort-list">
       
     
      <div>
           问题分类：
        <span @click="getProblem(null)" :class="!actSort ? 'act' : ''"
          >全部</span
        >
        <span
          v-for="item in sortList"
          :key="item.id"
          @click="getProblem(item.identity)"
          :class="actSort == item.identity ? 'act' : ''"
          >{{ item.name }}</span
        >
      </div>
        <el-input
         style="width:200px;margin-right:10px"
      v-model="keyword"
      clearable
      @change="getProblem(actSort)"
      @clear="getProblem(null)"
      size="large"
      placeholder="请搜索"
      :suffix-icon="Search"
    />
    </div>
    <div class="list" v-loading="loading">
      <div class="item" v-for="item in quesList" :key="item.id">
        <div>
          <div class="title">
           <b @click="toDetail(item)" >{{ item.title }}</b> 
            <div class="sort">
              <span v-for="mi in item.problem_categories" :key="mi.id">
               <b v-if="mi.category_basic">{{ mi.category_basic.name }}</b> 
              </span>
            </div>
          </div>
          <div class="msg">
            <span>创建时间：{{ item.created_at }}</span>
            <span>提交人数：{{ item.submit_num }}</span>
            <span>通过人数：{{ item.pass_num }}</span>
          </div>
        </div>
        <div class="edit">
          <el-icon @click="toDetail(item)" title="详情"><edit /></el-icon>
          <!-- <el-icon title="排行" @click="toRank(item)"><histogram /></el-icon> -->
          <el-icon title="提交列表" @click="toSubList(item)"><list /></el-icon>
        </div>
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
import { reactive, ref } from "@vue/reactivity";
import { Edit, Histogram, List } from "@element-plus/icons-vue";
import api from "../../api/api.js";
import { useRouter } from "vue-router";
const router = useRouter();
const loading=ref(false)
const quesList = ref([]);
const sortList = ref([]);
const pageSize=ref(10)
const currentPage=ref(1)
const total=ref(0)
const actSort = ref<null | number>(null);
const keyword=ref('')
const handleSizeChange = (val: number) => {
   getProblem(actSort.value)

}
const handleCurrentChange = (val: number) => {
  getProblem(actSort.value)
}

const getProblem = (sortId: number | null) => {
  actSort.value = sortId;
  loading.value=true
  api
    .getProblemList({
      category_identity: sortId,
      size:pageSize.value,
      page:currentPage.value,
      keyword:keyword.value
    })
    .then((res) => {
        loading.value=false
      if (res && res.data) {
        quesList.value = res.data.data.list;
        total.value=res.data.data.count
      }
      console.log(res);
    });
};
getProblem(null);
api.getSortList({}).then((res) => {
  if (res && res.data) {
    sortList.value = res.data.data.list;
  }
});
const toDetail = (item: any) => {
  router.push({
    path: "/questionDetail",
    query: item,
  });
};
const toRank = (item: any) => {
  router.push({
    path: "/topList",
    query: { identity: item.identity },
  });
};
const toSubList = (item: any) => {
  router.push({
    path: "/submitList",
    query: { identity: item.identity },
  });
};
</script>
<style scoped lang="scss">
.ques-list{
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
.item {
  padding: 20px;
  border-bottom: 1px solid #eee;
  display: flex;
  justify-content: space-between;
  align-items: center;
  .edit {
    width: 80px;
    display: flex;
    justify-content: space-between;
    color: #13b6f9;
    cursor: pointer;
  }
  .title {
    margin-bottom: 20px;
    display: flex;
    align-items: center;
    font-size: 18px;
    b{
        cursor: pointer;
    }
    span {
      background-color: #d3ebff;
      font-size: 12px;
      margin: 0 10px;
      padding: 4px;
      border-radius: 5px;
      border: 1px solid #62a9ff;
      color: #62a9ff;
    }
  }
  .msg {
    font-size: 14px;
    color: #999;
    span {
      margin-right: 20px;
    }
  }
}
.sort-list {
  display: flex;
  align-items: center;
  justify-content: space-between;
//   background-color: #0087bf;
//   color: white;
border-bottom: 1px solid #0087bf;
border-top: 1px solid #0087bf;
  padding: 20px;
  span {
    color: #13b6f9;
    margin-right: 20px;
    cursor: pointer;
    &.act {
      color: black;
    }
  }
}
</style>