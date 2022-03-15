<template>
<div class="h-full w-full flex flex-col items-center">
    <div class="h-5 pb-10 pt-7 w-1/2 md:w-3/4 sm:w-4/5 min:w-5/6">
         <div class="float-left text-xl pt-3">Text ClipBoard</div>
        <div class="float-right">
            <button
            type="button"
            class="inline-block align-top px-6 py-1.5 bg-blue-600 text-white font-medium text-sm leading-6 uppercase rounded shadow-md hover:bg-blue-700 hover:shadow-lg focus:bg-blue-700 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-blue-800 active:shadow-lg transition duration-150 ease-in-out"
            @click="createText">Upload</button>
        </div>
    </div>
    <div class="text-gray-70 h-3/4 w-1/2 md:w-3/4 sm:w-4/5 min:w-5/6 m-1">
        <textarea v-model="textContent" style="height: 100%; width: 100%;" class="appearance-none border border-gray-300 py-4 px-4 bg-white text-gray-700 placeholder-gray-400 rounded-lg text-base focus:outline-none focus:ring-2 focus:ring-purple-600 focus:border-transparent" id="content" placeholder="Enter your content" name="content" rows="5" cols="40"></textarea>
    </div>
</div>
</template>
<script lang="ts">
import {defineComponent} from "vue";
import axios from "axios";
import {showTips} from "../utils/tips"
export default defineComponent({
  name: "TextArea",
  data(){
      return {
          textContent: "",
      }
  },
  methods: {
      createText: function() {
        let txt = this.textContent;
        if(txt == null || txt == undefined || txt.length == 0) {
            showTips({
                content: "Content is empty!",
                title: "Warning",
                type: "warning"
            })
            return;
        }
        let data = new FormData();
        data.append("content", txt);
        axios.post("text/add", data).then(res => {
            if(res.status ==200 &&  res.data.retcode == 0) {
                console.log("上传成功");
                showTips({
                    content: "Send text to clipboard successfully!",
                    title: "Success",
                    type: "success"
                })
            }else{
                console.log("上传失败");
                console.log(res);
                showTips({
                    content: "Failed: " + res.data.message,
                    title: "Error",
                    type: "error"
                })
            }
        }).catch(err=>{
            console.log(err);
            showTips({
                content: "Failed: " + err,
                title: "Error",
                type: "error"
            })
        });
    }
  },
});
</script>

<style scoped>
</style>