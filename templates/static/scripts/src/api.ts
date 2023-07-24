import {
    user, 
    server_response,
    post
} from "./interfaces"
import { notify_modal } from "./controllers/notifications"


const root = "http://localhost:10000/api/"
const POST = "POST"
const GET = "GET"
const PUT = "PUT"
const DELETE = "DELETE"

const show_response = (r:server_response) => {
    notify_modal(r.message_type, r.message, "Nice Shirt!", "")
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
    show_response(response)

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
    show_response(response)

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
    show_response(response)

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
    show_response(response)

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
    show_response(response)

    return response
}

// posts
export const create_post = async (post:post, author:string):Promise<server_response> => {
    const opts = {
        method:POST,
            body:JSON.stringify(post),
            headers:{
                "Content-Type":"application/json"
            }
    }
    const result = await fetch(root+"posts/"+author, opts)
    const response:server_response = await result.json()
    show_response(response)

    return response
}

