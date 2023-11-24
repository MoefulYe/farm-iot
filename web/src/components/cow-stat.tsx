import { Field, FieldName, GetKeepAliveByUuid, type KeepAlive } from '../api/keep_alive'
import dayjs from 'dayjs'
import { NDatePicker, NPopover, NSelect, type SelectOption } from 'naive-ui/lib'
import { computed, defineComponent, onMounted, ref, watch, type Ref } from 'vue'
import { useRoute } from 'vue-router'
import { use } from 'echarts/core'
import { LineChart } from 'echarts/charts'
import {
  TooltipComponent,
  TitleComponent,
  ToolboxComponent,
  GridComponent,
  DataZoomComponent
} from 'echarts/components'
import { CanvasRenderer } from 'echarts/renderers'
import type { ComposeOption } from 'echarts/core'
import type { LineSeriesOption } from 'echarts/charts'
import {
  TooltipComponentOption,
  TitleComponentOption,
  ToolboxComponentOption,
  GridComponentOption,
  DataZoomComponentOption,
  LegendComponent,
  LegendComponentOption
} from 'echarts/components'
import { UniversalTransition } from 'echarts/features'
import VChart from 'vue-echarts'

type EChartsOption = ComposeOption<
  | TooltipComponentOption
  | TitleComponentOption
  | ToolboxComponentOption
  | GridComponentOption
  | DataZoomComponentOption
  | LineSeriesOption
  | LegendComponentOption
>

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

const fieldOpts: SelectOption[] = [
  { label: '经度', value: Field.longitude },
  { label: '维度', value: Field.latitude },
  { label: '健康', value: Field.health },
  { label: '体重', value: Field.weight }
]

export default defineComponent({
  name: 'StatView',
  setup: () => {
    use([
      TooltipComponent,
      TitleComponent,
      ToolboxComponent,
      GridComponent,
      DataZoomComponent,
      LineChart,
      CanvasRenderer,
      UniversalTransition,
      LegendComponent
    ])
    const uuid = useRoute().params.uuid as string
    const custom = ref<Range<number>>({
      start: dayjs().subtract(1, 'day').valueOf(),
      stop: dayjs().valueOf()
    })
    const rangeSelected = ref(RangeOptions.PastFifteenMinutes)
    const rangeOpts = rangeOptions(custom)
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
    const fieldSelected = ref<Field[]>([Field.health])
    const loading = ref(false)
    const data = ref<KeepAlive[]>([])
    const fetch = () => {
      loading.value = true
      GetKeepAliveByUuid(uuid, {
        fields: fieldSelected.value,
        ...rangeStr()
      }).then((ok) => {
        data.value = ok
        loading.value = false
      })
    }

    const chartOpts = computed<EChartsOption>(() => {
      return {
        tooltip: {
          trigger: 'axis',
          position: (pt) => [pt[0], '10%']
        },
        legend: {},
        xAxis: {
          type: 'time',
          boundaryGap: false
        },
        yAxis: {
          type: 'value',
          boundaryGap: [0, '100%'],
          max: 'dataMax',
          min: 'dataMin'
        },
        dataZoom: [{ type: 'inside', start: 0, end: 100 }, {}],
        series: fieldSelected.value.map((field) => {
          return {
            name: FieldName[field],
            type: 'line',
            symbol: 'none',
            smooth: true,
            data: data.value.map((item) => [item.time.valueOf(), item[field] as number])
          }
        })
      }
    })

    onMounted(() => fetch())
    watch([rangeSelected, fieldSelected], () => fetch())

    return () => (
      <div>
        <div class="flex">
          <NSelect
            class="w-48 p-2 inline"
            options={rangeOpts}
            v-model={[rangeSelected.value, 'value']}
          />
          <NSelect
            class="w-96 p-2 inline"
            multiple
            options={fieldOpts}
            value={fieldSelected.value}
            onUpdateValue={(val: Field[]) => {
              if (val.length > 0) {
                fieldSelected.value = val
              }
            }}
          />
        </div>
        <VChart option={chartOpts.value} class="h-96" loading={loading.value} />
      </div>
    )
  }
})
