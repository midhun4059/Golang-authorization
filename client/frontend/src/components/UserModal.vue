<template>
  <div class="modal fade show" id="modal-lg" style="display: block; padding-right: 15px;" aria-modal="true" role="dialog">
    <div class="modal-dialog modal-lg">
      <div class="modal-content">
        <div class="modal-header">
          <h4 class="modal-title">{{ isEditing ? "Change User Data" : "Add a new User" }}</h4>
          <button type="button" class="close" @click="closeModal" aria-label="Close">
           
          </button>
        </div>

        <div class="modal-body">
          <form @submit.prevent="handleUserSubmit">
           

            <div class="form-group">
              <label for="userName">Name</label>
              <input type="text" id="userName" class="form-control" v-model="userData.Username" required />
            </div>

            <div class="form-group">
              <label for="userEmail">Email</label>
              <input type="email" id="userEmail" class="form-control" v-model="userData.Email" required />
            </div>

            <div v-if="!isEditing" class="form-group" >
              <label for="userEmail">Password</label>
              <input type="password" id="userPassword" class="form-control" v-model="userData.Password" :required="!isEditing" :disabled="isEditing" />
            </div>

            <div class="form-group">
              <label for="EmployeePosition">Position</label>
              <select id="EmployeePosition" class="form-control" v-model.number="userData.RollID" required>
                <option value="0" disabled>Select Position</option>
                <option :value="1" >Super Admin</option>
                <option :value="2">Admin</option>
                <option :value="3">Front Office</option>
              </select>
            </div>
          </form>
        </div>

        <div class="modal-footer justify-content-between">
          <button type="button" class="btn btn-default" @click="closeModal">Close</button>
          <button type="button" class="btn btn-primary" @click="handleUserSubmit">
            {{ isEditing ? "Update User" : "Save User" }}
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import api from "/axios";

export default {
  props: {
    user: Object,
    isEditing: Boolean,
  },
  data() {
    return {
      userData: {
        ID: 0,
        Username:"",
        Email:"",
        EmployeePosition:"",
        Password:"",
        RollID:0,
      },
      
    };
  },
  watch: {
user: {
  immediate: true,
  handler(newUser) {
    if (newUser) {
      this.userData = { ...newUser };

      // Ensure RollID is always a number
      this.userData.RollID = Number(newUser.RollID) || 0;


    }
  },
},
},

  methods: {
    async handleUserSubmit() {
  try {
    const payload = {
          Username: this.userData.Username,
          Email: this.userData.Email,
          Password: this.userData.Password,
          RollID: this.userData.RollID, // Only update RollID, not EmployeePosition
      };

      if (this.isEditing) {
          await api.put(`/users/${this.userData.ID}`, payload); // Ensure ID is sent in the URL
      } else {
          await api.post("/register", payload);
      }

      this.$emit("userUpdated");
      this.closeModal();
  } catch (error) {
      console.error("Error saving user:", error);
  }
}
,
    closeModal() {
      this.$emit("close");
    },
  },
};
</script>

<style >
.modal {
  background: rgba(0, 0, 0, 0.8);
}

.disabled-field {
opacity: 0.6; /* Make it look disabled */
pointer-events: none; /* Prevent interaction */
}

</style>
