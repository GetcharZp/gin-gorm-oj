import http from './http.js'

 
export default{
	//  get请求示例
	getProblemList(param){//问题列表
		return http.get(`/problem-list`,param)
	},
	getProblemDetail(param){//问题详情
		return http.get(`/problem-detail`,param)
	},
	getSortList(param){//分类列表
		return http.get(`/category-list`,param)
	},
	sendCode(param){//发送验证码
		return http.postUncode(`/send-code`,param)
	},
	login(param){//登录
		return http.postUncode(`/login`,param)
	},
	// 文件上传
	uploadFile(param){
		return http.upFile('http://mz.xiqurongmei.com/a/material/uploadAll',param)
	},
	
	
 
	 
}
