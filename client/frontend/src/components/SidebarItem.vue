<script>
export default {
  props: {
    name: String,  // Menu name (Dashboard, User Table, Logout)
    icon: String,  // FontAwesome icon class
    path: String,  // Path for navigation
  },
  methods: {
    updateNavbar() {
      this.$emit("update-sidebar", this.name); // Emit event to update the Navbar title
    },
  },
};
</script>


<template>
  <li class="nav-item">
    <router-link :to="path" class="nav-link" @click="updateNavbar">
      <i class="nav-icon" :class="icon"></i>
      <p>{{ name }}</p>
    </router-link>
  </li>
</template>


<style scoped>
/* Highlight active menu */
.router-link-active {
  background: #007bff;
  color: white;
  border-radius: 5px;
}
</style>