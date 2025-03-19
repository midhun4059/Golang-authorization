<script>
import { ref, onMounted } from "vue";

export default {
  name: "Dashboard",
  setup() {
    const name = ref(""); // Reactive variable

    onMounted(() => {
      name.value = localStorage.getItem("username") || "Guest"; // Default value if null
    });

    return { name };
  },
};
</script>

<style scoped>
h1 {
  text-align: center;
  margin-top: 20px;
  color: #333;
}
</style>
  <template>
    <div>
      <h1>Welcome back, {{ name }}!</h1>
    </div>
  </template>