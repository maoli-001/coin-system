<template>
  <el-container class="home-container">
    <div class="content-wrapper">
      <!-- 头部欢迎 -->
      <div class="welcome-box">
        <h1 class="title">欢迎使用蓝鼠兑换</h1>
        <p class="description">
          一站式金融服务平台 · 实时汇率兑换 + 财经资讯阅读
        </p>
      </div>

      <!-- 功能模块卡片 -->
      <div class="func-list" >
        <div class="func-card">
          <div class="icon">💱</div>
          <h3>货币汇率兑换</h3>
          <p>多币种实时换算，快速完成货币兑换计算</p>
        </div>

        <div class="func-card">
          <div class="icon">📰</div>
          <h3>金融资讯阅读</h3>
          <p>精选财经新闻，了解实时行业动态</p>
        </div>

        <div class="func-card">
          <div class="icon">🔐</div>
          <h3>个人账号体系</h3>
          <p>登录解锁完整功能，支持文章点赞收藏</p>
        </div>
      </div>

      <p class="tip-text">请选择顶部导航栏功能进行操作</p>
    </div>
  </el-container>
</template>

<script setup lang="ts">
</script>

<style scoped>
.home-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  background: #f7f8fa;
  padding: 40px 20px;
  box-sizing: border-box;
}

.content-wrapper {
  text-align: center;
  max-width: 900px;
  width: 100%;
}

/* 欢迎区域 */
.welcome-box {
  margin-bottom: 48px;
}

.title {
  color: #2c3e50;
  font-size: 38px;
  font-weight: 600;
  margin-bottom: 12px;
}

.description {
  color: #7f8c8d;
  font-size: 17px;
}

/* 功能卡片列表 */
.func-list {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 24px;
  margin-bottom: 40px;
}

.func-card {
  background: #fff;
  padding: 28px 20px;
  border-radius: 16px;
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.06);
  transition: all 0.25s ease;
}

.func-card:hover {
  transform: translateY(-6px);
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.08);
}

.icon {
  font-size: 40px;
  margin-bottom: 16px;
}

.func-card h3 {
  font-size: 18px;
  color: #2c3e50;
  margin: 0 0 8px 0;
}

.func-card p {
  font-size: 14px;
  color: #999;
  line-height: 1.6;
  margin: 0;
}

/* 底部提示 */
.tip-text {
  color: #666;
  font-size: 15px;
}

/* 适配手机 */
@media (max-width: 768px) {
  .func-list {
    grid-template-columns: 1fr;
  }
  .title {
    font-size: 28px;
  }
}
</style>