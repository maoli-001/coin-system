<template>
  <div class="exchange-modern">
    <!-- 动态粒子背景 -->
    <div class="particles-bg"></div>
    
    <div class="exchange-container">
      <div class="exchange-header">
        <h1 class="exchange-title">
          <span class="gradient-text">货币兑换</span>
          <br/>精准换算
        </h1>
        <p class="exchange-description">
          支持50+种货币 · 实时汇率更新 · 智能兑换计算
        </p>
      </div>

      <div class="exchange-card">
        <div class="currency-section">
          <div class="currency-input">
            <label class="input-label">从货币</label>
            <div class="select-wrapper">
              <select 
                v-model="form.fromCurrency" 
                class="modern-select"
                :class="{ 'has-value': form.fromCurrency }"
              >
                <option value="" disabled>请选择货币</option>
                <option v-for="currency in currencies" :key="currency" :value="currency">
                  {{ currency }}
                </option>
              </select>
            </div>
          </div>

          <div class="swap-icon" @click="swapCurrencies">
            <span class="swap-arrow">⇄</span>
          </div>

          <div class="currency-input">
            <label class="input-label">兑换为</label>
            <div class="select-wrapper">
              <select 
                v-model="form.toCurrency" 
                class="modern-select"
                :class="{ 'has-value': form.toCurrency }"
              >
                <option value="" disabled>请选择目标货币</option>
                <option v-for="currency in currencies" :key="currency" :value="currency">
                  {{ currency }}
                </option>
              </select>
            </div>
          </div>
        </div>

        <div class="amount-section">
          <label class="input-label">金额</label>
          <input 
            v-model="form.amount"
            type="number"
            placeholder="请输入兑换金额"
            class="amount-input"
          />
        </div>

        <button class="exchange-btn" @click="exchange">
          <span>立即兑换</span>
          <svg width="18" height="18" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M5 12H19" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
            <path d="M12 5L19 12L12 19" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
          </svg>
        </button>
      </div>

      <!-- 兑换结果 -->
      <transition name="fade">
        <div v-if="result !== null" class="result-card">
          <div class="result-icon">💰</div>
          <div class="result-content">
            <div class="result-label">兑换结果</div>
            <div class="result-amount">{{ result.toFixed(2) }}</div>
            <div class="result-currency">{{ form.toCurrency }}</div>
            <div class="rate-hint">
              1 {{ form.fromCurrency }} = {{ (form.amount ? result / form.amount : 0).toFixed(4) }} {{ form.toCurrency }}
            </div>
          </div>
        </div>
      </transition>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import axios from '../axios';

interface ExchangeRate {
  fromCurrency: string;
  toCurrency: string;
  rate: number;
}

const form = ref({
  fromCurrency: '',
  toCurrency: '',
  amount: 0,
});

const result = ref<number | null>(null);
const currencies = ref<string[]>([]);
const rates = ref<ExchangeRate[]>([]);

const fetchCurrencies = async () => {
  try {
    const response = await axios.get<ExchangeRate[]>('/exchangeRates');
    rates.value = response.data;

    const allCurrencies: string[] = [];
    response.data.forEach(rate => {
      allCurrencies.push(rate.fromCurrency);
      allCurrencies.push(rate.toCurrency);
    });
    currencies.value = [...new Set(allCurrencies)];
  } catch (error) {
    console.log('Failed to load currencies', error);
  }
};

// 货币互换
const swapCurrencies = () => {
  const temp = form.value.fromCurrency;
  form.value.fromCurrency = form.value.toCurrency;
  form.value.toCurrency = temp;
  if (result.value !== null) {
    exchange();
  }
};

// 任意两个货币可以兑换 A->人民币->B（完全保留原有逻辑）
const exchange = () => {
  const getCnyRate = (currency: string) => {
    const item = rates.value.find(
      r => r.fromCurrency === currency && r.toCurrency === "人民币"
    );
    if (item) return item.rate;

    const reverse = rates.value.find(
      r => r.fromCurrency === "人民币" && r.toCurrency === currency
    );
    if (reverse) return 1 / reverse.rate;

    return 1;
  };

  const fromRate = getCnyRate(form.value.fromCurrency);
  const toRate = getCnyRate(form.value.toCurrency);

  if (fromRate && toRate && form.value.amount > 0) {
    result.value = form.value.amount * (fromRate / toRate);
  } else {
    result.value = null;
  }
};

onMounted(fetchCurrencies);
</script>

