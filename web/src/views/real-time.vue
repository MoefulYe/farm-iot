<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { ComposeOption, use } from 'echarts/core'
import { CanvasRenderer } from 'echarts/renderers'
import { ScatterSeriesOption, ScatterChart } from 'echarts/charts'
import VChart from 'vue-echarts'
import 'echarts/extension/bmap/bmap'
import { CENTER } from '../contanst'
import { Field, GetKeepAlive } from '../api/keep_alive'

const fetch = async () => {
    return await GetKeepAlive({
        start: "-5m",
        fields: [Field.longtitude, Field.latitude]
    }).then(data => data.map(
        (item) => ({
            name: item.id,
            value: [item.longtitude!, item.latitude!]
        })
    ))
}

onMounted(async () => {
    await fetch()
})

use([CanvasRenderer, ScatterChart])
type Opt = ComposeOption<ScatterSeriesOption>

const option = ref<Opt>({
    bmap: {
        center: CENTER,
        zoom: 100,
        roam: false,
        mapStyle: {
            styleJson: []
        }
    },
})
</script>

<template>
    <v-chart :option="option" />
</template>