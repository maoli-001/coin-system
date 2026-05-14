<template>
  <div class="user-center-modern">
    <!-- 动态粒子背景 -->
    <div class="particles-bg"></div>
    
    <div class="center-container">
      <!-- 用户信息卡片 -->
      <div class="user-profile-card">
        <div class="profile-header">
          <div class="user-avatar-large">
            {{ authStore.userInfo?.username?.charAt(0).toUpperCase() || 'U' }}
          </div>
          <h2 class="user-name">{{ authStore.userInfo?.username || '用户' }}</h2>
        </div>
        
        <div class="profile-stats">
          <div class="stat-item">
            <div class="stat-value">{{ likedArticles.length }}</div>
            <div class="stat-label">点赞的文章</div>
          </div>
        </div>
      </div>

      <!-- 点赞的文章列表 -->
      <div class="liked-articles-section">
        <div class="section-header">
          <h3>❤️ 我点赞的文章</h3>
          <span class="article-count">{{ likedArticles.length }} 篇</span>
        </div>

        <div v-if="loading" class="loading-state">
          <div class="loading-spinner"></div>
          <p>加载中...</p>
        </div>

        <div v-else-if="likedArticles.length === 0" class="empty-state">
          <div class="empty-icon">💔</div>
          <h4>暂无点赞记录</h4>
          <p>您还没有点赞过任何文章，快去浏览新闻吧！</p>
          <button class="browse-btn" @click="goToNews">浏览新闻</button>
        </div>

        <div v-else class="articles-list">
          <div 
            v-for="article in likedArticles" 
            :key="article.ID"
            class="article-item"
            @click="viewArticle(article.ID)"
          >
            <div class="article-info">
              <h4 class="article-title">{{ article.Title }}</h4>
              <p class="article-preview">{{ getPreview(article.Preview) }}</p>
              <div class="article-meta">
                <span class="meta-date">📅 {{ formatDate(article.CreatedAt) }}</span>
              </div>
            </div>
            <div class="article-arrow">
              <span>→</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue';
import { useRouter } from 'vue-router';
import { useAuthStore } from '../store/auth';
import axios from '../axios';
import { ElMessage } from 'element-plus';

const router = useRouter();
const authStore = useAuthStore();
const likedArticles = ref<any[]>([]);
const loading = ref(true);
const totalLikes = ref(0);

// 获取用户点赞的文章列表
const fetchLikedArticles = async () => {
  loading.value = true;
  try {
    const response = await axios.get('/user/likedArticles');
    console.log('点赞文章列表响应:', response.data);
    
    likedArticles.value = response.data.articles || [];
    
    // 计算累计点赞数（所有文章点赞数之和）
    totalLikes.value = likedArticles.value.reduce((sum, article) => {
      return sum + (article.Likes || 0);
    }, 0);
    
    console.log('点赞文章列表:', likedArticles.value);
  } catch (error: any) {
    console.error('获取点赞文章失败:', error);
    if (error.response?.status === 401) {
      ElMessage.error('请先登录');
      router.push({ name: 'Login' });
    } else {
      ElMessage.error('加载失败，请稍后重试');
    }
  } finally {
    loading.value = false;
  }
};

// 查看文章详情
const viewArticle = (articleId: string) => {
  router.push({ name: 'NewsDetail', params: { id: articleId } });
};

// 浏览新闻
const goToNews = () => {
  router.push({ name: 'News' });
};

// 获取文章预览（截取前100字）
const getPreview = (content: string) => {
  if (!content) return '暂无内容';
  return content.length > 100 ? content.substring(0, 100) + '...' : content;
};

// 格式化日期
const formatDate = (dateStr?: string) => {
  if (!dateStr) return '未知日期';
  const date = new Date(dateStr);
  const year = date.getFullYear();
  const month = String(date.getMonth() + 1).padStart(2, '0');
  const day = String(date.getDate()).padStart(2, '0');
  return `${year}.${month}.${day}`;
};

onMounted(() => {
  // 检查是否登录
  if (!authStore.isAuthenticated) {
    ElMessage.warning('请先登录');
    router.push({ name: 'Login' });
    return;
  }
  fetchLikedArticles();
});
</script>

<style scoped>
.user-center-modern {
  min-height: 100vh;
  position: relative;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  background-attachment: fixed;
  padding: 60px 20px;
}

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

.center-container {
  max-width: 1000px;
  margin: 0 auto;
  position: relative;
  z-index: 2;
}

/* 用户信息卡片 */
.user-profile-card {
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(10px);
  border-radius: 32px;
  padding: 40px;
  text-align: center;
  margin-bottom: 32px;
  border: 1px solid rgba(255, 255, 255, 0.3);
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.1);
}

.profile-header {
  margin-bottom: 24px;
}

