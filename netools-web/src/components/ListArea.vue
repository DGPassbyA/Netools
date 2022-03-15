<template>
    <div class="container mx-auto px-4 sm:px-8 max-w-3xl">
        <div class="py-8">
            <div class="-mx-4 sm:-mx-8 px-4 sm:px-8 py-4 overflow-x-auto">
                <div style="text-align: left;" class="text-lg font-semibold">
                    Text List
                </div>
                <div class="inline-block min-w-full shadow rounded-lg overflow-hidden">
                    <table class="min-w-full leading-normal">
                        <div v-if="textList.length == 0">No Text</div>
                        <ListItem v-for="t in textList" :timestamp="t.timestamp" :content="t.content" :size="t.content.length" type="text" :uuid="t.uuid"></ListItem>
                    </table>
                </div>
                <div style="text-align: left;" class="mt-8 text-lg font-semibold">
                    File List
                </div>
                <div class="inline-block min-w-full shadow rounded-lg overflow-hidden">
                    <table class="min-w-full leading-normal">
                        <div v-if="fileList.length == 0">No File</div>
                        <ListItem v-for="f in fileList" :timestamp="f.timestamp" :content="f.filename" :size="f.size" type="file" :uuid="f.uuid"></ListItem>
                    </table>
                </div>
            </div>
        </div>
    </div>
</template>


<script lang="ts">
import { defineComponent } from "vue";
import ListItem from "./ListItem.vue";
import axios from "axios";
import {showTips} from "../utils/tips"
export default defineComponent({
    name: "ListArea",
    components: { ListItem },
    data() {
        return {
            textList: new Array(),
            fileList: new Array(),
        }
    },
    mounted() {
        axios.get("/text/all").then(res => {
            if(res.status == 200 && res.data.retcode == 0){
                if(res.data.data == null){
                    return
                }
                this.textList = res.data.data;
            }else{
                showTips({
                    title: "Error",
                    content: "Get text list failed: " + res.data.message,
                    type: "error",
                })
            }
        }).catch(err => {
            showTips({
                title: "Error",
                content: "Get text list failed: " + err,
                type: "error",
            })
        })
        axios.get("/file/all").then(res => {
            if(res.status == 200 && res.data.retcode == 0){
                if(res.data.data == null){
                    return
                }
                this.fileList = res.data.data;
            }else{
                showTips({
                    title: "Error",
                    content: "Get file list failed: " + res.data.message,
                    type: "error",
                })
            }
        }).catch(err => {
            showTips({
                title: "Error",
                content: "Get text list failed: " + err,
                type: "error",
            })
        })
    },
    methods: {

    },
})
</script>

<style scoped>
</style>