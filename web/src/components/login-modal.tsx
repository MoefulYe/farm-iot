import { Login, Register, type LoginReq, type RegisterReq } from '@/api/token'
import {
  NCard,
  NForm,
  NModal,
  type FormInst,
  NFormItem,
  NInput,
  NConfigProvider,
  NButton,
  NDivider
} from 'naive-ui/lib'
import { defineComponent, ref, render } from 'vue'
import theme from '../assets/nord.naiveui.json'

const login = defineComponent({
  setup(props, { emit }) {
    const model = ref<LoginReq>({ username: '', passwd: '' })
    const formRef = ref<FormInst | null>(null)
    const submit = async () => {
      await Login(model.value)
    }
    const validation = {
      username: {
        required: true,
        message: '请输入用户名',
        trigger: 'blur'
      },
      password: {
        required: true,
        message: '请输入密码',
        trigger: 'blur'
      }
    }
    return () => (
      <div class="flex flex-col items-center">
        <div class="text-lg pb-4">登录</div>
        <NForm
          ref={formRef}
          labelWidth="80"
          model={model.value}
          rules={validation}
          labelPlacement="left"
        >
          <NFormItem label="用户名" path="username">
            <NInput
              value={model.value.username}
              onUpdateValue={(v) => (model.value.username = v)}
            />
          </NFormItem>
          <NFormItem label="密码" path="password">
            <NInput value={model.value.passwd} onUpdateValue={(v) => (model.value.passwd = v)} />
          </NFormItem>
        </NForm>
        <div>
          <NButton
            class="m-2"
            onClick={async () => {
              await submit()
              window.$message.success('login success')
              emit('close')
            }}
          >
            登录
          </NButton>
          <NButton class="m-2" onClick={() => emit('register')}>
            注册
          </NButton>
        </div>
      </div>
    )
  }
})

const register = defineComponent({
  setup(props, { emit }) {
    const model = ref<RegisterReq>({ username: '', passwd: '' })
    const formRef = ref<FormInst | null>(null)
    const submit = async () => {
      await Register(model.value)
    }
    const validation = {
      username: {
        required: true,
        message: '请输入用户名',
        trigger: 'blur'
      },
      password: {
        required: true,
        message: '请输入密码',
        trigger: 'blur'
      }
    }
    return () => (
      <div class="flex flex-col items-center">
        <div class="text-lg pb-4">注册</div>
        <NForm
          ref={formRef}
          labelWidth="80"
          model={model.value}
          rules={validation}
          labelPlacement="left"
        >
          <NFormItem label="用户名" path="username">
            <NInput
              value={model.value.username}
              onUpdateValue={(v) => (model.value.username = v)}
            />
          </NFormItem>
          <NFormItem label="密码" path="password">
            <NInput value={model.value.passwd} onUpdateValue={(v) => (model.value.passwd = v)} />
          </NFormItem>
        </NForm>
        <div>
          <NButton
            class="m-2"
            onClick={async () => {
              await submit()
              window.$message.success('register success')
              emit('close')
            }}
          >
            注册
          </NButton>
          <NButton class="m-2" onClick={() => emit('login')}>
            登录
          </NButton>
        </div>
      </div>
    )
  }
})

const Modal = defineComponent({
  setup() {
    const show = ref(true)
    const status = ref<'login' | 'register'>('login')
    return () => (
      <NConfigProvider themeOverrides={theme}>
        <NModal show={show.value}>
          <NCard class="w-96 rounded-md">
            {status.value === 'login' ? (
              <login
                onRegister={() => (status.value = 'register')}
                onClose={() => {
                  window.$router.go(-1)
                  show.value = false
                }}
              />
            ) : (
              <register
                onLogin={() => (status.value = 'login')}
                onClose={() => {
                  window.$router.go(-1)
                  show.value = false
                }}
              />
            )}
          </NCard>
        </NModal>
      </NConfigProvider>
    )
  }
})

export default async () => {
  render(<Modal />, document.body)
}
