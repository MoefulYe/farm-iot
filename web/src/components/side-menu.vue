<template>
  <NMenu
    :collapsed="collapsed"
    :collapsed-width="56"
    :collapsed-icon-size="22"
    :options="entries"
    @update-value="click"
  />
</template>

<script setup lang="tsx">
import { type MenuOption, NIcon, NMenu } from 'naive-ui/lib'
import { RouterLink } from 'vue-router'
import { useTokenStore } from '../stores/token'
import loginModal from './login-modal'
import { ref } from 'vue'
import { MenuFoldOutlined, MenuUnfoldOutlined } from '@vicons/antd'
import { Dashboard, EdgeCluster, Exit, Settings } from '@vicons/carbon'

const emit = defineEmits<{
  toggle: [value: boolean]
}>()

const collapsed = ref(true)

const click = async (entry: Entry) => {
  switch (entry) {
    case Entry.Quit:
      useTokenStore().clearToken()
      window.$message.success('exit success')
      loginModal()
      break
    case Entry.Toggle:
      collapsed.value = !collapsed.value
      emit('toggle', collapsed.value)
      break
  }
}

const entries: MenuOption[] = [
  {
    label: collapsed.value ? '展开' : '收起',
    key: Entry.Toggle,
    icon: collapsed.value ? renderIcon(<MenuUnfoldOutlined />) : renderIcon(<MenuFoldOutlined />)
  },
  {
    label: renderRouterLink('/dashboard', '概况'),
    key: Entry.Dashboard,
    icon: renderIcon(<Dashboard />)
  },
  {
    label: renderRouterLink('/cow', '牲畜管理'),
    key: Entry.Cows,
    icon: renderIcon(<EdgeCluster />)
  },
  {
    label: '设置',
    key: Entry.Settings,
    icon: renderIcon(<Settings />),
    children: [
      {
        label: '退出',
        key: Entry.Quit,
        icon: renderIcon(<Exit />)
      }
    ]
  }
]
</script>

<script lang="tsx">
const renderIcon = (icon: JSX.Element) => () => <NIcon>{icon}</NIcon>
const renderRouterLink = (to: string, label: string) => () => (
  <RouterLink to={to}>{label}</RouterLink>
)
enum Entry {
  Toggle,
  Dashboard,
  Settings,
  Cows,
  Quit
}
</script>
