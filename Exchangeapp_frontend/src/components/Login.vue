<template>
  <div class="login-modern">
    <!-- 动态粒子背景 -->
    <div class="particles-bg"></div>
    
    <div class="login-container">
      <div class="login-card">
        <div class="welcome-text">
          <h2>欢迎回来</h2>
          <p>登录您的账户，继续体验</p>
        </div>

        <form class="login-form" @submit.prevent="login">
          <div class="form-group">
            <label class="form-label">
              <span class="label-icon">👤</span>
              用户名
            </label>
            <input
              v-model="form.username"
              type="text"
              placeholder="请输入用户名"
              class="form-input"
              :class="{ 'has-value': form.username }"
            />
          </div>

          <div class="form-group">
            <label class="form-label">
              <span class="label-icon">🔒</span>
              密码
            </label>
            <input
              v-model="form.password"
              type="password"
              placeholder="请输入密码"
              class="form-input"
              :class="{ 'has-value': form.password }"
            />
          </div>

          <button type="submit" class="login-btn" :disabled="isLogining">
            <span class="btn-text">{{ isLogining ? '登录中...' : '立即登录' }}</span>
          </button>
        </form>

        <div class="register-hint">
          还没有账号？
          <span class="register-link" @click="goToRegister">立即注册</span>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { useAuthStore } from '../store/auth';
import { ElMessage } from 'element-plus';

const form = ref({
  username: '',
  password: '',
});

const isLogining = ref(false);
const authStore = useAuthStore();
const router = useRouter();

const login = async () => {
  if (isLogining.value) return;
  
  isLogining.value = true;
  try {
    await authStore.login(form.value.username, form.value.password);
    ElMessage.success('登录成功！');
    router.push({ name: 'News' });
  } catch (error) {
    ElMessage.error('登录失败，请检查用户名和密码。');
  } finally {
    isLogining.value = false;
  }
};

const goToRegister = () => {
  router.push({ name: 'Register' });
};
</script>

<style scoped>
/* ===== 与主页/其他页面同风格 ===== */
.login-modern {
  min-height: 90vh;
  position: relative;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  background-attachment: fixed;
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 20px;
}

/* 粒子背景 */
.particles-bg {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-image: 
    radial-gradient(circle at 10% 20%, rgba(255,255,255,0.1) 2px, transparent 2px),
    radial-gradient(circle at 90% 80%, rgba(255,255,255,0.08) 1px, transparent 1px);
  background-size: 50px 50px, 30px 30px;
  pointer-events: none;
  z-index: 0;
}

.login-container {
  width: 100%;
  max-width: 460px;
  position: relative;
  z-index: 2;
}

/* 登录卡片 */
.login-card {
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(10px);
  border-radius: 32px;
  padding: 48px 40px;
  border: 1px solid rgba(255, 255, 255, 0.3);
  box-shadow: 0 25px 50px rgba(0, 0, 0, 0.2);
}

/* 欢迎文字 */
.welcome-text {
  text-align: center;
  margin-bottom: 32px;
}

.welcome-text h2 {
  font-size: 28px;
  font-weight: 700;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  background-clip: text;
  -webkit-background-clip: text;
  color: transparent;
  margin: 0 0 8px 0;
}

.welcome-text p {
  font-size: 14px;
  color: #64748b;
  margin: 0;
}

/* 表单 */
.login-form {
  width: 90%;
}

.form-group {
  margin-bottom: 24px;
}

.form-label {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 13px;
  font-weight: 600;
  color: #1e293b;
  margin-bottom: 10px;
}

.label-icon {
  font-size: 14px;
}

.form-input {
  width: 100%;
  padding: 14px 18px;
  background: #f8fafc;
  border: 2px solid #e2e8f0;
  border-radius: 16px;
  font-size: 15px;
  transition: all 0.3s ease;
  outline: none;
}

.form-input:focus {
  border-color: #667eea;
  box-shadow: 0 0 0 3px rgba(102, 126, 234, 0.1);
}

.form-input.has-value {
  border-color: #667eea;
}

/* 登录按钮 */
.login-btn {
  width: 110%;
  padding: 16px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border: none;
  border-radius: 20px;
  font-size: 16px;
  font-weight: 600;
  color: white;
  cursor: pointer;
  transition: all 0.3s ease;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 10px;
  margin-top: 8px;
}

.login-btn:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 10px 25px rgba(102, 126, 234, 0.4);
}

.login-btn:disabled {
  opacity: 0.7;
  cursor: not-allowed;
}
.login-btn:hover:not(:disabled) .btn-arrow {
  transform: translateX(4px);
}

/* 注册提示 */
.register-hint {
  text-align: center;
  margin-top: 28px;
  font-size: 14px;
  color: #64748b;
}

.register-link {
  color: #667eea;
  font-weight: 600;
  cursor: pointer;
  margin-left: 6px;
  transition: color 0.2s ease;
}

.register-link:hover {
  color: #764ba2;
  text-decoration: underline;
}

/* 响应式适配 */
@media (max-width: 500px) {
  .login-card {
    padding: 32px 24px;
  }
  
  .logo-text {
    font-size: 24px;
  }
  
  .welcome-text h2 {
    font-size: 24px;
  }
  
  .form-input {
    padding: 12px 16px;
  }
  
  .login-btn {
    padding: 14px;
  }
}
</style>