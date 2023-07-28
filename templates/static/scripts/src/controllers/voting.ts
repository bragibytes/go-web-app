import { vote } from "../interfaces";
import { send_vote } from "../api";
import { element_exists } from "./config";

const vote_box = "vote-box"
const upvote_color = "orange"
const downvote_color = "purple"


const vote_box_elements = document.getElementsByClassName('vote-box')

const run = () => {
   [...vote_box_elements].forEach(box => {


        const upvote_button = box.querySelector(':scope > a[name="up"]') as HTMLAnchorElement
        const downvote_button = box.querySelector(':scope > a[name="down"]') as HTMLAnchorElement

        const on_click = (e:Event) => {
         
                const btn = e.currentTarget as HTMLAnchorElement
                console.log(btn)
                const valStr = btn.getAttribute("value") as string
                const val = JSON.parse(valStr) as number
                const parentID = btn.getAttribute("parent") as string
                const model_type = btn.getAttribute("model_type") as string
                
                const data:vote = {
                    value:val,
                    _parent:parentID
                }
                
                console.log('sending data...')
                console.log(model_type, val, parentID)
                send_vote(data, model_type)
                .then(res=>{
                    console.log("vote response", res)
                    const vote_type = btn.getAttribute('name') as string
                    switch(vote_type){
                        case 'up':
                            btn.classList.toggle(upvote_color)
                            if(res.message == "Update"){
                                console.log("vote was updated")
                                downvote_button.classList.remove(downvote_color)
                            }
                            break;
                        case 'down':
                            btn.classList.toggle(downvote_color)
                            if(res.message == "Update"){
                                console.log("vote was updated")
                                upvote_button.classList.remove(upvote_color)
                            }
                            break;
                    }
                })
    
            
        }

        upvote_button.addEventListener("click", on_click)
        downvote_button.addEventListener("click", on_click)
    })
   
}

export default run;