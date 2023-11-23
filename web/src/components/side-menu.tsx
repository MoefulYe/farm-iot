import { type MenuOption, NIcon, NMenu } from 'naive-ui/lib'
import { defineComponent, ref } from 'vue'
import { ClockCircleOutlined, UnorderedListOutlined } from '@vicons/antd'
import { Task, ChartArea, Wallet, Settings } from '@vicons/carbon'
import { ExitOutline, StatsChart } from '@vicons/ionicons5'
import { MenuUnfoldOutlined, MenuFoldOutlined } from '@vicons/antd'
import { Pig } from '@vicons/tabler'
import { RouterLink } from 'vue-router'
import { useTokenStore } from '@/stores/token'
import loginModal from './login-modal'

const renderIcon = (icon: any) => () => <NIcon>{icon}</NIcon>
const renderRouterLink = (to: string, label: string) => () => (
  <RouterLink to={to}>{label}</RouterLink>
)

enum Keys {
  toggle = 'toggle',
  realtime = 'realtime',
  settings = 'settings',
  cows = 'cows',
  quit = 'quit'
}

export default defineComponent({
  setup(props, { emit }) {
    const isCollapsed = ref(true)
    const menuItems: MenuOption[] = [
      {
        label: () => <>{isCollapsed.value == true ? '展开' : '收起'}</>,
        key: Keys.toggle,
        icon: () => (
          <>
            {isCollapsed.value == true ? (
              <NIcon>
                <MenuUnfoldOutlined />
              </NIcon>
            ) : (
              <NIcon>
                <MenuFoldOutlined />
              </NIcon>
            )}
          </>
        )
      },
      {
        label: renderRouterLink('/real-time', '实时'),
        key: Keys.realtime,
        icon: renderIcon(<ClockCircleOutlined />)
      },
      {
        label: renderRouterLink('/cow-info', '牲畜'),
        key: Keys.cows,
        icon: renderIcon(<Pig />)
      },
      {
        label: () => '设置',
        key: Keys.settings,
        icon: renderIcon(<Settings />),
        children: [
          {
            label: () => '退出',
            key: Keys.quit,
            icon: renderIcon(<ExitOutline />)
          }
        ]
      }
    ]

    const handleClick = async (key: Keys) => {
      switch (key) {
        case Keys.quit:
          useTokenStore().clearToken()
          window.$message.success('exit success')
          loginModal()
        case Keys.toggle:
          isCollapsed.value = !isCollapsed.value
          emit('toggle', isCollapsed.value)
          break
      }
    }

    return () => (
      <NMenu
        collapsed={isCollapsed.value}
        collapsedWidth={64}
        collapsedIconSize={22}
        options={menuItems}
        onUpdate:value={handleClick}
      />
    )
  }
})
