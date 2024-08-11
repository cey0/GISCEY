<template>
    <div class="header-container">
      <div class="header-wrapper">
        <div class="header-content">
          <button @click="toggleMenu" class="burger-menu px-5">
            <div :class="{'change': isOpen}" class="menu-bar">
              <span class="block w-6 h-0.5 bg-black mb-1 transition-all duration-300"></span>
              <span class="block w-6 h-0.5 bg-black mb-1 transition-all duration-300"></span>
              <span class="block w-6 h-0.5 bg-black transition-all duration-300"></span>
            </div>
          </button>
        </div>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        isOpen: false,
      };
    },
    methods: {
      toggleMenu() {
        this.isOpen = !this.isOpen;
        this.$emit('toggleSidebar', this.isOpen); 
      },
    },
  };
  </script>
  
  <style scoped>
  .menu-bar {
    display: inline-block;
    cursor: pointer;
  }
  
  .menu-bar.change span:nth-child(1) {
    transform: translateY(6px) rotate(45deg);
  }
  
  .menu-bar.change span:nth-child(2) {
    opacity: 0;
  }
  
  .menu-bar.change span:nth-child(3) {
    transform: translateY(-6px) rotate(-45deg);
  }
  </style>
  