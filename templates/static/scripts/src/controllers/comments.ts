import Swal from "sweetalert2"
import { comment, post } from "../interfaces";
import { create_comment } from "../api";
import { element_exists, json_data } from "./config";

const comment_creator = "comment-creator";
const comment_updater = "comment-updater";
const comment_deleter = "comment-deleter";

const comment_creator_handler = () => {
    const data = json_data() as post
    const on_click = (e:Event) => {
        e.preventDefault()
        Swal.fire({
            title: "Create Comment",
            html: `
                <div class="container">
                    <div class="row">
                        <textarea id="content" class="swal2-input" placeholder="Content..."></textarea>
                    </div>
                </div>
            `,
            confirmButtonText: "Comment",
            showCancelButton: true,
            preConfirm: async () => {
                const content_input = Swal.getPopup()!.querySelector("#content") as HTMLInputElement;
                // Retrieve user input and handle data
                const content: string = content_input ? content_input.value:""
                // Do something with the newUsername and newEmail, e.g., send it to the server
                const new_comment:comment = {
                    content:content,
                    _parent: data._id
                }
                console.log(new_comment)
                console.log(data._id)
                create_comment(new_comment)
                
            },
        });
    }

    document.getElementById(comment_creator)!.addEventListener("click", on_click)
}

const run = () => {
    if(element_exists(comment_creator)) {
        comment_creator_handler()
    }
}

export default run