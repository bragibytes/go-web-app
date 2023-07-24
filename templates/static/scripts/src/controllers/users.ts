const update_button_element = document.getElementById("update-button") as HTMLButtonElement
const delete_button_element = document.getElementById("delete-button") as HTMLButtonElement

import {element_exists, store} from "./config";
import Swal from "sweetalert2"
import {update_user, delete_user, logout_user} from "../api";
import {user} from "../interfaces"
import { notify } from "./notifications";

const update_button_handler = () => {
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
                const new_username = new_username_input ? new_username_input.value:store.get_user()?.name
                const new_email = new_email_input ? new_email_input.value:store.get_user()?.email
                // Do something with the newUsername and newEmail, e.g., send it to the server
                const user_to_update:user = {
                    name:new_username,
                    email:new_email,
                }
                const res = await update_user(user_to_update)
                if(res.message_type === "success"){
                    logout_user()
                }
            },
        });
    }

    update_button_element.addEventListener("click", on_click)
}
const delete_button_handler = () => {
    const on_click = (e:Event) => {
        delete_user()
        .then(res=>{
            if(res.message_type === "success"){
                window.location.replace("/")

            }
        })
    }

    delete_button_element.addEventListener("click", on_click)
}
const run = () => {
    console.log("looking for update and delete buttons")
    if(element_exists("update-button")){
        console.log("found update button")
        update_button_handler()
    }
    if(element_exists("delete-button")){
        console.log("found the delete button")
        delete_button_handler()
    }
}

export default run