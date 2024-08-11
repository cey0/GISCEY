<template>
  <div :class="['sidebar-container', { '-translate-x-full': !isOpen }]" class="bg-customBrown ml-5 rounded-lg transform transition-transform duration-300 ease-in-out">
    <div class="sidebar-wrapper">
      <div class="sidebar-content flex flex-col justify-center align-middle  gap-5">
        <router-link to="/" class="text-white rounded-xl content p-5 text-lg font-bold mb-4">Beranda</router-link>
        <router-link to="/maps" class="text-white content rounded-xl p-5 text-lg font-bold mb-4">Peta</router-link>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    isOpen: {
      type: Boolean,
      default: false
    }
  }
};
</script>

<style scoped>
.sidebar-container {
  height: 88vh;
  width: 250px;
}
.content{
  background-color: #5FBDFF;
}
</style>
