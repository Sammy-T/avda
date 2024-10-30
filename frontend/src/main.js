import './style.css';
import { mount } from 'svelte';
import App from './lib/App.svelte';

const app = mount(App, {
  target: document.getElementById('app')
});

export default app;
