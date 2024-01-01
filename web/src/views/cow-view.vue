<script setup lang="ts">
import { NCard, NButton, NTabs, NTabPane } from 'naive-ui/lib'
import { useRoute } from 'vue-router'
import CowStat from '../components/cow-stat'
import CowCoord from '../components/cow-coord'
import { ref, watch } from 'vue'
import { CowInfoWithChildren, fetchCowInfoByUuid } from '../api/cow'
import { onMounted } from 'vue'
import { NULL } from '../contansts'
import { KillCow as killCow } from '../api/cow'

const route = useRoute()
const data = ref<CowInfoWithChildren>()
const uuid = ref<string>(route.params.uuid as string)

watch(
  () => route.params.uuid,
  () => {
    fetch()
    uuid.value = route.params.uuid as string
  }
)

const fetch = () => fetchCowInfoByUuid(route.params.uuid as string).then((ok) => (data.value = ok))
const kill = () =>
  killCow([data.value?.id!]).then(() => {
    fetch()
    window.$message.success('操作成功')
  })
onMounted(fetch)
</script>

<template>
  <div class="p-2 grow flex flex-col" v-if="data !== undefined">
    <NCard class="m-1 shadow-sm">
      <template #header> 基本信息 </template>
      <template #default>
        <div>编号：{{ data.id }}</div>
        <div>出生日期：{{ data.born_at }}</div>
        <div v-if="data.parent !== NULL">
          母亲：
          <RouterLink
            class="no-underline text-[#434c5e] hover:text-[#81a1c1]"
            v-if="data.parent !== NULL"
            :to="{
              name: 'cow',
              params: {
                uuid: data.parent
              }
            }"
            >{{ data.parent }}</RouterLink
          >
        </div>
        <div v-if="data.edges.children !== undefined">
          孩子：
          <RouterLink
            class="no-underline text-[#434c5e] hover:text-[#81a1c1]"
            v-for="{ id } in data.edges.children"
            :to="{
              name: 'cow',
              params: {
                uuid: id
              }
            }"
            >{{ id }}，</RouterLink
          >
        </div>
        <div v-if="data.dead_at !== undefined">
          <div>死亡日期：{{ data.dead_at }}</div>
          <div>死亡原因：{{ data.reason }}</div>
        </div>
        <div v-else>
          <div>指令下发：</div>
          <NButton class="mt-2" @click="kill"> 杀死该牲畜 </NButton>
        </div>
      </template>
    </NCard>
    <NCard class="m-1 grow shadow-sm">
      <template #header>数据统计</template>
      <template #default>
        <NTabs>
          <NTabPane name="统计">
            <CowStat />
          </NTabPane>
          <NTabPane name="位置">
            <CowCoord />
          </NTabPane>
        </NTabs>
      </template>
    </NCard>
  </div>
</template>
