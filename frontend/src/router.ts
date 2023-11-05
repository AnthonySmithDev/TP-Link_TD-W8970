import { createRouter, createWebHistory } from "vue-router";

export default createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: "/",
      name: "login",
      component: () => import("./components/Login.vue"),
    },
    {
      path: "/admin",
      name: "admin",
      component: () => import("./components/Admin.vue"),
      children: [
        {
          path: "/wireless",
          name: "wireless",
          component: () => import("./components/Wireless.vue"),
        },
        {
          path: "/system",
          name: "system",
          component: () => import("./components/System.vue"),
        },
      ],
    },
  ],
});
