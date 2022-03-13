<script setup lang="ts">
// This starter template is using Vue 3 <script setup> SFCs
// Check out https://v3.vuejs.org/api/sfc-script-setup.html#sfc-script-setup
import Header from './components/Header.vue'
import TextArea from './components/TextArea.vue';
import FileArea from './components/FileArea.vue';
</script>

<template>
  <div class="bg-gray-50 h-screen flex flex-col">
    <Header></Header>
    <FileArea></FileArea>
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
