import Vue from 'vue'
import VueRouter from 'vue-router'
import App from '../App'

Vue.use(VueRouter)

const routes = [
  {
    path: '',
    name: 'Base',
    component: App,
    children:[
      {
        path: '/',
        redirect: '/dashboard'
      },
      {
        name: "Dashboard",
        path: "dashboard",
        component: () => import("../views/Dashboard" )
      },
      {
        name: "Transaction",
        path: "transaction",
        component: () => import("../views/Transaction" )
      }
    ]
  },
]

const router = new VueRouter({
  routes
})

export default router
