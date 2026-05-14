<template>
  <div class="home-modern">
    <!-- 动态粒子背景 -->
    <div class="particles-bg"></div>
    
    <div class="hero-section">
      <div class="hero-content">
        <h1 class="hero-title">
          <span class="gradient-text">货币兑换</span>
          <br/>智能汇率中枢
        </h1>
        <p class="hero-description">
          实时外汇数据 · 社区专业讨论 · 智能兑换计算
        </p>
        <div class="hero-stats">
          <div class="stat-item">
            <div class="stat-number">24/7</div>
            <div class="stat-label">实时更新</div>
          </div>
          <div class="stat-item">
            <div class="stat-number">50+</div>
            <div class="stat-label">货币对</div>
          </div>
          <div class="stat-item">
            <div class="stat-number">10k+</div>
            <div class="stat-label">社区成员</div>
          </div>
        </div>
      </div>
      <div class="hero-graphic">
        <div class="floating-chart"></div>
      </div>
    </div>

    <div class="features-section">
      <div class="container">
        <div class="section-header">
          <h2>核心功能</h2>
          <p>一站式金融服务，满足您的所有需求</p>
        </div>
        
        <div class="features-grid">
          <div class="feature-card" @click="navigateToExchange">
            <div class="feature-icon">💱</div>
            <h3>货币汇率兑换</h3>
            <p>实时汇率换算，支持50+种货币，快速完成兑换计算</p>
            <div class="feature-link">立即体验 →</div>
          </div>

          <div class="feature-card" @click="navigateToNews">
            <div class="feature-icon">📰</div>
            <h3>金融资讯阅读</h3>
            <p>精选财经新闻，了解实时行业动态，把握市场先机</p>
            <div class="feature-link">阅读资讯 →</div>
          </div>

          <div class="feature-card" @click="navigateToProfile">
            <div class="feature-icon">🔐</div>
            <h3>个人账号体系</h3>
            <p>登录解锁完整功能，支持文章点赞收藏，个性化体验</p>
            <div class="feature-link">进入个人中心 →</div>
          </div>
        </div>
      </div>
    </div>

    <div class="cta-section">
      <div class="cta-content">
        <h2>准备好开始交易了吗？</h2>
        <p>加入蓝鼠社区，与专业交易者一起成长</p>
        <button class="cta-button" @click="startTrading">开始使用 →</button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import router from "../router/index"
const navigateToExchange = () => {
  router.push('/exchange')
}

const navigateToNews = () => {
  router.push('/news')
}


const navigateToProfile = () => {
  router.push('/login')
}

const startTrading = () => {
  // 可以复用 navigateToExchange 或其他逻辑
  navigateToExchange()
}
</script>

<style scoped>
/* ===== CSS 变量定义 ===== */
.home-modern {
  --primary: #2563eb;
  --primary-dark: #1d4ed8;
  --primary-glow: rgba(37, 99, 235, 0.15);
  --dark-bg: #0f172a;
  --text-main: #1e293b;
  --text-secondary: #64748b;
  --border-light: #e2e8f0;
  
  min-height: 100vh;
  position: relative;
  /* 关键修复：直接设置渐变背景，确保白色文字可见 */
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  background-attachment: fixed;
}

/* 粒子背景效果 - 在渐变之上增加层次感 */
.particles-bg {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-image: 
    radial-gradient(circle at 10% 20%, rgba(255,255,255,0.15) 2px, transparent 2px),
    radial-gradient(circle at 90% 80%, rgba(255,255,255,0.1) 1px, transparent 1px);
  background-size: 50px 50px, 30px 30px;
  pointer-events: none;
  z-index: 0;
}

/* Hero 区域 - 白色文字在深色背景上清晰可见 */
.hero-section {
  display: flex;
  justify-content: space-between;
  align-items: center;
  max-width: 1200px;
  margin: 0 auto;
  padding: 80px 40px;
  position: relative;
  z-index: 2;
  gap: 60px;
  flex-wrap: wrap;
}

.hero-content {
  flex: 1;
  min-width: 300px;
}

.hero-title {
  font-size: 56px;
  font-weight: 800;
  line-height: 1.2;
  margin-bottom: 24px;
  color: white; /* 确保主标题白色 */
}

