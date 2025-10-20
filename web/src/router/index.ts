import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/home/HomeView.vue'
import EchoView from '../views/echo/EchoView.vue'
import NotFoundView from '../views/404/NotFoundView.vue'
import { useUserStore } from '@/stores/user'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
      meta: {
        title: 'Home',
        optionalAuth: true,
      },
    },
    {
      path: '/widget',
      name: 'widget',
      component: () => import('../views/widget/WidgetView.vue'),
    },
    {
      path: '/panel',
      name: 'panel',
      component: () => import('../views/panel/PanelView.vue'),
      redirect: '/panel/dashboard',
      meta: {
        requiresAuth: true,
      },
      children: [
        {
          path: 'dashboard',
          name: 'panel-dashboard',
          component: () => import('../views/panel/modules/TheDashboard.vue'),
        },
        {
          path: 'setting',
          name: 'panel-setting',
          component: () => import('../views/panel/modules/TheSetting.vue'),
        },
        {
          path: 'user',
          name: 'panel-user',
          component: () => import('../views/panel/modules/TheUser.vue'),
        },
        {
          path: 'storage',
          name: 'panel-storage',
          component: () => import('../views/panel/modules/TheStorage.vue'),
        },
        {
          path: 'sso',
          name: 'panel-sso',
          component: () => import('../views/panel/modules/TheSSO.vue'),
        },
        {
          path: 'extension',
          name: 'panel-extension',
          component: () => import('../views/panel/modules/TheExtension.vue'),
        },
        {
          path: 'advance',
          name: 'panel-advance',
          component: () => import('../views/panel/modules/TheAdvance.vue'),
        },
      ],
      // beforeEnter: (to, from, next) => {
      //   const userStore = useUserStore()
      //   if (userStore.isLogin) {
      //     next()
      //   } else {
      //     next({ name: 'auth' })
      //   }
      // },
    },
    {
      path: '/auth',
      name: 'auth',
      component: () => import('../views/auth/AuthView.vue'),
    },
    {
      path: '/connect',
      name: 'connect',
      component: () => import('../views/connect/ConnectView.vue'),
    },
    {
      path: '/echo/:echoId',
      name: 'echo',
      component: EchoView,
    },
    // {
    //   path: '/fediverse',
    //   name: 'fediverse',
    //   component: () => import('../views/fediverse/FediverseView.vue'),
    // },
    {
      path: '/:pathMatch(.*)*',
      name: 'not-found',
      component: NotFoundView,
    },
  ],
})

// 全局路由守卫
router.beforeEach(async (to, from, next) => {
  const userStore = useUserStore()

  // 等待用户信息初始化完成
  if (!userStore.initialized) {
    await userStore.init()
  }

  const token = localStorage.getItem('token')
  const needRedirect = localStorage.getItem('needLoginRedirect')

  //  强制鉴权页面或 token 无效
  if (
    (to.meta.requiresAuth && !userStore.isLogin) ||
    (to.meta.optionalAuth && token && !userStore.isLogin && needRedirect === 'true')
  ) {
    localStorage.removeItem('needLoginRedirect')
    localStorage.removeItem('token')
    return next({ name: 'auth' })
  }

  next()
})

export default router
