import {createStore} from 'vuex'
const store = createStore({
    state:{
        collapse:false,
        isLogin:false,
        token:null,
        username:null
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
        state.username=data
    },
    logout(state,data){
        state.isLogin=false

    }
    }

})
export default store