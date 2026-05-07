<template>
  <div class="exchange-page">
    <div class="exchange-wrapper">
      <h2 class="title">💱 货币实时兑换</h2>
      <p class="subtitle">快速查询全球货币汇率，一键完成换算</p>

      <el-card shadow="hover" class="exchange-card">
        <el-form :model="form" class="exchange-form">
          <el-form-item label="从货币" label-width="80px">
            <el-select
              v-model="form.fromCurrency"
              placeholder="请选择货币"
              style="width: 100%"
              size="default"
            >
              <el-option
                v-for="currency in currencies"
                :key="currency"
                :label="currency"
                :value="currency"
              />
            </el-select>
          </el-form-item>

          <el-form-item label="兑换为" label-width="80px">
            <el-select
              v-model="form.toCurrency"
              placeholder="请选择目标货币"
              style="width: 100%"
              size="default"
            >
              <el-option
                v-for="currency in currencies"
                :key="currency"
                :label="currency"
                :value="currency"
              />
            </el-select>
          </el-form-item>

          <el-form-item label="金额" label-width="80px">
            <el-input
              v-model="form.amount"
              type="number"
              placeholder="请输入兑换金额"
              style="width: 100%"
              size="default"
            />
          </el-form-item>

          <el-form-item style="margin-top: 10px">
            <el-button
              type="primary"
              @click="exchange"
              class="exchange-btn"
              size="default"
              block
            >
              立即兑换
            </el-button>
          </el-form-item>
        </el-form>
      </el-card>

      <!-- 兑换结果 -->
      <div v-if="result !== null" class="result-card">
        <div class="result-title">兑换结果</div>
        <div class="result-value">{{ result.toFixed(2) }}</div>
        <div class="result-tip">{{ form.toCurrency }} 金额</div>
      </div>
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

const exchange = () => {
  const rate = rates.value.find(
    (rate) => rate.fromCurrency === form.value.fromCurrency && rate.toCurrency === form.value.toCurrency
  )?.rate;

  if (rate) {
    result.value = form.value.amount * rate;
  } else {
    result.value = null;
  }
};

onMounted(fetchCurrencies);
</script>

<style scoped>
/* 页面整体布局 */
.exchange-page {
  min-height: 100vh;
  padding: 40px 20px;
  background: #f7f8fa;
  display: flex;
  justify-content: center;
  align-items: flex-start;
}

/* 内容容器 */
.exchange-wrapper {
  width: 100%;
  max-width: 520px;
}

/* 标题 */
.title {
  font-size: 26px;
  font-weight: 600;
  color: #2c3e50;
  text-align: center;
  margin: 0 0 8px 0;
}

.subtitle {
  font-size: 14px;
  color: #7f8c8d;
  text-align: center;
  margin: 0 0 30px 0;
}

/* 兑换卡片 */
.exchange-card {
  border-radius: 16px;
  padding: 28px;
  border: none;
  background: #ffffff;
}

/* 表单间距 */
.exchange-form {
  width: 100%;
}

/* 按钮美化 */
.exchange-btn {
  height: 44px;
  width:425px;
  font-size: 16px;
  font-weight: 500;
  border-radius: 8px;
  background: linear-gradient(135deg, #409eff 0%, #69b1ff 100%);
  border: none;
}

/* 结果卡片 */
.result-card {
  margin-top: 24px;
  padding: 28px;
  background: #ffffff;
  border-radius: 16px;
  text-align: center;
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.06);
}

.result-title {
  font-size: 14px;
  color: #7f8c8d;
  margin-bottom: 8px;
}

.result-value {
  font-size: 36px;
  font-weight: bold;
  color: #27ae60;
  line-height: 1.2;
}

.result-tip {
  font-size: 13px;
  color: #95a5a6;
  margin-top: 6px;
}

/* 元素间距 */
:deep(.el-form-item) {
  margin-bottom: 20px;
}

/* 输入框/选择器圆角 */
:deep(.el-input__wrapper),
:deep(.el-select__wrapper) {
  border-radius: 10px !important;
  box-shadow: none;
  border: 1px solid #e4e7ed;
}

/* 聚焦效果 */
:deep(.el-input__wrapper.is-focus),
:deep(.el-select__wrapper.is-focus) {
  box-shadow: 0 0 0 2px rgba(64, 158, 255, 0.1);
  border-color: #409eff;
}
</style>