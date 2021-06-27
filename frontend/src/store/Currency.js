import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default {
    state:{
        currencies: []
    },
    mutations:{
        UPDATE_CURRENCIES: function ( state, payload ){
            state.currencies = payload;
        },
        ADD_CURRENCY: function ( state, payload ){
            state.currencies.push( payload );
        }
    },
    actions:{
        updateCurrencies: async function( context, payload ){
            context.commit("UPDATE_CURRENCIES", payload );
        },
        addCurrency: async function( context, payload ){
            context.commit( "ADD_CURRENCY", payload );
        }
    },
    getters:{
        getCurrencies: ( state ) => {
            return state.currencies;
        }
    }
}

