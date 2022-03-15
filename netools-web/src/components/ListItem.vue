<template>
    <tr>
        <td class="px-5 py-5 border-b border-gray-200 bg-white text-sm">
            <div class="flex items-center">
                <div class="flex-shrink-0">
                    <a href="#" class="block relative">
                        <svg
                            xmlns="http://www.w3.org/2000/svg"
                            v-if="type == 'text'"
                            class="h-6 w-6"
                            fill="none"
                            viewBox="0 0 24 24"
                            stroke="currentColor"
                            stroke-width="2"
                        >
                            <path
                                stroke-linecap="round"
                                stroke-linejoin="round"
                                d="M7 8h10M7 12h4m1 8l-4-4H5a2 2 0 01-2-2V6a2 2 0 012-2h14a2 2 0 012 2v8a2 2 0 01-2 2h-3l-4 4z"
                            />
                        </svg>
                        <svg
                            xmlns="http://www.w3.org/2000/svg"
                            v-else-if="type == 'file'"
                            class="h-6 w-6"
                            fill="none"
                            viewBox="0 0 24 24"
                            stroke="currentColor"
                            stroke-width="2"
                        >
                            <path
                                stroke-linecap="round"
                                stroke-linejoin="round"
                                d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"
                            />
                        </svg>
                    </a>
                </div>
                <div class="ml-3">
                    <p class="text-gray-900 whitespace-no-wrap">{{type=="text"?(content?.substring(0,15)):content}}</p>
                </div>
            </div>
        </td>
        <td class="px-5 py-5 border-b border-gray-200 bg-white text-sm">
            <p class="text-gray-900 whitespace-no-wrap">{{type=="text"?(size+"字符"):Math.floor((size as number)/1024)+"KB"}}</p>
        </td>
        <td class="px-5 py-5 border-b border-gray-200 bg-white text-sm">
            <p class="text-gray-900 whitespace-no-wrap">{{getTime((timestamp as number))}}</p>
        </td>
        <td class="px-5 py-5 border-b border-gray-200 bg-white text-sm cursor-pointer" @click="actionClick"
        >
            <svg
                v-if="type == 'text'"
                xmlns="http://www.w3.org/2000/svg"
                class="h-6 w-6"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
                stroke-width="2"
            >
                <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    d="M8 5H6a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2v-1M8 5a2 2 0 002 2h2a2 2 0 002-2M8 5a2 2 0 012-2h2a2 2 0 012 2m0 0h2a2 2 0 012 2v3m2 4H10m0 0l3-3m-3 3l3 3"
                />
            </svg>
            <svg
                v-else-if="type == 'file'"
                xmlns="http://www.w3.org/2000/svg"
                class="h-6 w-6"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
                stroke-width="2"
            >
                <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    d="M8 7H5a2 2 0 00-2 2v9a2 2 0 002 2h14a2 2 0 002-2V9a2 2 0 00-2-2h-3m-1 4l-3 3m0 0l-3-3m3 3V4"
                />
            </svg>
        </td>
    </tr>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import axios from 'axios';
import {showTips} from "../utils/tips"
export default defineComponent({
    name: "ListItem",
    props: {
        type: String,
        content: String,
        uuid: String,
        size: Number,
        timestamp: Number,
    },
    data() {
        return {

        }
    },
    methods: {
    textCopy: function (){
        this.$copyText(this.content).then( function(e:any){
            // alert('复制成功')
            console.log(e)
            showTips({
                title: "Success",
                content: "Copy to clipboard successfully!",
                type: "success"
            })
        }, function (e: any) {
            // alert('复制失败')
            console.log(e)
            showTips({
                title: "Error",
                content: e.message,
                type: "error"
            })
        })
    },
    fileDownload: function () {
        window.open(axios.defaults.baseURL + "file/download?uuid="+this.uuid);
    },
    actionClick: function() {
        if(this.type == "text"){
            return
        }else if(this.type == "file"){
            this.fileDownload();
        }
    },
    getTime: function(t:number) {
        let date = new Date(t*1000)
        let M = date.getMonth();
        let D = date.getDate();
        let Hour = date.getHours();
        let Min = date.getMinutes();
        let Sec = date.getSeconds();
        return M + "-" + D + " " + Hour + ":" + Min + ":" + Sec;
    }
    },
})
</script>

<style scoped>
</style>