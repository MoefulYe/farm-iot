<template>
  <NDataTable
    :remote="true"
    :columns="columns"
    :data="cows"
    :pagination="pagination"
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
</template>

<script setup lang="ts">
import { DataTableColumns, NDataTable, NDropdown, DropdownOption } from 'naive-ui/lib'
import { CowInfo } from '../api/cow'
import { nextTick, onMounted, ref } from 'vue'
import { getCowInfo } from '../api/cow'

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

const fetch = async () => {
  const { data, cnt } = await getCowInfo({
    page: pagination.value.page,
    size: pagination.value.pageSize
  })
  cows.value = data
  pagination.value.itemCount = cnt
}

onMounted(async () => {
  await fetch()
})

const x = ref(0)
const y = ref(0)
const showDropdown = ref(false)
const selected = ref(0)

const onSelected = (key: string) => {
  if (key === 'stat') {
    window.$router.push({
      name: 'stat',
      params: {
        uuid: cows.value[selected.value].id
      }
    })
  }
}
</script>

<script lang="ts">
const columns: DataTableColumns<CowInfo> = [
  { title: '名字', key: 'id' },
  {
    title: '出生时间',
    key: 'born_at',
    render(row) {
      return row.born_at.toISOString()
    }
  },
  {
    title: '死亡时间',
    key: 'dead_at',
    render(row) {
      return row.dead_at?.toISOString() ?? ''
    }
  },
  {
    title: '死因',
    key: 'reason',
    render(row) {
      return row.reason ?? ''
    }
  }
]

const dropdownItems: DropdownOption[] = [{ label: '统计', key: 'stat' }]
</script>
