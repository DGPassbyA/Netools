<script setup lang="ts">
// This starter template is using Vue 3 <script setup> SFCs
// Check out https://v3.vuejs.org/api/sfc-script-setup.html#sfc-script-setup
import Header from './components/Header.vue'

</script>

<template>
  <div class="absolute overflow-x-hidden h-screen w-screen z-0">
    <div class="bg-gray-50 h-screen flex flex-col">
      <Header></Header>
      <router-view></router-view>
    </div>
  </div>
</template>

<style>
#app {
  font-family: Helvetica, Avenir, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  height: 100%;
}
</style>
