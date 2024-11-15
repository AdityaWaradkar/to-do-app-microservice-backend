import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';

export default defineConfig({
  plugins: [react()],
  base: '/to-do-app-microservice-backend/',  // Correct base URL for GitHub Pages
});
