import Vue from 'vue'
import Vuex from 'vuex'
import createPersistedState from 'vuex-persistedstate'
import Currency from "./Currency"
import Message from "@/store/Message";
Vue.use(Vuex)

const store = new Vuex.Store({
  plugins: [createPersistedState()],
  modules: {
    currency: Currency,
    message: Message
  }
})

export default store
