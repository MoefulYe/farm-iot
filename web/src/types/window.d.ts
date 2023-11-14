import { MessageApiInjection } from 'naive-ui/lib/message/src/MessageProvider'
import { LoadingBarApiInjection } from 'naive-ui/lib/loading-bar/src/LoadingBarProvider'
import { Router } from 'vue-router'

declare global {
  interface Window {
    $message: MessageApiInjection
    $loading: LoadingBarApiInjection
    $router: Router
  }
}
