import {user, server_response} from "./interfaces";
import notie, {AlertType} from 'notie'
import Swal, {SweetAlertIcon} from "sweetalert2"

const path = "http://localhost:8080/api/users/auth";
const login_form_element = document.getElementById("login-form") as HTMLFormElement
const logout_button_element = document.getElementById("logout-button") as HTMLButtonElement

const notify = (msg:string, msgType:string) => {
    notie.alert({
        type: msgType as AlertType|undefined, // optional, default = 4, enum: [1, 2, 3, 4, 5, 'success', 'warning', 'error', 'info', 'neutral']
        text: msg,
    })
}
const notify_modal = (icon:string, title:string, text:string, footer:string) => {
    Swal.fire({
        icon: icon as SweetAlertIcon,
        title: title,
        text: text,
        footer: footer
    })
}
const element_exists = (n:string):boolean => {
    const ele = document.getElementById(n)
    return ele == null ? false:true
}

const init_login_form = () => {

    const username = ():HTMLInputElement => {
       return login_form_element.querySelector(`[name="username"]`) as HTMLInputElement
    }
    const password = ():HTMLInputElement => {
        return login_form_element.querySelector(`[name="password"]`) as HTMLInputElement
    }
    const onSubmit = async (e:Event)=> {
        e.preventDefault()
        let data: user = {
            name:username().value,
            password:password().value
        }
        const opts = {
            method:"POST",
            body:JSON.stringify(data),
            headers:{
                "Content-Type":"application/json"
            }
        }
        const result = await fetch(path, opts)
        const response:server_response = await result.json()
        notify(response.message, response.message_type)
        console.log(response)

    }
    login_form_element.addEventListener("submit", onSubmit)
    logout_button_element.addEventListener("click", async (e)=>{
        const opts = {
            method:"PUT",
            headers:{
                "Content-Type":"application/json",
            }
        }
        const result = await fetch(path, opts)
        const res = await result.json()
        notify_modal(res.message_type, "", res.message, "")
    })
}

if(element_exists("login-form")){
    init_login_form()
}



