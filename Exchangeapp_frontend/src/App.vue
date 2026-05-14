<template>
  <div class="app-layout">
    <!-- 现代化导航栏 -->
    <nav class="modern-nav">
      <div class="nav-container">
        <!-- 左侧 Logo -->
        <div class="logo-area" @click="goToHome">
          <span class="logo-text">蓝鼠兑换</span>
        </div>

        <!-- 中间导航菜单 -->
        <div class="nav-menu">
          <div 
            class="nav-item"
            :class="{ active: activeIndex === 'Home' }"
            @click="navigateTo('Home')"
          >
            <span class="nav-icon">🏠</span>
            <span>首页</span>
          </div>
          <div 
            class="nav-item"
            :class="{ active: activeIndex === 'CurrencyExchange' }"
            @click="navigateTo('CurrencyExchange')"
          >
            <span class="nav-icon">💱</span>
            <span>兑换货币</span>
          </div>
          <div 
            class="nav-item"
            :class="{ active: activeIndex === 'News' }"
            @click="navigateTo('News')"
          >
            <span class="nav-icon">📰</span>
            <span>查看新闻</span>
          </div>
          <div 
            v-if="!authStore.isAuthenticated"
            class="nav-item"
            :class="{ active: activeIndex === 'Login' }"
            @click="navigateTo('Login')"
          >
            <span class="nav-icon">🔐</span>
            <span>登录</span>
          </div>
          <div 
            v-if="!authStore.isAuthenticated"
            class="nav-item"
            :class="{ active: activeIndex === 'Register' }"
            @click="navigateTo('Register')"
          >
            <span class="nav-icon">📝</span>
            <span>注册</span>
          </div>
          <div 
            v-if="authStore.isAuthenticated"
            class="nav-item logout-item"
            @click="handleLogout"
          >
            <span class="nav-icon">🚪</span>
            <span>退出</span>
          </div>
        </div>

        <!-- 右侧用户中心 -->
        <div class="user-center" v-if="authStore.isAuthenticated" @click="navigateTo('UserCenter')">
          <div class="user-avatar">
            {{ authStore.userInfo?.username?.charAt(0).toUpperCase() || 'U' }}
          </div>
          <div class="user-info">
            <span class="user-name">{{ authStore.userInfo?.username }}</span>
            <span class="user-role">普通会员</span>
          </div>
        </div>

        <!-- 未登录显示占位 -->
        <div class="user-placeholder" v-else></div>
      </div>
    </nav>

    <!-- 主内容区域 -->
    <main class="main-content">
      <router-view></router-view>
    </main>
  </div>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue';
import { useRouter, useRoute } from 'vue-router';
import { useAuthStore } from './store/auth';

const router = useRouter();
const route = useRoute();
const authStore = useAuthStore();

// 当前激活的菜单项（与路由名称完全一致）
const activeIndex = ref('Home');

// 更新激活状态
const updateActiveIndex = () => {
  const routeName = route.name?.toString() || 'Home';
  activeIndex.value = routeName;
};

// 监听路由变化
watch(route, () => {
  updateActiveIndex();
}, { immediate: true });

// 导航跳转
const navigateTo = (routeName: string) => {
  router.push({ name: routeName });
};

// 退出登录
const handleLogout = () => {
  authStore.logout();
  router.push({ name: 'Home' });
};

// 点击 Logo 回到首页
const goToHome = () => {
  router.push({ name: 'Home' });
};
</script>

<style scoped>
/* ===== 现代化导航栏布局 ===== */
.app-layout {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}

/* 导航栏 */
.modern-nav {
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(10px);
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.08);
  position: sticky;
  top: 0;
  z-index: 1000;
  border-bottom: 1px solid rgba(102, 126, 234, 0.1);
}

.nav-container {
  max-width: 1400px;
  margin: 0 auto;
  padding: 0 24px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 64px;
}

/* Logo 区域 */
.logo-area {
  display: flex;
  align-items: center;
  gap: 10px;
  cursor: pointer;
  transition: transform 0.2s ease;
}

.logo-area:hover {
  transform: scale(1.02);
}

.logo-text {
  font-size: 20px;
  font-weight: 700;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  background-clip: text;
  -webkit-background-clip: text;
  color: transparent;
}

/* 导航菜单 */
.nav-menu {
  display: flex;
  align-items: center;
  gap: 8px;
  background: #f1f5f9;
  padding: 6px;
  border-radius: 50px;
}

.nav-item {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 8px 20px;
  border-radius: 40px;
  font-size: 14px;
  font-weight: 500;
  color: #475569;
  cursor: pointer;
  transition: all 0.3s ease;
}

.nav-icon {
  font-size: 16px;
}

.nav-item:hover {
  background: rgba(102, 126, 234, 0.1);
  color: #667eea;
}

.nav-item.active {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  box-shadow: 0 4px 12px rgba(102, 126, 234, 0.3);
}

.nav-item.active .nav-icon {
  filter: drop-shadow(0 0 2px rgba(255,255,255,0.5));
}

.logout-item {
  background: rgba(239, 68, 68, 0.1);
  color: #ef4444;
}

.logout-item:hover {
  background: rgba(239, 68, 68, 0.2);
  color: #dc2626;
}

/* 用户中心（已登录） */
.user-center {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 6px 16px 6px 10px;
  background: linear-gradient(135deg, #f8fafc 0%, #f1f5f9 100%);
  border-radius: 50px;
  cursor: pointer;
  transition: all 0.3s ease;
  border: 1px solid rgba(102, 126, 234, 0.2);
}

.user-center:hover {
  background: linear-gradient(135deg, #f1f5f9 0%, #e2e8f0 100%);
  transform: translateY(-1px);
}

.user-avatar {
  width: 36px;
  height: 36px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 16px;
  font-weight: 600;
  color: white;
}

.user-info {
  display: flex;
  flex-direction: column;
}

.user-name {
  font-size: 14px;
  font-weight: 600;
  color: #1e293b;
}

.user-role {
  font-size: 10px;
  color: #94a3b8;
}

/* 未登录占位（保持布局平衡） */
.user-placeholder {
  width: 100px;
}

/* 主内容区域 */
.main-content {
  flex: 1;
}

/* 响应式适配 */
@media (max-width: 900px) {
  .nav-container {
    padding: 0 16px;
  }
  
  .nav-item span:last-child {
    display: none;
  }
  
  .nav-item {
    padding: 8px 14px;
  }
  
  .nav-icon {
    font-size: 20px;
  }
  
  .logo-text {
    display: none;
  }
  
  .user-info {
    display: none;
  }
  
  .user-center {
    padding: 6px 10px;
  }
  
  .user-placeholder {
    width: 50px;
  }
}

@media (max-width: 600px) {
  .nav-menu {
    gap: 4px;
  }
  
  .nav-item {
    padding: 6px 10px;
  }
  
  .modern-nav .nav-container {
    height: 56px;
  }
}
</style>