
export default {
    state:{
        selectedCurrency: null,
        currencies: []
    },
    mutations:{
        UPDATE_CURRENCIES: function ( state, payload ){
            state.currencies = [];
            for( let item of payload ){
                item.balance = item.balance/100.0;
                item.usd_change = item.usd_change/100.0;
                state.currencies.push( item );
            }
        },
        ADD_CURRENCY: function ( state, payload ){
            payload.balance = payload.balance/100.0;
            payload.usd_change = payload.usd_change/100.0;
            state.currencies.push( payload );
        },
        REMOVE_CURRENCY: function ( state, payload ) {
            for( let index = 0; index < state.currencies.length; index ++ ){
                if( state.currencies[ index ].ID === payload.ID ){
                    state.currencies.splice( index, 1 );
                    return
                }
            }
        },
        UPDATE_BALANCE: function ( state, payload ){
            for( let i = 0; i < state.currencies.length; i++ ){
                if( state.currencies[i].ID === payload.currency_id ){
                    state.currencies[ i ].balance += ( payload.amount / 100.0 );
                    break;
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
        },
        updateBalance: async function( context, payload ){
            context.commit( "UPDATE_BALANCE", payload )
        }
    },
    getters:{
        getCurrencies: ( state ) => {
            return state.currencies;
        },
        getCurrencyFromId: ( state ) => (id)=>{
            return state.currencies.find( thing => thing.ID === id );
        }
    }
}

