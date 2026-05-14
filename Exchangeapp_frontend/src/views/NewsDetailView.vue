<template>
  <div class="news-detail-modern">
    <!-- 动态粒子背景 -->
    <div class="particles-bg"></div>
    
    <div class="detail-container">
      <!-- 返回按钮 -->
      <button class="back-btn" @click="goBack">
        <span class="back-arrow">←</span>
        返回资讯列表
      </button>

      <!-- 文章详情卡片 -->
      <div v-if="article" class="article-card">
        <h1 class="article-title">{{ article.Title }}</h1>
        
        <div class="article-meta">
          <span class="meta-item">
            <span class="meta-icon">📅</span>
            {{ formatDate(article.CreatedAt) }}
          </span>
          <span class="meta-item">
            <span class="meta-icon">❤️</span>
            {{ likes }} 点赞数
          </span>
        </div>

        <div class="article-divider"></div>

        <div class="article-content">{{ article.Content }}</div>

        <!-- 点赞区域 -->
        <div class="like-section">
          <button 
            class="like-btn" 
            @click="likeArticle"
            :class="{ 'liked': hasLiked }"
            :disabled="isLiking"
          >
            <span class="like-icon">{{ hasLiked ? '❤️' : '🤍' }}</span>
            <span class="like-text">{{ isLiking ? '处理中...' : (hasLiked ? '已点赞' : '点赞') }}</span>
          </button>
          <div class="like-stats">
            <span class="stats-number">{{ likes }}</span>
            <span class="stats-label">点赞数</span>
          </div>
        </div>
      </div>

      <!-- 未登录/无文章展示 -->
      <div v-else class="empty-card">
        <div class="empty-icon">📭</div>
        <h3 class="empty-title">无法阅读文章</h3>
        <p class="empty-desc">您需要登录/注册后才能查看文章详情</p>
        <button class="empty-btn" @click="goToLogin">去登录</button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue";
import { useRoute, useRouter } from "vue-router";
import axios from "../axios";
import { ElMessage } from "element-plus";
import { useAuthStore } from '../store/auth';
import type { Article } from "../types/Article";

const article = ref<Article | null>(null);
const route = useRoute();
const router = useRouter();
const authStore = useAuthStore();
const likes = ref<number>(0);
const hasLiked = ref(false);
const isLiking = ref(false);
const articleId = route.params.id as string;

// 获取文章详情
const fetchArticle = async () => {
  try {
    const res = await axios.get<Article>(`/articles/${articleId}`);
    article.value = res.data;
  } catch (err) {
    console.error("获取文章失败", err);
  }
};

// 获取点赞数 + 是否已点赞（适配后端返回的 isLiked 字段）
const fetchLike = async () => {
  try {
    const res = await axios.get(`/articles/${articleId}/like`);
    console.log("获取点赞信息:", res.data);
    
    // 后端返回: { likes: "5", isLiked: true }
    likes.value = Number(res.data.likes) || 0;
    hasLiked.value = res.data.isLiked || false;
  } catch (err) {
    console.log("获取点赞状态失败", err);
  }
};

// 点赞 / 取消点赞（二合一）
const likeArticle = async () => {
  // 检查是否登录
  if (!authStore.isAuthenticated) {
    ElMessage.warning('请先登录后再点赞');
    router.push({ name: 'Login' });
    return;
  }
  
  if (isLiking.value) return;
  isLiking.value = true;

  try {
    const res = await axios.post(`/articles/${articleId}/like`);
    console.log("点赞操作响应:", res.data);
    
    // 后端返回: { message: "...", likes: 5, liked: true/false }
    likes.value = Number(res.data.likes) || 0;
    hasLiked.value = res.data.liked || false;

    if (hasLiked.value) {
      ElMessage.success("点赞成功 ✅");
    } else {
      ElMessage.info("已取消点赞 ❌");
    }
  } catch (err: any) {
    console.error("点赞操作失败:", err);
    if (err.response?.status === 401) {
      ElMessage.error("请先登录");
      router.push({ name: 'Login' });
    } else {
      ElMessage.error("操作失败，请重试");
    }
  } finally {
    isLiking.value = false;
  }
};

const goBack = () => router.back();
const goToLogin = () => router.push("/login");

const formatDate = (d?: string) => {
  if (!d) return "未知日期";
  const date = new Date(d);
  return `${date.getFullYear()}.${String(date.getMonth()+1).padStart(2, '0')}.${String(date.getDate()).padStart(2, '0')}`;
};

onMounted(() => {
  fetchArticle();
  fetchLike();
});
</script>

