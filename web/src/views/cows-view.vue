<template>
  <div class="p-2 grow flex flex-col">
    <NCard class="m-1 shadow-sm grow">
      <div class="flex">
        <NSelect class="mt-2 mb-4 mr-1 w-24" :options="filterOpts" v-model:value="filter" />
        <NButton class="mt-2 mb-4 mx-1" @click="() => kill(checked)">杀死选中牲畜</NButton>
        <NButton class="mt-2 mb-4 mx-1" @click="() => spawn()">新增牲畜</NButton>
      </div>
      <NDataTable
        :remote="true"
        :columns="columns"
        :data="cows"
        :pagination="pagination"
        :row-key="(row: CowInfo) => row.id"
        :row-props="
          (_, idx) => {
            return {
              onContextmenu: (e: MouseEvent) => {
                e.preventDefault()
                showDropdown = false
                x = e.clientX
                y = e.clientY
                selected = idx
                nextTick(() => {
                  showDropdown = true
                })
              }
            }
          }
        "
        @update-checked-row-keys="(ids) => checked = ids as string[]"
      />
      <NDropdown
        placement="bottom-start"
        trigger="manual"
        :x="x"
        :y="y"
        :options="dropdownItems"
        :show="showDropdown"
        @clickoutside="() => (showDropdown = false)"
        @select="onSelected"
      />
    </NCard>
  </div>
</template>

<script setup lang="tsx">
import { DataTableColumns, NDataTable, NDropdown, DropdownOption, NCard } from 'naive-ui/lib'
import { CowInfo, CowQueryFilter, KillCow, spawnCow } from '../api/cow'
import { nextTick, onMounted, ref, watch } from 'vue'
import { fetchCowInfo } from '../api/cow'
import { NULL } from '../contansts'
import { RouterLink } from 'vue-router'
import { NButton, NIcon, NSelect, SelectOption } from 'naive-ui'
import { Information, OperationsField } from '@vicons/carbon'

const cows = ref<CowInfo[]>([])
const pagination = ref({
  page: 1,
  pageSize: 10,
  pageSizes: [2, 5, 10, 20, 40],
  itemCount: 0,
  showSizePicker: true,
  onUpdatePage(page: number) {
    pagination.value.page = page
    fetch()
  },
  onUpdatePageSize(pageSize: number) {
    pagination.value.pageSize = pageSize
    pagination.value.page = 1
    fetch()
  }
})
const x = ref(0)
const y = ref(0)
const showDropdown = ref(false)
const selected = ref(0)
const checked = ref<string[]>([])
const filter = ref(CowQueryFilter.Alive)

const fetch = async () => {
  const { data, cnt } = await fetchCowInfo({
    page: pagination.value.page,
    size: pagination.value.pageSize,
    filter: filter.value
  })
  cows.value = data
  pagination.value.itemCount = cnt
}

onMounted(() => fetch())

watch(filter, () => fetch())

const onSelected = (key: Entry) => {
  switch (key) {
    case Entry.Detail:
      window.$router.push({
        name: 'cow',
        params: {
          uuid: cows.value[selected.value].id
        }
      })
      break
    case Entry.Kill:
      kill([cows.value[selected.value].id])
      break
  }
}

const kill = (cows: string[]) =>
  KillCow(cows).then(() => {
    fetch()
    window.$message.success('操作成功')
  })

const spawn = () => spawnCow().then(() => window.$message.success('操作成功'))

const dropdownItems: DropdownOption[] = [
  {
    label: '详情',
    key: Entry.Detail,
    icon: () => (
      <NIcon>
        <Information />
      </NIcon>
    )
  },
  {
    label: '操作',
    key: Entry.Operation,
    icon: () => (
      <NIcon>
        <OperationsField />
      </NIcon>
    ),
    children: [
      {
        label: '杀死该牲畜',
        key: Entry.Kill,
        disabled:
          cows.value.at(selected.value) !== undefined &&
          cows.value.at(selected.value)?.dead_at === undefined
      }
    ]
  }
]
</script>

<script lang="tsx">
const columns: DataTableColumns<CowInfo> = [
  {
    type: 'selection'
  },
  {
    title: '名字',
    key: 'id',
    render: ({ id: uuid }) => (
      <RouterLink
        class="no-underline text-[#434c5e] hover:text-[#81a1c1]"
        to={{
          name: 'cow',
          params: {
            uuid
          }
        }}
      >
        {uuid}
      </RouterLink>
    )
  },
  {
    title: '出生时间',
    key: 'born_at'
  },
  {
    title: '母亲',
    key: 'parent',
    render: ({ parent: uuid }) =>
      uuid === NULL ? (
        ''
      ) : (
        <RouterLink
          class="no-underline text-[#434c5e] hover:text-[#81a1c1]"
          to={{
            name: 'cow',
            params: {
              uuid
            }
          }}
        >
          {uuid}
        </RouterLink>
      )
  },
  {
    title: '死亡时间',
    key: 'dead_at',
    render: (row) => row.dead_at ?? ''
  },
  {
    title: '死因',
    key: 'reason',
    render: (row) => row.reason ?? ''
  }
]

enum Entry {
  Detail,
  Operation,
  Kill
}

const filterOpts: SelectOption[] = [
  { label: CowQueryFilter.Alive, value: CowQueryFilter.Alive },
  { label: CowQueryFilter.Dead, value: CowQueryFilter.Dead },
  { label: CowQueryFilter.All, value: CowQueryFilter.All }
]
</script>