.gradient-text {
  background: linear-gradient(135deg, #fff 0%, #a5d8ff 100%);
  background-clip: text;
  -webkit-background-clip: text;
  color: transparent; /* 渐变文字保留透明效果 */
}

.hero-description {
  font-size: 18px;
  color: rgba(255, 255, 255, 0.9);
  margin-bottom: 40px;
  line-height: 1.6;
}

.hero-stats {
  display: flex;
  gap: 48px;
}

.stat-item {
  text-align: left;
}

.stat-number {
  font-size: 32px;
  font-weight: 700;
  color: white;
  margin-bottom: 8px;
}

.stat-label {
  font-size: 14px;
  color: rgba(255, 255, 255, 0.7);
}

.hero-graphic {
  flex: 1;
  min-width: 300px;
  height: 300px;
  position: relative;
}

.floating-chart {
  width: 100%;
  height: 100%;
  background: linear-gradient(135deg, rgba(255,255,255,0.2), rgba(255,255,255,0.05));
  border-radius: 30px;
  backdrop-filter: blur(10px);
  animation: float 6s ease-in-out infinite;
  position: relative;
}

.floating-chart::before {
  content: '📈';
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  font-size: 80px;
  opacity: 0.5;
}

@keyframes float {
  0%, 100% { transform: translateY(0px); }
  50% { transform: translateY(-20px); }
}

/* 功能区域 */
.features-section {
  background: white;
  padding: 80px 40px;
  position: relative;
  z-index: 2;
  border-radius: 40px 40px 0 0; /* 与渐变背景过渡柔和 */
}

.container {
  max-width: 1200px;
  margin: 0 auto;
}

.section-header {
  text-align: center;
  margin-bottom: 60px;
}

.section-header h2 {
  font-size: 36px;
  font-weight: 700;
  color: var(--text-main);
  margin-bottom: 12px;
}

.section-header p {
  font-size: 18px;
  color: var(--text-secondary);
}

.features-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(320px, 1fr));
  gap: 32px;
}

.feature-card {
  background: white;
  padding: 40px 32px;
  border-radius: 20px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.08);
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  cursor: pointer;
  border: 1px solid var(--border-light);
  position: relative;
  overflow: hidden;
}

.feature-card::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 4px;
  background: linear-gradient(90deg, var(--primary), var(--primary-dark));
  transform: scaleX(0);
  transition: transform 0.3s ease;
}

.feature-card:hover::before {
  transform: scaleX(1);
}

.feature-card:hover {
  transform: translateY(-8px);
  box-shadow: 0 12px 40px rgba(37, 99, 235, 0.15);
  border-color: transparent;
}

.feature-icon {
  font-size: 48px;
  margin-bottom: 24px;
}

.feature-card h3 {
  font-size: 24px;
  font-weight: 600;
  margin-bottom: 12px;
  color: var(--text-main);
}

.feature-card p {
  color: var(--text-secondary);
  line-height: 1.6;
  margin-bottom: 24px;
}

.feature-link {
  color: var(--primary);
  font-weight: 600;
  font-size: 14px;
  transition: transform 0.2s ease;
  display: inline-block;
}

.feature-card:hover .feature-link {
  transform: translateX(8px);
}

/* CTA 区域 */
.cta-section {
  margin: 0 0px 40px 0px;
  padding: 80px 40px;
  text-align: center;
  position: relative;
  z-index: 2;
  overflow: hidden;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.2);
}

.cta-section::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(45deg, rgba(37, 99, 235, 0.1), transparent);
}

.cta-content {
  position: relative;
  z-index: 2;
}

.cta-section h2 {
  font-size: 36px;
  font-weight: 700;
  color: white;
  margin-bottom: 16px;
}

.cta-section p {
  font-size: 18px;
  color: rgba(255, 255, 255, 0.8);
  margin-bottom: 32px;
}

.cta-button {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  border: none;
  padding: 14px 36px;
  font-size: 16px;
  font-weight: 600;
  border-radius: 50px;
  cursor: pointer;
  transition: all 0.3s ease;
}

.cta-button:hover {
  transform: translateY(-2px);
  box-shadow: 0 10px 25px rgba(102, 126, 234, 0.3);
}

/* 响应式适配 */
@media (max-width: 768px) {
  .hero-section {
    flex-direction: column;
    padding: 40px 20px;
    text-align: center;
  }
  
  .hero-title {
    font-size: 36px;
  }
  
  .hero-stats {
    justify-content: center;
  }
  
  .stat-item {
    text-align: center;
  }
  
  .features-section {
    padding: 40px 20px;
  }
  
  .section-header h2 {
    font-size: 28px;
  }
  
  .section-header p {
    font-size: 16px;
  }
  
  .features-grid {
    gap: 20px;
  }
  
  .feature-card {
    padding: 28px 20px;
  }
  
  .feature-card h3 {
    font-size: 20px;
  }
  
  .cta-section {
    padding: 40px 20px;
    margin: 0 16px 32px 16px;
  }
  
  .cta-section h2 {
    font-size: 28px;
  }
  
  .cta-section p {
    font-size: 16px;
  }
}

@media (max-width: 480px) {
  .hero-stats {
    flex-direction: column;
    gap: 20px;
  }
  
  .hero-graphic {
    height: 200px;
  }
  
  .floating-chart::before {
    font-size: 60px;
  }
}
</style>