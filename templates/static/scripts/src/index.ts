import {user, server_response} from "./interfaces";
import notie, {AlertType} from 'notie'

const path = "http://localhost:8080/api/users/auth";
const login_form_element = document.getElementById("login-form") as HTMLFormElement

const element_exists = (n:string):boolean => {
    const ele = document.getElementById(n)
    return ele == null ? false:true
}



const init_login_form = () => {

    const username = ():HTMLInputElement => {
       return login_form_element.namedItem("name") as HTMLInputElement
    }
    const password = ():HTMLInputElement => {
        return login_form_element.elements.namedItem("password") as HTMLInputElement
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
        notie.alert({
            type: response.message_type as AlertType, // optional, default = 4, enum: [1, 2, 3, 4, 5, 'success', 'warning', 'error', 'info', 'neutral']
            text: response.message,
        })
    }
    login_form_element.addEventListener("submit", onSubmit)
}

if(element_exists("login-form")){
    init_login_form()
}



