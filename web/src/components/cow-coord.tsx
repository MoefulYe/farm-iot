import { computed, defineComponent } from 'vue'
import { ComposeOption, use } from 'echarts/core'
import { CanvasRenderer } from 'echarts/renderers'
import { TooltipComponent, TooltipComponentOption } from 'echarts/components'
import { CustomChart, CustomSeriesOption } from 'echarts/charts'
import { CENTER, POLYGON } from '../contansts'
import VChart from 'vue-echarts'
import 'echarts/extension/bmap/bmap'
import { useRoute } from 'vue-router'
import { Dayjs } from 'dayjs'
import { GetKeepAliveByUuid } from '../api/keep_alive'

type ChartOpts = ComposeOption<CustomSeriesOption | TooltipComponentOption>
type Record = {
  time: Dayjs
  value: [number, number]
}

export default defineComponent({
  name: 'CowCoord',
  setup() {
    use([CanvasRenderer, CustomChart, TooltipComponent])
    const uuid = useRoute().params.uuid as string
    const fetch = async () => GetKeepAliveByUuid(uuid, {})
    const chartOpts = computed<ChartOpts>(() => {
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
            type: 'custom',
            coordinateSystem: 'bmap',
            renderItem: (_, api) => {
              const points = POLYGON.map((p) => api.coord(p))
              return {
                type: 'polygon',
                shape: { points },
                style: { fill: '#8fbcbb', opacity: 0.1 }
              }
            },
            animation: false,
            silent: true,
            data: POLYGON,
            z: -10
          }
        ]
      }
    })
    return () => (
      <div class="h-full">
        <VChart class="h-full" option={chartOpts.value} />
      </div>
    )
  }
})
