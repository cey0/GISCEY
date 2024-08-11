<template>
  <AppLayout>
    <router-view />
  </AppLayout>
</template>

<script>
import AppLayout from './layouts/AppLayout.vue'

export default {
  name: 'App',
  components: {
    AppLayout
  }
}
</script>

<style>
/* Add any global styles here */
</style>