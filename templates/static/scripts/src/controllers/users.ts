const delete_button = "delete-button"
const update_button = "update-button"

const update_button_element = document.getElementById(update_button) as HTMLButtonElement
const delete_button_element = document.getElementById(delete_button) as HTMLButtonElement

import Swal from "sweetalert2"
import {update_user, delete_user} from "../api";
import {user} from "../interfaces"
import { element_exists, chemical_x } from "./config";

const update_button_handler = () => {
    const data = chemical_x() as user
    const on_click = () => {
        // Show SweetAlert2 modal with input fields
        Swal.fire({
            title: "Update User",
            html: `
                <div class="container">
                    <div class="row">
                        <input id="username" class="swal2-input" placeholder="New Username">
                        <input id="email" class="swal2-input" placeholder="New Email">
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
        delete_user()
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
}

export default run