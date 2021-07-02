
export default {
    state:{
        selectedCurrency: null,
        currencies: []
    },
    mutations:{
        UPDATE_CURRENCIES: function ( state, payload ){
            state.currencies = [];
            for( var item of payload ){
                item.balance = item.balance/100.0;
                item.usd_change = item.usd_change/100.0;
                state.currencies.push( item );
            }
        },
        ADD_CURRENCY: function ( state, payload ){
            state.currencies.push( payload );
        },
        REMOVE_CURRENCY: function ( state, payload ) {
            for( let index = 0; index < state.currencies.length; index ++ ){
                if( state.currencies[ index ].ID === payload.ID ){
                    state.currencies.splice( index, 1 );
                    return
                }
            }
        }
    },
    actions:{
        updateCurrencies: async function( context, payload ){
            context.commit("UPDATE_CURRENCIES", payload );
        },
        addCurrency: async function( context, payload ){
            context.commit( "ADD_CURRENCY", payload );
        },
        removeCurrency: async function( context, payload ){
            context.commit( "REMOVE_CURRENCY", payload );
        }
    },
    getters:{
        getCurrencies: ( state ) => {
            return state.currencies;
        }
    }
}

