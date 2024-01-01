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
      <p class="text-lg">三日内收支情况</p>
      <VChart class="h-96" :option="lineOpts" autoresize />
    </NCard>
    <NCard class="grow shadow-sm m-2">
      <VChart :option="mapOpts" autoresize @mousedown="({ name }) => gotoCowView(name)" />
    </NCard>
  </div>
</template>

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
  GridComponentOption,
  DataZoomInsideComponent,
  DataZoomComponentOption,
  LegendComponent,
  LegendComponentOption,
  TitleComponent,
  TitleComponentOption
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
import { CENTER, POLYGON } from '../contansts'
import { Field, fetchHeartbeat } from '../api/heartbeat'
import BmapTheme from '../assets/bmap.theme.json'

type MapOpts = ComposeOption<
  ScatterSeriesOption | CustomSeriesOption | TooltipComponentOption | TitleComponentOption
>
type LineOpts = ComposeOption<
  | LineSeriesOption
  | TooltipComponentOption
  | GridComponentOption
  | DataZoomComponentOption
  | LegendComponentOption
>
use([
  CanvasRenderer,
  ScatterChart,
  CustomChart,
  TooltipComponent,
  LineChart,
  GridComponent,
  DataZoomInsideComponent,
  LegendComponent,
  TitleComponent
])

const mapOpts = computed<MapOpts>(() => ({
  title: {
    text: '实时牲畜分布',
    left: 'center'
  },
  bmap: {
    center: CENTER,
    zoom: 16,
    roam: true,
    mapStyle: {
      styleJson: BmapTheme
    }
  },
  tooltip: {
    trigger: 'item'
  },
  series: [
    {
      type: 'scatter',
      coordinateSystem: 'bmap',
      data: cows.value.map((cow) => ({
        value: [cow.longitude, cow.latitude, cow.weight],
        name: cow.id
      })),
      itemStyle: {
        color: '#3b4252'
      },
      tooltip: {
        formatter: ({ name, value }) => {
          const [x, y] = value as [number, number]
          return `名字: ${name}<br/> 位置: (${x.toFixed(4)}, ${y.toFixed(4)})`
        }
      },
      symbolSize: (value: [number, number, number], _) => weight2size(value[2])
    },
    {
      type: 'custom',
      coordinateSystem: 'bmap',
      renderItem: (_, api) => {
        let points = POLYGON.map((p) => api.coord(p))
        return {
          type: 'polygon',
          shape: { points },
          style: { stroke: '#81a1c11c', fill: '#81a1c103' }
        }
      },
      animation: false,
      silent: true,
      data: POLYGON,
      z: -10
    }
  ]
}))

const cowCnt = ref(0)
const balances = ref<Balance[]>([])
const cows = ref<Cow[]>([])

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
  dataZoom: {
    type: 'inside'
  },
  legend: {},
  series: [
    {
      type: 'line',
      name: '收入',
      data: balances.value.map((balance) => [balance.when, balance.in]),
      smooth: true,
      showSymbol: false,
      areaStyle: {
        opacity: 0.3
      }
    },
    {
      type: 'line',
      name: '支出',
      data: balances.value.map((balance) => [balance.when, balance.out]),
      smooth: true,
      showSymbol: false,
      areaStyle: {
        opacity: 0.3
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
    fields: [Field.Longitude, Field.Latitude, Field.Weight]
  }).then((ok) => (cows.value = <Cow[]>ok))
})
</script>

<script lang="ts">
interface Cow {
  id: string
  longitude: number
  latitude: number
  weight: number
}

const formatNumber = (num: number): string =>
  num.toFixed(2).replace(/\d+/, (n) => n.replace(/(\d)(?=(\d{3})+$)/g, ($1) => $1 + ','))

const formatInteger = (num: number): string =>
  num.toString().replace(/\d+/, (n) => n.replace(/(\d)(?=(\d{3})+$)/g, ($1) => $1 + ','))

const threeDaysAgo = () =>
  dayjs().set('hour', 0).set('minute', 0).set('second', 0).set('millisecond', 0).subtract(3, 'day')

const gotoCowView = (uuid: string) => {
  window.$router.push({ name: 'cow', params: { uuid } })
}

const gotoBalanceView = () => window.$router.push({ name: 'balance' })

const weight2size = (weight: number) => {
  weight = Math.min(weight, 150)
  return weight / 15 + 5
}
</script>
