import {createApp} from "vue";
import Tips from "../components/Tips.vue";
export const showTips =  function(props:any){
    let duration = props.duration;
    if(duration == undefined || duration <= 0){
        duration = 2000;
        props.duration = duration;
    }
    let app = createApp(Tips, props);
    let divElement = document.createElement("div");
    divElement.setAttribute("class", "absolute overflow-x-hidden h-40 w-screen")
    document.body.appendChild(divElement);
    app.mount(divElement);
    setTimeout(()=> {
        app.unmount();
          document.body.removeChild(divElement);
    },800 + duration);
}