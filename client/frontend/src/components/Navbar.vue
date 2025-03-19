<template>
  <nav class="main-header navbar navbar-expand navbar-white navbar-light">
    <ul class="navbar-nav">
      <li class="nav-item">
        <a class="nav-link" data-widget="pushmenu" href="#"><i class="fas fa-bars"></i></a>
      </li>
      <li class="nav-item d-none d-sm-inline-block">
        <a href="#" class="nav-link">{{ currentMenu }}</a>  <!-- Display selected menu -->
      </li>
    </ul>
  </nav>
</template>

<script>
export default {
  props: {
    currentMenu: String,  // Receive dynamic menu name
  },
};
</script>