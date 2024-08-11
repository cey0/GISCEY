<template>
  <div class="app-layout">
    <header class="header mt-3 rounded-xl bg-customBrown">
      <HeaderComponent @toggleSidebar="toggleSidebar" class="header" />
    </header>

    <div class="main-content">
      <SidebarComponent :isOpen="isSidebarOpen" class="sidebar" />
      <main class="content">
        <slot></slot>
      </main>
    </div>

    <footer class="footer">
      <!-- Add footer content here -->
    </footer>
  </div>
</template>

<script>
import SidebarComponent from '../components/SidebarComponent.vue';
import HeaderComponent from '../components/HeaderComponent.vue';

export default {
  name: 'AppLayout',
  components: {
    SidebarComponent,
    HeaderComponent,
  },
  data() {
    return {
      isSidebarOpen: false,
    };
  },
  methods: {
    toggleSidebar(isOpen) {
      this.isSidebarOpen = isOpen;
    },
  },
};
</script>

<style scoped>
.app-layout {
  display: flex;
  flex-direction: column;
  min-height: 100vh;
}

.main-content {
  display: flex;
  flex: 1;
  justify-content: center;
  align-items: stretch; /* Adjust to make content take up full height */
  padding-top: 10px;
}

.sidebar {
  width: 250px;
}

.content {
  flex: 1;
  padding: 20px;
}


.header {
  width: 98vw;
  margin: 10px auto;
}

.footer {
  width: 100%;
  padding: 10px;
  background-color: #f0f0f0;
}


</style>
