<template>
  <div class="news-page">
    <div class="container">
      <h1 class="page-title">📰 资讯中心</h1>

      <div v-if="articles && articles.length" class="article-list">
        <el-card
          v-for="article in articles"
          :key="article.ID"
          class="article-card"
          shadow="hover"
        >
          <div class="card-body">
            <h2 class="article-title">{{ article.Title }}</h2>
            <p class="article-preview">{{ article.Preview }}</p>
            <div class="action-row">
              <el-button type="primary" text @click="viewDetail(article.ID)">
                阅读更多
              </el-button>
            </div>
          </div>
        </el-card>
      </div>

      <div v-else class="no-data-box">
        <el-empty description="您必须登录/注册才可以查看文章" />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { ElMessage, ElEmpty } from 'element-plus';
import axios from '../axios';
import { useAuthStore } from '../store/auth';
import type { Article } from "../types/Article";

const articles = ref<Article[]>([]);
const router = useRouter();
const authStore = useAuthStore();

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

onMounted(fetchArticles);
</script>

<style scoped>
/* 页面整体 */
.news-page {
  width: 100%;
  min-height: 100vh;
  background: #f7f8fa;
  padding: 40px 20px;
}

/* 内容容器 */
.container {
  max-width: 800px;
  margin: 0 auto;
}

/* 页面标题 */
.page-title {
  font-size: 28px;
  font-weight: 600;
  color: #2c3e50;
  text-align: center;
  margin-bottom: 36px;
}

/* 文章列表 */
.article-list {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

/* 文章卡片 */
.article-card {
  border-radius: 16px;
  border: none;
  overflow: hidden;
}

.card-body {
  padding: 24px;
}

/* 文章标题 */
.article-title {
  font-size: 20px;
  font-weight: 600;
  color: #2c3e50;
  margin: 0 0 12px 0;
  line-height: 1.4;
}

/* 预览内容 */
.article-preview {
  font-size: 15px;
  color: #7f8c8d;
  line-height: 1.7;
  margin: 0 0 16px 0;
}

/* 按钮行 */
.action-row {
  text-align: right;
}

/* 无数据 */
.no-data-box {
  margin-top: 60px;
  text-align: center;
}
</style>