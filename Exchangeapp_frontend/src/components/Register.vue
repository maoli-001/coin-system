<template>
  <div class="register-page">
    <div class="register-card">
      <div class="card-header">
        <h2>📊 汇率兑换 & 金融资讯</h2>
        <p>创建您的账户，开始使用</p>
      </div>

      <el-form :model="form" class="register-form" @submit.prevent="register">
        <el-form-item label="用户名" label-width="80px">
          <el-input
            v-model="form.username"
            placeholder="请输入用户名"
            size="default"
            style="width: 100%"
          />
        </el-form-item>

        <el-form-item label="密码" label-width="80px">
          <el-input
            v-model="form.password"
            type="password"
            placeholder="请输入密码"
            size="default"
            style="width: 100%"
          />
        </el-form-item>

        <el-form-item class="center-btn">
          <el-button
            type="primary"
            native-type="submit"
            class="submit-btn"
            size="default"
            block
          >
            立即注册
          </el-button>
        </el-form-item>
      </el-form>
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

const authStore = useAuthStore();
const router = useRouter();

const register = async () => {
  try {
    await authStore.register(form.value.username, form.value.password);
    router.push({ name: 'News' });
  } catch {
    ElMessage.error('注册失败，请重试。');
  }
};
</script>

<style scoped>
/* 页面整体 */
.register-page {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 80vh;
  background: #f7f8fa;
  padding: 20px;
}

/* 注册卡片 */
.register-card {
  width: 100%;
  max-width: 420px;
  background: #ffffff;
  border-radius: 16px;
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.08);
  padding: 36px;
}

/* 标题 */
.card-header {
  text-align: center;
  margin-bottom: 32px;
}

.card-header h2 {
  font-size: 24px;
  font-weight: 600;
  color: #2c3e50;
  margin: 0 0 8px 0;
}

.card-header p {
  font-size: 14px;
  color: #7f8c8d;
  margin: 0;
}

/* 表单 */
.register-form {
  width: 100%;
}

/* 按钮 */
.center-btn {
  margin-top: 12px;
}

.submit-btn {
  height: 44px;
  width:425px;
  font-size: 16px;
  border-radius: 10px;
  background: linear-gradient(135deg, #409eff 0%, #69b1ff 100%);
  border: none;
}

/* 输入框美化 */
:deep(.el-input__wrapper) {
  border-radius: 10px !important;
  box-shadow: none;
  border: 1px solid #e4e7ed;
}

:deep(.el-input__wrapper.is-focus) {
  border-color: #409eff;
  box-shadow: 0 0 0 2px rgba(64, 158, 255, 0.1);
}
</style>