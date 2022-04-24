import http from './http.js'

 
export default{
	//  get请求示例
	getProblemList(param){//问题列表
		return http.get(`/problem-list`,param)
	},
	getSortList(param){//分类列表
		return http.get(`/admin/category-list`,param)
	},
	// 文件上传
	uploadFile(param){
		return http.upFile('http://mz.xiqurongmei.com/a/material/uploadAll',param)
	},
	
	
 
	 
}