<style scoped>
/* ===== 与主页/其他页面同风格 ===== */
.news-detail-modern {
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

.detail-container {
  max-width: 850px;
  margin: 0 auto;
  position: relative;
  z-index: 2;
}

/* 返回按钮 */
.back-btn {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  background: rgba(255, 255, 255, 0.2);
  backdrop-filter: blur(10px);
  border: 1px solid rgba(255, 255, 255, 0.3);
  padding: 10px 20px;
  border-radius: 40px;
  font-size: 14px;
  font-weight: 500;
  color: white;
  cursor: pointer;
  margin-bottom: 28px;
  transition: all 0.3s ease;
}

.back-btn:hover {
  background: rgba(255, 255, 255, 0.3);
  transform: translateX(-4px);
}

.back-arrow {
  font-size: 18px;
}

/* 文章卡片 */
.article-card {
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(10px);
  border-radius: 32px;
  padding: 48px;
  border: 1px solid rgba(255, 255, 255, 0.3);
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.15);
}

/* 文章标题 */
.article-title {
  font-size: 32px;
  font-weight: 800;
  background: linear-gradient(135deg, #1e293b 0%, #334155 100%);
  background-clip: text;
  -webkit-background-clip: text;
  color: transparent;
  margin: 0 0 20px 0;
  line-height: 1.3;
}

/* 文章元信息 */
.article-meta {
  display: flex;
  gap: 24px;
  margin-bottom: 20px;
  flex-wrap: wrap;
}

.meta-item {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  font-size: 13px;
  color: #64748b;
}

.meta-icon {
  font-size: 14px;
}

/* 分割线 */
.article-divider {
  height: 1px;
  background: linear-gradient(90deg, transparent, #e2e8f0, transparent);
  margin: 20px 0 28px 0;
}

/* 文章内容 */
.article-content {
  font-size: 17px;
  color: #334155;
  line-height: 1.85;
  margin-bottom: 40px;
  white-space: pre-wrap;
  word-break: break-word;
}

/* 点赞区域 */
.like-section {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding-top: 24px;
  border-top: 1px solid #e2e8f0;
  flex-wrap: wrap;
  gap: 16px;
}

.like-btn {
  display: inline-flex;
  align-items: center;
  gap: 10px;
  background: linear-gradient(135deg, #f1f5f9 0%, #e2e8f0 100%);
  border: none;
  padding: 12px 28px;
  border-radius: 50px;
  font-size: 15px;
  font-weight: 600;
  color: #475569;
  cursor: pointer;
  transition: all 0.3s ease;
}

.like-btn:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 6px 15px rgba(0, 0, 0, 0.1);
}

.like-btn:disabled {
  opacity: 0.7;
  cursor: not-allowed;
}

.like-btn.liked {
  background: linear-gradient(135deg, #ff6b6b 0%, #ee5a5a 100%);
  color: white;
}

.like-icon {
  font-size: 18px;
}

.like-stats {
  text-align: right;
}

.stats-number {
  display: block;
  font-size: 28px;
  font-weight: 800;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  background-clip: text;
  -webkit-background-clip: text;
  color: transparent;
  line-height: 1.2;
}

.stats-label {
  font-size: 12px;
  color: #94a3b8;
}

/* 空状态卡片 */
.empty-card {
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(10px);
  border-radius: 32px;
  padding: 60px 40px;
  text-align: center;
  border: 1px solid rgba(255, 255, 255, 0.3);
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.1);
}

.empty-icon {
  font-size: 64px;
  margin-bottom: 20px;
  opacity: 0.7;
}

.empty-title {
  font-size: 24px;
  font-weight: 600;
  color: #1e293b;
  margin: 0 0 12px 0;
}

.empty-desc {
  font-size: 16px;
  color: #64748b;
  margin: 0 0 28px 0;
}

.empty-btn {
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

.empty-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 8px 20px rgba(102, 126, 234, 0.3);
}

/* 响应式适配 */
@media (max-width: 768px) {
  .news-detail-modern {
    padding: 40px 16px;
  }
  
  .article-card {
    padding: 28px 24px;
  }
  
  .article-title {
    font-size: 24px;
  }
  
  .article-content {
    font-size: 15px;
    line-height: 1.75;
  }
  
  .like-section {
    flex-direction: column;
    align-items: stretch;
    text-align: center;
  }
  
  .like-stats {
    text-align: center;
  }
  
  .empty-card {
    padding: 40px 24px;
  }
  
  .empty-title {
    font-size: 20px;
  }
}

@media (max-width: 480px) {
  .article-card {
    padding: 20px 18px;
  }
  
  .article-title {
    font-size: 20px;
  }
  
  .stats-number {
    font-size: 24px;
  }
  
  .like-btn {
    justify-content: center;
  }
}
</style>