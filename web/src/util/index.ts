export const importRemoteJs = (url: string) => {
  return new Promise((resolve, reject) => {
    const script = document.createElement('script')
    script.src = url
    script.onload = () => {
      resolve(url)
    }
    script.onerror = (event, source, lineno, colno, error) => {
      reject(error)
    }
    document.body.appendChild(script)
  })
}
