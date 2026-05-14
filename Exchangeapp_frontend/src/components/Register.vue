<template>
  <div class="register-modern">
    <!-- 动态粒子背景 -->
    <div class="particles-bg"></div>
    
    <div class="register-container">
      <div class="register-card">
        <div class="welcome-text">
          <h2>创建账户</h2>
          <p>加入蓝鼠社区，开启您的金融之旅</p>
        </div>

        <form class="register-form" @submit.prevent="register">
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

          <button type="submit" class="register-btn" :disabled="isRegistering">
            <span class="btn-text">{{ isRegistering ? '注册中...' : '立即注册' }}</span>
          </button>
        </form>

        <div class="login-hint">
          已有账号？
          <span class="login-link" @click="goToLogin">立即登录</span>
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

const isRegistering = ref(false);
const authStore = useAuthStore();
const router = useRouter();

const register = async () => {
  if (isRegistering.value) return;
  
  // 简单的前端验证
  if (!form.value.username.trim()) {
    ElMessage.warning('请输入用户名');
    return;
  }
  if (form.value.password.length < 6) {
    ElMessage.warning('密码长度不能少于6位');
    return;
  }
  
  isRegistering.value = true;
  try {
    await authStore.register(form.value.username, form.value.password);
    ElMessage.success('注册成功！正在跳转...');
    router.push({ name: 'News' });
  } catch (error: any) {
    if (error.response?.status === 409) {
      ElMessage.error('用户名已存在，请换一个试试');
    } else {
      ElMessage.error('注册失败，请重试。');
    }
  } finally {
    isRegistering.value = false;
  }
};

const goToLogin = () => {
  router.push({ name: 'Login' });
};
</script>

<style scoped>
/* ===== 与登录页面同风格 ===== */
.register-modern {
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

.register-container {
  width: 100%;
  max-width: 460px;
  position: relative;
  z-index: 2;
}

/* 注册卡片 */
.register-card {
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
.register-form {
  width: 100%;
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
  width: 90%;
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

/* 注册按钮 */
.register-btn {
  width: 100%;
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

.register-btn:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 10px 25px rgba(102, 126, 234, 0.4);
}

.register-btn:disabled {
  opacity: 0.7;
  cursor: not-allowed;
}

.register-btn:hover:not(:disabled) .btn-arrow {
  transform: translateX(4px);
}

/* 登录提示 */
.login-hint {
  text-align: center;
  margin-top: 28px;
  font-size: 14px;
  color: #64748b;
}

.login-link {
  color: #667eea;
  font-weight: 600;
  cursor: pointer;
  margin-left: 6px;
  transition: color 0.2s ease;
}

.login-link:hover {
  color: #764ba2;
  text-decoration: underline;
}

/* 响应式适配 */
@media (max-width: 500px) {
  .register-card {
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
  
  .register-btn {
    padding: 14px;
  }
}
</style>