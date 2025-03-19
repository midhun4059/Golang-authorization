<template>
  <div id="app">
    <!-- Navbar and Sidebar appear only when logged in -->
    <Navbar v-if="isLoggedIn" :currentMenu="selectedSidebar" />
    <Sidebar v-if="isLoggedIn" @update-navbar="selectedSidebar = $event" />

    <main :class="{ 'with-sidebar': isLoggedIn }">
      <router-view></router-view> <!-- Dynamically load the content here -->
    </main>

    <Footer v-if="isLoggedIn"></Footer>
  </div>
</template>

<script>
import Navbar from "./components/Navbar.vue";
import Sidebar from "./components/Sidebar.vue";
import Footer from "./components/Footer.vue";


export default {
  components: {
    Navbar,
    Sidebar,
    Footer,
  },
  data() {
    return {
      selectedSidebar: "Dashboard",  // Default menu
      isLoggedIn: localStorage.getItem("access_token") !== null, // Initial check
    };
  },
  
  mounted() {
    window.addEventListener("storage", this.syncLoginState); // Listen for localStorage changes
  },
  beforeUnmount() {
    window.removeEventListener("storage", this.syncLoginState);
  },
  watch: {
    $route() {
      // Automatically check login state on route change
      this.isLoggedIn = localStorage.getItem("access_token") !== null;
    },
  },
  methods: {
    syncLoginState() {
      this.isLoggedIn = localStorage.getItem("access_token") !== null;
    },
    logout() {
      localStorage.removeItem("access_token");
      this.isLoggedIn = false;
      this.$router.push("/login");
    },
  },
};
</script>

<style>
/* Default style when sidebar is hidden */
main {
  padding: px;
  margin-left: 0; /* No margin when sidebar is hidden */
}

/* Apply left margin only when the user is logged in */
main.with-sidebar {
  margin-left: 230px; /* Adjust based on sidebar width */
}
</style>