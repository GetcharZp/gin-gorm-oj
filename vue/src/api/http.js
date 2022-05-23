/****   http.js   ****/
// 导入封装好的axios实例
import request from './request'
import qs from 'qs'
const http ={
    /**
     * methods: 请求
     * @param url 请求地址 
     * @param params 请求参数
     */
    get(url,params){
        const config = {
            method: 'get',
            url:url
        }
        if(params) config.params = params
        return request(config)
    },
    post(url,params){
        const config = {
            method: 'post',
            url:url
        }
        if(params) config.data = params
        return request(config)
    },
    postJson(url,params){
        const config = {
            method: 'post',
            url:url,
            headers:{
			    'Content-Type': 'multipart/form-data'
			}
        }
        console.log(params)
        if(params) config.data =params
        return request(config)
    },
    postSB(url,params){
		const config = {
		    method: 'post',
		    url:url,
			headers:{
			    'Content-Type': 'application/x-www-form-urlencoded'
			}
		}
		
		if(params) config.data = qs.stringify(params,{arrayFormat:'repeat'})
		
		return request(config)
	},
	postUncode(url,params){
		const config = {
		    method: 'post',
		    url:url,
			headers:{
			    'Content-Type': 'application/x-www-form-urlencoded'
			}
		}
		
		if(params) config.data = qs.stringify(params)
		
		return request(config)
	},
	upFile(url,data){
		return request({
			method: 'post',
			url: url,
			data: data,
			headers:{
					"Content-Type": "multipart/form-data"
				}
		}) 
	},
    put(url,params){
        const config = {
            method: 'put',
            url:url,
            headers:{
			    'Content-Type': 'application/x-www-form-urlencoded'
			}
        }
        
        if(params) config.data = qs.stringify(params)
        return request(config)
    },
    putJson(url,params){
        const config = {
            method: 'put',
            url:url,
            
        }
        
        if(params) config.data =params
        return request(config)
    },
    delete(url,params){
        const config = {
            method: 'delete',
            url:url,
            
        }
        if(params) config.params = params
        return request(config)
    }
}
//导出
export default http