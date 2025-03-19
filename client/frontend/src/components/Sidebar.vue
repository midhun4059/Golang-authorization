<script>
import SidebarItem from "./SidebarItem.vue";

export default {
  components: {
    SidebarItem,
  },
  methods: {
    updateNavbarText(newText) {  //passing data to parent sidebar component
      this.$emit("update-navbar", newText); // Emit event to App.vue
    },
  },
};
</script>


<template>
  <aside class="main-sidebar sidebar-dark-primary elevation-4">
    <a href="#" class="brand-link">
      <img src="/dist/img/AdminLTELogo.png" alt="Logo" class="brand-image img-circle elevation-3">
      <span class="brand-text font-weight-light">User Authorization</span>
    </a>

    <div class="sidebar">
      <nav class="mt-2">
        <ul class="nav nav-pills nav-sidebar flex-column">
          <!-- Dashboard -->
          <SidebarItem name="Dashboard" icon="fas fa-bars" path="/dashboard" @update-sidebar="updateNavbarText" />

          <!-- User Table -->
          <SidebarItem name="User Table" icon="fas fa-user" path="/user-table" @update-sidebar="updateNavbarText" />

          <!-- Logout -->
          <!-- the update navbar sends data from  sidebar item to sidebar -->
          <SidebarItem name="Logout" icon="fas fa-sign-out-alt" path="/logout" @update-sidebar="updateNavbarText" />
        </ul>
      </nav>
    </div>
  </aside>
</template>
