<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import { ComposeOption, use } from 'echarts/core'
import { CanvasRenderer } from 'echarts/renderers'
import { TooltipComponent, TooltipComponentOption } from 'echarts/components'
import { ScatterSeriesOption, ScatterChart, CustomChart, CustomSeriesOption } from 'echarts/charts'
import VChart from 'vue-echarts'
import 'echarts/extension/bmap/bmap'
import { CENTER, POLYGON } from '../contanst'
import { Field, GetKeepAlive } from '../api/keep_alive'

type Opt = ComposeOption<ScatterSeriesOption | CustomSeriesOption | TooltipComponentOption>
use([CanvasRenderer, ScatterChart, CustomChart, TooltipComponent])
type Record = { name: string; value: [number, number] }

const fetch = async (): Promise<void> => {
  loading.value = true
  return await GetKeepAlive({
    start: '-5m',
    fields: [Field.longitude, Field.latitude]
  }).then((ok) => {
    data.value = ok.map((item) => ({
      name: item.id,
      value: [item.longitude!, item.latitude!]
    }))
    loading.value = false
  })
}

onMounted(async () => {
  await fetch()
})

const loadingOpts = {
  text: 'loading...',
  color: '#a3be8c',
  mask: 'rgba(255, 255, 255, 0.4)'
}
const data = ref<Record[]>([])
const loading = ref(false)
const opt = computed<Opt>(() => {
  return {
    bmap: {
      center: CENTER,
      zoom: 16,
      roam: true,
      mapStyle: {
        styleJson: []
      }
    },
    tooltip: {
      trigger: 'item'
    },
    series: [
      {
        type: 'scatter',
        coordinateSystem: 'bmap',
        data: data.value,
        symbolSize: 18,
        label: {
          formatter: '{a}',
          position: 'right'
        },
        itemStyle: {
          color: '#d08770'
        }
      },
      {
        type: 'custom',
        coordinateSystem: 'bmap',
        renderItem: (_, api) => {
          let points = POLYGON.map((p) => api.coord(p))
          return { type: 'polygon', shape: { points }, style: { fill: '#8fbcbb', opacity: 0.1 } }
        },
        animation: false,
        silent: true,
        data: POLYGON,
        z: -10
      }
    ]
  }
})
</script>

<template>
  <v-chart :option="opt" :loading="loading" :loading-options="loadingOpts" />
</template>
