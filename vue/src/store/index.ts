import {createStore} from 'vuex'
const store = createStore({
    state:{
        collapse:false,
        isLogin:false,
    },
    mutations:{
        changeCollapse(state,data){
                console.log(data)
                state.collapse=data
        },
        loginSucc(state,data){
            
            state.isLogin=true
    },
    logout(state,data){
        state.isLogin=false

    }
    }

})
export default store