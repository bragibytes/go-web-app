import Swal from "sweetalert2"
import { 
    create_post,
    update_post,
    delete_post
 } from "../api";
import { post, server_response } from "../interfaces";
import { element_exists, json_data } from "./config";


const post_creator = "post-creator"
const post_updater = "post-updater"
const post_deleter = "post-deleter"



const post_creator_element = document.getElementById(post_creator) as HTMLFormElement;
const post_updater_element = document.getElementById(post_updater) as HTMLButtonElement;
const post_deleter_element = document.getElementById(post_deleter) as HTMLButtonElement;


const post_creator_handler = () => {

    const title = () => {
        return post_creator_element.querySelector('[name="title"]') as HTMLInputElement;
    }
    const content = () => {
        return post_creator_element.querySelector('[name="content"]') as HTMLInputElement;
    }
    const clear_inputs = () => {
        title().value = ""
        content().value = ""
    }
    const on_submit = async (e:Event) => {
        e.preventDefault();
        const data:post = {
            title: title().value,
            content: content().value,
        }
        create_post(data)
        .then((res:server_response)=>{
            if(res.message_type == "success"){
                clear_inputs()
            }
        })
    }
    post_creator_element.addEventListener('submit', on_submit)
}
const post_updater_handler = () => {
   
    const on_click = async (e:Event) => {
        e.preventDefault();
        const data = json_data() as post;
        Swal.fire({
            title: "Update Post",
            html: `
                <div class="container">
                    <div class="row">
                        <input id="title" class="swal2-input" placeholder="New Title">
                        <textarea id="content" class="swal2-input" placeholder="New Content"></textarea>
                    </div>
                </div>
            `,
            confirmButtonText: "Update",
            showCancelButton: true,
            preConfirm: async () => {
                const new_title_input = Swal.getPopup()!.querySelector("#title") as HTMLInputElement;
                const new_content_input = Swal.getPopup()!.querySelector("#content") as HTMLInputElement;
                // Retrieve user input and handle data
                const new_title: string = new_title_input ? new_title_input.value:""
                const new_content: string = new_content_input ? new_content_input.value:""
                // Do something with the newUsername and newEmail, e.g., send it to the server
                const post_update:post = {
                    title:new_title,
                    content:new_content,
                }
                update_post(data._id! ,post_update)
            },
        });
    }
    post_updater_element.addEventListener('click', on_click)
}
const post_deleter_handler = async () => {
    const data = json_data() as post;
    const on_click = (e:Event) => {
        delete_post(data._id!)
    }
    post_deleter_element.addEventListener('click', on_click)
}

const run = () => {
    if(element_exists(post_creator)){
        post_creator_handler()
    }
    if(element_exists(post_updater)){
        post_updater_handler()
    }
}

export default run;