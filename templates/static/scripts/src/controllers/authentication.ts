import {user} from "../interfaces";
import {element_exists} from "./config";
import { login_user, logout_user, register_user } from "../api";

const login_form_element = document.getElementById("login-form") as HTMLFormElement
const logout_button_element = document.getElementById("logout-button") as HTMLButtonElement
const register_form_element = document.getElementById("register-form") as HTMLFormElement

const login_form_handler = () => {

    const username = (): HTMLInputElement => {
        return login_form_element.querySelector(`[name="username"]`) as HTMLInputElement
    }
    const password = (): HTMLInputElement => {
        return login_form_element.querySelector(`[name="password"]`) as HTMLInputElement
    }
    const clear_inputs = () => {
        username().value = ""
        password().value = ""
    }
    const onSubmit = async (e: Event) => {
        e.preventDefault()
        let data: user = {
            name: username().value,
            password: password().value
        }
        login_user(data)
        .then(res=>{
            if(res.message_type === "success"){
                window.location.href = "/profile"
                clear_inputs()
            }
        })
        
            
    }
    login_form_element.addEventListener("submit", onSubmit)
}

const logout_button_handler = () => {

    const on_click = async (e:Event) => {
        
        logout_user()
        .then(res=>{
            if(res.message_type === "success"){
                window.location.replace("/")
            }
        })
    }
    logout_button_element.addEventListener("click", on_click)
}

const register_form_handler = () => {
    const username = (): HTMLInputElement => {
        return login_form_element.querySelector(`[name="username"]`) as HTMLInputElement
    }
    const password = (): HTMLInputElement => {
        return login_form_element.querySelector(`[name="password"]`) as HTMLInputElement
    }
    const email = (): HTMLInputElement => {
        return login_form_element.querySelector(`[name="email"]`) as HTMLInputElement
    }
    const confirm_password = (): HTMLInputElement => {
        return login_form_element.querySelector(`[name="confirm_password"]`) as HTMLInputElement
    }
    const on_submit = async (e:Event) => {
        e.preventDefault()
        const data: user = {
            name:username().value,
            email:email().value,
            password:password().value,
            confirm_password:confirm_password().value
        }
        register_user(data)
    }
    register_form_element.addEventListener("submit", on_submit)
}

const run = () => {
    if(element_exists("login-form")){
        login_form_handler()
    }
    if(element_exists("register-form")){
        register_form_handler()
    }
    if(element_exists("logout-button")){
        logout_button_handler()
    }
}

export default run