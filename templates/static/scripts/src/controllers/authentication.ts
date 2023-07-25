const login_form = "login-form"
const register_form = "register-form"
const logout_button = "logout-button"

const login_form_element = document.getElementById(login_form) as HTMLFormElement
const logout_button_element = document.getElementById(logout_button) as HTMLButtonElement
const register_form_element = document.getElementById(register_form) as HTMLFormElement

import {user} from "../interfaces";
import { login_user, logout_user, register_user } from "../api";
import { element_exists } from "./config";

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
    const on_submit = async (e: Event) => {
        e.preventDefault()
        let data: user = {
            name: username().value,
            password: password().value
        }
        login_user(data) 
    }
    login_form_element.addEventListener("submit", on_submit)
}

const logout_button_handler = () => {

    const on_click = async (e:Event) => {
        
        logout_user()
    }
    logout_button_element.addEventListener("click", on_click)
}

const register_form_handler = () => {
    const username = (): HTMLInputElement => {
        return register_form_element.querySelector(`[name="username"]`) as HTMLInputElement
    }
    const password = (): HTMLInputElement => {
        return register_form_element.querySelector(`[name="password"]`) as HTMLInputElement
    }
    const email = (): HTMLInputElement => {
        return register_form_element.querySelector(`[name="email"]`) as HTMLInputElement
    }
    const confirm_password = (): HTMLInputElement => {
        return register_form_element.querySelector(`[name="confirm-password"]`) as HTMLInputElement
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
    if(element_exists(login_form)) {
        login_form_handler()
    }
    if(element_exists(register_form)){
        register_form_handler()
    }
    if(element_exists(logout_button)){
        logout_button_handler()
    }
}

export default run