.user-avatar-large {
  width: 100px;
  height: 100px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 42px;
  font-weight: 600;
  color: white;
  margin: 0 auto 20px;
  box-shadow: 0 8px 20px rgba(102, 126, 234, 0.3);
}

.user-name {
  font-size: 28px;
  font-weight: 700;
  color: #1e293b;
  margin: 0 0 8px 0;
}

.profile-stats {
  display: flex;
  justify-content: center;
  align-items: center;
  padding-top: 24px;
  border-top: 1px solid #e2e8f0;
  gap: 40px;
}

.stat-value {
  font-size: 32px;
  font-weight: 800;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  background-clip: text;
  -webkit-background-clip: text;
  color: transparent;
  line-height: 1.2;
}

.stat-label {
  font-size: 13px;
  color: #64748b;
  margin-top: 6px;
}

/* 点赞文章区域 */
.liked-articles-section {
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(10px);
  border-radius: 32px;
  padding: 32px;
  border: 1px solid rgba(255, 255, 255, 0.3);
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.1);
}

.section-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;
  padding-bottom: 16px;
  border-bottom: 1px solid #e2e8f0;
}

.section-header h3 {
  font-size: 20px;
  font-weight: 600;
  color: #1e293b;
  margin: 0;
}

.article-count {
  font-size: 14px;
  color: #64748b;
  background: #f1f5f9;
  padding: 4px 12px;
  border-radius: 20px;
}

/* 文章列表 */
.articles-list {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.article-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 20px;
  background: #f8fafc;
  border-radius: 20px;
  cursor: pointer;
  transition: all 0.3s ease;
  border: 1px solid #e2e8f0;
}

.article-item:hover {
  transform: translateX(6px);
  background: white;
  border-color: #667eea;
  box-shadow: 0 6px 16px rgba(102, 126, 234, 0.12);
}

.article-info {
  flex: 1;
}

.article-title {
  font-size: 18px;
  font-weight: 600;
  color: #1e293b;
  margin: 0 0 8px 0;
  line-height: 1.4;
}

.article-preview {
  font-size: 14px;
  color: #64748b;
  margin: 0 0 12px 0;
  line-height: 1.5;
}

.article-meta {
  display: flex;
  gap: 20px;
  font-size: 12px;
  color: #94a3b8;
}

.article-arrow {
  font-size: 24px;
  color: #cbd5e1;
  transition: all 0.2s ease;
  margin-left: 16px;
}

.article-item:hover .article-arrow {
  transform: translateX(4px);
  color: #667eea;
}

/* 加载状态 */
.loading-state {
  text-align: center;
  padding: 60px;
}

.loading-spinner {
  width: 44px;
  height: 44px;
  margin: 0 auto 20px;
  border: 3px solid #e2e8f0;
  border-top-color: #667eea;
  border-radius: 50%;
  animation: spin 0.8s linear infinite;
}

.loading-state p {
  color: #64748b;
  font-size: 14px;
}

@keyframes spin {
  to { transform: rotate(360deg); }
}

/* 空状态 */
.empty-state {
  text-align: center;
  padding: 60px;
}

.empty-icon {
  font-size: 72px;
  margin-bottom: 20px;
  opacity: 0.7;
}

.empty-state h4 {
  font-size: 20px;
  font-weight: 600;
  color: #1e293b;
  margin: 0 0 10px 0;
}

.empty-state p {
  font-size: 14px;
  color: #64748b;
  margin: 0 0 28px 0;
}

.browse-btn {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  border: none;
  padding: 12px 30px;
  border-radius: 40px;
  cursor: pointer;
  font-size: 15px;
  font-weight: 500;
  transition: all 0.3s ease;
}

.browse-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 8px 20px rgba(102, 126, 234, 0.3);
}

/* 响应式 */
@media (max-width: 768px) {
  .user-center-modern {
    padding: 40px 16px;
  }
  
  .user-profile-card {
    padding: 28px 20px;
  }
  
  .user-avatar-large {
    width: 80px;
    height: 80px;
    font-size: 34px;
  }
  
  .user-name {
    font-size: 22px;
  }
  
  .stat-value {
    font-size: 24px;
  }
  
  .liked-articles-section {
    padding: 24px 20px;
  }
  
  .article-title {
    font-size: 16px;
  }
  
  .article-preview {
    font-size: 13px;
  }
  
  .profile-stats {
    gap: 24px;
  }
}

@media (max-width: 500px) {
  .article-item {
    flex-direction: column;
    align-items: flex-start;
  }
  
  .article-arrow {
    align-self: flex-end;
    margin-top: 12px;
  }
  
  .profile-stats {
    flex-direction: column;
    gap: 16px;
  }
  
  .stat-divider {
    display: none;
  }
}
</style>