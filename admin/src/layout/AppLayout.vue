<template>
  <div id="app-layout">
    <header>
      <!-- Header content -->
    </header>
    <aside>
      <!-- Sidebar content -->
    </aside>
    <main>
      <router-view />
    </main>
    <footer>
      <!-- Footer content -->
    </footer>
  </div>
</template>

<script>
export default {
  name: 'AppLayout',
};
</script>

<style scoped>
#app-layout {
  display: flex;
  flex-direction: column;
  height: 100vh;
}

header {
  /* Header styles */
}

aside {
  /* Sidebar styles */
}

main {
  flex: 1;
  /* Main content styles */
}

footer {
  /* Footer styles */
}
</style>
