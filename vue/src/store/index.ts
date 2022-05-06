import {createStore} from 'vuex'
const store = createStore({
    state:{
        collapse:false,
        isLogin:false,
        token:null,
        username:null,
        is_admin:0
    },
    mutations:{
        changeCollapse(state,data){
                console.log(data)
                state.collapse=data
        },
        loginSucc(state,data){
            state.token=data
            state.isLogin=true
            
    },
    setUser(state,data){
        state.username=data.username
        state.is_admin=data.is_admin
    },
    logout(state,data){
        state.isLogin=false

    }
    }

})
export default store