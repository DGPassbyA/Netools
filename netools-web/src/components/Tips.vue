<template>
    <div
      :class="cssShow"
      class="transition-all z-40 overflow-x-hidden ease-in-out duration-1000 border-l-4 p-4 absolute top-12"
      role="alert"
    >
      <p class="font-bold">{{title}}</p>
      <p>{{(content as String).length > 45 ? content : (content + "&nbsp".repeat(45-(content as String).length))}}</p>
    </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
export default defineComponent({
    name: "Tips",
    props: {
        title: String,
        content: String,
        type: String,
        duration:Number,
    },
    data() {
        return {
            cssShow: {
                "opacity-0 -right-20": true,
                "opacity-1 right-0": false,
                "bg-green-200 border-green-600 text-green-600": false,
                "bg-yellow-200 border-yellow-600 text-yellow-600": false,
                "bg-red-200 border-red-600 text-red-600": false
            },
        }
    },
    created() {
        switch(this.type){
            case undefined:
            case "success":
                this.cssShow["bg-green-200 border-green-600 text-green-600"] =  true;
                break;
            case "warning":
                this.cssShow["bg-yellow-200 border-yellow-600 text-yellow-600"] = true;
                break;
            case "error":
                this.cssShow["bg-red-200 border-red-600 text-red-600"] = true;
                break;
        }
        let d = 0;
        if(this.duration == undefined || this.duration <= 0) {
            d = 1700;
        }else{
            d = this.duration;
        }
        this.cssShow["opacity-0 -right-20"] = true;
        this.cssShow["opacity-1 right-0"] = false;
        setTimeout(()=>{
            this.cssShow["opacity-0 -right-20"] = false;
            this.cssShow["opacity-1 right-0"] = true;
        },100)
        setTimeout(()=>{
            this.cssShow["opacity-0 -right-20"] = true;
            this.cssShow["opacity-1 right-0"] = false;
        }, 100+d)
    },
    methods: {
    }
})
</script>

<style scoped>
</style>