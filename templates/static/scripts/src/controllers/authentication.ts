const login_button = "login-button"
const register_button = "register-button"
const logout_button = "logout-button"

const login_button_element = document.getElementById(login_button) as HTMLAnchorElement
const logout_button_element = document.getElementById(logout_button) as HTMLAnchorElement
const register_button_element = document.getElementById(register_button) as HTMLAnchorElement

import Swal from "sweetalert2"
import {user} from "../interfaces";
import { login_user, logout_user, register_user } from "../api";
import { element_exists } from "./config";

const login_button_handler = () => {

    
    const on_click = async (e: Event) => {
        e.preventDefault()
        Swal.fire({
            title: "Login",
            html: `
                <form class="container">
                    <div class="form-group row">
                        <input id="username" class="swal2-input" placeholder="Username...">
                    </div>
                    <div class="form-group row">
                        <input type="password" id="password" class="swal2-input" placeholder="Password...">
                    </div>
                </form>
            `,
            confirmButtonText: "Log In!",
            showCancelButton: true,
            allowEnterKey:true,
            preConfirm: async () => {
                const username_input = Swal.getPopup()!.querySelector("#username") as HTMLInputElement;
                const password_input = Swal.getPopup()!.querySelector("#password") as HTMLInputElement;

                // Retrieve user input and handle data
                const username: string = username_input ? username_input.value:""
                const password: string = password_input ? password_input.value:""
                // Do something with the newUsername and newEmail, e.g., send it to the server
                const user_to_login:user = {
                    name:username,
                    password:password,
                }
                login_user(user_to_login)
            },
        });
        
    }
    login_button_element.addEventListener("click", on_click)
}

const logout_button_handler = () => {

    const on_click = async (e:Event) => {
        
        logout_user()
    }
    logout_button_element.addEventListener("click", on_click)
}

const register_button_handler = () => {
  
    const on_click = async (e:Event) => {
        Swal.fire({
            title: "Create User",
            html: `
                <div class="container">
                    <div class="row">
                        <input id="username" class="swal2-input" placeholder="Username...">
                        <input id="email" class="swal2-input" placeholder="Email">
                        <input type="password" id="password" class="swal2-input" placeholder="Password...">
                        <input type="password" id="confirm-password" class="swal2-input" placeholder="Password Again...">
                    </div>
                </div>
            `,
            confirmButtonText: "Sign Up!",
            showCancelButton: true,
            preConfirm: async () => {
                const username_input = Swal.getPopup()!.querySelector("#username") as HTMLInputElement;
                const email_input = Swal.getPopup()!.querySelector("#email") as HTMLInputElement;
                const password_input = Swal.getPopup()!.querySelector("#password") as HTMLInputElement;
                const confirm_password_input = Swal.getPopup()!.querySelector("#confirm-password") as HTMLInputElement;
                
                // Do something with the newUsername and newEmail, e.g., send it to the server
                const new_user:user = {
                    name:username_input.value,
                    email:email_input.value,
                    password:password_input.value,
                    confirm_password:confirm_password_input.value
                }
                register_user(new_user)
            },
        });
    }
    register_button_element.addEventListener("click", on_click);
}

const run = () => {
    console.log("running authentication")
   
    if(element_exists(login_button)) {
        console.log("login form exists")
        login_button_handler()
    }
    if(element_exists(register_button)){
        console.log("register button exists")
        register_button_handler()
    }
    if(element_exists(logout_button)){
        console.log("logout button exists")
        logout_button_handler()
    }
}

export default run