<script lang="ts">
import {defineComponent} from "vue";
export default defineComponent({
  name: "Header",
  data(){
      return {
          textGray: {
              'text-gray-300': true,
              'text-gray-800': false,
          },
          textBlack: {
              'text-gray-300': false,
              'text-gray-800': true,
          },
          textColor1: {},
          textColor2: {},
          textColor3: {},
          smallMenu: {
              "hidden": true,
          },
          isShow: false,
      }
  },
  mounted() {
      this.textColor1 = this.textBlack;
      this.textColor2 = this.textGray;
      this.textColor3 = this.textGray;
  },
  methods: {
      setActive: function(n : number) {
        switch(n){
            case 1:
                this.textColor2 = this.textColor3 = this.textGray;
                this.textColor1 = this.textBlack;
                break;
            case 2:
                this.textColor1 = this.textColor3 = this.textGray;
                this.textColor2 = this.textBlack;
                break;
            case 3:
                this.textColor2 = this.textColor1 = this.textGray;
                this.textColor3 = this.textBlack;
                break;
        }
      },
    showMenu: function() {
        console.log(this.isShow)
        this.smallMenu.hidden = this.isShow;
        this.isShow = !this.isShow;
    }
  },
});
</script>

<template>
    <div>
        <nav class="bg-white dark:bg-gray-800 shadow ">
            <div class="max-w-7xl mx-auto px-8">
                <div class="flex items-center justify-between h-16">
                    <div class="flex items-center">
                        <a class="flex-shrink-0" href="/">
                            <img class="h-8 w-8" src="../assets/logo.svg" alt="ClipBoard"/>
                        </a>
                        <div class="hidden md:block">
                            <div class="ml-10 flex items-baseline space-x-4">
                            <router-link to="/text">
                                <a
                                    class="text-gray-300 hover:text-gray-800 dark:hover:text-white px-3 py-2 rounded-md text-lg font-medium transition duration-500 ease-in-out"
                                    :class = "textColor1"
                                    @click="setActive(1)"
                                    href="/#/text"
                                >Text</a></router-link>
                                <router-link to="/file">
                                <a
                                    class="text-gray-300 hover:text-gray-800 dark:hover:text-white px-3 py-2 rounded-md text-lg font-medium transition duration-500 ease-in-out"
                                    :class = "textColor2"
                                    href="/#/file"
                                    @click="setActive(2)"
                                >File</a></router-link>
                                <router-link to="/list">
                                <a
                                    class="text-gray-300 hover:text-gray-800 dark:hover:text-white px-3 py-2 rounded-md text-lg font-medium transition duration-500 ease-in-out"
                                    :class = "textColor3"
                                    href="/#/list"
                                    @click="setActive(3)"
                                >List</a></router-link>
                            </div>
                        </div>
                    </div>
                    <div class="block">
                        <div class="ml-4 flex items-center md:ml-6"></div>
                    </div>
                    <div class="-mr-2 flex md:hidden">
                        <button
                            class="text-gray-800 dark:text-white hover:text-gray-300 inline-flex items-center justify-center p-2 rounded-md focus:outline-none"
                            @click="showMenu"
                        >
                            <svg
                                width="20"
                                height="20"
                                fill="currentColor"
                                class="h-8 w-8"
                                viewBox="0 0 1792 1792"
                                xmlns="http://www.w3.org/2000/svg"
                            >
                                <path
                                    d="M1664 1344v128q0 26-19 45t-45 19h-1408q-26 0-45-19t-19-45v-128q0-26 19-45t45-19h1408q26 0 45 19t19 45zm0-512v128q0 26-19 45t-45 19h-1408q-26 0-45-19t-19-45v-128q0-26 19-45t45-19h1408q26 0 45 19t19 45zm0-512v128q0 26-19 45t-45 19h-1408q-26 0-45-19t-19-45v-128q0-26 19-45t45-19h1408q26 0 45 19t19 45z"
                                />
                            </svg>
                        </button>
                    </div>
                </div>
            </div>
            <div class="md:hidden" :class="smallMenu">
                <div class="px-2 pt-2 pb-3 space-y-1 sm:px-3">
                    <a
                        class=" hover:text-gray-800 dark:hover:text-white block px-3 py-2 rounded-md text-base transition duration-500 ease-in-out"
                        href="/#/text"
                        :class = "textColor1"
                        @click="setActive(1)"
                    >Text</a>
                    <a
                        class=" hover:text-gray-800 dark:hover:text-white block px-3 py-2 rounded-md text-base font-medium transition duration-500 ease-in-out"
                        href="/#/file"
                        :class = "textColor2"
                        @click="setActive(2)"
                    >File</a>
                    <a
                        class=" hover:text-gray-800 dark:hover:text-white block px-3 py-2 rounded-md text-base font-medium transition duration-500 ease-in-out"
                        href="/#/list"
                        :class = "textColor3"
                        @click="setActive(3)"
                    >List</a>
                </div>
            </div>
        </nav>
    </div>
</template>


<style scoped>
</style>

