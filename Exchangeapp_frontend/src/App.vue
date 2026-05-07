<template>
  <el-container>
    <el-header class="header-bar">
      <!-- 左侧菜单 -->
      <el-menu
        :default-active="activeIndex"
        class="el-menu-demo"
        mode="horizontal"
        :ellipsis="true"
        @select="handleSelect"
      >
        <el-menu-item index="home">首页</el-menu-item>
        <el-menu-item index="currencyExchange">兑换货币</el-menu-item>
        <el-menu-item index="news">查看新闻</el-menu-item>
        <el-menu-item index="login" v-if="!authStore.isAuthenticated">登录</el-menu-item>
        <el-menu-item index="register" v-if="!authStore.isAuthenticated">注册</el-menu-item>
        <el-menu-item index="logout" v-if="authStore.isAuthenticated">退出</el-menu-item>
      </el-menu>

      <!-- 右侧用户中心 ↓↓↓ 这是我加的 -->
      <div class="user-center" v-if="authStore.isAuthenticated">
        <el-icon class="user-icon"><User /></el-icon>
        <span class="username">{{ authStore.userInfo?.username }}</span>
      </div>
    </el-header>

    <el-main>
      <router-view></router-view>
    </el-main>
  </el-container>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue';
import { useRouter, useRoute } from 'vue-router';
import { useAuthStore } from './store/auth';
import { User } from '@element-plus/icons-vue'; // 引入用户图标

const router = useRouter();
const route = useRoute();
const authStore = useAuthStore();
const activeIndex = ref(route.name?.toString() || 'home');

watch(route, (newRoute) => {
  activeIndex.value = newRoute.name?.toString() || 'home';
});

const handleSelect = (key: string) => {
  if (key === 'logout') {
    authStore.logout();
    router.push({ name: 'Home' });
  } else {
    router.push({ name: key.charAt(0).toUpperCase() + key.slice(1) });
  }
};
</script>

<style scoped>
/* 头部布局：左侧菜单 + 右侧用户 */
.header-bar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0 20px;
  background: #fff;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.06);
}

.el-menu-demo {
  line-height: 60px;
  border-bottom: none;
  flex: 1;
}

/* 用户中心样式 */
.user-center {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 0 12px;
  height: 36px;
  background: #f5f7fa;
  border-radius: 18px;
  font-size: 14px;
  color: #2c3e50;
}

.user-icon {
  font-size: 16px;
  color: #409eff;
}

.username {
  font-weight: 500;
}
</style>