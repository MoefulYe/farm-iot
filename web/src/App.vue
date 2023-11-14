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
import theme from './assets/theme.json'
import TopEntry from './components/top-entry.vue'
</script>
