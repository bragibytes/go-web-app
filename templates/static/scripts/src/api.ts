import {
    user, 
    server_response,
    post,
    vote,
    comment
} from "./interfaces"
import Swal, { SweetAlertIcon, SweetAlertResult } from "sweetalert2"


const root = "http://localhost:10000/api/"
const POST = "POST"
const GET = "GET"
const PUT = "PUT"
const DELETE = "DELETE"
const success_timeout = 2000


const successful = (r:server_response) => {
    return r.message_type == "success"
}
const teapot = (r:Response) => {
    return r.status == 418
}
const swal_type = (r:server_response):SweetAlertIcon => {
    return r.message_type as SweetAlertIcon
}
const swal_success = (r:server_response):Promise<SweetAlertResult> => {
    return Swal.fire({
        icon:swal_type(r),
        title:r.message,
        showConfirmButton:false,
        timer:success_timeout
    })
}
const swal_error = (r:server_response):Promise<SweetAlertResult> => {
    return Swal.fire({
        icon:swal_type(r),
        title:r.message,
    })
}
 // users
export const update_user = async (update:user):Promise<server_response> => {
    const opts = {
        method:PUT,
        body:JSON.stringify(update),
        headers:{
            "Content-Type":"application/json"
        }
    }
    const result = await fetch(root+"users", opts)
    const response:server_response = await result.json()
    
    if(successful(response)) {
        swal_success(response)
        .then(()=>{window.location.reload()})
    }else{
        swal_error(response)
    }

    return response
}
export const delete_user = async ():Promise<server_response> => {
    const opts = {
        method:DELETE,
        headers:{
            "Content-Type":"application/json"
        }
    }
    const result = await fetch(root+"users", opts)
    const response:server_response = await result.json()
    
    if(successful(response)){
        logout_user()
    }else{
        swal_error(response)
    }

    return response
}
export const login_user = async (user:user):Promise<server_response> => {
    const opts = {
        method: "POST",
        body: JSON.stringify(user),
        headers: {
            "Content-Type": "application/json"
        }
    }
    const result = await fetch(root+"users/auth", opts)
    const response: server_response = await result.json()
    
    if(successful(response)){
        swal_success(response)
        .then(() => {window.location.href = "/profile"})   
    }else{
        swal_error(response)
    }

    return response
}
export const logout_user = async (): Promise<server_response> => {
    const opts = {
        method:"PUT",
        headers:{
            "Content-Type":"application/json",
        }
    }
    const result = await fetch(root+"users/auth", opts)
    const response: server_response = await result.json()

    
    if(successful(response)){
        swal_success(response)
        .then(() => {window.location.replace("/")})
    }else{
        swal_error(response)
    }

    return response
}
export const register_user = async (user:user):Promise<server_response> => {
    const opts = {
        method:POST,
        body:JSON.stringify(user),
        headers:{
            "Content-Type":"application/json",
        }
    }
    const result = await fetch(root+"users", opts)
    const response = await result.json()
    
    if(successful(response)){
        swal_success(response)
        .then(() => {window.location.href = "/profile" })
    }else if(teapot(result)){
        Swal.fire({
            icon:response.message_type,
            title:"Validation Errors",
            html:response.data.join("\n")
        })
    }else {
        swal_error(response)
    }

    return response
}
// posts
export const create_post = async (post:post):Promise<server_response> => {
    const opts = {
        method:POST,
            body:JSON.stringify(post),
            headers:{
                "Content-Type":"application/json"
            }
    }
    const result = await fetch(root+"posts", opts)
    const response:server_response = await result.json()

    if(successful(response)){
        swal_success(response)
        .then(() => {
            window.location.reload()
        })
    }else{
        swal_error(response)
    }

    return response
}
export const delete_post = async (id:string):Promise<server_response> => {
    const opts = {
        method:DELETE,
    }
    const result = await fetch(root+"posts/"+id, opts)
    const response:server_response = await result.json()

    if(successful(response)){
        swal_success(response)
        .then(() => {window.location.replace("/board")})   
        
    }else{
        swal_error(response)
    }

    return response
}
export const update_post = async (id:string, data:post):Promise<server_response> => {
    const opts = {
        method:PUT,
        body:JSON.stringify(data),
        headers:{
            "Content-Type":"application/json"
        }
    }
    const result = await fetch(root+"posts/"+id, opts)
    const response:server_response = await result.json()

    if(successful(response)){
        swal_success(response)
        .then(() => {window.location.reload()})
    }else{
        swal_error(response)
    }

    return response
}
// comments
export const create_comment = async (data:comment):Promise<server_response> => {
    const opts = {
        method:POST,
        body:JSON.stringify(data),
        headers:{
            "Content-Type":"application/json"
        }
    }
    const result = await fetch(root+"comments", opts)
    const response:server_response = await result.json()

    if(successful(response)){
        swal_success(response)
        .then(() => {window.location.reload()})
    }else{
        swal_error(response)
    }

    return response
}
// votes
export const vote_on_post = async (vote:vote):Promise<server_response> => {
    const opts = {
        method:POST,
            body:JSON.stringify(vote),
            headers:{
                "Content-Type":"application/json"
            }
    }
    const result = await fetch(root+"posts/vote", opts)
    const response:server_response = await result.json()
    
    !successful(response) ? swal_error(response):window.location.reload()

    return response
}

