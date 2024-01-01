<script setup lang="ts">
import { NCard, NDivider, NStatistic } from 'naive-ui/lib'
import { computed, onMounted, ref } from 'vue'
import { CowQueryFilter, fetchCowInfo } from '../api/cow'
import { type Balance, fetchBalance } from '../api/balance'
import dayjs from 'dayjs'

import { ComposeOption, use } from 'echarts/core'
import { CanvasRenderer } from 'echarts/renderers'
import {
  TooltipComponent,
  TooltipComponentOption,
  GridComponent,
  GridComponentOption
} from 'echarts/components'
import {
  ScatterSeriesOption,
  ScatterChart,
  CustomChart,
  CustomSeriesOption,
  LineChart,
  LineSeriesOption
} from 'echarts/charts'
import VChart from 'vue-echarts'
import 'echarts/extension/bmap/bmap'
import { CENTER, COW_SVG_PATH, POLYGON } from '../contansts'
import { Field, fetchHeartbeat } from '../api/heartbeat'
import { ECElementEvent } from 'echarts/types/dist/shared'
import BmapTheme from '../assets/bmap.theme.json'

type MapOpts = ComposeOption<ScatterSeriesOption | CustomSeriesOption | TooltipComponentOption>
type LineOpts = ComposeOption<LineSeriesOption | TooltipComponentOption | GridComponentOption>
use([CanvasRenderer, ScatterChart, CustomChart, TooltipComponent, LineChart, GridComponent])

// const fetch = async (): Promise<void> => {
//   loading.value = true
//   return await fetchHeartbeat({
//     start: '-5m',
//     fields: [Field.longitude, Field.latitude]
//   }).then((ok) => {
//     data.value = ok.map((item) => ({
//       name: item.id,
//       value: [item.longitude!, item.latitude!]
//     }))
//     loading.value = false
//   })
// }

// onMounted(async () => {
//   await fetch()
// })

// const loadingOpts = {
//   text: 'loading...',
//   color: '#a3be8c',
//   mask: 'rgba(255, 255, 255, 0.4)'
// }
// const data = ref<Record[]>([])
// const loading = ref(false)
// const opt = computed<Opt>(() => {
//   return {
//     bmap: {
//       center: CENTER,
//       zoom: 16,
//       roam: true,
//       mapStyle: {
//         styleJson: BmapTheme
//       }
//     },
//     tooltip: {
//       trigger: 'item'
//     },
//     series: [
//       {
//         type: 'scatter',
//         coordinateSystem: 'bmap',
//         data: data.value,
//         symbol: COW_SVG_PATH,
//         symbolSize: 36,
//         label: {
//           formatter: (params) => params.name.substring(0, 5) + '...',
//           position: 'right',
//           show: true
//         },
//         itemStyle: {
//           color: '#3b4252'
//         },
//         tooltip: {
//           formatter: ({ name, value }) => {
//             const [x, y] = value as [number, number]
//             return `名字: ${name}<br/> 位置: (${x.toFixed(4)}, ${y.toFixed(4)})`
//           }
//         }
//       },
//       {
//         type: 'custom',
//         coordinateSystem: 'bmap',
//         renderItem: (_, api) => {
//           let points = POLYGON.map((p) => api.coord(p))
//           return { type: 'polygon', shape: { points }, style: { fill: '#81a1c1', opacity: 0.1 } }
//         },
//         animation: false,
//         silent: true,
//         data: POLYGON,
//         z: -10
//       }
//     ]
//   }
// })

// const onChartClick = (params: ECElementEvent) => {
//   window.$router.push({ name: 'stat', params: { uuid: params.name } })
// }

const cowCnt = ref(0)
const balances = ref<Balance[]>([])
const income = computed(() => balances.value.reduce((acc, cur) => acc + cur.in - cur.out, 0))

const lineOpts = computed<LineOpts>(() => ({
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
  series: [
    {
      type: 'line',
      name: '收入',
      data: balances.value.map((balance) => [balance.when, balance.in]),
      smooth: true,
      showSymbol: false,
      areaStyle: {
        opacity: 0.1
      }
    },
    {
      type: 'line',
      name: '支出',
      data: balances.value.map((balance) => [balance.when, balance.out]),
      smooth: true,
      showSymbol: false,
      areaStyle: {
        opacity: 0.1
      }
    }
  ]
}))

onMounted(() => {
  fetchCowInfo({
    filter: CowQueryFilter.Alive,
    size: 0,
    page: 1
  }).then(({ cnt }) => (cowCnt.value = cnt))
  fetchBalance({
    from: threeDaysAgo().format()
  }).then((value) => (balances.value = value))
  fetchHeartbeat({
    start: '-5m',
    fields: [Field.longitude, Field.latitude]
  })
})
</script>

<script lang="ts">
const formatNumber = (num: number): string =>
  num.toFixed(2).replace(/\d+/, (n) => n.replace(/(\d)(?=(\d{3})+$)/g, ($1) => $1 + ','))

const formatInteger = (num: number): string =>
  num.toString().replace(/\d+/, (n) => n.replace(/(\d)(?=(\d{3})+$)/g, ($1) => $1 + ','))

const threeDaysAgo = () =>
  dayjs().set('hour', 0).set('minute', 0).set('second', 0).set('millisecond', 0).subtract(3, 'day')
</script>

<template>
  <div class="p-2 grow flex">
    <NCard class="shadow-sm m-2">
      欢迎来到智能牧场管理系统
      <NDivider />
      <div class="flex justify-around">
        <NStatistic label="存栏量" class="inline-block">
          <template #default>
            {{ formatInteger(cowCnt) }}
          </template>
          <template #suffix>
            <span>头</span>
          </template>
        </NStatistic>
        <NStatistic label="近三日收支" class="inline-block">
          <template #default>
            {{ formatNumber(income) }}
          </template>
          <template #suffix>
            <span>元</span>
          </template>
        </NStatistic>
      </div>
      <NDivider />
      <VChart class="h-96" :option="lineOpts" autoresize />
    </NCard>
    <NCard class="grow shadow-sm m-2"></NCard>
  </div>
</template>
