<template>
    <div>
      <h1>Restaurant Map</h1>
      <MapComponent />
    </div>
  </template>
  
  <script>
  import MapComponent from '../components/MapComponent.vue';
  
  export default {
    name: 'HomeView',
    components: {
      MapComponent,
    },
  };
  </script>
  
  <style>
  /* Add any global styles here */
  </style>
  