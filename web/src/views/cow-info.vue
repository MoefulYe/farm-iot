<template>
    <NDataTable :remote="true" :columns="columns" :data="cows" :pagination="pagination" />
</template>

<script setup lang="ts">
import { DataTableColumns, NDataTable } from 'naive-ui/lib'
import { CowInfo } from '../api/cow'
import { onMounted, ref } from 'vue'
import { GetCowInfo } from '../api/cow'

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
    const { data, cnt } = await GetCowInfo({
        page: pagination.value.page,
        size: pagination.value.pageSize
    })
    cows.value = data
    pagination.value.itemCount = cnt
}

onMounted(async () => {
    await fetch()
})
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
</script>
