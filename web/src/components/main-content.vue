<script setup lang="ts">
import { NCard } from 'naive-ui'
import { RouterView } from 'vue-router'
</script>

<template>
  <div class="p-3 grow flex flex-col">
    <NCard id="content" class="grow">
      <RouterView />
    </NCard>
  </div>
</template>

<style lang="scss" scoped>
#content {
  box-shadow: 0px 1px 10px 0px rgba($color: #000000, $alpha: 0.05);
}
</style>
