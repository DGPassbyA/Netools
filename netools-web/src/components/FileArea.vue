<template>
<div class="h-full w-full flex flex-col items-center">
    <div class="h-5 pb-10 pt-7 lg:w-2/3 md:w-4/5 sm:w-4/5 min:w-5/6 flex">
        <div class="h-5 text-xl pt-3">File Upload</div>
        <div class="h-5 flex-grow"></div>
        <div class="h-5">
            <button
            type="button"
            class="inline-block align-top px-6 py-1.5 bg-blue-600 text-white font-medium text-sm leading-6 uppercase rounded shadow-md hover:bg-blue-700 hover:shadow-lg focus:bg-blue-700 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-blue-800 active:shadow-lg transition duration-150 ease-in-out"
            @click="uploadFile">Upload</button>
        </div>
    </div>
    <div class="text-gray-70 h-1/3 lg:w-2/3 md:w-4/5 sm:w-4/5 min:w-5/6 m-1 border-dashed border-4 border-light-blue-500 flex flex-col justify-center items-center">
        <div class="mb-3 font-semibold text-gray-900">Drag and drop your files anywhere or</div>
        <input class="hidden" type="file" name="upload" id="upload" @change="uploadFile" multiple/>
        <button onclick="upload.click()" class="mt-2 rounded-sm px-3 py-1 bg-gray-200 hover:bg-gray-300 focus:shadow-outline focus:outline-none">Upload a file</button>
    </div>
    <div class="h-5 pb-10 pt-1 lg:w-2/3 md:w-4/5 sm:w-4/5 min:w-5/6">
        <div class="float-left text-xl pt-3 pb-2" style="width:100%;text-align: left">Upload Files</div>
        <div>
            <FileChip v-for="key,file in files" :content="file" @deleteItem="deleteChip"></FileChip>
        </div>
    </div>
</div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import FileChip  from './FileChip.vue';
export default defineComponent({
    name: "FileArea",
    components: { FileChip },
    data() {
        return {
            files: new Map(),
        }
    },
    methods: {
        uploadFile: function(e) {
            let fileList : FileList = e.srcElement.files;
            Array.from(fileList).forEach(file => {
                let randomStr =  Math.random().toString(16).slice(-8);
                this.files.set(file.name + "-" + randomStr, file);
            });
        },
        deleteChip: function(e : String) {
            console.log(this.files);
            this.files.delete(e);
            console.log(this.files);
        }
    },
});
</script>

<style scoped>
</style>