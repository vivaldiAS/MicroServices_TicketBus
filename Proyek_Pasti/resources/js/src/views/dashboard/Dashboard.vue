<template>
  <div>
    <!-- <DashboardPenumpang></DashboardPenumpang> -->
    <div v-if="mountedUserRole === '2'">
      <DashboardPenumpang></DashboardPenumpang>
    </div>
    <div v-if="mountedUserRole === '3'">
      <DashboardDriver></DashboardDriver>
    </div>
    <div v-if="mountedUserRole === '4'">
      <DashboardAdmin></DashboardAdmin>
    </div>
    <div v-if="mountedUserRole === '1'">
      <DashboardAdminLoket></DashboardAdminLoket>
    </div>
    <div v-if="mountedUserRole === '5'">
      <DashboardDireksi></DashboardDireksi>
    </div>
  </div>
</template>
<script>

import DashboardAdmin from './DashboardAdmin.vue';
import DashboardPenumpang from './DashboardPenumpang.vue';
import DashboardDriver from './DashboardDriver.vue';
import DashboardAdminLoket from './DashboardAdminLoket.vue';
import DashboardDireksi from './DashboardDireksi.vue'

export default {
  data() {
    return {
      mountedUserRole: '' // Deklarasikan mountedUserRole dalam data
    }
  },
  components: {
    DashboardAdmin,
    DashboardPenumpang,
    DashboardDriver,
    DashboardAdminLoket,
    DashboardDireksi
  },
  computed: {
    userRole() {
      return this.$store.state.userRole
    },
  },
  mounted() {
    this.mountedUserRole = localStorage.getItem('userRole'); // Inisialisasi mountedUserRole di dalam mounted
    console.log('Role yang di mounted:', this.mountedUserRole);
  }
}
</script>
