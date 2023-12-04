<script setup lang="ts">
import { NTabs, NTabPane } from 'naive-ui/lib'
import CowStat from '../components/cow-stat'
import CowCoord from '../components/cow-coord'
</script>

<template>
  <NTabs
    class="w-full h-full flex flex-col"
    type="line"
    animated
    pane-wrapper-style="flex-grow: 1;"
    pane-style="height: 100%;"
  >
    <NTabPane name="统计">
      <CowStat />
    </NTabPane>
    <NTabPane name="位置">
      <CowCoord />
    </NTabPane>
  </NTabs>
</template>
