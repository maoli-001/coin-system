<template>
  <div class="news-modern">
    <!-- 动态粒子背景 -->
    <div class="particles-bg"></div>
    
    <div class="news-container">
      <div class="news-header">
        <h1 class="news-title">
          <span class="gradient-text">资讯中心</span>
          <br/>掌握市场动态
        </h1>
        <p class="news-description">
          精选财经新闻 · 实时行业动态 · 专业深度解读
        </p>
      </div>

      <div v-if="articles && articles.length" class="article-list">
        <div
          v-for="article in articles"
          :key="article.ID"
          class="article-card"
          @click="viewDetail(article.ID)"
        >
          <div class="card-content">
            <div class="card-header-area">
              <div class="article-icon">📄</div>
              <div class="article-meta">
                <span class="meta-tag">最新资讯</span>
                <span class="meta-date">{{ formatDate(article.CreatedAt) }}</span>
              </div>
            </div>
            <h2 class="article-title">{{ article.Title }}</h2>
            <p class="article-preview">{{ article.Preview }}</p>
            <div class="card-footer">
              <span class="read-more">
                阅读全文
                <span class="arrow">→</span>
              </span>
            </div>
          </div>
        </div>
      </div>

      <div v-else class="no-data-card">
        <div class="no-data-icon">📭</div>
        <h3 class="no-data-title">暂无文章</h3>
        <p class="no-data-desc">您需要登录/注册后才能查看资讯内容</p>
        <button class="login-hint-btn" @click="goToLogin">去登录</button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { ElMessage } from 'element-plus';
import axios from '../axios';
import { useAuthStore } from '../store/auth';
import type { Article } from "../types/Article";

const articles = ref<Article[]>([]);
const router = useRouter();
const authStore = useAuthStore();

// 格式化日期（简单处理）
const formatDate = (dateStr?: string) => {
  if (!dateStr) return '近日';
  const date = new Date(dateStr);
  const now = new Date();
  const diff = now.getTime() - date.getTime();
  const days = Math.floor(diff / (1000 * 60 * 60 * 24));
  
  if (days === 0) return '今天';
  if (days === 1) return '昨天';
  if (days < 7) return `${days}天前`;
  if (days < 30) return `${Math.floor(days / 7)}周前`;
  return `${Math.floor(days / 30)}个月前`;
};

const fetchArticles = async () => {
  try {
    const response = await axios.get<Article[]>('/articles');
    articles.value = response.data;
  } catch (error) {
    console.error('Failed to load articles:', error);
  }
};

const viewDetail = (id: string) => {
  if (!authStore.isAuthenticated) {
    ElMessage.error('请先登录后再查看');
    return;
  }
  router.push({ name: 'NewsDetail', params: { id } });
};

const goToLogin = () => {
  router.push({ name: 'Login' });
};

onMounted(fetchArticles);
</script>

