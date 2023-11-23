<template>
  <NConfigProvider :theme-overrides="theme">
    <NMessageProvider>
      <NLoadingBarProvider>
        <TopEntry class="w-screen h-screen" />
      </NLoadingBarProvider>
    </NMessageProvider>
  </NConfigProvider>
</template>

<script setup lang="ts">
import { NConfigProvider, NMessageProvider, NLoadingBarProvider } from 'naive-ui/lib'
import theme from './assets/nord.naiveui.json'
import TopEntry from './components/top-entry.vue'
import { provide } from 'vue'
import { THEME_KEY } from 'vue-echarts'
import { registerTheme } from 'echarts'
import echartsNordTheme from './assets/nord.echarts.json'
registerTheme('nord', echartsNordTheme)
provide(THEME_KEY, 'nord')
</script>
