import { MenuOption, NIcon, NMenu } from 'naive-ui/lib'
import { defineComponent, ref } from 'vue'
import { UnorderedListOutlined } from '@vicons/antd'
import { Task, ChartArea, Wallet } from '@vicons/carbon'
import { MenuUnfoldOutlined, MenuFoldOutlined } from '@vicons/antd'
import { RouterLink } from 'vue-router'

const renderIcon = (icon: any) => () => <NIcon>{icon}</NIcon>
const renderRouterLink = (to: string, label: string) => () => (
  <RouterLink to={to}>{label}</RouterLink>
)

enum Keys {
  toggle = 'toggle',
  list = 'list',
  todo = 'todo',
  balance = 'balance',
  chart = 'chart'
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
        label: renderRouterLink('/todo', '待办'),
        key: Keys.todo,
        icon: renderIcon(<Task />)
      },
      {
        label: renderRouterLink('/group', '清单'),
        key: Keys.list,
        icon: renderIcon(<UnorderedListOutlined />)
      },
      {
        label: () => '收支',
        key: Keys.balance,
        icon: renderIcon(<Wallet />)
      },
      {
        label: () => '统计',
        key: Keys.chart,
        icon: renderIcon(<ChartArea />)
      }
    ]

    const handleClick = async (key: Keys) => {
      switch (key) {
        case Keys.toggle:
          isCollapsed.value = !isCollapsed.value
          emit('toggle', isCollapsed.value)
          break
        case Keys.todo:
          break
        case Keys.list:
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
