import {server_response, user} from "../interfaces";
const POST = "POST"
const PUT = "PUT"
const auth_path = "http://localhost:8080/api/users/auth";
const create_path = "http://localhost:8080/api/users";

const login_form_element = document.getElementById("login-form") as HTMLFormElement
const logout_button_element = document.getElementById("logout-button") as HTMLButtonElement
const register_form_element = document.getElementById("register-form") as HTMLFormElement

const element_exists = (n:string):boolean => {
    const ele = document.getElementById(n)
    return ele == null ? false:true
}
import {notify_modal} from "./notifications";

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
        const opts = {
            method: "POST",
            body: JSON.stringify(data),
            headers: {
                "Content-Type": "application/json"
            }
        }
        const result = await fetch(auth_path, opts)
        const response: server_response = await result.json()
        notify_modal(response.message_type, response.message, response.data , response.code.toString())
            .then((res)=>{
                if(response.message_type == "success"){
                    window.location.href = "/profile"
                }
            })

        clear_inputs()
    }
    login_form_element.addEventListener("submit", onSubmit)
}

const logout_button_handler = async () => {
    const opts = {
        method:"PUT",
        headers:{
            "Content-Type":"application/json",
        }
    }
    const result = await fetch(auth_path, opts)
    const response: server_response = await result.json()

    notify_modal(response.message_type, response.message, response.data, response.code.toString())
        .then(res=>{
            window.location.replace("/")
            console.log(res)
        })
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
        const opts = {
            method:POST,
            body:JSON.stringify(data),
            headers:{
                "Content-Type":"application/json",
            }
        }
        const result = await fetch(create_path, opts)
        const response = await result.json()
        notify_modal(response.message_type,response.message, response.data, response.code)
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
    if(logout_button_element !== null){
        logout_button_element.addEventListener("click", logout_button_handler)
    }
}

export default run