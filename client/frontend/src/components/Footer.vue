<script>
export default {
  name: "Footer",
};
</script>

<template>
  <footer class="main-footer">
    <strong>Copyright &copy; 2023-2024 
      <a href="https://iai.com">IAI</a>.
    </strong> 
    All rights reserved.
  </footer>
</template>

<style scoped>
.main-footer {
  position: fixed;
  bottom: 0;
  left: 0;
  width: 100%;
  background: #f8f9fa;
  text-align: center;
  padding: 10px;
  border-top: 1px solid #ddd;
  font-size: 14px;
}
</style>