import { createApp } from 'vue'
import App from './App.vue'
import NumberDisplay from './components/NumberDisplay.vue' // Adjust the path according to your file structure

const app = createApp(App)
app.component('number-display', NumberDisplay)
app.mount('#app')