<style scoped>
/* ===== 与主页/兑换页同风格 ===== */
.news-modern {
  --primary: #2563eb;
  --primary-dark: #1d4ed8;
  
  min-height: 100vh;
  position: relative;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  background-attachment: fixed;
  padding: 60px 20px;
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

.news-container {
  max-width: 900px;
  margin: 0 auto;
  position: relative;
  z-index: 2;
}

/* 头部区域 */
.news-header {
  text-align: center;
  margin-bottom: 48px;
}

.badge {
  display: inline-block;
  background: rgba(255, 255, 255, 0.2);
  backdrop-filter: blur(10px);
  padding: 6px 18px;
  border-radius: 50px;
  font-size: 13px;
  font-weight: 500;
  color: white;
  margin-bottom: 20px;
}

.news-title {
  font-size: 48px;
  font-weight: 800;
  line-height: 1.2;
  margin-bottom: 16px;
  color: white;
}

.gradient-text {
  background: linear-gradient(135deg, #fff 0%, #a5d8ff 100%);
  background-clip: text;
  -webkit-background-clip: text;
  color: transparent;
}

.news-description {
  font-size: 16px;
  color: rgba(255, 255, 255, 0.85);
  letter-spacing: 0.3px;
}

/* 文章列表 */
.article-list {
  display: flex;
  flex-direction: column;
  gap: 24px;
}

/* 文章卡片 */
.article-card {
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(10px);
  border-radius: 24px;
  padding: 28px 32px;
  cursor: pointer;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  border: 1px solid rgba(255, 255, 255, 0.3);
  box-shadow: 0 4px 15px rgba(0, 0, 0, 0.05);
}

.article-card:hover {
  transform: translateY(-4px);
  box-shadow: 0 20px 35px rgba(0, 0, 0, 0.15);
  border-color: rgba(102, 126, 234, 0.3);
  background: white;
}

/* 卡片头部区域 */
.card-header-area {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-bottom: 20px;
}

.article-icon {
  font-size: 28px;
  filter: drop-shadow(0 2px 4px rgba(0, 0, 0, 0.1));
}

.article-meta {
  display: flex;
  gap: 12px;
  align-items: center;
  flex-wrap: wrap;
}

.meta-tag {
  background: linear-gradient(135deg, #667eea15 0%, #764ba215 100%);
  padding: 4px 12px;
  border-radius: 30px;
  font-size: 12px;
  font-weight: 500;
  color: #667eea;
}

.meta-date {
  font-size: 12px;
  color: #94a3b8;
}

/* 文章标题 */
.article-title {
  font-size: 22px;
  font-weight: 700;
  color: #1e293b;
  margin: 0 0 12px 0;
  line-height: 1.4;
  transition: color 0.2s ease;
}

.article-card:hover .article-title {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  background-clip: text;
  -webkit-background-clip: text;
  color: transparent;
}

/* 预览内容 */
.article-preview {
  font-size: 15px;
  color: #64748b;
  line-height: 1.65;
  margin: 0 0 20px 0;
}

/* 卡片底部 */
.card-footer {
  text-align: right;
}

.read-more {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  font-size: 14px;
  font-weight: 500;
  color: #667eea;
  transition: gap 0.2s ease;
}

.arrow {
  transition: transform 0.2s ease;
}

.article-card:hover .read-more .arrow {
  transform: translateX(4px);
}

/* 无数据卡片（登录提示） */
.no-data-card {
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(10px);
  border-radius: 32px;
  padding: 60px 40px;
  text-align: center;
  border: 1px solid rgba(255, 255, 255, 0.3);
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.1);
}

.no-data-icon {
  font-size: 64px;
  margin-bottom: 20px;
  opacity: 0.7;
}

.no-data-title {
  font-size: 24px;
  font-weight: 600;
  color: #1e293b;
  margin: 0 0 12px 0;
}

.no-data-desc {
  font-size: 16px;
  color: #64748b;
  margin: 0 0 28px 0;
}

.login-hint-btn {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  border: none;
  padding: 12px 32px;
  font-size: 15px;
  font-weight: 600;
  border-radius: 40px;
  cursor: pointer;
  transition: all 0.3s ease;
}

.login-hint-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 8px 20px rgba(102, 126, 234, 0.3);
}

/* 加载状态（可选） */
.loading-state {
  display: flex;
  justify-content: center;
  padding: 60px;
}

/* 响应式适配 */
@media (max-width: 768px) {
  .news-modern {
    padding: 40px 16px;
  }
  
  .news-title {
    font-size: 32px;
  }
  
  .news-description {
    font-size: 14px;
  }
  
  .article-card {
    padding: 20px 24px;
  }
  
  .article-title {
    font-size: 18px;
  }
  
  .article-preview {
    font-size: 14px;
  }
  
  .card-header-area {
    flex-wrap: wrap;
  }
  
  .no-data-card {
    padding: 40px 24px;
  }
  
  .no-data-title {
    font-size: 20px;
  }
}

@media (max-width: 480px) {
  .article-meta {
    width: 100%;
  }
  
  .article-icon {
    font-size: 24px;
  }
}
</style>