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
	getRankList(param){//排行榜
		return http.get(`/rank-list`,param)
	},
	getSubmitList(param){//提交列表
		return http.get(`/submit-list`,param)
	},
	sendCode(param){//发送验证码
		return http.postUncode(`/send-code`,param)
	},
	login(param){//登录
		return http.postUncode(`/login`,param)
	},
	register(param){//注册
		return http.postUncode(`/register`,param)
	},
	delSort(param){//删除
		return http.delete(`/admin/category-delete`,param)
	},
	addSort(param){//分类创建
		return http.postUncode(`/admin/category-create`,param)
	},
	addProblem(param){//问题创建
		return http.post(`/admin/problem-create`,param)
	},
	editProblem(param){//问题编辑
		return http.putJson(`/admin/problem-modify`,param)
	},
	editSort(param){//分类编辑
		return http.put(`/admin/category-modify`,param)
	},
	submitCode(param,id){//提交代码
		return http.postJson(`/user/submit?problem_identity=${id}`,param)
	},
	// 文件上传
	uploadFile(param){
		return http.upFile('',param)
	},
	
	
 
	 
}
