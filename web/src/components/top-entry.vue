<template>
  <NLayout class="w-full h-full flex" has-sider>
    <NLayoutSider
      id="menu"
      collapse-mode="width"
      :collapsed-width="64"
      :width="240"
      :collapsed="isCollapsed"
      @collapse="() => (isCollapsed = true)"
      @expand="() => (isCollapsed = false)"
    >
      <SideMenu :collapsed="isCollapsed" @toggle="(toggle) => (isCollapsed = toggle)" />
    </NLayoutSider>
    <NLayout class="grow" content-style="display: flex; flex-direction: column;">
      <NLayoutContent
        class="grow flex flex-col"
        content-style="display: flex; flex-direction: column; flex-grow: 1;"
      >
        <MainContent />
      </NLayoutContent>
    </NLayout>
  </NLayout>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { NLayout, NLayoutContent, NLayoutSider, useMessage, useLoadingBar } from 'naive-ui/lib'
import { onMounted } from 'vue'
import SideMenu from './side-menu'
import MainContent from './main-content.vue'
import { useRouter } from 'vue-router'

const isCollapsed = ref(true)

onMounted(async () => {
  window.$loading = useLoadingBar()
  window.$message = useMessage()
  window.$router = useRouter()
})
</script>

<style lang="scss" scoped>
#menu {
  box-shadow: 2px 0px 10px 0px rgba($color: #000000, $alpha: 0.05);
}
</style>
