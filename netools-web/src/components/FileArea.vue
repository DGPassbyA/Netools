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
        <input class="hidden" type="file" name="upload" id="upload" @change="addFile" multiple/>
        <button onclick="upload.click()" class="mt-2 rounded-sm px-3 py-1 bg-gray-200 hover:bg-gray-300 focus:shadow-outline focus:outline-none">Upload a file</button>
    </div>
    <div class="h-5 pb-10 pt-1 lg:w-2/3 md:w-4/5 sm:w-4/5 min:w-5/6">
        <div class="float-left text-xl pt-3 pb-2" style="width:100%;text-align: left">Upload Files</div>
        <div class="flex flex-wrap" style="width:100%;text-align: left">
            <FileChip v-for="file of files.entries()" :content="file[0]" @deleteItem="deleteChip" class="mx-2 my-1"></FileChip>
        </div>
    </div>
</div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import FileChip  from './FileChip.vue';
import axios from 'axios';
import {showTips} from "../utils/tips";
export default defineComponent({
    name: "FileArea",
    components: { FileChip },
    data() {
        return {
            files: new Map(),
        }
    },
    methods: {
        addFile: function(e : Event) {
            let files = (<HTMLInputElement>e.target).files;
            if(files == null) return;
            let fileList : FileList = files;
            console.log(fileList);
            Array.from(fileList).forEach(file => {
                let randomStr =  Math.random().toString(16).slice(-8);
                this.files.set(file.name + "-" + randomStr, file);
            });
        },
        uploadFile: function() {
            if(this.files.size == 0){
                showTips({
                    content: "Files list is empty!",
                    title: "Warning",
                    type: "warning",
                });
                return;
            }
            let all = this.files.size;
            let promises = [];
            for(let [key, value] of this.files) {
                promises.push(new Promise((resolve, reject) => {
                    let formData = new FormData();
                    formData.append("file", value);
                    axios.post("/file/upload", formData).then(res => {
                        if(res.status == 200 && res.data.retcode == 0){
                            resolve(key);
                        }else{
                            resolve("");
                        }
                    }).catch(err=>{
                        resolve(err)
                    });
                }))
            }
            Promise.all(promises).then(values => {
                let count = 0;
                let err = "";
                values.forEach(name => {
                    if(name instanceof Error){
                        err = name.message;
                    }else if(name !== "" || name.length != 0) {
                        count++;
                        this.files.delete(name);
                    }
                })
                if(count !== all || err !== ""){
                    showTips({
                        content: "Failed: " + (all-count) + " files & Error: " + err,
                        title: "Error",
                        type: "error"
                    });
                    return;
                }
                showTips({
                    content: "Send file successfully: " + count + "/" + all,
                    title: "Success",
                    type: "success"
                });
            });
        },
        deleteChip: function(e : String) {
            this.files.delete(e);
        }
    },
});
</script>

<style scoped>
</style>