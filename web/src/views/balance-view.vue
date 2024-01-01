<template>
  <div class="p-2 grow flex flex-col">
    <NCard class="m-1 shadow-sm">
      <template #header> 收支情况统计 </template>
      <NScrollbar class="max-h-64">
        <template #default>
          <NDatePicker class="p-2" type="datetimerange" v-model:value="range" />
          <NDataTable :data="data" :columns="columns" />
        </template>
      </NScrollbar>
    </NCard>
    <NCard class="m-1 shadow-sm grow">
      <VChart :option="opts" autoresize />
    </NCard>
  </div>
</template>

<script setup lang="ts">
import { NCard, NDatePicker, NDataTable, NScrollbar, DataTableColumn } from 'naive-ui'
import { onMounted, ref, watch, computed } from 'vue'
import dayjs from 'dayjs'
import { Balance, fetchBalance } from '../api/balance'

import { ComposeOption, use } from 'echarts/core'
import { CanvasRenderer } from 'echarts/renderers'
import {
  TooltipComponent,
  TooltipComponentOption,
  GridComponent,
  GridComponentOption,
  DataZoomInsideComponent,
  DataZoomComponentOption,
  LegendComponent,
  LegendComponentOption
} from 'echarts/components'
import { LineChart, LineSeriesOption } from 'echarts/charts'
import VChart from 'vue-echarts'
type Opts = ComposeOption<
  | LineSeriesOption
  | TooltipComponentOption
  | GridComponentOption
  | DataZoomComponentOption
  | LegendComponentOption
>
use([
  CanvasRenderer,
  TooltipComponent,
  LineChart,
  GridComponent,
  DataZoomInsideComponent,
  LegendComponent
])

const range = ref<[number, number]>([
  dayjs()
    .set('hour', 0)
    .set('minute', 0)
    .set('second', 0)
    .set('millisecond', 0)
    .subtract(3, 'day')
    .valueOf(),
  dayjs().valueOf()
])
const data = ref<Balance[]>([])
watch(range, () => fetch())
onMounted(() => fetch())

const opts = computed<Opts>(() => ({
  xAxis: {
    type: 'time',
    axisLabel: {
      show: false
    }
  },
  yAxis: {
    type: 'value'
  },
  tooltip: {
    trigger: 'axis'
  },
  dataZoom: {
    type: 'inside'
  },
  legend: {},
  series: [
    {
      type: 'line',
      name: '收入',
      data: data.value.map((balance) => [balance.when, balance.in]),
      smooth: true,
      showSymbol: false,
      areaStyle: {
        opacity: 0.3
      }
    },
    {
      type: 'line',
      name: '支出',
      data: data.value.map((balance) => [balance.when, balance.out]),
      smooth: true,
      showSymbol: false,
      areaStyle: {
        opacity: 0.3
      }
    }
  ]
}))

const columns: DataTableColumn<Balance>[] = [
  {
    title: '时间',
    key: 'when'
  },
  {
    title: '收入',
    key: 'in',
    render: (row) => format(row.in)
  },
  {
    title: '支出',
    key: 'out',
    render: (row) => format(row.out)
  },
  {
    title: '收支',
    key: 'in-out',
    render: (row) => format(row.in - row.out)
  }
]

const fetch = () =>
  fetchBalance({
    from: dayjs(range.value[0]).format(),
    to: dayjs(range.value[1]).format()
  }).then((ok) => (data.value = ok))

const format = (num: number): string =>
  num.toFixed(2).replace(/\d+/, (n) => n.replace(/(\d)(?=(\d{3})+$)/g, ($1) => $1 + ','))
</script>
