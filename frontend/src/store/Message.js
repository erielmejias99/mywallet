
export default {
    state:{
        logs: []
    },
    mutations:{
        ADD_LOG: function ( state, payload ) {
            state.logs.push( payload );
        },
        CLEAR_LOGS: function (state){
            state.logs = []
        }
    },
    actions:{
        log: async function( context, payload ){
            context.commit("ADD_LOG", payload );
        },
        clearLogs: async function( context ){
            context.commit("CLEAR_LOGS");
        },
    },
    getters:{
        getLogs: ( state ) => {
            return state.logs;
        },
        lastLog: ( state ) =>{
            return state.logs[ state.logs.length - 1 ];
        },
        areLogs: (state )=>{
            return state.logs.length > 0;
        }
    }
}

