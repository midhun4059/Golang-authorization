import { createRouter, createWebHistory } from "vue-router";
// import Login from "@/components/Login.vue";
// import Dashboard from "@/components/Dashboard.vue";
import Login from "../components/Login.vue";
import Dashboard from "../components/Dashboard.vue";
import UserTable from "../components/UserTable.vue";
import Sidebar from "../components/Sidebar.vue";


const routes = [
  { path: "/", component: Login },
 
  {
    path: "/dashboard",
    component: Dashboard,Sidebar,
    
    beforeEnter: (to, from, next) => {
      const token = localStorage.getItem("access_token");
      if (!token) {
        next("/");
      } else {
        next();
      }
    },
  },
  {
    path: "/user-table",
    name: "UserTable",
    component: UserTable,
  },{
    path: "/logout",
    name: "Logout",
    beforeEnter: (to, from, next) => {
      localStorage.clear()    // Clear token on logout
     
   
      next("/"); // Redirect to login page
    },
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;