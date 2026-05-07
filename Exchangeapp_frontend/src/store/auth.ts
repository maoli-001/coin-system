import { defineStore } from 'pinia';
import { ref, computed } from 'vue';
import axios from '../axios';

interface UserInfo {
  username: string;
}

export const useAuthStore = defineStore('auth', () => {
  // 初始化 token
  const token = ref<string | null>(localStorage.getItem('token'));

  // 初始化用户信息
  const userInfo = ref<UserInfo | null>(null);

  // 计算属性
  const isAuthenticated = computed(() => !!token.value);

  // ========== 初始化时从 localStorage 恢复用户信息 ==========
  const savedUsername = localStorage.getItem('username');
  if (token.value && savedUsername) {
    userInfo.value = { username: savedUsername };
  }

  // 登录
  const login = async (username: string, password: string) => {
    try {
      const response = await axios.post('/auth/login', { username, password });
      token.value = response.data.token;

      // 存用户名
      userInfo.value = { username };
      localStorage.setItem('token', token.value);
      localStorage.setItem('username', username); // 关键

      // 设置请求头
      axios.defaults.headers.common['Authorization'] = `Bearer ${token.value}`;
    } catch (error) {
      throw new Error('登录失败');
    }
  };

  // 注册
  const register = async (username: string, password: string) => {
    try {
      const response = await axios.post('/auth/register', { username, password });
      token.value = response.data.token;

      // 存用户名
      userInfo.value = { username };
      localStorage.setItem('token', token.value);
      localStorage.setItem('username', username); // 关键

      axios.defaults.headers.common['Authorization'] = `Bearer ${token.value}`;
    } catch (error) {
      throw new Error('注册失败');
    }
  };

  // 退出登录
  const logout = () => {
    token.value = null;
    userInfo.value = null;

    localStorage.removeItem('token');
    localStorage.removeItem('username'); 
    delete axios.defaults.headers.common['Authorization'];
  };

  return {
    token,
    userInfo,
    isAuthenticated,
    login,
    register,
    logout,
  };
});