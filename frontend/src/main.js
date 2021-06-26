import 'core-js/stable';
import 'regenerator-runtime/runtime';
import '@mdi/font/css/materialdesignicons.css';
import Vue from 'vue';
import Vuetify from 'vuetify';
import 'vuetify/dist/vuetify.min.css';

Vue.use(Vuetify);

import App from './App.vue';

Vue.config.productionTip = false;
Vue.config.devtools = true;

import Wails from '@wailsapp/runtime';
import router from './router'
import '@mdi/font/css/materialdesignicons.min.css'

Wails.Init(() => {
	new Vue({
        vuetify: new Vuetify({
			icons: {
				iconfont: 'mdi'
			},
		}),

        router,
        render: h => h(App)
    }).$mount('#app');
});
