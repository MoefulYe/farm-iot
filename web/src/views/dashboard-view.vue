<script setup lang="ts">
// import { computed, onMounted, ref } from 'vue'
// import { ComposeOption, use } from 'echarts/core'
// import { CanvasRenderer } from 'echarts/renderers'
// import { TooltipComponent, TooltipComponentOption } from 'echarts/components'
// import { ScatterSeriesOption, ScatterChart, CustomChart, CustomSeriesOption } from 'echarts/charts'
// import VChart from 'vue-echarts'
// import 'echarts/extension/bmap/bmap'
// import { CENTER, COW_SVG_PATH, POLYGON } from '../contansts'
// import { Field, fetchHeartbeat } from '../api/heartbeat'
// import { ECElementEvent } from 'echarts/types/dist/shared'
// import BmapTheme from '../assets/bmap.theme.json'

// type Opt = ComposeOption<ScatterSeriesOption | CustomSeriesOption | TooltipComponentOption>
// use([CanvasRenderer, ScatterChart, CustomChart, TooltipComponent])
// type Record = { name: string; value: [number, number] }

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
</script>

<template>ss</template>
