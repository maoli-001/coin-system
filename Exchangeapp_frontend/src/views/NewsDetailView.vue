<template>
  <div class="news-detail-page">
    <div class="container">
      <!-- 文章详情 -->
      <el-card v-if="article" class="article-card" shadow="hover">
        <div class="card-body">
          <h1 class="article-title">{{ article.Title }}</h1>
          <p class="article-content">{{ article.Content }}</p>

          <!-- 点赞区域 -->
          <div class="like-row">
            <el-button type="primary" @click="likeArticle">❤️ 点赞</el-button>
            <span class="like-count">点赞数：{{ likes }}</span>
          </div>
        </div>
      </el-card>

      <!-- 未登录展示 -->
      <div v-else class="empty-box">
        <el-empty description="您必须登录/注册才可以阅读文章" />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue";
import { useRoute } from "vue-router";
import axios from "../axios";
import { ElEmpty } from "element-plus";
import type { Article, Like } from "../types/Article";

const article = ref<Article | null>(null);
const route = useRoute();
const likes = ref<number>(0);

const { id } = route.params;

const fetchArticle = async () => {
  try {
    const response = await axios.get<Article>(`/articles/${id}`);
    article.value = response.data;
  } catch (error) {
    console.error("Failed to load article:", error);
  }
};

const likeArticle = async () => {
  try {
    const res = await axios.post<Like>(`articles/${id}/like`);
    likes.value = res.data.likes;
    await fetchLike();
  } catch (error) {
    console.log("Error Liking article:", error);
  }
};

const fetchLike = async () => {
  try {
    const res = await axios.get<Like>(`articles/${id}/like`);
    likes.value = res.data.likes;
  } catch (error) {
    console.log("Error fetching likes:", error);
  }
};

onMounted(fetchArticle);
onMounted(fetchLike);
</script>

<style scoped>
/* 页面整体 */
.news-detail-page {
  width: 100%;
  min-height: 100vh;
  background: #f7f8fa;
  padding: 40px 20px;
}

/* 容器居中 */
.container {
  max-width: 800px;
  margin: 0 auto;
}

/* 文章卡片 */
.article-card {
  border-radius: 16px;
  border: none;
  overflow: hidden;
}

.card-body {
  padding: 32px;
}

/* 文章标题 */
.article-title {
  font-size: 26px;
  font-weight: 600;
  color: #2c3e50;
  text-align: center;
  margin: 0 0 24px 0;
  line-height: 1.4;
}

/* 文章内容 */
.article-content {
  font-size: 16px;
  color: #34495e;
  line-height: 1.8;
  margin: 0 0 32px 0;
  white-space: pre-wrap;
}

/* 点赞栏 */
.like-row {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding-top: 16px;
  border-top: 1px solid #eee;
}

.like-count {
  font-size: 15px;
  color: #7f8c8d;
}

/* 空状态 */
.empty-box {
  margin-top: 60px;
  text-align: center;
}
</style>