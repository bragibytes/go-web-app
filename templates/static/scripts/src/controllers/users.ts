import Swal from "sweetalert2"
const delete_button = "delete-button"
const update_button = "update-button"
const bio_creator = "bio-creator"

const update_button_element = document.getElementById(update_button) as HTMLButtonElement
const delete_button_element = document.getElementById(delete_button) as HTMLButtonElement
const bio_creator_element = document.getElementById(bio_creator) as HTMLFormElement

import {update_user, delete_user, update_user_bio} from "../api";
import {user} from "../interfaces"
import { element_exists, chemical_x } from "./config";



const bio_creator_handler = () => {

    const bio = ():string => {
        const x = bio_creator_element.querySelector("[name='bio']") as HTMLTextAreaElement
        const n = x.value as string
        return n
    }

    const on_submit = async (e:Event) => {
        e.preventDefault()
        const data:user = {
            bio: bio()
        }
        update_user_bio(data)
    }

    bio_creator_element.addEventListener("submit", on_submit)
}
const update_button_handler = () => {
    const data = chemical_x() as user
    const on_click = () => {
        // Show SweetAlert2 modal with input fields
        Swal.fire({
            title: "Update User",
            html: `
                <div class="container">
                    <div class="row">
                        <input id="username" class="swal2-input" value='${data.name}'>
                        <input id="email" class="swal2-input" value='${data.email}'>
                    </div>
                </div>
            `,
            confirmButtonText: "Update",
            showCancelButton: true,
            preConfirm: async () => {
                const new_username_input = Swal.getPopup()!.querySelector("#username") as HTMLInputElement;
                const new_email_input = Swal.getPopup()!.querySelector("#email") as HTMLInputElement;
                // Retrieve user input and handle data
                const new_username: string = new_username_input ? new_username_input.value:""
                const new_email: string = new_email_input ? new_email_input.value:""
                // Do something with the newUsername and newEmail, e.g., send it to the server
                const user_to_update:user = {
                    name:new_username,
                    email:new_email,
                }
                update_user(user_to_update)
            },
        });
    }

    update_button_element.addEventListener("click", on_click)
}
const delete_button_handler = () => {
    const on_click = (e:Event) => {
        Swal.fire({
            title: "Are you sure?",
            text: "You won't be able to revert this!",
            icon: "warning",
            showCancelButton: true,
            confirmButtonColor: "#3085d6",
            preConfirm: () => {
                delete_user()
            }
        })
        
    }

    delete_button_element.addEventListener("click", on_click)
}
const run = () => {
    if(element_exists(update_button)){
        update_button_handler()
    }
    if(element_exists(delete_button)){
        delete_button_handler()
    } 
    if(element_exists(bio_creator)){
        bio_creator_handler()
    }
}

export default run