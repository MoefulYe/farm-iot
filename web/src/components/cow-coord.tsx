import { computed, defineComponent, ref, watch, type Ref, onMounted } from 'vue'
import { type ComposeOption, use } from 'echarts/core'
import { CanvasRenderer } from 'echarts/renderers'
import { TooltipComponent, type TooltipComponentOption } from 'echarts/components'
import {
  CustomChart,
  type CustomSeriesOption,
  LinesChart,
  type LinesSeriesOption
} from 'echarts/charts'
import { CENTER, POLYGON } from '../contansts'
import VChart from 'vue-echarts'
import 'echarts/extension/bmap/bmap'
import { useRoute } from 'vue-router'
import dayjs from 'dayjs'
import { Field, fetchHeartbeatByUuid } from '../api/heartbeat'
import { NPopover, type SelectOption, NDatePicker, NSelect } from 'naive-ui'
import BmapTheme from '../assets/bmap.theme.json'

type ChartOpts = ComposeOption<CustomSeriesOption | TooltipComponentOption | LinesSeriesOption>
interface Range<T extends number | string> {
  start: T
  stop?: T
}

enum RangeOptions {
  Custom = 'custom',
  PastOneMinute = '-1m',
  PastFiveMinutes = '-5m',
  PastFifteenMinutes = '-15m',
  PastOneHour = '-1h',
  PastThreeHours = '-3h',
  PastSixHours = '-6h',
  PastTwelveHours = '-12h',
  PastOneDay = '-1d',
  PastTwoDays = '-2d',
  PastSevenDays = '-7d',
  PastOneMonth = '-30d'
}

const rangeOptions = (custom: Ref<Range<number>>): SelectOption[] => [
  {
    label: '自定义',
    value: RangeOptions.Custom,
    render: ({ node }) => (
      <NPopover trigger="hover" placement="right-start">
        {{
          trigger: () => node,
          default: () => (
            <div class="w-56">
              开始时间
              <NDatePicker class="p-2" type="datetime" v-model={[custom.value.start, 'value']} />
              结束时间
              <NDatePicker class="p-2" type="datetime" v-model={[custom.value.stop, 'value']} />
            </div>
          )
        }}
      </NPopover>
    )
  },
  {
    label: '过去1分钟',
    value: RangeOptions.PastOneMinute
  },
  {
    label: '过去5分钟',
    value: RangeOptions.PastFiveMinutes
  },
  {
    label: '过去15分钟',
    value: RangeOptions.PastFifteenMinutes
  },
  {
    label: '过去1小时',
    value: RangeOptions.PastOneHour
  },
  {
    label: '过去3小时',
    value: RangeOptions.PastThreeHours
  },
  {
    label: '过去6小时',
    value: RangeOptions.PastSixHours
  },
  {
    label: '过去12小时',
    value: RangeOptions.PastTwelveHours
  },
  {
    label: '过去1天',
    value: RangeOptions.PastOneDay
  },
  {
    label: '过去2天',
    value: RangeOptions.PastTwoDays
  },
  {
    label: '过去7天',
    value: RangeOptions.PastSevenDays
  },
  {
    label: '过去1个月',
    value: RangeOptions.PastOneMonth
  }
]

export default defineComponent({
  name: 'CowCoord',
  setup() {
    use([CanvasRenderer, CustomChart, TooltipComponent, LinesChart])
    const route = useRoute()
    const data = ref<[number, number][] | undefined>(undefined)
    const loading = ref(true)
    const chartOpts = computed<ChartOpts>(() => {
      return {
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
            type: 'lines',
            coordinateSystem: 'bmap',
            data:
              data.value != undefined
                ? [
                    {
                      coords: data.value,
                      lineStyle: {
                        color: '#a3be8c',
                        width: 2
                      },
                      name: `${(route.params.uuid as string).substring(0, 6)}...的运动轨迹`
                    }
                  ]
                : undefined
          }
        ]
      }
    })

    const custom = ref<Range<number>>({
      start: dayjs().subtract(1, 'day').valueOf(),
      stop: dayjs().valueOf()
    })
    const rangeOpts = rangeOptions(custom)
    const rangeSelected = ref(RangeOptions.PastFifteenMinutes)
    const rangeStr = (): Range<string> => {
      if (rangeSelected.value == RangeOptions.Custom) {
        return {
          start: dayjs(custom.value.start).format('YYYY-MM-DDTHH:mm:ssZ'),
          stop: dayjs(custom.value.stop).format('YYYY-MM-DDTHH:mm:ssZ')
        }
      } else {
        return {
          start: rangeSelected.value
        }
      }
    }
    const fetch = async () => {
      loading.value = true
      fetchHeartbeatByUuid(route.params.uuid as string, {
        fields: [Field.Longitude, Field.Latitude],
        ...rangeStr()
      }).then((ok) => {
        data.value = ok.map((item) => [item.longitude!, item.latitude!])
        loading.value = false
      })
    }
    onMounted(() => fetch())
    watch([rangeSelected, route], () => fetch())
    return () => (
      <div class="h-full">
        <NSelect class="w-48 pb-4" v-model={[rangeSelected.value, 'value']} options={rangeOpts} />
        <VChart class="h-96" option={chartOpts.value} loading={loading.value} />
      </div>
    )
  }
})