<style scoped>
/* ===== 与主页同风格 ===== */
.exchange-modern {
  --primary: #2563eb;
  --primary-dark: #1d4ed8;
  --dark-bg: #0f172a;
  
  min-height: 100vh;
  position: relative;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  background-attachment: fixed;
  padding: 60px 20px;
  display: flex;
  justify-content: center;
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

.exchange-container {
  max-width: 550px;
  width: 100%;
  position: relative;
  z-index: 2;
}

/* 头部区域 */
.exchange-header {
  text-align: center;
  margin-bottom: 48px;
}

.exchange-title {
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

.exchange-description {
  font-size: 16px;
  color: rgba(255, 255, 255, 0.85);
  letter-spacing: 0.3px;
}

/* 兑换卡片 */
.exchange-card {
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(10px);
  border-radius: 32px;
  padding: 40px 32px;
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.15);
  border: 1px solid rgba(255, 255, 255, 0.3);
}

/* 货币选择区域 */
.currency-section {
  display: flex;
  gap: 16px;
  align-items: flex-end;
  margin-bottom: 28px;
}

.currency-input {
  flex: 1;
}

.input-label {
  display: block;
  font-size: 13px;
  font-weight: 600;
  color: #1e293b;
  margin-bottom: 8px;
  letter-spacing: 0.3px;
}

.select-wrapper {
  position: relative;
}

.modern-select {
  width: 100%;
  padding: 14px 16px;
  background: #f8fafc;
  border: 2px solid #e2e8f0;
  border-radius: 16px;
  font-size: 15px;
  font-weight: 500;
  color: #1e293b;
  cursor: pointer;
  transition: all 0.3s ease;
  appearance: none;
  background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='16' height='16' viewBox='0 0 24 24' fill='none' stroke='%2364748b' stroke-width='2' stroke-linecap='round' stroke-linejoin='round'%3E%3Cpolyline points='6 9 12 15 18 9'%3E%3C/polyline%3E%3C/svg%3E");
  background-repeat: no-repeat;
  background-position: right 16px center;
}

.modern-select:hover {
  border-color: #cbd5e1;
}

.modern-select:focus {
  outline: none;
  border-color: #667eea;
  box-shadow: 0 0 0 3px rgba(102, 126, 234, 0.1);
}

.modern-select.has-value {
  border-color: #667eea;
}

/* 互换图标 */
.swap-icon {
  width: 48px;
  height: 48px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border-radius: 16px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: all 0.3s ease;
  color: white;
  margin-bottom: 2px;
}

.swap-arrow {
  font-size: 24px;
  font-weight: 600;
  color: white;
  display: inline-block;
  transition: transform 0.3s ease;
}

.swap-icon:hover .swap-arrow {
  transform: rotate(180deg);
}

/* 金额输入区域 */
.amount-section {
  margin-bottom: 32px;
}

.amount-input {
  width: 90%;
  padding: 14px 18px;
  background: #f8fafc;
  border: 2px solid #e2e8f0;
  border-radius: 16px;
  font-size: 16px;
  font-weight: 500;
  color: #1e293b;
  transition: all 0.3s ease;
}

.amount-input:focus {
  outline: none;
  border-color: #667eea;
  box-shadow: 0 0 0 3px rgba(102, 126, 234, 0.1);
}

.amount-input::placeholder {
  color: #94a3b8;
  font-weight: normal;
}

/* 兑换按钮 */
.exchange-btn {
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
}

.exchange-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 10px 25px rgba(102, 126, 234, 0.4);
}

.exchange-btn:active {
  transform: translateY(0);
}

/* 结果卡片 */
.result-card {
  margin-top: 28px;
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(10px);
  border-radius: 24px;
  padding: 32px;
  display: flex;
  align-items: center;
  gap: 20px;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.1);
  border: 1px solid rgba(255, 255, 255, 0.3);
}

.result-icon {
  font-size: 48px;
  filter: drop-shadow(0 4px 6px rgba(0, 0, 0, 0.1));
}

.result-content {
  flex: 1;
}

.result-label {
  font-size: 13px;
  color: #64748b;
  margin-bottom: 6px;
  letter-spacing: 0.5px;
}

.result-amount {
  font-size: 42px;
  font-weight: 800;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  background-clip: text;
  -webkit-background-clip: text;
  color: transparent;
  line-height: 1.1;
  margin-bottom: 4px;
}

.result-currency {
  font-size: 14px;
  color: #475569;
  font-weight: 500;
  margin-bottom: 8px;
}

.rate-hint {
  font-size: 12px;
  color: #94a3b8;
}

/* 动画 */
.fade-enter-active,
.fade-leave-active {
  transition: all 0.3s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
  transform: translateY(-10px);
}

/* 响应式 */
@media (max-width: 600px) {
  .exchange-modern {
    padding: 40px 16px;
  }
  
  .exchange-title {
    font-size: 32px;
  }
  
  .exchange-card {
    padding: 28px 20px;
  }
  
  .currency-section {
    flex-direction: column;
    gap: 20px;
  }
  
  .swap-icon {
    align-self: center;
    transform: rotate(90deg);
    width: 44px;
    height: 44px;
  }
  
  .swap-icon:hover {
    transform: rotate(90deg) scale(1.05);
  }
  
  .result-card {
    padding: 24px;
  }
  
  .result-amount {
    font-size: 32px;
  }
  
  .result-icon {
    font-size: 36px;
  }
}

/* 移除 input number 箭头 */
.amount-input::-webkit-inner-spin-button,
.amount-input::-webkit-outer-spin-button {
  opacity: 0.5;
}
</style>